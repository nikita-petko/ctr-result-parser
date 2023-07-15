package descriptions

import "fmt"

type ecDescription int32

const (
	ecDescriptionTooManyContents ecDescription = iota + 128
	ecDescriptionTooManyTitles
	ecDescriptionNotInitialized
	ecDescriptionTitleInvalid
	ecDescriptionInvalidSelectionFilter
	ecDescriptionNotInService
	ecDescriptionNotImplemented
	ecDescriptionDataTitleNotFound
	ecDescriptionDataTitleContentInfoNotFound
	ecDescriptionDataTitleContentNotFound
	ecDescriptionMetaDataTitleNotFound
	ecDescriptionMetaDataInvalidFormat
	ecDescriptionMetaDataIconNotFound
	ecDescriptionMetaDataContentNotFound
	ecDescriptionOutOfFilterMemory
	ecDescriptionCatalogInvalidArgument
	ecDescriptionSessionPreparedTitleNotFound
	ecDescriptionAppletCloseApplicationRequested
	ecDescriptionParentalControlShopUseRestricted
	ecDescriptionParentalControlWrongPinCode
	ecDescriptionDataTitleNotOwned
	ecDescriptionServiceTitleNotOwned
	ecDescriptionAccountNotCreated
	ecDescriptionNotSupported
	ecDescriptionTicketEnvelopeSizeInvalid
	ecDescriptionDataTitleInUse
	ecDescriptionOutOfCatalogMemory
	ecDescriptionDataTitleContentNotDeletable
	ecDescriptionTooManyItems
	ecDescriptionPriceFormatInvalid
	ecDescriptionUpdatedVersionRetrieved ecDescription = iota + 192 - 30
	ecDescriptionSessionInvalid
	ecDescriptionNewSessionRequired
	ecDescriptionEcardStatusRedeemedError ecDescription = iota + 256 - 33
	ecDescriptionEcardStatusRevokedError
	ecDescriptionEcardStatusInactiveError
	ecDescriptionEcardStatusUnknownError
	ecDescriptionEcardTypeLegacyPointsError
	ecDescriptionEcardInvalidIdError
	ecDescriptionEcardTypeCacheError
	ecDescriptionEcardTypeUnknownError
	ecDescriptionEcardNoRedeemableItemsError
	ecDescriptionFsMediaWriteProtectedError
	ecDescriptionInfraReportError
	ecDescriptionInfraReceiveError
	ecDescriptionInfraInvalidResponseError
	ecDescriptionInfraNeedsUpdateError
	ecDescriptionInfraTimeoutError
	ecDescriptionInfraBusyError
	ecDescriptionInfraInMaintenanceError
	ecDescriptionNetworkTimeoutError
	ecDescriptionNetworkInMaintenanceError
	ecDescriptionOutOfCatalogMemoryError
	ecDescriptionOutOfSystemMemoryError
	ecDescriptionFsMediaAccessFailedError
	ecDescriptionNoEnoughSpaceError
	ecDescriptionAcNotConnectedError
	ecDescriptionAppletUnknownError
	ecDescriptionUnknownError
	ecDescriptionEcAppletNotFoundError
	ecDescriptionInfraNeedsReconnectError ecDescription = iota + 320 - 62
	ecDescriptionAppletCancelled          ecDescription = iota + 1 - 61
	ecDescriptionAppletHomeButton
	ecDescriptionAppletSoftwareReset
	ecDescriptionAppletPowerButton
	ecDescriptionAppletNotSupportedError
	ecDescriptionAppletNotEnoughSpaceToDownloadContent
	ecDescriptionAppletNecessaryAttributeNotFound
	ecDescriptionAppletMetadataNotFound
	ecDescriptionAppletTooManyContents
	ecDescriptionAppletAlreadyPurchased
	ecDescriptionAppletItemUnpurchasable
	ecDescriptionAppletItemNotFound
	ecDescriptionAppletAlreadyDownloaded
	ecDescriptionAppletRedeemableItemNotFound
	ecDescriptionAppletAlreadyRedeemedItem
	ecDescriptionAppletShopServerError
	ecDescriptionAppletImportFailed
	ecDescriptionAppletWifiOffError
	ecDescriptionAppletFailedConnectError
	ecDescriptionAppletEcardUnavailable
	ecDescriptionAppletInvalidRivtoken
	ecDescriptionAppletInvalidParameter
	ecDescriptionAppletFailed
	ecDescriptionAppletFailedError
	ecDescriptionAppletWrongPinCodeError
	ecDescriptionAppletDuplicateContentIndex
	ecDescriptionAppletFailedLoginError
	ecDescriptionAppletDataTitleInUse
	ecDescriptionAppletNeedsDatatitleUpdate = iota + 64 - 89
	ecDescriptionAppletDatatitleNotImported
	ecDescriptionAppletRequiresReConnect
	ecDescriptionAppletNeedsBalanceUpdate
	ecDescriptionAppletNeedsSystemUpdate = iota + 96 - 93
	ecDescriptionAppletStandbyMode
	ecDescriptionAppletShopServiceTerminated
	ecDescriptionAppletShopServiceUnavailable
	ecDescriptionAppletCountryChanged
	ecDescriptionAppletNotAgreeEulaError
	ecDescriptionAppletTooManyDatatitles
	ecDescriptionAppletNotEnoughSpaceToDownloadDatatitle
	ecDescriptionAppletSdNotInserted
	ecDescriptionAppletSdWriteProtected
	ecDescriptionAppletSdBroken
	ecDescriptionAppletSdEjected
	ecDescriptionAppletInvalidVersion
	ecDescriptionAppletSdAccessError
)

var ecDescriptionToString = map[ecDescription]string{
	ecDescriptionTooManyContents:                         "TooManyContents",
	ecDescriptionTooManyTitles:                           "TooManyTitles",
	ecDescriptionNotInitialized:                          "NotInitialized",
	ecDescriptionTitleInvalid:                            "TitleInvalid",
	ecDescriptionInvalidSelectionFilter:                  "InvalidSelectionFilter",
	ecDescriptionNotInService:                            "NotInService",
	ecDescriptionNotImplemented:                          "NotImplemented",
	ecDescriptionDataTitleNotFound:                       "DataTitleNotFound",
	ecDescriptionDataTitleContentInfoNotFound:            "DataTitleContentInfoNotFound",
	ecDescriptionDataTitleContentNotFound:                "DataTitleContentNotFound",
	ecDescriptionMetaDataTitleNotFound:                   "MetaDataTitleNotFound",
	ecDescriptionMetaDataInvalidFormat:                   "MetaDataInvalidFormat",
	ecDescriptionMetaDataIconNotFound:                    "MetaDataIconNotFound",
	ecDescriptionMetaDataContentNotFound:                 "MetaDataContentNotFound",
	ecDescriptionOutOfFilterMemory:                       "OutOfFilterMemory",
	ecDescriptionCatalogInvalidArgument:                  "CatalogInvalidArgument",
	ecDescriptionSessionPreparedTitleNotFound:            "SessionPreparedTitleNotFound",
	ecDescriptionAppletCloseApplicationRequested:         "AppletCloseApplicationRequested",
	ecDescriptionParentalControlShopUseRestricted:        "ParentalControlShopUseRestricted",
	ecDescriptionParentalControlWrongPinCode:             "ParentalControlWrongPinCode",
	ecDescriptionDataTitleNotOwned:                       "DataTitleNotOwned",
	ecDescriptionServiceTitleNotOwned:                    "ServiceTitleNotOwned",
	ecDescriptionAccountNotCreated:                       "AccountNotCreated",
	ecDescriptionNotSupported:                            "NotSupported",
	ecDescriptionTicketEnvelopeSizeInvalid:               "TicketEnvelopeSizeInvalid",
	ecDescriptionDataTitleInUse:                          "DataTitleInUse",
	ecDescriptionOutOfCatalogMemory:                      "OutOfCatalogMemory",
	ecDescriptionDataTitleContentNotDeletable:            "DataTitleContentNotDeletable",
	ecDescriptionTooManyItems:                            "TooManyItems",
	ecDescriptionPriceFormatInvalid:                      "PriceFormatInvalid",
	ecDescriptionUpdatedVersionRetrieved:                 "UpdatedVersionRetrieved",
	ecDescriptionSessionInvalid:                          "SessionInvalid",
	ecDescriptionNewSessionRequired:                      "NewSessionRequired",
	ecDescriptionEcardStatusRedeemedError:                "EcardStatusRedeemedError",
	ecDescriptionEcardStatusRevokedError:                 "EcardStatusRevokedError",
	ecDescriptionEcardStatusInactiveError:                "EcardStatusInactiveError",
	ecDescriptionEcardStatusUnknownError:                 "EcardStatusUnknownError",
	ecDescriptionEcardTypeLegacyPointsError:              "EcardTypeLegacyPointsError",
	ecDescriptionEcardInvalidIdError:                     "EcardInvalidIdError",
	ecDescriptionEcardTypeCacheError:                     "EcardTypeCacheError",
	ecDescriptionEcardTypeUnknownError:                   "EcardTypeUnknownError",
	ecDescriptionEcardNoRedeemableItemsError:             "EcardNoRedeemableItemsError",
	ecDescriptionFsMediaWriteProtectedError:              "FsMediaWriteProtectedError",
	ecDescriptionInfraReportError:                        "InfraReportError",
	ecDescriptionInfraReceiveError:                       "InfraReceiveError",
	ecDescriptionInfraInvalidResponseError:               "InfraInvalidResponseError",
	ecDescriptionInfraNeedsUpdateError:                   "InfraNeedsUpdateError",
	ecDescriptionInfraTimeoutError:                       "InfraTimeoutError",
	ecDescriptionInfraBusyError:                          "InfraBusyError",
	ecDescriptionInfraInMaintenanceError:                 "InfraInMaintenanceError",
	ecDescriptionNetworkTimeoutError:                     "NetworkTimeoutError",
	ecDescriptionNetworkInMaintenanceError:               "NetworkInMaintenanceError",
	ecDescriptionOutOfCatalogMemoryError:                 "OutOfCatalogMemoryError",
	ecDescriptionOutOfSystemMemoryError:                  "OutOfSystemMemoryError",
	ecDescriptionFsMediaAccessFailedError:                "FsMediaAccessFailedError",
	ecDescriptionNoEnoughSpaceError:                      "NoEnoughSpaceError",
	ecDescriptionAcNotConnectedError:                     "AcNotConnectedError",
	ecDescriptionAppletUnknownError:                      "AppletUnknownError",
	ecDescriptionUnknownError:                            "UnknownError",
	ecDescriptionEcAppletNotFoundError:                   "EcAppletNotFoundError",
	ecDescriptionInfraNeedsReconnectError:                "InfraNeedsReconnectError",
	ecDescriptionAppletCancelled:                         "AppletCancelled",
	ecDescriptionAppletHomeButton:                        "AppletHomeButton",
	ecDescriptionAppletSoftwareReset:                     "AppletSoftwareReset",
	ecDescriptionAppletPowerButton:                       "AppletPowerButton",
	ecDescriptionAppletNotSupportedError:                 "AppletNotSupportedError",
	ecDescriptionAppletNotEnoughSpaceToDownloadContent:   "AppletNotEnoughSpaceToDownloadContent",
	ecDescriptionAppletNecessaryAttributeNotFound:        "AppletNecessaryAttributeNotFound",
	ecDescriptionAppletMetadataNotFound:                  "AppletMetadataNotFound",
	ecDescriptionAppletTooManyContents:                   "AppletTooManyContents",
	ecDescriptionAppletAlreadyPurchased:                  "AppletAlreadyPurchased",
	ecDescriptionAppletItemUnpurchasable:                 "AppletItemUnpurchasable",
	ecDescriptionAppletItemNotFound:                      "AppletItemNotFound",
	ecDescriptionAppletAlreadyDownloaded:                 "AppletAlreadyDownloaded",
	ecDescriptionAppletRedeemableItemNotFound:            "AppletRedeemableItemNotFound",
	ecDescriptionAppletAlreadyRedeemedItem:               "AppletAlreadyRedeemedItem",
	ecDescriptionAppletShopServerError:                   "AppletShopServerError",
	ecDescriptionAppletImportFailed:                      "AppletImportFailed",
	ecDescriptionAppletWifiOffError:                      "AppletWifiOffError",
	ecDescriptionAppletFailedConnectError:                "AppletFailedConnectError",
	ecDescriptionAppletEcardUnavailable:                  "AppletEcardUnavailable",
	ecDescriptionAppletInvalidRivtoken:                   "AppletInvalidRivtoken",
	ecDescriptionAppletInvalidParameter:                  "AppletInvalidParameter",
	ecDescriptionAppletFailed:                            "AppletFailed",
	ecDescriptionAppletFailedError:                       "AppletFailedError",
	ecDescriptionAppletWrongPinCodeError:                 "AppletWrongPinCodeError",
	ecDescriptionAppletDuplicateContentIndex:             "AppletDuplicateContentIndex",
	ecDescriptionAppletFailedLoginError:                  "AppletFailedLoginError",
	ecDescriptionAppletDataTitleInUse:                    "AppletDataTitleInUse",
	ecDescriptionAppletNeedsDatatitleUpdate:              "AppletNeedsDatatitleUpdate",
	ecDescriptionAppletDatatitleNotImported:              "AppletDatatitleNotImported",
	ecDescriptionAppletRequiresReConnect:                 "AppletRequiresReConnect",
	ecDescriptionAppletNeedsBalanceUpdate:                "AppletNeedsBalanceUpdate",
	ecDescriptionAppletNeedsSystemUpdate:                 "AppletNeedsSystemUpdate",
	ecDescriptionAppletStandbyMode:                       "AppletStandbyMode",
	ecDescriptionAppletShopServiceTerminated:             "AppletShopServiceTerminated",
	ecDescriptionAppletShopServiceUnavailable:            "AppletShopServiceUnavailable",
	ecDescriptionAppletCountryChanged:                    "AppletCountryChanged",
	ecDescriptionAppletNotAgreeEulaError:                 "AppletNotAgreeEulaError",
	ecDescriptionAppletTooManyDatatitles:                 "AppletTooManyDatatitles",
	ecDescriptionAppletNotEnoughSpaceToDownloadDatatitle: "AppletNotEnoughSpaceToDownloadDatatitle",
	ecDescriptionAppletSdNotInserted:                     "AppletSdNotInserted",
	ecDescriptionAppletSdWriteProtected:                  "AppletSdWriteProtected",
	ecDescriptionAppletSdBroken:                          "AppletSdBroken",
	ecDescriptionAppletSdEjected:                         "AppletSdEjected",
	ecDescriptionAppletInvalidVersion:                    "AppletInvalidVersion",
	ecDescriptionAppletSdAccessError:                     "AppletSdAccessError",
}

var ecDescriptionDescription = map[ecDescription]string{
	ecDescriptionTooManyContents:                         "Too many contents.",
	ecDescriptionTooManyTitles:                           "Too many titles.",
	ecDescriptionNotInitialized:                          "Not initialized.",
	ecDescriptionTitleInvalid:                            "Title invalid.",
	ecDescriptionInvalidSelectionFilter:                  "Invalid selection filter.",
	ecDescriptionNotInService:                            "Not in service.",
	ecDescriptionNotImplemented:                          "Not implemented.",
	ecDescriptionDataTitleNotFound:                       "Data title not found.",
	ecDescriptionDataTitleContentInfoNotFound:            "Data title content info not found.",
	ecDescriptionDataTitleContentNotFound:                "Data title content not found.",
	ecDescriptionMetaDataTitleNotFound:                   "Meta data title not found.",
	ecDescriptionMetaDataInvalidFormat:                   "Meta data invalid format.",
	ecDescriptionMetaDataIconNotFound:                    "Meta data icon not found.",
	ecDescriptionMetaDataContentNotFound:                 "Meta data content not found.",
	ecDescriptionOutOfFilterMemory:                       "Out of filter memory.",
	ecDescriptionCatalogInvalidArgument:                  "Catalog invalid argument.",
	ecDescriptionSessionPreparedTitleNotFound:            "Session prepared title not found.",
	ecDescriptionAppletCloseApplicationRequested:         "Applet close application requested.",
	ecDescriptionParentalControlShopUseRestricted:        "Parental control shop use restricted.",
	ecDescriptionParentalControlWrongPinCode:             "Parental control wrong pin code.",
	ecDescriptionDataTitleNotOwned:                       "Data title not owned.",
	ecDescriptionServiceTitleNotOwned:                    "Service title not owned.",
	ecDescriptionAccountNotCreated:                       "Account not created.",
	ecDescriptionNotSupported:                            "Not supported.",
	ecDescriptionTicketEnvelopeSizeInvalid:               "Ticket envelope size invalid.",
	ecDescriptionDataTitleInUse:                          "Data title in use.",
	ecDescriptionOutOfCatalogMemory:                      "Out of catalog memory.",
	ecDescriptionDataTitleContentNotDeletable:            "Data title content not deletable.",
	ecDescriptionTooManyItems:                            "Too many items.",
	ecDescriptionPriceFormatInvalid:                      "Price format invalid.",
	ecDescriptionUpdatedVersionRetrieved:                 "Updated version retrieved.",
	ecDescriptionSessionInvalid:                          "Session invalid.",
	ecDescriptionNewSessionRequired:                      "New session required.",
	ecDescriptionEcardStatusRedeemedError:                "E-CARD status redeemed error.",
	ecDescriptionEcardStatusRevokedError:                 "E-CARD status revoked error.",
	ecDescriptionEcardStatusInactiveError:                "E-CARD status inactive error.",
	ecDescriptionEcardStatusUnknownError:                 "E-CARD status unknown error.",
	ecDescriptionEcardTypeLegacyPointsError:              "E-CARD type legacy points error.",
	ecDescriptionEcardInvalidIdError:                     "E-CARD invalid ID error.",
	ecDescriptionEcardTypeCacheError:                     "E-CARD type cache error.",
	ecDescriptionEcardTypeUnknownError:                   "E-CARD type unknown error.",
	ecDescriptionEcardNoRedeemableItemsError:             "E-CARD no redeemable items error.",
	ecDescriptionFsMediaWriteProtectedError:              "FS media write protected error.",
	ecDescriptionInfraReportError:                        "Infra report error.",
	ecDescriptionInfraReceiveError:                       "Infra receive error.",
	ecDescriptionInfraInvalidResponseError:               "Infra invalid response error.",
	ecDescriptionInfraNeedsUpdateError:                   "Infra needs update error.",
	ecDescriptionInfraTimeoutError:                       "Infra timeout error.",
	ecDescriptionInfraBusyError:                          "Infra busy error.",
	ecDescriptionInfraInMaintenanceError:                 "Infra in maintenance error.",
	ecDescriptionNetworkTimeoutError:                     "Network timeout error.",
	ecDescriptionNetworkInMaintenanceError:               "Network in maintenance error.",
	ecDescriptionOutOfCatalogMemoryError:                 "Out of catalog memory error.",
	ecDescriptionOutOfSystemMemoryError:                  "Out of system memory error.",
	ecDescriptionFsMediaAccessFailedError:                "FS media access failed error.",
	ecDescriptionNoEnoughSpaceError:                      "No enough space error.",
	ecDescriptionAcNotConnectedError:                     "AC not connected error.",
	ecDescriptionAppletUnknownError:                      "Applet unknown error.",
	ecDescriptionUnknownError:                            "Unknown error.",
	ecDescriptionEcAppletNotFoundError:                   "EC applet not found error.",
	ecDescriptionInfraNeedsReconnectError:                "Infra needs reconnect error.",
	ecDescriptionAppletCancelled:                         "Applet cancelled.",
	ecDescriptionAppletHomeButton:                        "Applet home button.",
	ecDescriptionAppletSoftwareReset:                     "Applet software reset.",
	ecDescriptionAppletPowerButton:                       "Applet power button.",
	ecDescriptionAppletNotSupportedError:                 "Applet not supported error.",
	ecDescriptionAppletNotEnoughSpaceToDownloadContent:   "Applet not enough space to download content.",
	ecDescriptionAppletNecessaryAttributeNotFound:        "Applet necessary attribute not found.",
	ecDescriptionAppletMetadataNotFound:                  "Applet metadata not found.",
	ecDescriptionAppletTooManyContents:                   "Applet too many contents.",
	ecDescriptionAppletAlreadyPurchased:                  "Applet already purchased.",
	ecDescriptionAppletItemUnpurchasable:                 "Applet item unpurchasable.",
	ecDescriptionAppletItemNotFound:                      "Applet item not found.",
	ecDescriptionAppletAlreadyDownloaded:                 "Applet already downloaded.",
	ecDescriptionAppletRedeemableItemNotFound:            "Applet redeemable item not found.",
	ecDescriptionAppletAlreadyRedeemedItem:               "Applet already redeemed item.",
	ecDescriptionAppletShopServerError:                   "Applet shop server error.",
	ecDescriptionAppletImportFailed:                      "Applet import failed.",
	ecDescriptionAppletWifiOffError:                      "Applet wifi off error.",
	ecDescriptionAppletFailedConnectError:                "Applet failed connect error.",
	ecDescriptionAppletEcardUnavailable:                  "Applet E-CARD unavailable.",
	ecDescriptionAppletInvalidRivtoken:                   "Applet invalid RIV token.",
	ecDescriptionAppletInvalidParameter:                  "Applet invalid parameter.",
	ecDescriptionAppletFailed:                            "Applet failed.",
	ecDescriptionAppletFailedError:                       "Applet failed error.",
	ecDescriptionAppletWrongPinCodeError:                 "Applet wrong pin code error.",
	ecDescriptionAppletDuplicateContentIndex:             "Applet duplicate content index.",
	ecDescriptionAppletFailedLoginError:                  "Applet failed login error.",
	ecDescriptionAppletDataTitleInUse:                    "Applet data title in use.",
	ecDescriptionAppletNeedsDatatitleUpdate:              "Applet needs data title update.",
	ecDescriptionAppletDatatitleNotImported:              "Applet data title not imported.",
	ecDescriptionAppletRequiresReConnect:                 "Applet requires reconnect.",
	ecDescriptionAppletNeedsBalanceUpdate:                "Applet needs balance update.",
	ecDescriptionAppletNeedsSystemUpdate:                 "Applet needs system update.",
	ecDescriptionAppletStandbyMode:                       "Applet standby mode.",
	ecDescriptionAppletShopServiceTerminated:             "Applet shop service terminated.",
	ecDescriptionAppletShopServiceUnavailable:            "Applet shop service unavailable.",
	ecDescriptionAppletCountryChanged:                    "Applet country changed.",
	ecDescriptionAppletNotAgreeEulaError:                 "Applet not agree EULA error.",
	ecDescriptionAppletTooManyDatatitles:                 "Applet too many data titles.",
	ecDescriptionAppletNotEnoughSpaceToDownloadDatatitle: "Applet not enough space to download data title.",
	ecDescriptionAppletSdNotInserted:                     "Applet SD not inserted.",
	ecDescriptionAppletSdWriteProtected:                  "Applet SD write protected.",
	ecDescriptionAppletSdBroken:                          "Applet SD broken.",
	ecDescriptionAppletSdEjected:                         "Applet SD ejected.",
	ecDescriptionAppletInvalidVersion:                    "Applet invalid version.",
	ecDescriptionAppletSdAccessError:                     "Applet SD access error.",
}

func (d ecDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", ecDescriptionToString[d], d, ecDescriptionDescription[d])
}
