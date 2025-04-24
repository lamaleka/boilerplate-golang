package main

import (
	"flag"

	"github.com/lamaleka/boilerplate-golang/internal/config"

	"github.com/pressly/goose/v3"
)

func main() {
	viper := config.NewViper()
	log := config.NewLogger(viper.Log)
	db := config.NewDatabase(viper.Db.App, log)
	dbInstance, _ := db.DB()
	defer dbInstance.Close()

	// Set the dialect to MSSQL
	if err := goose.SetDialect("mssql"); err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	// Define command-line flags
	upFlag := flag.Bool("up", false, "Apply migrations")
	downFlag := flag.Bool("down", false, "Roll back migrations")
	refreshFlag := flag.Bool("refresh", false, "Refresh migrations")
	resetFlag := flag.Bool("reset", false, "Reset migrations")

	flag.Parse()
	migrationDirectory := "././db/migrations"

	if *upFlag {
		if err := goose.Up(dbInstance, migrationDirectory); err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}
		log.Println("Migrations applied successfully!")
	} else if *downFlag {
		if err := goose.Down(dbInstance, migrationDirectory); err != nil {
			log.Fatalf("Failed to roll back migrations: %v", err)
		}
		log.Println("Migrations rolled back successfully!")
	} else if *refreshFlag {
		if err := goose.DownTo(dbInstance, migrationDirectory, 0); err != nil {
			log.Fatalf("Failed to refresh migrations: %v", err)
		}
		if err := goose.Up(dbInstance, migrationDirectory); err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}
		log.Println("Migrations refreshed successfully!")
	} else if *resetFlag {
		if err := goose.DownTo(dbInstance, migrationDirectory, 0); err != nil {
			log.Fatalf("Failed to reset migrations: %v", err)
		}
		log.Println("Migrations reset successfully!")
	} else {
		log.Println("No action specified. Use --up or --down.")
	}

}
