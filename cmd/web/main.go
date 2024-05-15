package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/zleitz/snippetbox/cmd/web/config"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	app := &config.Application{
		Logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home(app))
	mux.HandleFunc("GET /snippet/view/{id}", snippetView(app))
	mux.HandleFunc("GET /snippet/create", snippetCreate(app))
	mux.HandleFunc("POST /snippet/create", snippetCreatePost(app))

	app.Logger.Info("starting server", slog.String("addr", *addr))

	err := http.ListenAndServe(*addr, mux)
	app.Logger.Error(err.Error())
	os.Exit(1)
}
