package descriptions

import "fmt"

type mcuDescription int32

const (
	mcuDescriptionInvalidAddressOrScale mcuDescription = iota + 1
)

var mcuDescriptionToString = map[mcuDescription]string{
	mcuDescriptionInvalidAddressOrScale: "InvalidAddressOrScale",
}

var mcuDescriptionDescription = map[mcuDescription]string{
	mcuDescriptionInvalidAddressOrScale: "Invalid address or abnormal scale value is specified.",
}

func (d mcuDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", mcuDescriptionToString[d], d, mcuDescriptionDescription[d])
}
