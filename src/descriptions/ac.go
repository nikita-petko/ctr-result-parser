package descriptions

import "fmt"

type acDescription int32

const (
	acDescriptionWanConnected                  acDescription = 50
	acDescriptionLanConnected                  acDescription = 51
	acDescriptionUnnecessaryHotspotLogout      acDescription = 52
	acDescriptionProcessing                    acDescription = 70
	acDescriptionFailedStartup                 acDescription = 100
	acDescriptionFailedConnectAp               acDescription = 101
	acDescriptionFailedDhcp                    acDescription = 102
	acDescriptionConflictIpAddress             acDescription = 103
	acDescriptionInvalidKeyValue               acDescription = 104
	acDescriptionUnsupportAuthAlgorithm        acDescription = 105
	acDescriptionDenyUsbAp                     acDescription = 106
	acDescriptionInvalidDns                    acDescription = 150
	acDescriptionInvalidProxy                  acDescription = 151
	acDescriptionFailedConnTest                acDescription = 152
	acDescriptionUnsupportHotspot              acDescription = 200
	acDescriptionFailedHotspotAuthentification acDescription = 201
	acDescriptionFailedHotspotConnTest         acDescription = 202
	acDescriptionUnsupportPlace                acDescription = 203
	acDescriptionFailedHotspotLogout           acDescription = 204
	acDescriptionAlreadyConnectUnsupportAp     acDescription = 205
	acDescriptionFailedScan                    acDescription = 300
	acDescriptionAlreadyConnecting             acDescription = 301
	acDescriptionNotConnecting                 acDescription = 302
	acDescriptionAlreadyExclusive              acDescription = 303
	acDescriptionNotExclusive                  acDescription = 304
	acDescriptionInvalidLocation               acDescription = 305
	acDescriptionNotAgreeEula                  acDescription = 900
	acDescriptionWifiOff                       acDescription = 901
	acDescriptionBrokenNand                    acDescription = 902
	acDescriptionBrokenWireless                acDescription = 903
)

var acDescriptionToString = map[acDescription]string{
	acDescriptionWanConnected:                  "WanConnected",
	acDescriptionLanConnected:                  "LanConnected",
	acDescriptionUnnecessaryHotspotLogout:      "UnnecessaryHotspotLogout",
	acDescriptionProcessing:                    "Processing",
	acDescriptionFailedStartup:                 "FailedStartup",
	acDescriptionFailedConnectAp:               "FailedConnectAp",
	acDescriptionFailedDhcp:                    "FailedDhcp",
	acDescriptionConflictIpAddress:             "ConflictIpAddress",
	acDescriptionInvalidKeyValue:               "InvalidKeyValue",
	acDescriptionUnsupportAuthAlgorithm:        "UnsupportAuthAlgorithm",
	acDescriptionDenyUsbAp:                     "DenyUsbAp",
	acDescriptionInvalidDns:                    "InvalidDns",
	acDescriptionInvalidProxy:                  "InvalidProxy",
	acDescriptionFailedConnTest:                "FailedConnTest",
	acDescriptionUnsupportHotspot:              "UnsupportHotspot",
	acDescriptionFailedHotspotAuthentification: "FailedHotspotAuthentification",
	acDescriptionFailedHotspotConnTest:         "FailedHotspotConnTest",
	acDescriptionUnsupportPlace:                "UnsupportPlace",
	acDescriptionFailedHotspotLogout:           "FailedHotspotLogout",
	acDescriptionAlreadyConnectUnsupportAp:     "AlreadyConnectUnsupportAp",
	acDescriptionFailedScan:                    "FailedScan",
	acDescriptionAlreadyConnecting:             "AlreadyConnecting",
	acDescriptionNotConnecting:                 "NotConnecting",
	acDescriptionAlreadyExclusive:              "AlreadyExclusive",
	acDescriptionNotExclusive:                  "NotExclusive",
	acDescriptionInvalidLocation:               "InvalidLocation",
	acDescriptionNotAgreeEula:                  "NotAgreeEula",
	acDescriptionWifiOff:                       "WifiOff",
	acDescriptionBrokenNand:                    "BrokenNand",
	acDescriptionBrokenWireless:                "BrokenWireless",
}

var acDescriptionDescription = map[acDescription]string{
	acDescriptionWanConnected:                  "WAN connection established",
	acDescriptionLanConnected:                  "LAN connection established",
	acDescriptionUnnecessaryHotspotLogout:      "Hotspot authentication logout is unnecessary",
	acDescriptionProcessing:                    "Connecting",
	acDescriptionFailedStartup:                 "Wireless device initialization failed",
	acDescriptionFailedConnectAp:               "Failed to connect to the access point",
	acDescriptionFailedDhcp:                    "Failed to obtain IP address",
	acDescriptionConflictIpAddress:             "IP address conflict",
	acDescriptionInvalidKeyValue:               "Incorrect encryption key",
	acDescriptionUnsupportAuthAlgorithm:        "Unsupported authentication mode",
	acDescriptionDenyUsbAp:                     "A connection to a Nintendo Wi-Fi USB Connector was denied",
	acDescriptionInvalidDns:                    "Failed to resolve name",
	acDescriptionInvalidProxy:                  "Failed to connect to the proxy server",
	acDescriptionFailedConnTest:                "Failed to connect to the HTTP server",
	acDescriptionUnsupportHotspot:              "Connection from unsupported access point",
	acDescriptionFailedHotspotAuthentification: "Failed hotspot authentication",
	acDescriptionFailedHotspotConnTest:         "Failed to connect to to HTTP server after hotspot authentication",
	acDescriptionUnsupportPlace:                "Application cannot use the Internet from present physical location",
	acDescriptionFailedHotspotLogout:           "Failed hotspot authentication logout",
	acDescriptionAlreadyConnectUnsupportAp:     "Already connected to an unsupported access point",
	acDescriptionFailedScan:                    "Scan failed",
	acDescriptionAlreadyConnecting:             "Already connecting",
	acDescriptionNotConnecting:                 "Not connected",
	acDescriptionAlreadyExclusive:              "Already exclusive",
	acDescriptionNotExclusive:                  "Not exclusive",
	acDescriptionInvalidLocation:               "Invalid location",
	acDescriptionNotAgreeEula:                  "The user has not agreed to the EULA",
	acDescriptionWifiOff:                       "Disabled",
	acDescriptionBrokenNand:                    "The NAND device is damaged or malfunctioning",
	acDescriptionBrokenWireless:                "The wireless device is damaged or malfunctioning",
}

func (d acDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", acDescriptionToString[d], d, acDescriptionDescription[d])
}
