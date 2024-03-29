package main

import (
	"github.com/siparisa/ServiceCatalog/internal"
	"github.com/siparisa/ServiceCatalog/internal/db"
	"log"
)

func main() {
	dbg, err := db.InitDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	sqlDB, err := dbg.DB()
	if err != nil {
		log.Fatalf("failed to get DB connection: %v", err)
	}
	defer sqlDB.Close()

	r := internal.SetupRouter(dbg)
	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
