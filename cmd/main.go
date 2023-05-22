package main

import (
	"fmt"
	"github.com/siparisa/ServiceCatalog/internal"
	"github.com/siparisa/ServiceCatalog/internal/db"
	"github.com/siparisa/ServiceCatalog/internal/entity"
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

	// Insert a new serviceHandler record
	name1 := "s12"
	serviceHandler := entity.Service{
		Name:        &name1,
		Description: "A sample serviceHandler",
		Versions:    []string{"v1", "v2"},
	}
	result := dbg.Table("services").Create(&serviceHandler)
	if result.Error != nil {
		fmt.Println("errirr1111")
		panic(result.Error)
	}

	//// Retrieve the inserted serviceHandler
	//var retrievedService db.Service
	//result = dbg.Table("services").First(&retrievedService)
	//if result.Error != nil {
	//	fmt.Println("errirr222")
	//	panic(result.Error)
	//}
	//
	r := internal.SetupRouter(dbg)
	r.Run(":8080")

	fmt.Println("Hello, World!")
}
