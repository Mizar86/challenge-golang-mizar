package application

import (
	"carpooling-service/internal/domain"
	"errors"
)

// CarService provides fuctionalities for managing cars and assigning groups to cars
type CarService struct {
	cars   []*domain.Car
	groups map[int]*domain.Group
	queue  []*domain.Group
	carGroupMap map[int]int
}

// NewCarService initializes a new instance of CarService with a list of cars.
func NewCarService(cars []*domain.Car) *CarService {
	return &CarService{
		cars:   cars,
		groups: make(map[int]*domain.Group),
		queue:  []*domain.Group{},
		carGroupMap: make(map[int]int),
	}
}


// AssignGroupToCar assigns a group to the first car with enough capacity.
// If no car is available, the group is added to the waiting queue.
func (cs *CarService) AssignGroupToCar(group *domain.Group) error {
    for _, car := range cs.cars {
        if car.CanAccommodate(group.People) {
            car.AddGroup(group.People)
            cs.groups[group.ID] = group
            cs.carGroupMap[group.ID] = car.ID
            return nil
        }
    }

    cs.queue = append(cs.queue, group)
    cs.groups[group.ID] = group
    return errors.New("no available car for the group, added to queue")
}

// DropOffGroup removes a group from its assigned car and Updates the available capaciity.
// Returns an error if the group is not found.
func (cs *CarService) DropOffGroup(groupID int) error {
    group, exists := cs.groups[groupID]
    if !exists {
        return errors.New("group not found")
    }
    delete(cs.groups, groupID)

    carID, assigned := cs.carGroupMap[groupID]
    if !assigned {
        return nil
    }

    for _, car := range cs.cars {
        if car.ID == carID {
            car.RemoveGroup(group.People)
            delete(cs.carGroupMap, groupID)
            return nil
        }
    }
    return nil
}


// LocateGroup finds the car where a specific group is currently located.
// Returns nil if the group is found in the queue or not assigned to any car.
func (cs *CarService) LocateGroup(groupID int) (*domain.Car, error) {
    carID, exists := cs.carGroupMap[groupID]
    if !exists {
        return nil, errors.New("group not found")
    }

    for _, car := range cs.cars {
        if car.ID == carID {
            return car, nil
        }
    }
    return nil, errors.New("car not found for the group")
}

// ResetCars resets the list of cars, clears all groups, and resets the queue.
// Useful for scenarios where the fleet is updated.
func (cs *CarService) ResetCars(newCars []*domain.Car) {
	cs.cars = newCars
	cs.groups = make(map[int]*domain.Group)
	cs.queue = []*domain.Group{}
}