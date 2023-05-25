package main

import (
	"github.com/siparisa/ServiceCatalog/internal/db"
	"github.com/siparisa/ServiceCatalog/migrations"
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

	err = migrations.RollbackServicesTable(dbg)
	if err != nil {
		log.Fatalf("failed to rollback services table: %v", err)
	}

}
