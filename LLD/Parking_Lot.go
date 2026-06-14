package main

import "fmt"

type VehicleType int

const (
	Bike VehicleType = iota
	Car
	Cycle
)

type Vehicle struct {
	Number int
	Type   VehicleType
}

type ParkingSpot struct {
	ID         int
	SpotType   VehicleType
	IsOccupied bool
	Vehicle    *Vehicle
}

func (p *ParkingSpot) Park(v *Vehicle) {
	p.IsOccupied = true
	p.Vehicle = v
}

func (p *ParkingSpot) Unpark() {
	p.IsOccupied = false
	p.Vehicle = nil
}

type Ticket struct {
	ID      int
	VNo     int
	SpotID  int
}

type ParkingLot struct {
	Spots   []*ParkingSpot
	Tickets map[int]*Ticket
}

func (p *ParkingLot) ParkVehicle(v *Vehicle) *Ticket {

	for _, spot := range p.Spots {

		if !spot.IsOccupied &&
			spot.SpotType == v.Type {

			spot.Park(v)

			ticket := &Ticket{
				ID:     len(p.Tickets) + 1,
				VNo:    v.Number,
				SpotID: spot.ID,
			}

			p.Tickets[v.Number] = ticket

			return ticket
		}
	}

	return nil
}

func (p *ParkingLot) UnparkVehicle(v *Vehicle) *Ticket {

	for _, spot := range p.Spots {

		if spot.IsOccupied &&
			spot.Vehicle != nil &&
			spot.Vehicle.Number == v.Number {

			spot.Unpark()

			ticket := p.Tickets[v.Number]

			delete(p.Tickets, v.Number)

			return ticket
		}
	}

	return nil
}

func (p *ParkingLot) AvailableSpots(vType VehicleType) int {

	count := 0

	for _, spot := range p.Spots {

		if !spot.IsOccupied &&
			spot.SpotType == vType {

			count++
		}
	}

	return count
}

func main() {

	spots := []*ParkingSpot{
		{ID: 1, SpotType: Bike},
		{ID: 2, SpotType: Cycle},
		{ID: 3, SpotType: Car},
		{ID: 4, SpotType: Car},
	}

	parkingLot := &ParkingLot{
		Spots:   spots,
		Tickets: make(map[int]*Ticket),
	}

	fmt.Println("Parking Vehicles")

	car := &Vehicle{
		Number: 1,
		Type:   Car,
	}

	bike := &Vehicle{
		Number: 11,
		Type:   Bike,
	}

	cycle := &Vehicle{
		Number: 111,
		Type:   Cycle,
	}

	fmt.Println(parkingLot.ParkVehicle(car))
	fmt.Println(parkingLot.ParkVehicle(bike))
	fmt.Println(parkingLot.ParkVehicle(cycle))

	fmt.Println()

	fmt.Println("Available Car Spots:",
		parkingLot.AvailableSpots(Car))

	fmt.Println()

	fmt.Println("Unparking Car")

	fmt.Println(parkingLot.UnparkVehicle(car))

	fmt.Println()

	fmt.Println("Available Car Spots:",
		parkingLot.AvailableSpots(Car))
}