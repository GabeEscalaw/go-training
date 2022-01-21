package speed

// TODO: define the 'Car' type struct

type Car struct {
	battery int 
	batteryDrain int
	speed int
	distance int
}

type Track struct {
	distance int
}


// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	newCar := Car{battery: 100,batteryDrain: batteryDrain, speed: speed}
	return newCar
}

// TODO: define the 'Track' type struct

// NewTrack created a new track
func NewTrack(distance int) Track {
	newTrack := Track{distance: distance}
	return newTrack
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	newDrive := Car{battery: car.battery, distance: car.distance, speed: car.speed, batteryDrain: car.batteryDrain}
	
	if car.battery > car.batteryDrain {
		newDrive = Car{battery: car.battery - car.batteryDrain, distance: car.distance + car.speed, speed: car.speed, batteryDrain: car.batteryDrain}
	} else {
		newDrive = Car{battery: car.battery, distance: car.distance, speed: car.speed, batteryDrain: car.batteryDrain}
	}
	
	return newDrive
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	newCar := Car{battery: car.battery, distance: car.distance, speed: car.speed, batteryDrain: car.batteryDrain}

	newTrack := Track{distance: track.distance}

	for newCar.battery > newCar.batteryDrain {
		newCar = Drive(newCar)
	}

	if newCar.distance >= newTrack.distance {
		return true
	} else {
		return false
	}
}
