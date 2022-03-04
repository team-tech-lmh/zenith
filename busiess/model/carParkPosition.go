package model

// 汽车和车位关系
type CarParkPositionRelWithCar struct {
	CarParkPositionRel
	Car *Car `json:"car"`
}

type CarParkPositionRelWithParkPosition struct {
	CarParkPositionRel
	ParkPosition *Car `json:"parkPosition"`
}

type CarParkPositionRel struct{}

func CarParkPositionRelSync() error {
	return nil
}

func CarParkPositionRelListForCar(carID string) (*[]CarParkPositionRelWithParkPosition, error) {
	return nil, nil
}

func CarParkPositionRelListForParkPosition(parkPositionID string) (*[]CarParkPositionRelWithCar, error) {
	return nil, nil
}
