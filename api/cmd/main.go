package main

import (
	"log"
	"order/db"
	"order/internal/dish"
	"order/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	dishRep := dish.NewRepository(dbConn.GetDB())
	dishSvc := dish.NewService(dishRep)
	dishHandler := dish.NewHandler(dishSvc)

	router.InitRouter(dishHandler)
	err = router.Start("0.0.0.0:8081")
	if err != nil {
		return
	}
}
