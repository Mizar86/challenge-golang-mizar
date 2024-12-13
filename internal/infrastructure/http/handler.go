package http

import (
	"carpooling-service/internal/application"
	"carpooling-service/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handlers struct {
	CarService *application.CarService

}

func NewHandlers(carService *application.CarService) *Handlers {
	return &Handlers{CarService: carService}
}


func (h *Handlers) Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "service is ready"})

}


func (h *Handlers) UpdateEVs(c *gin.Context) {
	var newCars []domain.Car
	if err := c.ShouldBindJSON(&newCars); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	carPointers := make([]*domain.Car, len(newCars))

	for i, car := range newCars {
		carPointers[i] = &car
	}

	h.CarService.ResetCars(carPointers)
	c.JSON(http.StatusOK, gin.H{"message": "EVs list updated"})
}

func (h *Handlers) AssignGroup(c *gin.Context) {
	var group domain.Group
	if err := c.ShouldBindJSON(&group); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.CarService.AssignGroupToCar(&group)
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{"message": err.Error()})

	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Group assigned to a car"})
	}
}

func (h *Handlers) DropOffGroup(c *gin.Context) {
	var req struct {
		ID int `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.CarService.DropOffGroup(req.ID)
	if err != nil {
		if err.Error() == "group not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusNoContent, nil)
	}
}

func (h *Handlers) LocateGroup(c *gin.Context) {
	var req struct {
		ID int `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car, err := h.CarService.LocateGroup(req.ID)
	if err != nil {
		if err.Error() == "group not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		
	} else if car == nil {
		c.JSON(http.StatusNoContent, nil)
	} else {
		c.JSON(http.StatusOK, gin.H{"car": car})
	}
}
