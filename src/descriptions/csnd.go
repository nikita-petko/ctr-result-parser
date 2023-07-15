package descriptions

import "fmt"

type csndDescription int32

const (
	csndDescriptionSleep csndDescription = iota + 1
	csndDescriptionInferiorPriority
)

var csndDescriptionToString = map[csndDescription]string{
	csndDescriptionSleep:            "Sleep",
	csndDescriptionInferiorPriority: "InferiorPriority",
}

var csndDescriptionDescription = map[csndDescription]string{
	csndDescriptionSleep:            "The operation is not supported in the current state.",
	csndDescriptionInferiorPriority: "The operation is not supported in the current state.",
}

func (d csndDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", csndDescriptionToString[d], d, csndDescriptionDescription[d])
}
