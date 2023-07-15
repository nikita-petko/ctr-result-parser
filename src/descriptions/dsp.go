package descriptions

import "fmt"

type dspDescription int32

const (
	dspDescriptionOk dspDescription = iota
	dspDescriptionComponentNotLoaded
)

var dspDescriptionToString = map[dspDescription]string{
	dspDescriptionOk:                 "OK",
	dspDescriptionComponentNotLoaded: "ComponentNotLoaded",
}

var dspDescriptionDescription = map[dspDescription]string{
	dspDescriptionOk:                 "The operation was successful.",
	dspDescriptionComponentNotLoaded: "The component is not loaded.",
}

func (d dspDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", dspDescriptionToString[d], d, dspDescriptionDescription[d])
}
