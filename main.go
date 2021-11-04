package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"
    "time"
)

const version   string = "1.0.0"

type config struct {
    port    int
    apiKey  string
}

type application struct {
    version     string
    config      config
    infoLog     *log.Logger
    errorLog    *log.Logger
}

func (app *application) serve() error {
    server := &http.Server{
            Addr:               fmt.Sprintf(":%d", app.config.port),
        Handler:            app.routes(),
        IdleTimeout:        5 * time.Second,
        ReadHeaderTimeout:  3 * time.Second,
        ReadTimeout:        5 * time.Second,
        WriteTimeout:       5 * time.Second,
    }

    app.infoLog.Printf("Starting HTTP server on port %d\n", app.config.port)

    return server.ListenAndServe()
}

func main() {
    var (
        cfg config
        err error
    )

    infoLog := log.New(os.Stdout, "INFO ", log.Ldate | log.Ltime)
    errorLog := log.New(os.Stdout, "ERROR ", log.Ldate | log.Ltime | log.Lshortfile)

    cfg.port, err = strconv.Atoi(os.Getenv("PORT"))
    if err != nil {
        errorLog.Fatal(err)
    }
    cfg.apiKey = os.Getenv("API_KEY")

    app := &application{
        version:    version,
        config:     cfg,
        infoLog:    infoLog,
        errorLog:   errorLog,
    }

    if err = app.serve(); err != nil {
        errorLog.Fatal(err)
    }
}
