package descriptions

import "fmt"

type sndDescription int32

const (
	sndDescriptionSndNoDspComponentLoaded sndDescription = iota + 1
)

var sndDescriptionToString = map[sndDescription]string{
	sndDescriptionSndNoDspComponentLoaded: "SndNoDspComponentLoaded",
}

var sndDescriptionDescription = map[sndDescription]string{
	sndDescriptionSndNoDspComponentLoaded: "No DSP component is loaded.",
}

func (d sndDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", sndDescriptionToString[d], d, sndDescriptionDescription[d])
}
