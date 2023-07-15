package descriptions

import "fmt"

type friendsDescription int32

const (
	friendsDescriptionFacilityShift       uint32 = 5
	friendsDescriptionReturnCodeUndefined uint32 = (1 << (friendsDescriptionFacilityShift + 1)) - 1

	friendsDescriptionFacilityCore           friendsDescription = 1 << friendsDescriptionFacilityShift
	friendsDescriptionFacilityDdl            friendsDescription = 2 << friendsDescriptionFacilityShift
	friendsDescriptionFacilityRendezvous     friendsDescription = 3 << friendsDescriptionFacilityShift
	friendsDescriptionFacilityPythonCore     friendsDescription = 4 << friendsDescriptionFacilityShift
	friendsDescriptionFacilityTransport      friendsDescription = 5 << friendsDescriptionFacilityShift
	friendsDescriptionFacilityDoCore         friendsDescription = 6 << friendsDescriptionFacilityShift
	friendsDescriptionFacilityFpd            friendsDescription = 7 << friendsDescriptionFacilityShift
	friendsDescriptionFacilityAuthentication friendsDescription = 8 << friendsDescriptionFacilityShift

	friendsDescriptionCoreSuccess               friendsDescription = friendsDescriptionFacilityCore + 0
	friendsDescriptionCoreSuccessPending        friendsDescription = friendsDescriptionFacilityCore + 1
	friendsDescriptionCoreUnknown               friendsDescription = friendsDescriptionFacilityCore + 2
	friendsDescriptionCoreNotImplemented        friendsDescription = friendsDescriptionFacilityCore + 3
	friendsDescriptionCoreInvalidPointer        friendsDescription = friendsDescriptionFacilityCore + 4
	friendsDescriptionCoreOperationAborted      friendsDescription = friendsDescriptionFacilityCore + 5
	friendsDescriptionCoreException             friendsDescription = friendsDescriptionFacilityCore + 6
	friendsDescriptionCoreAccessDenied          friendsDescription = friendsDescriptionFacilityCore + 7
	friendsDescriptionCoreInvalidHandle         friendsDescription = friendsDescriptionFacilityCore + 8
	friendsDescriptionCoreInvalidIndex          friendsDescription = friendsDescriptionFacilityCore + 9
	friendsDescriptionCoreOutOfMemory           friendsDescription = friendsDescriptionFacilityCore + 10
	friendsDescriptionCoreInvalidArgument       friendsDescription = friendsDescriptionFacilityCore + 11
	friendsDescriptionCoreTimeout               friendsDescription = friendsDescriptionFacilityCore + 12
	friendsDescriptionCoreInitializationFailure friendsDescription = friendsDescriptionFacilityCore + 13
	friendsDescriptionCoreCallInitiationFailure friendsDescription = friendsDescriptionFacilityCore + 14
	friendsDescriptionCoreRegistrationError     friendsDescription = friendsDescriptionFacilityCore + 15
	friendsDescriptionCoreBufferOverflow        friendsDescription = friendsDescriptionFacilityCore + 16
	friendsDescriptionCoreInvalidLockState      friendsDescription = friendsDescriptionFacilityCore + 17
	friendsDescriptionCoreUndefined             friendsDescription = friendsDescriptionFacilityCore + friendsDescription(friendsDescriptionReturnCodeUndefined)

	friendsDescriptionDdlSuccess          friendsDescription = friendsDescriptionFacilityDdl + 0
	friendsDescriptionDdlInvalidSignature friendsDescription = friendsDescriptionFacilityDdl + 1
	friendsDescriptionDdlIncorrectVersion friendsDescription = friendsDescriptionFacilityDdl + 2
	friendsDescriptionDdlUndefined        friendsDescription = friendsDescriptionFacilityDdl + friendsDescription(friendsDescriptionReturnCodeUndefined)

	friendsDescriptionRendezvousSuccess                           friendsDescription = friendsDescriptionFacilityRendezvous + 0
	friendsDescriptionRendezvousConnectionFailure                 friendsDescription = friendsDescriptionFacilityRendezvous + 1
	friendsDescriptionRendezvousNotAuthenticated                  friendsDescription = friendsDescriptionFacilityRendezvous + 2
	friendsDescriptionRendezvousInvalidUsername                   friendsDescription = friendsDescriptionFacilityRendezvous + 3
	friendsDescriptionRendezvousInvalidPassword                   friendsDescription = friendsDescriptionFacilityRendezvous + 4
	friendsDescriptionRendezvousUsernameAlreadyExists             friendsDescription = friendsDescriptionFacilityRendezvous + 5
	friendsDescriptionRendezvousAccountDisabled                   friendsDescription = friendsDescriptionFacilityRendezvous + 6
	friendsDescriptionRendezvousAccountExpired                    friendsDescription = friendsDescriptionFacilityRendezvous + 7
	friendsDescriptionRendezvousConcurrentLoginDenied             friendsDescription = friendsDescriptionFacilityRendezvous + 8
	friendsDescriptionRendezvousEncryptionFailure                 friendsDescription = friendsDescriptionFacilityRendezvous + 9
	friendsDescriptionRendezvousInvalidPid                        friendsDescription = friendsDescriptionFacilityRendezvous + 10
	friendsDescriptionRendezvousMaxConnectionsReached             friendsDescription = friendsDescriptionFacilityRendezvous + 11
	friendsDescriptionRendezvousInvalidGid                        friendsDescription = friendsDescriptionFacilityRendezvous + 12
	friendsDescriptionRendezvousInvalidThreadId                   friendsDescription = friendsDescriptionFacilityRendezvous + 13
	friendsDescriptionRendezvousInvalidOperationInLiveEnvironment friendsDescription = friendsDescriptionFacilityRendezvous + 14
	friendsDescriptionRendezvousDuplicateEntry                    friendsDescription = friendsDescriptionFacilityRendezvous + 15
	friendsDescriptionRendezvousControlScriptFailure              friendsDescription = friendsDescriptionFacilityRendezvous + 16
	friendsDescriptionRendezvousClassNotFound                     friendsDescription = friendsDescriptionFacilityRendezvous + 17
	friendsDescriptionRendezvousSessionVoid                       friendsDescription = friendsDescriptionFacilityRendezvous + 18
	friendsDescriptionRendezvousLspGatewayUnreachable             friendsDescription = friendsDescriptionFacilityRendezvous + 19
	friendsDescriptionRendezvousDdlMismatch                       friendsDescription = friendsDescriptionFacilityRendezvous + 20
	friendsDescriptionRendezvousInvalidFtpInfo                    friendsDescription = friendsDescriptionFacilityRendezvous + 21
	friendsDescriptionRendezvousSessionFull                       friendsDescription = friendsDescriptionFacilityRendezvous + 22
	friendsDescriptionRendezvousUndefined                         friendsDescription = friendsDescriptionFacilityRendezvous + friendsDescription(friendsDescriptionReturnCodeUndefined)

	friendsDescriptionPythonCoreSuccess          friendsDescription = friendsDescriptionFacilityPythonCore + 0
	friendsDescriptionPythonCoreException        friendsDescription = friendsDescriptionFacilityPythonCore + 1
	friendsDescriptionPythonCoreTypeError        friendsDescription = friendsDescriptionFacilityPythonCore + 2
	friendsDescriptionPythonCoreIndexError       friendsDescription = friendsDescriptionFacilityPythonCore + 3
	friendsDescriptionPythonCoreInvalidReference friendsDescription = friendsDescriptionFacilityPythonCore + 4
	friendsDescriptionPythonCoreCallFailure      friendsDescription = friendsDescriptionFacilityPythonCore + 5
	friendsDescriptionPythonCoreMemoryError      friendsDescription = friendsDescriptionFacilityPythonCore + 6
	friendsDescriptionPythonCoreKeyError         friendsDescription = friendsDescriptionFacilityPythonCore + 7
	friendsDescriptionPythonCoreOperationError   friendsDescription = friendsDescriptionFacilityPythonCore + 8
	friendsDescriptionPythonCoreConversionError  friendsDescription = friendsDescriptionFacilityPythonCore + 9
	friendsDescriptionPythonCoreValidationError  friendsDescription = friendsDescriptionFacilityPythonCore + 10
	friendsDescriptionPythonCoreUndefined        friendsDescription = friendsDescriptionFacilityPythonCore + friendsDescription(friendsDescriptionReturnCodeUndefined)

	friendsDescriptionTransportSuccess                       friendsDescription = friendsDescriptionFacilityTransport + 0
	friendsDescriptionTransportUnknown                       friendsDescription = friendsDescriptionFacilityTransport + 1
	friendsDescriptionTransportConnectionFailure             friendsDescription = friendsDescriptionFacilityTransport + 2
	friendsDescriptionTransportInvalidUrl                    friendsDescription = friendsDescriptionFacilityTransport + 3
	friendsDescriptionTransportInvalidKey                    friendsDescription = friendsDescriptionFacilityTransport + 4
	friendsDescriptionTransportInvalidUrlType                friendsDescription = friendsDescriptionFacilityTransport + 5
	friendsDescriptionTransportDuplicateEndpoint             friendsDescription = friendsDescriptionFacilityTransport + 6
	friendsDescriptionTransportIoError                       friendsDescription = friendsDescriptionFacilityTransport + 7
	friendsDescriptionTransportTimeout                       friendsDescription = friendsDescriptionFacilityTransport + 8
	friendsDescriptionTransportConnectionReset               friendsDescription = friendsDescriptionFacilityTransport + 9
	friendsDescriptionTransportIncorrectRemoteAuthentication friendsDescription = friendsDescriptionFacilityTransport + 10
	friendsDescriptionTransportServerRequestError            friendsDescription = friendsDescriptionFacilityTransport + 11
	friendsDescriptionTransportDecompressionFailure          friendsDescription = friendsDescriptionFacilityTransport + 12
	friendsDescriptionTransportCongestedEndpoint             friendsDescription = friendsDescriptionFacilityTransport + 13
	friendsDescriptionTransportUpnpCannotInit                friendsDescription = friendsDescriptionFacilityTransport + 14
	friendsDescriptionTransportUpnpCannotAddMapping          friendsDescription = friendsDescriptionFacilityTransport + 15
	friendsDescriptionTransportNatPmpCannotInit              friendsDescription = friendsDescriptionFacilityTransport + 16
	friendsDescriptionTransportNatPmpCannotAddMapping        friendsDescription = friendsDescriptionFacilityTransport + 17
	friendsDescriptionTransportUndefined                     friendsDescription = friendsDescriptionFacilityTransport + friendsDescription(friendsDescriptionReturnCodeUndefined)

	friendsDescriptionDoCoreSuccess                    friendsDescription = friendsDescriptionFacilityDoCore + 0
	friendsDescriptionDoCoreCallPostponed              friendsDescription = friendsDescriptionFacilityDoCore + 1
	friendsDescriptionDoCoreStationNotReached          friendsDescription = friendsDescriptionFacilityDoCore + 2
	friendsDescriptionDoCoreTargetStationDisconnect    friendsDescription = friendsDescriptionFacilityDoCore + 3
	friendsDescriptionDoCoreLocalStationLeaving        friendsDescription = friendsDescriptionFacilityDoCore + 4
	friendsDescriptionDoCoreObjectNotFound             friendsDescription = friendsDescriptionFacilityDoCore + 5
	friendsDescriptionDoCoreInvalidRole                friendsDescription = friendsDescriptionFacilityDoCore + 6
	friendsDescriptionDoCoreCallTimeout                friendsDescription = friendsDescriptionFacilityDoCore + 7
	friendsDescriptionDoCoreRmcDispatchFailed          friendsDescription = friendsDescriptionFacilityDoCore + 8
	friendsDescriptionDoCoreMigrationInProgress        friendsDescription = friendsDescriptionFacilityDoCore + 9
	friendsDescriptionDoCoreNoAuthority                friendsDescription = friendsDescriptionFacilityDoCore + 10
	friendsDescriptionDoreCoreNoTargetStationSpecified friendsDescription = friendsDescriptionFacilityDoCore + 11
	friendsDescriptionDoCoreJoinFailed                 friendsDescription = friendsDescriptionFacilityDoCore + 12
	friendsDescriptionDoCoreJoinDenied                 friendsDescription = friendsDescriptionFacilityDoCore + 13
	friendsDescriptionDoCoreConnectivityTestFailed     friendsDescription = friendsDescriptionFacilityDoCore + 14
	friendsDescriptionDoCoreUnknown                    friendsDescription = friendsDescriptionFacilityDoCore + 15
	friendsDescriptionDoCoreUnfreedReferences          friendsDescription = friendsDescriptionFacilityDoCore + 16
	friendsDescriptionDoCoreUndefined                  friendsDescription = friendsDescriptionFacilityDoCore + friendsDescription(friendsDescriptionReturnCodeUndefined)

	friendsDescriptionFpdSuccess                friendsDescription = friendsDescriptionFacilityFpd + 0
	friendsDescriptionRmcNotCalled              friendsDescription = friendsDescriptionFacilityFpd + 1
	friendsDescriptionDaemonNotInitialized      friendsDescription = friendsDescriptionFacilityFpd + 2
	friendsDescriptionDaemonAlreadyInitialized  friendsDescription = friendsDescriptionFacilityFpd + 3
	friendsDescriptionNotConnected              friendsDescription = friendsDescriptionFacilityFpd + 4
	friendsDescriptionConnected                 friendsDescription = friendsDescriptionFacilityFpd + 5
	friendsDescriptionInitializationFailure     friendsDescription = friendsDescriptionFacilityFpd + 6
	friendsDescriptionOutOfmemory               friendsDescription = friendsDescriptionFacilityFpd + 7
	friendsDescriptionRmcFailed                 friendsDescription = friendsDescriptionFacilityFpd + 8
	friendsDescriptionInvalidArgument           friendsDescription = friendsDescriptionFacilityFpd + 9
	friendsDescriptionInvalidLocalAccountId     friendsDescription = friendsDescriptionFacilityFpd + 10
	friendsDescriptionInvalidPrincipalId        friendsDescription = friendsDescriptionFacilityFpd + 11
	friendsDescriptionInvalidLocalFriendCode    friendsDescription = friendsDescriptionFacilityFpd + 12
	friendsDescriptionLocalAccountNotExists     friendsDescription = friendsDescriptionFacilityFpd + 13
	friendsDescriptionLocalAccountNotLoaded     friendsDescription = friendsDescriptionFacilityFpd + 14
	friendsDescriptionLocalAccountAlreadyLoaded friendsDescription = friendsDescriptionFacilityFpd + 15
	friendsDescriptionFriendAlreadyExists       friendsDescription = friendsDescriptionFacilityFpd + 16
	friendsDescriptionFriendNotExists           friendsDescription = friendsDescriptionFacilityFpd + 17
	friendsDescriptionFriendNumMax              friendsDescription = friendsDescriptionFacilityFpd + 18
	friendsDescriptionNotFriend                 friendsDescription = friendsDescriptionFacilityFpd + 19
	friendsDescriptionFileIoError               friendsDescription = friendsDescriptionFacilityFpd + 20
	friendsDescriptionP2pInternetProhibited     friendsDescription = friendsDescriptionFacilityFpd + 21
	friendsDescriptionUnknown                   friendsDescription = friendsDescriptionFacilityFpd + 22
	friendsDescriptionInvalidState              friendsDescription = friendsDescriptionFacilityFpd + 23
	friendsDescriptionMiiNotExists              friendsDescription = friendsDescriptionFacilityFpd + 24
	friendsDescriptionAddFriendProhibited       friendsDescription = friendsDescriptionFacilityFpd + 25
	friendsDescriptionInvalidReference          friendsDescription = friendsDescriptionFacilityFpd + 26
	friendsDescriptionFpdUndefined              friendsDescription = friendsDescriptionFacilityFpd + friendsDescription(friendsDescriptionReturnCodeUndefined)

	friendsDescriptionAuthenticationSuccess              friendsDescription = friendsDescriptionFacilityAuthentication + 0
	friendsDescriptionAuthenticationNasAuthenticateError friendsDescription = friendsDescriptionFacilityAuthentication + 1
	friendsDescriptionAuthenticationTokenParseError      friendsDescription = friendsDescriptionFacilityAuthentication + 2
	friendsDescriptionAuthenticationHttpConnectionError  friendsDescription = friendsDescriptionFacilityAuthentication + 3
	friendsDescriptionAuthenticationHttpDnsError         friendsDescription = friendsDescriptionFacilityAuthentication + 4
	friendsDescriptionAuthenticationHttpGetProxySetting  friendsDescription = friendsDescriptionFacilityAuthentication + 5
	friendsDescriptionAuthenticationTokenExpired         friendsDescription = friendsDescriptionFacilityAuthentication + 6
	friendsDescriptionAuthenticationValidationFailed     friendsDescription = friendsDescriptionFacilityAuthentication + 7
	friendsDescriptionAuthenticationInvalidParam         friendsDescription = friendsDescriptionFacilityAuthentication + 8
	friendsDescriptionAuthenticationPrincipalIdUnmatched friendsDescription = friendsDescriptionFacilityAuthentication + 9
	friendsDescriptionAuthenticationMoveCountUnmatch     friendsDescription = friendsDescriptionFacilityAuthentication + 10
	friendsDescriptionAuthenticationUnderMaintenance     friendsDescription = friendsDescriptionFacilityAuthentication + 11
	friendsDescriptionAuthenticationUnsupportedVersion   friendsDescription = friendsDescriptionFacilityAuthentication + 12
	friendsDescriptionAuthenticationUnknown              friendsDescription = friendsDescriptionFacilityAuthentication + 13
	friendsDescriptionAuthenticationUndefined            friendsDescription = friendsDescriptionFacilityAuthentication + friendsDescription(friendsDescriptionReturnCodeUndefined)

	friendsDescriptionFacilityCtr friendsDescription = 0 << friendsDescriptionFacilityShift

	friendsDescriptionInvalidFriendCode friendsDescription = iota + friendsDescriptionFacilityCtr - 149
	friendsDescriptionNotLoggedIn
	friendsDescriptionNotFriendsResult
	friendsDescriptionUndefinedFacility
	friendsDescriptionAcNotConnected
)

var friendsDescriptionToString = map[friendsDescription]string{
	friendsDescriptionCoreSuccess:               "CoreSuccess",
	friendsDescriptionCoreSuccessPending:        "CoreSuccessPending",
	friendsDescriptionCoreUnknown:               "CoreUnknown",
	friendsDescriptionCoreNotImplemented:        "CoreNotImplemented",
	friendsDescriptionCoreInvalidPointer:        "CoreInvalidPointer",
	friendsDescriptionCoreOperationAborted:      "CoreOperationAborted",
	friendsDescriptionCoreException:             "CoreException",
	friendsDescriptionCoreAccessDenied:          "CoreAccessDenied",
	friendsDescriptionCoreInvalidHandle:         "CoreInvalidHandle",
	friendsDescriptionCoreInvalidIndex:          "CoreInvalidIndex",
	friendsDescriptionCoreOutOfMemory:           "CoreOutOfMemory",
	friendsDescriptionCoreInvalidArgument:       "CoreInvalidArgument",
	friendsDescriptionCoreTimeout:               "CoreTimeout",
	friendsDescriptionCoreInitializationFailure: "CoreInitializationFailure",
	friendsDescriptionCoreCallInitiationFailure: "CoreCallInitiationFailure",
	friendsDescriptionCoreRegistrationError:     "CoreRegistrationError",
	friendsDescriptionCoreBufferOverflow:        "CoreBufferOverflow",
	friendsDescriptionCoreInvalidLockState:      "CoreInvalidLockState",
	friendsDescriptionCoreUndefined:             "CoreUndefined",

	friendsDescriptionDdlSuccess:          "DdlSuccess",
	friendsDescriptionDdlInvalidSignature: "DdlInvalidSignature",
	friendsDescriptionDdlIncorrectVersion: "DdlIncorrectVersion",
	friendsDescriptionDdlUndefined:        "DdlUndefined",

	friendsDescriptionRendezvousSuccess:                           "RendezvousSuccess",
	friendsDescriptionRendezvousConnectionFailure:                 "RendezvousConnectionFailure",
	friendsDescriptionRendezvousNotAuthenticated:                  "RendezvousNotAuthenticated",
	friendsDescriptionRendezvousInvalidUsername:                   "RendezvousInvalidUsername",
	friendsDescriptionRendezvousInvalidPassword:                   "RendezvousInvalidPassword",
	friendsDescriptionRendezvousUsernameAlreadyExists:             "RendezvousUsernameAlreadyExists",
	friendsDescriptionRendezvousAccountDisabled:                   "RendezvousAccountDisabled",
	friendsDescriptionRendezvousAccountExpired:                    "RendezvousAccountExpired",
	friendsDescriptionRendezvousConcurrentLoginDenied:             "RendezvousConcurrentLoginDenied",
	friendsDescriptionRendezvousEncryptionFailure:                 "RendezvousEncryptionFailure",
	friendsDescriptionRendezvousInvalidPid:                        "RendezvousInvalidPid",
	friendsDescriptionRendezvousMaxConnectionsReached:             "RendezvousMaxConnectionsReached",
	friendsDescriptionRendezvousInvalidGid:                        "RendezvousInvalidGid",
	friendsDescriptionRendezvousInvalidThreadId:                   "RendezvousInvalidThreadId",
	friendsDescriptionRendezvousInvalidOperationInLiveEnvironment: "RendezvousInvalidOperationInLiveEnvironment",
	friendsDescriptionRendezvousDuplicateEntry:                    "RendezvousDuplicateEntry",
	friendsDescriptionRendezvousControlScriptFailure:              "RendezvousControlScriptFailure",
	friendsDescriptionRendezvousClassNotFound:                     "RendezvousClassNotFound",
	friendsDescriptionRendezvousSessionVoid:                       "RendezvousSessionVoid",
	friendsDescriptionRendezvousLspGatewayUnreachable:             "RendezvousLspGatewayUnreachable",
	friendsDescriptionRendezvousDdlMismatch:                       "RendezvousDdlMismatch",
	friendsDescriptionRendezvousInvalidFtpInfo:                    "RendezvousInvalidFtpInfo",
	friendsDescriptionRendezvousSessionFull:                       "RendezvousSessionFull",
	friendsDescriptionRendezvousUndefined:                         "RendezvousUndefined",

	friendsDescriptionPythonCoreSuccess:          "PythonCoreSuccess",
	friendsDescriptionPythonCoreException:        "PythonCoreException",
	friendsDescriptionPythonCoreTypeError:        "PythonCoreTypeError",
	friendsDescriptionPythonCoreIndexError:       "PythonCoreIndexError",
	friendsDescriptionPythonCoreInvalidReference: "PythonCoreInvalidReference",
	friendsDescriptionPythonCoreCallFailure:      "PythonCoreCallFailure",
	friendsDescriptionPythonCoreMemoryError:      "PythonCoreMemoryError",
	friendsDescriptionPythonCoreKeyError:         "PythonCoreKeyError",
	friendsDescriptionPythonCoreOperationError:   "PythonCoreOperationError",
	friendsDescriptionPythonCoreConversionError:  "PythonCoreConversionError",
	friendsDescriptionPythonCoreValidationError:  "PythonCoreValidationError",
	friendsDescriptionPythonCoreUndefined:        "PythonCoreUndefined",

	friendsDescriptionTransportSuccess:                       "TransportSuccess",
	friendsDescriptionTransportUnknown:                       "TransportUnknown",
	friendsDescriptionTransportConnectionFailure:             "TransportConnectionFailure",
	friendsDescriptionTransportInvalidUrl:                    "TransportInvalidUrl",
	friendsDescriptionTransportInvalidKey:                    "TransportInvalidKey",
	friendsDescriptionTransportInvalidUrlType:                "TransportInvalidUrlType",
	friendsDescriptionTransportDuplicateEndpoint:             "TransportDuplicateEndpoint",
	friendsDescriptionTransportIoError:                       "TransportIoError",
	friendsDescriptionTransportTimeout:                       "TransportTimeout",
	friendsDescriptionTransportConnectionReset:               "TransportConnectionReset",
	friendsDescriptionTransportIncorrectRemoteAuthentication: "TransportIncorrectRemoteAuthentication",
	friendsDescriptionTransportServerRequestError:            "TransportServerRequestError",
	friendsDescriptionTransportDecompressionFailure:          "TransportDecompressionFailure",
	friendsDescriptionTransportCongestedEndpoint:             "TransportCongestedEndpoint",
	friendsDescriptionTransportUpnpCannotInit:                "TransportUpnpCannotInit",
	friendsDescriptionTransportUpnpCannotAddMapping:          "TransportUpnpCannotAddMapping",
	friendsDescriptionTransportNatPmpCannotInit:              "TransportNatPmpCannotInit",
	friendsDescriptionTransportNatPmpCannotAddMapping:        "TransportNatPmpCannotAddMapping",
	friendsDescriptionTransportUndefined:                     "TransportUndefined",

	friendsDescriptionDoCoreSuccess:                    "DoCoreSuccess",
	friendsDescriptionDoCoreCallPostponed:              "DoCoreCallPostponed",
	friendsDescriptionDoCoreStationNotReached:          "DoCoreStationNotReached",
	friendsDescriptionDoCoreTargetStationDisconnect:    "DoCoreTargetStationDisconnect",
	friendsDescriptionDoCoreLocalStationLeaving:        "DoCoreLocalStationLeaving",
	friendsDescriptionDoCoreObjectNotFound:             "DoCoreObjectNotFound",
	friendsDescriptionDoCoreInvalidRole:                "DoCoreInvalidRole",
	friendsDescriptionDoCoreCallTimeout:                "DoCoreCallTimeout",
	friendsDescriptionDoCoreRmcDispatchFailed:          "DoCoreRmcDispatchFailed",
	friendsDescriptionDoCoreMigrationInProgress:        "DoCoreMigrationInProgress",
	friendsDescriptionDoCoreNoAuthority:                "DoCoreNoAuthority",
	friendsDescriptionDoreCoreNoTargetStationSpecified: "DoCoreNoTargetStationSpecified",
	friendsDescriptionDoCoreJoinFailed:                 "DoCoreJoinFailed",
	friendsDescriptionDoCoreJoinDenied:                 "DoCoreJoinDenied",
	friendsDescriptionDoCoreConnectivityTestFailed:     "DoCoreConnectivityTestFailed",
	friendsDescriptionDoCoreUnknown:                    "DoCoreUnknown",
	friendsDescriptionDoCoreUnfreedReferences:          "DoCoreUnfreedReferences",
	friendsDescriptionDoCoreUndefined:                  "DoCoreUndefined",

	friendsDescriptionFpdSuccess:                "FpdSuccess",
	friendsDescriptionRmcNotCalled:              "RmcNotCalled",
	friendsDescriptionDaemonNotInitialized:      "DaemonNotInitialized",
	friendsDescriptionDaemonAlreadyInitialized:  "DaemonAlreadyInitialized",
	friendsDescriptionNotConnected:              "NotConnected",
	friendsDescriptionConnected:                 "Connected",
	friendsDescriptionInitializationFailure:     "InitializationFailure",
	friendsDescriptionOutOfmemory:               "OutOfmemory",
	friendsDescriptionRmcFailed:                 "RmcFailed",
	friendsDescriptionInvalidArgument:           "InvalidArgument",
	friendsDescriptionInvalidLocalAccountId:     "InvalidLocalAccountId",
	friendsDescriptionInvalidPrincipalId:        "InvalidPrincipalId",
	friendsDescriptionInvalidLocalFriendCode:    "InvalidLocalFriendCode",
	friendsDescriptionLocalAccountNotExists:     "LocalAccountNotExists",
	friendsDescriptionLocalAccountNotLoaded:     "LocalAccountNotLoaded",
	friendsDescriptionLocalAccountAlreadyLoaded: "LocalAccountAlreadyLoaded",
	friendsDescriptionFriendAlreadyExists:       "FriendAlreadyExists",
	friendsDescriptionFriendNotExists:           "FriendNotExists",
	friendsDescriptionFriendNumMax:              "FriendNumMax",
	friendsDescriptionNotFriend:                 "NotFriend",
	friendsDescriptionFileIoError:               "FileIoError",
	friendsDescriptionP2pInternetProhibited:     "P2pInternetProhibited",
	friendsDescriptionUnknown:                   "Unknown",
	friendsDescriptionInvalidState:              "InvalidState",
	friendsDescriptionMiiNotExists:              "MiiNotExists",
	friendsDescriptionAddFriendProhibited:       "AddFriendProhibited",
	friendsDescriptionInvalidReference:          "InvalidReference",
	friendsDescriptionFpdUndefined:              "FpdUndefined",

	friendsDescriptionAuthenticationSuccess:              "AuthenticationSuccess",
	friendsDescriptionAuthenticationNasAuthenticateError: "AuthenticationNasAuthenticateError",
	friendsDescriptionAuthenticationTokenParseError:      "AuthenticationTokenParseError",
	friendsDescriptionAuthenticationHttpConnectionError:  "AuthenticationHttpConnectionError",
	friendsDescriptionAuthenticationHttpDnsError:         "AuthenticationHttpDnsError",
	friendsDescriptionAuthenticationHttpGetProxySetting:  "AuthenticationHttpGetProxySetting",
	friendsDescriptionAuthenticationTokenExpired:         "AuthenticationTokenExpired",
	friendsDescriptionAuthenticationValidationFailed:     "AuthenticationValidationFailed",
	friendsDescriptionAuthenticationInvalidParam:         "AuthenticationInvalidParam",
	friendsDescriptionAuthenticationPrincipalIdUnmatched: "AuthenticationPrincipalIdUnmatched",
	friendsDescriptionAuthenticationMoveCountUnmatch:     "AuthenticationMoveCountUnmatch",
	friendsDescriptionAuthenticationUnderMaintenance:     "AuthenticationUnderMaintenance",
	friendsDescriptionAuthenticationUnsupportedVersion:   "AuthenticationUnsupportedVersion",
	friendsDescriptionAuthenticationUnknown:              "AuthenticationUnknown",
	friendsDescriptionAuthenticationUndefined:            "AuthenticationUndefined",

	friendsDescriptionInvalidFriendCode: "InvalidFriendCode",
	friendsDescriptionNotLoggedIn:       "NotLoggedIn",
	friendsDescriptionNotFriendsResult:  "NotFriendsResult",
	friendsDescriptionUndefinedFacility: "UndefinedFacility",
	friendsDescriptionAcNotConnected:    "AcNotConnected",
}

var friendsDescriptionDescription = map[friendsDescription]string{
	friendsDescriptionCoreSuccess:               "The friends-core operation was successful.",
	friendsDescriptionCoreSuccessPending:        "The friends-core operation was successful, but the result is pending.",
	friendsDescriptionCoreUnknown:               "An unknown error occurred in the friends-core operation.",
	friendsDescriptionCoreNotImplemented:        "The friends-core operation is not implemented.",
	friendsDescriptionCoreInvalidPointer:        "The pointer passed to the friends-core operation is invalid.",
	friendsDescriptionCoreOperationAborted:      "The friends-core operation was aborted.",
	friendsDescriptionCoreException:             "An exception occurred in the friends-core operation.",
	friendsDescriptionCoreAccessDenied:          "Access was denied in the friends-core operation.",
	friendsDescriptionCoreInvalidHandle:         "The handle passed to the friends-core operation is invalid.",
	friendsDescriptionCoreInvalidIndex:          "The index passed to the friends-core operation is invalid.",
	friendsDescriptionCoreOutOfMemory:           "The friends-core operation ran out of memory.",
	friendsDescriptionCoreInvalidArgument:       "The argument passed to the friends-core operation is invalid.",
	friendsDescriptionCoreTimeout:               "The friends-core operation timed out.",
	friendsDescriptionCoreInitializationFailure: "The friends-core operation failed to initialize.",
	friendsDescriptionCoreCallInitiationFailure: "The friends-core operation failed to initiate a call.",
	friendsDescriptionCoreRegistrationError:     "The friends-core operation failed to register.",
	friendsDescriptionCoreBufferOverflow:        "The friends-core operation overflowed a buffer.",
	friendsDescriptionCoreInvalidLockState:      "The lock state passed to the friends-core operation is invalid.",
	friendsDescriptionCoreUndefined:             "The friends-core operation returned an undefined error.",

	friendsDescriptionDdlSuccess:          "The friends-ddl operation was successful.",
	friendsDescriptionDdlInvalidSignature: "The signature passed to the friends-ddl operation is invalid.",
	friendsDescriptionDdlIncorrectVersion: "The version passed to the friends-ddl operation is incorrect.",
	friendsDescriptionDdlUndefined:        "The friends-ddl operation returned an undefined error.",

	friendsDescriptionRendezvousSuccess:                           "The friends-rendezvous operation was successful.",
	friendsDescriptionRendezvousConnectionFailure:                 "The friends-rendezvous operation failed to connect.",
	friendsDescriptionRendezvousNotAuthenticated:                  "The friends-rendezvous operation is not authenticated.",
	friendsDescriptionRendezvousInvalidUsername:                   "The username passed to the friends-rendezvous operation is invalid.",
	friendsDescriptionRendezvousInvalidPassword:                   "The password passed to the friends-rendezvous operation is invalid.",
	friendsDescriptionRendezvousUsernameAlreadyExists:             "The username passed to the friends-rendezvous operation already exists.",
	friendsDescriptionRendezvousAccountDisabled:                   "The account passed to the friends-rendezvous operation is disabled.",
	friendsDescriptionRendezvousAccountExpired:                    "The account passed to the friends-rendezvous operation is expired.",
	friendsDescriptionRendezvousConcurrentLoginDenied:             "The concurrent login passed to the friends-rendezvous operation is denied.",
	friendsDescriptionRendezvousEncryptionFailure:                 "The encryption passed to the friends-rendezvous operation failed.",
	friendsDescriptionRendezvousInvalidPid:                        "The pid passed to the friends-rendezvous operation is invalid.",
	friendsDescriptionRendezvousMaxConnectionsReached:             "The max connections passed to the friends-rendezvous operation is reached.",
	friendsDescriptionRendezvousInvalidGid:                        "The gid passed to the friends-rendezvous operation is invalid.",
	friendsDescriptionRendezvousInvalidThreadId:                   "The thread id passed to the friends-rendezvous operation is invalid.",
	friendsDescriptionRendezvousInvalidOperationInLiveEnvironment: "The operation passed to the friends-rendezvous operation is invalid in live environment.",
	friendsDescriptionRendezvousDuplicateEntry:                    "The entry passed to the friends-rendezvous operation is duplicate.",
	friendsDescriptionRendezvousControlScriptFailure:              "The control script passed to the friends-rendezvous operation failed.",
	friendsDescriptionRendezvousClassNotFound:                     "The class passed to the friends-rendezvous operation is not found.",
	friendsDescriptionRendezvousSessionVoid:                       "The session passed to the friends-rendezvous operation is void.",
	friendsDescriptionRendezvousLspGatewayUnreachable:             "The lsp gateway passed to the friends-rendezvous operation is unreachable.",
	friendsDescriptionRendezvousDdlMismatch:                       "The ddl passed to the friends-rendezvous operation is mismatch.",
	friendsDescriptionRendezvousInvalidFtpInfo:                    "The ftp info passed to the friends-rendezvous operation is invalid.",
	friendsDescriptionRendezvousSessionFull:                       "The session passed to the friends-rendezvous operation is full.",
	friendsDescriptionRendezvousUndefined:                         "The friends-rendezvous operation returned an undefined error.",

	friendsDescriptionPythonCoreSuccess:          "The friends-pythoncore operation was successful.",
	friendsDescriptionPythonCoreException:        "An exception occurred in the friends-pythoncore operation.",
	friendsDescriptionPythonCoreTypeError:        "The type passed to the friends-pythoncore operation is invalid.",
	friendsDescriptionPythonCoreIndexError:       "The index passed to the friends-pythoncore operation is invalid.",
	friendsDescriptionPythonCoreInvalidReference: "The reference passed to the friends-pythoncore operation is invalid.",
	friendsDescriptionPythonCoreCallFailure:      "The call passed to the friends-pythoncore operation failed.",
	friendsDescriptionPythonCoreMemoryError:      "The memory passed to the friends-pythoncore operation is invalid.",
	friendsDescriptionPythonCoreKeyError:         "The key passed to the friends-pythoncore operation is invalid.",
	friendsDescriptionPythonCoreOperationError:   "The operation passed to the friends-pythoncore operation is invalid.",
	friendsDescriptionPythonCoreConversionError:  "The conversion passed to the friends-pythoncore operation is invalid.",
	friendsDescriptionPythonCoreValidationError:  "The validation passed to the friends-pythoncore operation is invalid.",
	friendsDescriptionPythonCoreUndefined:        "The friends-pythoncore operation returned an undefined error.",

	friendsDescriptionTransportSuccess:                       "The friends-transport operation was successful.",
	friendsDescriptionTransportUnknown:                       "An unknown error occurred in the friends-transport operation.",
	friendsDescriptionTransportConnectionFailure:             "The connection passed to the friends-transport operation failed.",
	friendsDescriptionTransportInvalidUrl:                    "The url passed to the friends-transport operation is invalid.",
	friendsDescriptionTransportInvalidKey:                    "The key passed to the friends-transport operation is invalid.",
	friendsDescriptionTransportInvalidUrlType:                "The url type passed to the friends-transport operation is invalid.",
	friendsDescriptionTransportDuplicateEndpoint:             "The endpoint passed to the friends-transport operation is duplicate.",
	friendsDescriptionTransportIoError:                       "The io passed to the friends-transport operation is invalid.",
	friendsDescriptionTransportTimeout:                       "The timeout passed to the friends-transport operation is invalid.",
	friendsDescriptionTransportConnectionReset:               "The connection reset passed to the friends-transport operation is invalid.",
	friendsDescriptionTransportIncorrectRemoteAuthentication: "The remote authentication passed to the friends-transport operation is incorrect.",
	friendsDescriptionTransportServerRequestError:            "The server request passed to the friends-transport operation is invalid.",
	friendsDescriptionTransportDecompressionFailure:          "The decompression passed to the friends-transport operation failed.",
	friendsDescriptionTransportCongestedEndpoint:             "The endpoint passed to the friends-transport operation is congested.",
	friendsDescriptionTransportUpnpCannotInit:                "The upnp passed to the friends-transport operation cannot init.",
	friendsDescriptionTransportUpnpCannotAddMapping:          "The upnp passed to the friends-transport operation cannot add mapping.",
	friendsDescriptionTransportNatPmpCannotInit:              "The nat pmp passed to the friends-transport operation cannot init.",
	friendsDescriptionTransportNatPmpCannotAddMapping:        "The nat pmp passed to the friends-transport operation cannot add mapping.",
	friendsDescriptionTransportUndefined:                     "The friends-transport operation returned an undefined error.",

	friendsDescriptionDoCoreSuccess:                    "The friends-docore operation was successful.",
	friendsDescriptionDoCoreCallPostponed:              "The call passed to the friends-docore operation is postponed.",
	friendsDescriptionDoCoreStationNotReached:          "The station passed to the friends-docore operation is not reached.",
	friendsDescriptionDoCoreTargetStationDisconnect:    "The target station passed to the friends-docore operation is disconnected.",
	friendsDescriptionDoCoreLocalStationLeaving:        "The local station passed to the friends-docore operation is leaving.",
	friendsDescriptionDoCoreObjectNotFound:             "The object passed to the friends-docore operation is not found.",
	friendsDescriptionDoCoreInvalidRole:                "The role passed to the friends-docore operation is invalid.",
	friendsDescriptionDoCoreCallTimeout:                "The call passed to the friends-docore operation timed out.",
	friendsDescriptionDoCoreRmcDispatchFailed:          "The rmc passed to the friends-docore operation dispatch failed.",
	friendsDescriptionDoCoreMigrationInProgress:        "The migration passed to the friends-docore operation is in progress.",
	friendsDescriptionDoCoreNoAuthority:                "The authority passed to the friends-docore operation is invalid.",
	friendsDescriptionDoreCoreNoTargetStationSpecified: "The target station passed to the friends-docore operation is not specified.",
	friendsDescriptionDoCoreJoinFailed:                 "The join passed to the friends-docore operation failed.",
	friendsDescriptionDoCoreJoinDenied:                 "The join passed to the friends-docore operation is denied.",
	friendsDescriptionDoCoreConnectivityTestFailed:     "The connectivity test passed to the friends-docore operation failed.",
	friendsDescriptionDoCoreUnknown:                    "An unknown error occurred in the friends-docore operation.",
	friendsDescriptionDoCoreUnfreedReferences:          "The references passed to the friends-docore operation are unfreed.",
	friendsDescriptionDoCoreUndefined:                  "The friends-docore operation returned an undefined error.",

	friendsDescriptionFpdSuccess:                "The friends-fpd operation was successful.",
	friendsDescriptionRmcNotCalled:              "The rmc passed to the friends-fpd operation is not called.",
	friendsDescriptionDaemonNotInitialized:      "The daemon passed to the friends-fpd operation is not initialized.",
	friendsDescriptionDaemonAlreadyInitialized:  "The daemon passed to the friends-fpd operation is already initialized.",
	friendsDescriptionNotConnected:              "The connection passed to the friends-fpd operation is not connected.",
	friendsDescriptionConnected:                 "The connection passed to the friends-fpd operation is connected.",
	friendsDescriptionInitializationFailure:     "The initialization passed to the friends-fpd operation failed.",
	friendsDescriptionOutOfmemory:               "The memory passed to the friends-fpd operation is out of memory.",
	friendsDescriptionRmcFailed:                 "The rmc passed to the friends-fpd operation failed.",
	friendsDescriptionInvalidArgument:           "The argument passed to the friends-fpd operation is invalid.",
	friendsDescriptionInvalidLocalAccountId:     "The local account id passed to the friends-fpd operation is invalid.",
	friendsDescriptionInvalidPrincipalId:        "The principal id passed to the friends-fpd operation is invalid.",
	friendsDescriptionInvalidLocalFriendCode:    "The local friend code passed to the friends-fpd operation is invalid.",
	friendsDescriptionLocalAccountNotExists:     "The local account passed to the friends-fpd operation does not exist.",
	friendsDescriptionLocalAccountNotLoaded:     "The local account passed to the friends-fpd operation is not loaded.",
	friendsDescriptionLocalAccountAlreadyLoaded: "The local account passed to the friends-fpd operation is already loaded.",
	friendsDescriptionFriendAlreadyExists:       "The friend passed to the friends-fpd operation already exists.",
	friendsDescriptionFriendNotExists:           "The friend passed to the friends-fpd operation does not exist.",
	friendsDescriptionFriendNumMax:              "The friend passed to the friends-fpd operation is max.",
	friendsDescriptionNotFriend:                 "The friend passed to the friends-fpd operation is not a friend.",
	friendsDescriptionFileIoError:               "The io passed to the friends-fpd operation is invalid.",
	friendsDescriptionP2pInternetProhibited:     "The internet passed to the friends-fpd operation is prohibited.",
	friendsDescriptionUnknown:                   "An unknown error occurred in the friends-fpd operation.",
	friendsDescriptionInvalidState:              "The state passed to the friends-fpd operation is invalid.",
	friendsDescriptionMiiNotExists:              "The mii passed to the friends-fpd operation does not exist.",
	friendsDescriptionAddFriendProhibited:       "The friend passed to the friends-fpd operation is prohibited.",
	friendsDescriptionInvalidReference:          "The reference passed to the friends-fpd operation is invalid.",
	friendsDescriptionFpdUndefined:              "The friends-fpd operation returned an undefined error.",

	friendsDescriptionAuthenticationSuccess:              "The friends-authentication operation was successful.",
	friendsDescriptionAuthenticationNasAuthenticateError: "The network authentication server passed to the friends-authentication operation is invalid.",
	friendsDescriptionAuthenticationTokenParseError:      "The token passed to the friends-authentication operation is invalid.",
	friendsDescriptionAuthenticationHttpConnectionError:  "The http connection passed to the friends-authentication operation is invalid.",
	friendsDescriptionAuthenticationHttpDnsError:         "The http dns passed to the friends-authentication operation is invalid.",
	friendsDescriptionAuthenticationHttpGetProxySetting:  "The http get proxy passed to the friends-authentication operation is invalid.",
	friendsDescriptionAuthenticationTokenExpired:         "The token passed to the friends-authentication operation is expired.",
	friendsDescriptionAuthenticationValidationFailed:     "The validation passed to the friends-authentication operation is invalid.",
	friendsDescriptionAuthenticationInvalidParam:         "The param passed to the friends-authentication operation is invalid.",
	friendsDescriptionAuthenticationPrincipalIdUnmatched: "The principal id passed to the friends-authentication operation is unmatched.",
	friendsDescriptionAuthenticationMoveCountUnmatch:     "The move count passed to the friends-authentication operation is unmatch.",
	friendsDescriptionAuthenticationUnderMaintenance:     "The maintenance passed to the friends-authentication operation is under maintenance.",
	friendsDescriptionAuthenticationUnsupportedVersion:   "The version passed to the friends-authentication operation is unsupported.",
	friendsDescriptionAuthenticationUnknown:              "An unknown error occurred in the friends-authentication operation.",
	friendsDescriptionAuthenticationUndefined:            "The friends-authentication operation returned an undefined error.",

	friendsDescriptionInvalidFriendCode: "The friend code passed to the friends operation is invalid.",
	friendsDescriptionNotLoggedIn:       "The login passed to the friends operation is not logged in.",
	friendsDescriptionNotFriendsResult:  "The result passed to the friends operation is not friends.",
	friendsDescriptionUndefinedFacility: "The facility passed to the friends operation is undefined.",
	friendsDescriptionAcNotConnected:    "The connection passed to the friends operation is not connected.",
}

func (d friendsDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", friendsDescriptionToString[d], d, friendsDescriptionDescription[d])
}
