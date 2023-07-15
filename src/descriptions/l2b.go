package descriptions

import "fmt"

type l2bDescription int32

const (
	l2bDescriptionIsSleeping l2bDescription = iota + 1
	l2bDescriptionInvalidL2bNo
)

var l2bDescriptionToString = map[l2bDescription]string{
	l2bDescriptionIsSleeping:   "IsSleeping",
	l2bDescriptionInvalidL2bNo: "InvalidL2bNo",
}

var l2bDescriptionDescription = map[l2bDescription]string{
	l2bDescriptionIsSleeping:   "The L2B is sleeping.",
	l2bDescriptionInvalidL2bNo: "The L2B number is invalid.",
}

func (d l2bDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", l2bDescriptionToString[d], d, l2bDescriptionDescription[d])
}
