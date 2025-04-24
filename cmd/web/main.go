package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lamaleka/boilerplate-golang/internal/app"
	"github.com/lamaleka/boilerplate-golang/internal/config"
)

func main() {
	viper := config.NewViper()
	log := config.NewLogger(viper.Log)
	db := config.NewDatabase(viper.Db.App, log)
	validator := config.NewValidator()

	e := config.NewEcho(log)
	app.Bootstrap(&app.BootstrapConfig{
		DB:        db,
		App:       e,
		Log:       log,
		Validator: validator,
		Viper:     viper,
	})
	exitCh := make(chan os.Signal, 1)
	signal.Notify(
		exitCh,
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	webPort := viper.Web.Port
	go func() {
		defer func() { exitCh <- syscall.SIGTERM }()
		err := e.Start(fmt.Sprintf(":%d", webPort))
		if err != nil {
			log.Println(err)
		}
	}()
	<-exitCh
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	dbInstance, _ := db.DB()
	if err := dbInstance.Close(); err != nil {
		e.Logger.Fatal("Error closing database connection:", err)
	}

	fmt.Println("Server and database connection closed gracefully")
}
