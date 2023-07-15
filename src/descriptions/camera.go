package descriptions

import "fmt"

type cameraDescription int32

const (
	cameraDescriptionCameraIsSleeping cameraDescription = iota + 1
	cameraDescriptionCameraFatalError
)

var cameraDescriptionToString = map[cameraDescription]string{
	cameraDescriptionCameraIsSleeping: "CameraIsSleeping",
	cameraDescriptionCameraFatalError: "CameraFatalError",
}

var cameraDescriptionDescription = map[cameraDescription]string{
	cameraDescriptionCameraIsSleeping: "Camera is sleeping",
	cameraDescriptionCameraFatalError: "Camera fatal error",
}

func (d cameraDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", cameraDescriptionToString[d], d, cameraDescriptionDescription[d])
}
