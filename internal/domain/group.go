package domain

// Group represents a group of people requesting a ride.
type Group struct {
	ID   int `json:"id" binding:"required"`
    People int `json:"people" binding:"required,min=1"`
}
