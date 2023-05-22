package main

import (
	"fmt"
	"github.com/siparisa/ServiceCatalog/internal/db"
	"github.com/siparisa/ServiceCatalog/migrations"
)

func main() {
	dbg, err := db.InitDB()
	if err != nil {
		panic("failed to connect to database")
	}
	// defer db.Close()

	err = migrations.MigrateServicesTable(dbg)
	if err != nil {
		panic(err)
	}

	// Insert a new service record
	service := db.Service{
		Name:        "My Service",
		Description: "A sample service",
		Versions:    []string{"v1", "v2"},
	}
	result := dbg.Table("services").Create(&service)
	if result.Error != nil {
		fmt.Println("errirr1111")
		panic(result.Error)
	}

	// Retrieve the inserted service
	var retrievedService db.Service
	result = dbg.Table("services").First(&retrievedService)
	if result.Error != nil {
		fmt.Println("errirr222")
		panic(result.Error)
	}

	fmt.Println("Inserted Service:")
	fmt.Println(retrievedService.ID, retrievedService.Name, retrievedService.Description, retrievedService.Versions)

	fmt.Println("Hello, World!")
}
