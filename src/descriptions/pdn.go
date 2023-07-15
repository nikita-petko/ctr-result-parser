package descriptions

import "fmt"

type pdnDescription int32

const (
	pdnDescriptionClockNotSupplied pdnDescription = iota + 1
)

var pdnDescriptionToString = map[pdnDescription]string{
	pdnDescriptionClockNotSupplied: "ClockNotSupplied",
}

var pdnDescriptionDescription = map[pdnDescription]string{
	pdnDescriptionClockNotSupplied: "Clock not supplied",
}

func (d pdnDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", pdnDescriptionToString[d], d, pdnDescriptionDescription[d])
}
