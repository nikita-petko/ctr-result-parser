package descriptions

import "fmt"

type plDescription int32

const (
	plDescriptionSharedfontNotFound plDescription = iota + 2
	plDescriptionGamecoinDataReset  plDescription = iota + 99
	plDescriptionLackOfGamecoin
)

var plDescriptionToString = map[plDescription]string{
	plDescriptionSharedfontNotFound: "SharedfontNotFound",
	plDescriptionGamecoinDataReset:  "GamecoinDataReset",
	plDescriptionLackOfGamecoin:     "LackOfGamecoin",
}

var plDescriptionDescription = map[plDescription]string{
	plDescriptionSharedfontNotFound: "The shared font was not found.",
	plDescriptionGamecoinDataReset:  "The data was processed successfully, but was reset because there was a problem in the save data.",
	plDescriptionLackOfGamecoin:     "The user attempted to use more game coins than they own.",
}

func (d plDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", plDescriptionToString[d], d, plDescriptionDescription[d])
}
