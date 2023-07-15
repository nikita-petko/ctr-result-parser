package descriptions

import "fmt"

type nimDescription int32

const (
	nimDescriptionOk nimDescription = iota
	nimDescriptionNotTerminatedString
	nimDescriptionTooManySessions
	nimDescriptionUnknownEcError
	nimDescriptionUnknownState
	nimDescriptionAlreadyDownloaded
	nimDescriptionTerminating
	nimDescriptionBossFailed
	nimDescriptionUnsupportedHttpStatusCode
	nimDescriptionHashChanged
	nimDescriptionOverAgeLimit
	nimDescriptionNeedAgreeLatestEula
	nimDescriptionInvalidCountryCode
	nimDescriptionInvalidSerialNumber
	nimDescriptionNupCapacityOver
	nimDescriptionNoTitle
	nimDescriptionInvalidTitleList
	nimDescriptionInvalidTitle
	nimDescriptionInvalidSaveData
	nimDescriptionSuspending
	nimDescriptionInvalidCombination
	nimDescriptionInvalidTitleVersion
	nimDescriptionInvalidData
	nimDescriptionCannotSetIvs
	nimDescriptionAcNotConnected
	nimDescriptionNeedGetIvs
	nimDescriptionCannotGetIvs
	nimDescriptionInvalidLanguageCode
	nimDescriptionInvalidDownloadType
	nimDescriptionNotPrepared
	nimDescriptionNotSupported
	nimDescriptionTitleNotDownloaded
	nimDescriptionOutOfFilterMemory
	nimDescriptionOutOfDownloadTaskList
	nimDescriptionBufferNotEnough
	nimDescriptionAccountNotCreated
	nimDescriptionOutOfCatalogMemory
	nimDescriptionStandbyMode

	nimDescriptionEcErrorOk                  nimDescription = 100
	nimDescriptionEcErrorFail                nimDescription = 101
	nimDescriptionEcErrorNotSupported        nimDescription = 102
	nimDescriptionEcErrorInvalid             nimDescription = 104
	nimDescriptionEcErrorNomem               nimDescription = 105
	nimDescriptionEcErrorNotFound            nimDescription = 106
	nimDescriptionEcErrorNotBusy             nimDescription = 107
	nimDescriptionEcErrorBusy                nimDescription = 108
	nimDescriptionEcErrorOutmem              nimDescription = 110
	nimDescriptionEcErrorWsReport            nimDescription = 115
	nimDescriptionEcErrorEcard               nimDescription = 117
	nimDescriptionEcErrorOverflow            nimDescription = 118
	nimDescriptionEcErrorNetContent          nimDescription = 119
	nimDescriptionEcErrorMinReplenish        nimDescription = 132
	nimDescriptionEcErrorMaxBalance          nimDescription = 133
	nimDescriptionEcErrorWsResp              nimDescription = 134
	nimDescriptionEcErrorCanceled            nimDescription = 138
	nimDescriptionEcErrorAlready             nimDescription = 139
	nimDescriptionEcErrorTimeout             nimDescription = 140
	nimDescriptionEcErrorInit                nimDescription = 141
	nimDescriptionEcErrorWsRecv              nimDescription = 143
	nimDescriptionEcErrorNeedUpdate          nimDescription = 149
	nimDescriptionEcErrorConfig              nimDescription = 153
	nimDescriptionEcErrorBalanceNotCleared   nimDescription = 159
	nimDescriptionEcErrorStreamWrite         nimDescription = 160
	nimDescriptionEcErrorStreamWriteSize     nimDescription = 161
	nimDescriptionEcErrorEciBusy             nimDescription = 164
	nimDescriptionEcErrorEciStandby          nimDescription = 167
	nimDescriptionEcErrorInvalidToken        nimDescription = 168
	nimDescriptionEcErrorConnect             nimDescription = 169
	nimDescriptionEcErrorMigrateLimitReached nimDescription = 173
	nimDescriptionEcErrorAgeRestricted       nimDescription = 177
	nimDescriptionEcErrorAlreadyOwn          nimDescription = 179
	nimDescriptionEcErrorInsufficientFunds   nimDescription = 181
	nimDescriptionEcErrorResultModule        nimDescription = 299
)

var nimDescriptionToString = map[nimDescription]string{
	nimDescriptionOk:                        "Ok",
	nimDescriptionNotTerminatedString:       "NotTerminatedString",
	nimDescriptionTooManySessions:           "TooManySessions",
	nimDescriptionUnknownEcError:            "UnknownEcError",
	nimDescriptionUnknownState:              "UnknownState",
	nimDescriptionAlreadyDownloaded:         "AlreadyDownloaded",
	nimDescriptionTerminating:               "Terminating",
	nimDescriptionBossFailed:                "BossFailed",
	nimDescriptionUnsupportedHttpStatusCode: "UnsupportedHttpStatusCode",
	nimDescriptionHashChanged:               "HashChanged",
	nimDescriptionOverAgeLimit:              "OverAgeLimit",
	nimDescriptionNeedAgreeLatestEula:       "NeedAgreeLatestEula",
	nimDescriptionInvalidCountryCode:        "InvalidCountryCode",
	nimDescriptionInvalidSerialNumber:       "InvalidSerialNumber",
	nimDescriptionNupCapacityOver:           "NupCapacityOver",
	nimDescriptionNoTitle:                   "NoTitle",
	nimDescriptionInvalidTitleList:          "InvalidTitleList",
	nimDescriptionInvalidTitle:              "InvalidTitle",
	nimDescriptionInvalidSaveData:           "InvalidSaveData",
	nimDescriptionSuspending:                "Suspending",
	nimDescriptionInvalidCombination:        "InvalidCombination",
	nimDescriptionInvalidTitleVersion:       "InvalidTitleVersion",
	nimDescriptionInvalidData:               "InvalidData",
	nimDescriptionCannotSetIvs:              "CannotSetIvs",
	nimDescriptionAcNotConnected:            "AcNotConnected",
	nimDescriptionNeedGetIvs:                "NeedGetIvs",
	nimDescriptionCannotGetIvs:              "CannotGetIvs",
	nimDescriptionInvalidLanguageCode:       "InvalidLanguageCode",
	nimDescriptionInvalidDownloadType:       "InvalidDownloadType",
	nimDescriptionNotPrepared:               "NotPrepared",
	nimDescriptionNotSupported:              "NotSupported",
	nimDescriptionTitleNotDownloaded:        "TitleNotDownloaded",
	nimDescriptionOutOfFilterMemory:         "OutOfFilterMemory",
	nimDescriptionOutOfDownloadTaskList:     "OutOfDownloadTaskList",
	nimDescriptionBufferNotEnough:           "BufferNotEnough",
	nimDescriptionAccountNotCreated:         "AccountNotCreated",
	nimDescriptionOutOfCatalogMemory:        "OutOfCatalogMemory",
	nimDescriptionStandbyMode:               "StandbyMode",

	nimDescriptionEcErrorOk:                  "EcErrorOk",
	nimDescriptionEcErrorFail:                "EcErrorFail",
	nimDescriptionEcErrorNotSupported:        "EcErrorNotSupported",
	nimDescriptionEcErrorInvalid:             "EcErrorInvalid",
	nimDescriptionEcErrorNomem:               "EcErrorNomem",
	nimDescriptionEcErrorNotFound:            "EcErrorNotFound",
	nimDescriptionEcErrorNotBusy:             "EcErrorNotBusy",
	nimDescriptionEcErrorBusy:                "EcErrorBusy",
	nimDescriptionEcErrorOutmem:              "EcErrorOutmem",
	nimDescriptionEcErrorWsReport:            "EcErrorWsReport",
	nimDescriptionEcErrorEcard:               "EcErrorEcard",
	nimDescriptionEcErrorOverflow:            "EcErrorOverflow",
	nimDescriptionEcErrorNetContent:          "EcErrorNetContent",
	nimDescriptionEcErrorMinReplenish:        "EcErrorMinReplenish",
	nimDescriptionEcErrorMaxBalance:          "EcErrorMaxBalance",
	nimDescriptionEcErrorWsResp:              "EcErrorWsResp",
	nimDescriptionEcErrorCanceled:            "EcErrorCanceled",
	nimDescriptionEcErrorAlready:             "EcErrorAlready",
	nimDescriptionEcErrorTimeout:             "EcErrorTimeout",
	nimDescriptionEcErrorInit:                "EcErrorInit",
	nimDescriptionEcErrorWsRecv:              "EcErrorWsRecv",
	nimDescriptionEcErrorNeedUpdate:          "EcErrorNeedUpdate",
	nimDescriptionEcErrorConfig:              "EcErrorConfig",
	nimDescriptionEcErrorBalanceNotCleared:   "EcErrorBalanceNotCleared",
	nimDescriptionEcErrorStreamWrite:         "EcErrorStreamWrite",
	nimDescriptionEcErrorStreamWriteSize:     "EcErrorStreamWriteSize",
	nimDescriptionEcErrorEciBusy:             "EcErrorEciBusy",
	nimDescriptionEcErrorEciStandby:          "EcErrorEciStandby",
	nimDescriptionEcErrorInvalidToken:        "EcErrorInvalidToken",
	nimDescriptionEcErrorConnect:             "EcErrorConnect",
	nimDescriptionEcErrorMigrateLimitReached: "EcErrorMigrateLimitReached",
	nimDescriptionEcErrorAgeRestricted:       "EcErrorAgeRestricted",
	nimDescriptionEcErrorAlreadyOwn:          "EcErrorAlreadyOwn",
	nimDescriptionEcErrorInsufficientFunds:   "EcErrorInsufficientFunds",
	nimDescriptionEcErrorResultModule:        "EcErrorResultModule",
}

var nimDescriptionDescription = map[nimDescription]string{
	nimDescriptionOk:                        "Success.",
	nimDescriptionNotTerminatedString:       "The string is not terminated.",
	nimDescriptionTooManySessions:           "Too many sessions.",
	nimDescriptionUnknownEcError:            "Unknown EC error.",
	nimDescriptionUnknownState:              "Unknown state.",
	nimDescriptionAlreadyDownloaded:         "The title is already downloaded.",
	nimDescriptionTerminating:               "Terminating.",
	nimDescriptionBossFailed:                "Background Online System Service failed.",
	nimDescriptionUnsupportedHttpStatusCode: "Unsupported HTTP status code.",
	nimDescriptionHashChanged:               "The hash value has changed.",
	nimDescriptionOverAgeLimit:              "The user is over the age limit.",
	nimDescriptionNeedAgreeLatestEula:       "The user needs to agree to the latest EULA.",
	nimDescriptionInvalidCountryCode:        "The country code is invalid.",
	nimDescriptionInvalidSerialNumber:       "The serial number is invalid.",
	nimDescriptionNupCapacityOver:           "The NUP capacity is over.",
	nimDescriptionNoTitle:                   "The title does not exist.",
	nimDescriptionInvalidTitleList:          "The title list is invalid.",
	nimDescriptionInvalidTitle:              "The title is invalid.",
	nimDescriptionInvalidSaveData:           "The save data is invalid.",
	nimDescriptionSuspending:                "Suspending.",
	nimDescriptionInvalidCombination:        "The combination is invalid.",
	nimDescriptionInvalidTitleVersion:       "The title version is invalid.",
	nimDescriptionInvalidData:               "The data is invalid.",
	nimDescriptionCannotSetIvs:              "Cannot set IVS.",
	nimDescriptionAcNotConnected:            "The AC is not connected.",
	nimDescriptionNeedGetIvs:                "Need to get IVS.",
	nimDescriptionCannotGetIvs:              "Cannot get IVS.",
	nimDescriptionInvalidLanguageCode:       "The language code is invalid.",
	nimDescriptionInvalidDownloadType:       "The download type is invalid.",
	nimDescriptionNotPrepared:               "Not prepared.",
	nimDescriptionNotSupported:              "Not supported.",
	nimDescriptionTitleNotDownloaded:        "The title is not downloaded.",
	nimDescriptionOutOfFilterMemory:         "Out of filter memory.",
	nimDescriptionOutOfDownloadTaskList:     "Out of download task list.",
	nimDescriptionBufferNotEnough:           "The buffer is not enough.",
	nimDescriptionAccountNotCreated:         "The account is not created.",
	nimDescriptionOutOfCatalogMemory:        "Out of catalog memory.",
	nimDescriptionStandbyMode:               "Standby mode.",

	nimDescriptionEcErrorOk:                  "Success.",
	nimDescriptionEcErrorFail:                "Fail.",
	nimDescriptionEcErrorNotSupported:        "Not supported.",
	nimDescriptionEcErrorInvalid:             "Invalid.",
	nimDescriptionEcErrorNomem:               "No memory.",
	nimDescriptionEcErrorNotFound:            "Not found.",
	nimDescriptionEcErrorNotBusy:             "Not busy.",
	nimDescriptionEcErrorBusy:                "Busy.",
	nimDescriptionEcErrorOutmem:              "Out of memory.",
	nimDescriptionEcErrorWsReport:            "WS report.",
	nimDescriptionEcErrorEcard:               "E-card.",
	nimDescriptionEcErrorOverflow:            "Overflow.",
	nimDescriptionEcErrorNetContent:          "Net content.",
	nimDescriptionEcErrorMinReplenish:        "Minimum replenish.",
	nimDescriptionEcErrorMaxBalance:          "Maximum balance.",
	nimDescriptionEcErrorWsResp:              "WS response.",
	nimDescriptionEcErrorCanceled:            "Canceled.",
	nimDescriptionEcErrorAlready:             "Already.",
	nimDescriptionEcErrorTimeout:             "Timeout.",
	nimDescriptionEcErrorInit:                "Init.",
	nimDescriptionEcErrorWsRecv:              "WS receive.",
	nimDescriptionEcErrorNeedUpdate:          "Need update.",
	nimDescriptionEcErrorConfig:              "Config.",
	nimDescriptionEcErrorBalanceNotCleared:   "Balance not cleared.",
	nimDescriptionEcErrorStreamWrite:         "Stream write.",
	nimDescriptionEcErrorStreamWriteSize:     "Stream write size.",
	nimDescriptionEcErrorEciBusy:             "ECI busy.",
	nimDescriptionEcErrorEciStandby:          "ECI standby.",
	nimDescriptionEcErrorInvalidToken:        "Invalid token.",
	nimDescriptionEcErrorConnect:             "Connect.",
	nimDescriptionEcErrorMigrateLimitReached: "Migrate limit reached.",
	nimDescriptionEcErrorAgeRestricted:       "Age restricted.",
	nimDescriptionEcErrorAlreadyOwn:          "Already own.",
	nimDescriptionEcErrorInsufficientFunds:   "Insufficient funds.",
	nimDescriptionEcErrorResultModule:        "Result module.",
}

func (d nimDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", nimDescriptionToString[d], d, nimDescriptionDescription[d])
}
