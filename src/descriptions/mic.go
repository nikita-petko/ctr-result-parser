package descriptions

import "fmt"

type micDescription int32

const (
	micDescriptionMicShellClose micDescription = iota + 1
)

var micDescriptionToString = map[micDescription]string{
	micDescriptionMicShellClose: "MicShellClose",
}

var micDescriptionDescription = map[micDescription]string{
	micDescriptionMicShellClose: "The microphone cannot be used because the system is closed.",
}

func (d micDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", micDescriptionToString[d], d, micDescriptionDescription[d])
}
