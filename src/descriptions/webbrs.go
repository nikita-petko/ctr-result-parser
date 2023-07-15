package descriptions

import "fmt"

type webbrsDescription int32

const (
	webbrsDescriptionInvalidUrl webbrsDescription = iota + 1
	webbrsDescriptionUnavailableWebBrowser
	webbrsDescriptionLocalCommMode
)

var webbrsDescriptionToString = map[webbrsDescription]string{
	webbrsDescriptionInvalidUrl:            "InvalidUrl",
	webbrsDescriptionUnavailableWebBrowser: "UnavailableWebBrowser",
	webbrsDescriptionLocalCommMode:         "LocalCommMode",
}

var webbrsDescriptionDescription = map[webbrsDescription]string{
	webbrsDescriptionInvalidUrl:            "Invalid URL.",
	webbrsDescriptionUnavailableWebBrowser: "Unavailable.",
	webbrsDescriptionLocalCommMode:         "Cannot start during local communication.",
}

func (d webbrsDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", webbrsDescriptionToString[d], d, webbrsDescriptionDescription[d])
}
