package router

import (
	"api/internal/booking"
	"api/internal/hotel"
	"api/internal/review"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(hotelHandler *hotel.Handler, bookingHandler *booking.Handler, reviewHandler *review.Handler) {
	r = gin.Default()

	r.Use(CORS())

	r.GET("/hotels/:id", hotelHandler.GetHotel)
	r.GET("/hotels", hotelHandler.GetAll)
	r.POST("/bookings", bookingHandler.CreateBooking)
	r.GET("/reviews/:hotelId", reviewHandler.GetReview)
	r.POST("/reviews", reviewHandler.CreateReview)
}

func Start(addr string) error {
	return r.Run(addr)
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
