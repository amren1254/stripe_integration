package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/amren1254/stripe_integration/constant"
	"github.com/joho/godotenv"
)

const app_version = "0.0.1"

type config struct {
	port   string
	env    string
	api    string
	stripe struct {
		secret string
		key    string
	}
	secretKey string
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
}

func (app *application) Serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", app.config.port),
		Handler:           app.route(),
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
	app.infoLog.Printf("Starting HTTP server in %s mode on port %s", app.config.env, app.config.port)
	return srv.ListenAndServe()
}

func main() {
	err := godotenv.Load(constant.ENV_PATH)
	if err != nil {
		log.Println("unable to load env")
	}
	cfg := config{
		port: os.Getenv("APP_PORT"),
		env:  os.Getenv("ENV_MODE"),
		api:  "http://localhost:9001",
		stripe: struct {
			secret string
			key    string
		}{
			secret: os.Getenv("STRIPE_SECRET"),
			key:    os.Getenv("STRIPE_KEY"),
		},
		secretKey: os.Getenv("TOKEN_SECRET_KEY"),
	}
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  app_version,
	}
	log.Println(app)
	err = app.Serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}
