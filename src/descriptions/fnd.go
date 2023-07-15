package descriptions

import "fmt"

type fndDescription int32

const (
	fndDescriptionInvalidTlsIndex fndDescription = 33
)

var fndDescriptionToString = map[fndDescription]string{
	fndDescriptionInvalidTlsIndex: "InvalidTlsIndex",
}

var fndDescriptionDescription = map[fndDescription]string{
	fndDescriptionInvalidTlsIndex: "An unallocated TLS index was specified.",
}

func (d fndDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", fndDescriptionToString[d], d, fndDescriptionDescription[d])
}
