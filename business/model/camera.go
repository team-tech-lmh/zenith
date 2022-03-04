package model

type CameraWithParkPosition struct {
	Camera
	ParkPosition *ParkPosition `json:"parkPosition"`
}

type Camera struct{}

func CameraSync(projID string) error {
	return nil
}

func CameraList(projID string) (*[]Camera, error) {
	return nil, nil
}
