package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ericklima-ca/mailer_api/http_server"
	"github.com/ericklima-ca/mailer_api/mailer"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	var (
		HOST_SMTP = os.Getenv("HOST_SMTP")
	)

	ms := mailer.MailerService{
		HostPort: HOST_SMTP,
	}

	server := http_server.HTTPServer{
		MailerService: ms,
	}
	app := http_server.NewServer(server)

	httpSrv := &http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil &&
			errors.Is(err, http.ErrServerClosed) {
			log.Printf("Error on %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server exiting...")
}
