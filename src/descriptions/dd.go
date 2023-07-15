package descriptions

import "fmt"

type ddDescription int32

const (
	ddDescriptionInvalidSelection ddDescription = iota + 1
	ddDescriptionInnaccessiblePage
	ddDescriptionManualResetEventRequired
	ddDescriptionMaxHandle
)

var ddDescriptionToString = map[ddDescription]string{
	ddDescriptionInvalidSelection:         "InvalidSelection",
	ddDescriptionInnaccessiblePage:        "InnaccessiblePage",
	ddDescriptionManualResetEventRequired: "ManualResetEventRequired",
	ddDescriptionMaxHandle:                "MaxHandle",
}

var ddDescriptionDescription = map[ddDescription]string{
	ddDescriptionInvalidSelection:         "Bad selection",
	ddDescriptionInnaccessiblePage:        "Contains inaccessible pages",
	ddDescriptionManualResetEventRequired: "Manual reset event required",
	ddDescriptionMaxHandle:                "maximum number of handles reached",
}

func (d ddDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", ddDescriptionToString[d], d, ddDescriptionDescription[d])
}
