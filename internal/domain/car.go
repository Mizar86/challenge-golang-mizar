package domain


// Car represents a vehicle in the fleet.
type Car struct {
	ID int `json:"id" binding:"required,min=1"`
	Seats  int `json:"seats" binding:"required,min=1"`
	OccupiedSeats int 
}

// CanAccommodate checks if the car can accommodate a group of a given size.
func (c *Car) CanAccommodate(groupSize int) bool {
	return c.Seats-c.OccupiedSeats >= groupSize
}

// AddGroup assigns a group to the car if there is enough available space.

func (c *Car) AddGroup(groupSize int) {
	if c.CanAccommodate(groupSize) {
		c.OccupiedSeats += groupSize
	}
}

// RemoveGroup removes a group from the car, freeing up occupied seats.
func (c *Car) RemoveGroup(groupSize int) {
	c.OccupiedSeats -= groupSize
	if c.OccupiedSeats < 0 {
		c.OccupiedSeats = 0
	}
}