package model

// 汽车
type CarWithParkPosition struct {
	Car
	ParkPosition *ParkPosition `json:"position"`
	InvalidAt    int64         `json:"invalidAt"`
}

type Car struct{}

func CarSync(projID string) error {
	return nil
}

func CarList(projID string) (*[]Car, error) {
	return nil, nil
}
