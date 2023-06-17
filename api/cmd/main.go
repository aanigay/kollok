package main

import (
	"api/db"
	"api/internal/booking"
	"api/internal/hotel"
	"api/internal/review"
	"api/router"
	"log"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	hotelRep := hotel.NewRepository(dbConn.GetDB())
	hotelSvc := hotel.NewService(hotelRep)
	hotelHandler := hotel.NewHandler(hotelSvc)

	bookingRep := booking.NewRepository(dbConn.GetDB())
	bookingSvc := booking.NewService(bookingRep)
	bookingHandler := booking.NewHandler(bookingSvc)

	reviewRep := review.NewRepository(dbConn.GetDB())
	reviewSvc := review.NewService(reviewRep)
	reviewHandler := review.NewHandler(reviewSvc)

	router.InitRouter(hotelHandler, bookingHandler, reviewHandler)
	err = router.Start("0.0.0.0:8081")
	if err != nil {
		return
	}
}
