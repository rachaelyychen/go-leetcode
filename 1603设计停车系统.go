package main

type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big, medium, small}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.big == 0 {
			return false
		}
		this.big -= 1
	case 2:
		if this.medium == 0 {
			return false
		}
		this.medium -= 1
	case 3:
		if this.small == 0 {
			return false
		}
		this.small -= 1
	}
	return true
}
