package http

import (
	"carpooling-service/internal/application"
	"carpooling-service/internal/domain"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	//Define some example cars 
    cars := []*domain.Car{
        {ID: 1, Seats: 6},
        {ID: 2, Seats: 4},
        {ID: 3, Seats: 5},
    }

    // Create a car service using the list of cars.
    carService := application.NewCarService(cars)

    // Create the route handlers, passing the car service as a dependency
    handlers := NewHandlers(carService)

    // Defined routes.
	r.GET("/status", handlers.Status)
	r.PUT("/evs", handlers.UpdateEVs)
	r.POST("/journey", handlers.AssignGroup)
	r.POST("/dropoff", handlers.DropOffGroup)
	r.POST("/locate", handlers.LocateGroup)


	r.Run(":80") 
}
