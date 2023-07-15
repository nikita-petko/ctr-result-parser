package descriptions

import "fmt"

type ptmDescription int32

const (
	ptmDescriptionInvalidSystemtime ptmDescription = iota + 1
	ptmDescriptionNoalarm
	ptmDescriptionOverwritealarm
	ptmDescriptionFileerror
	ptmDescriptionNotSleeping
	ptmDescriptionInvalidTrigger
	ptmDescriptionMcuFatalError
)

var ptmDescriptionToString = map[ptmDescription]string{
	ptmDescriptionInvalidSystemtime: "InvalidSystemtime",
	ptmDescriptionNoalarm:           "Noalarm",
	ptmDescriptionOverwritealarm:    "Overwritealarm",
	ptmDescriptionFileerror:         "Fileerror",
	ptmDescriptionNotSleeping:       "NotSleeping",
	ptmDescriptionInvalidTrigger:    "InvalidTrigger",
	ptmDescriptionMcuFatalError:     "McuFatalError",
}

var ptmDescriptionDescription = map[ptmDescription]string{
	ptmDescriptionInvalidSystemtime: "The system time is invalid.",
	ptmDescriptionNoalarm:           "Indicates that the alarm has not been configured.",
	ptmDescriptionOverwritealarm:    "Indicates that the alarm has already been configured.",
	ptmDescriptionFileerror:         "Indicates that an error occurred while accessing the file.",
	ptmDescriptionNotSleeping:       "Indicates that the system is not in sleep mode.",
	ptmDescriptionInvalidTrigger:    "Indicates that the trigger is invalid.",
	ptmDescriptionMcuFatalError:     "Indicates that a fatal error occurred in the MCU.",
}

func (d ptmDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", ptmDescriptionToString[d], d, ptmDescriptionDescription[d])
}
