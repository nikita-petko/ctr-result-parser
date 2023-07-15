package descriptions

import "fmt"

type httpDescription int32

const (
	httpDescriptionNone httpDescription = iota
	httpDescriptionInvalidStatus
	httpDescriptionInvalidParam
	httpDescriptionNoMem
	httpDescriptionCreateEvent
	httpDescriptionCreateMutex
	httpDescriptionCreateQueue
	httpDescriptionCreateThread
	httpDescriptionConnectionNotInitialized httpDescription = iota + 100 - 8
	httpDescriptionAlreadyAssignHost
	httpDescriptionSession
	httpDescriptionClientProcessMax
	httpDescriptionIpcSessionMax
	httpDescriptionTimeout
	httpDescriptionMsgqSendLsn httpDescription = iota - 4
	httpDescriptionMsgqRecvLsn
	httpDescriptionMsgqRecvComm
	httpDescriptionConnNoMore httpDescription = iota + 20 - 17
	httpDescriptionConnNoSuch
	httpDescriptionConnStatus
	httpDescriptionConnAdd
	httpDescriptionConnCanceled
	httpDescriptionConnHostMax
	httpDescriptionConnProcessMax
	httpDescriptionReqUrl httpDescription = iota + 30 - 24
	httpDescriptionReqConnMsgPort
	httpDescriptionReqUnknownMethod
	httpDescriptionResHeader httpDescription = iota + 40 - 27
	httpDescriptionResNoNewline
	httpDescriptionResBodyBuf
	httpDescriptionResBodyBufShortage
	httpDescriptionPostAddedAnother httpDescription = iota + 50 - 31
	httpDescriptionPostBoundary
	httpDescriptionPostSend
	httpDescriptionPostUnknownEncType
	httpDescriptionPostNoData
	httpDescriptionSsl httpDescription = iota + 60 - 36
	httpDescriptionSslCertExist
	httpDescriptionSslNoCaCertStore httpDescription = iota + 200 - 38
	httpDescriptionSslNoClientCert
	httpDescriptionSslCaCertStoreMax
	httpDescriptionSslClientCertMax
	httpDescriptionSslFailToCreateCertStore
	httpDescriptionSslFailToCreateClientCert
	httpDescriptionSslCrlExist   httpDescription = iota + 62 - 44
	httpDescriptionSslNoCrlStore httpDescription = iota + 206 - 45
	httpDescriptionSslCrlStoreMax
	httpDescriptionSslFailToCreateCrlStore
	httpDescriptionSocDns httpDescription = iota + 70 - 48
	httpDescriptionSocSend
	httpDescriptionSocRecv
	httpDescriptionSocConn
	httpDescriptionSocKeepAliveDisconnected

	httpDescriptionGetProxySetting httpDescription = iota + 80 - 53
)

var httpDescriptionToString = map[httpDescription]string{
	httpDescriptionNone:                      "None",
	httpDescriptionInvalidStatus:             "InvalidStatus",
	httpDescriptionInvalidParam:              "InvalidParam",
	httpDescriptionNoMem:                     "NoMem",
	httpDescriptionCreateEvent:               "CreateEvent",
	httpDescriptionCreateMutex:               "CreateMutex",
	httpDescriptionCreateQueue:               "CreateQueue",
	httpDescriptionCreateThread:              "CreateThread",
	httpDescriptionConnectionNotInitialized:  "ConnectionNotInitialized",
	httpDescriptionAlreadyAssignHost:         "AlreadyAssignHost",
	httpDescriptionSession:                   "Session",
	httpDescriptionClientProcessMax:          "ClientProcessMax",
	httpDescriptionIpcSessionMax:             "IpcSessionMax",
	httpDescriptionTimeout:                   "Timeout",
	httpDescriptionMsgqSendLsn:               "MsgqSendLsn",
	httpDescriptionMsgqRecvLsn:               "MsgqRecvLsn",
	httpDescriptionMsgqRecvComm:              "MsgqRecvComm",
	httpDescriptionConnNoMore:                "ConnNoMore",
	httpDescriptionConnNoSuch:                "ConnNoSuch",
	httpDescriptionConnStatus:                "ConnStatus",
	httpDescriptionConnAdd:                   "ConnAdd",
	httpDescriptionConnCanceled:              "ConnCanceled",
	httpDescriptionConnHostMax:               "ConnHostMax",
	httpDescriptionConnProcessMax:            "ConnProcessMax",
	httpDescriptionReqUrl:                    "ReqUrl",
	httpDescriptionReqConnMsgPort:            "ReqConnMsgPort",
	httpDescriptionReqUnknownMethod:          "ReqUnknownMethod",
	httpDescriptionResHeader:                 "ResHeader",
	httpDescriptionResNoNewline:              "ResNoNewline",
	httpDescriptionResBodyBuf:                "ResBodyBuf",
	httpDescriptionResBodyBufShortage:        "ResBodyBufShortage",
	httpDescriptionPostAddedAnother:          "PostAddedAnother",
	httpDescriptionPostBoundary:              "PostBoundary",
	httpDescriptionPostSend:                  "PostSend",
	httpDescriptionPostUnknownEncType:        "PostUnknownEncType",
	httpDescriptionPostNoData:                "PostNoData",
	httpDescriptionSsl:                       "Ssl",
	httpDescriptionSslCertExist:              "SslCertExist",
	httpDescriptionSslNoCaCertStore:          "SslNoCaCertStore",
	httpDescriptionSslNoClientCert:           "SslNoClientCert",
	httpDescriptionSslCaCertStoreMax:         "SslCaCertStoreMax",
	httpDescriptionSslClientCertMax:          "SslClientCertMax",
	httpDescriptionSslFailToCreateCertStore:  "SslFailToCreateCertStore",
	httpDescriptionSslFailToCreateClientCert: "SslFailToCreateClientCert",
	httpDescriptionSslCrlExist:               "SslCrlExist",
	httpDescriptionSslNoCrlStore:             "SslNoCrlStore",
	httpDescriptionSslCrlStoreMax:            "SslCrlStoreMax",
	httpDescriptionSslFailToCreateCrlStore:   "SslFailToCreateCrlStore",
	httpDescriptionSocDns:                    "SocDns",
	httpDescriptionSocSend:                   "SocSend",
	httpDescriptionSocRecv:                   "SocRecv",
	httpDescriptionSocConn:                   "SocConn",
	httpDescriptionSocKeepAliveDisconnected:  "SocKeepAliveDisconnected",
	httpDescriptionGetProxySetting:           "GetProxySetting",
}

var httpDescriptionDescription = map[httpDescription]string{
	httpDescriptionNone:                      "No error",
	httpDescriptionInvalidStatus:             "Invalid status",
	httpDescriptionInvalidParam:              "Illegal parameter",
	httpDescriptionNoMem:                     "Dynamic memory allocation failed",
	httpDescriptionCreateEvent:               "Event generation failure",
	httpDescriptionCreateMutex:               "Mutex creation failure",
	httpDescriptionCreateQueue:               "Message queue creation failure",
	httpDescriptionCreateThread:              "Thread creation failure",
	httpDescriptionConnectionNotInitialized:  "Destination no allocated",
	httpDescriptionAlreadyAssignHost:         "Destination already allocated",
	httpDescriptionSession:                   "The IPC session is invalid",
	httpDescriptionClientProcessMax:          "A number of clients equivalent to the maximum number of simultaneous client processes is already being used.",
	httpDescriptionIpcSessionMax:             "The maximum number of simultaneous IPC session connections is already connected. (The number of clients and number of connections are both already at the maximum.)",
	httpDescriptionTimeout:                   "Timeout",
	httpDescriptionMsgqSendLsn:               "Failure sending to the listener thread message queue.",
	httpDescriptionMsgqRecvLsn:               "Failure receiving from the listener thread message queue.",
	httpDescriptionMsgqRecvComm:              "Failure receiving from the communication thread message queue.",
	httpDescriptionConnNoMore:                "The maximum number of registerable connection handles was exceeded.",
	httpDescriptionConnNoSuch:                "No such connection handle.",
	httpDescriptionConnStatus:                "The connection handle status is invalid.",
	httpDescriptionConnAdd:                   "Connection handle registration failed.",
	httpDescriptionConnCanceled:              "The connection handle was canceled.",
	httpDescriptionConnHostMax:               "Exceeded the maximum number of simultaneous connections to the same host.",
	httpDescriptionConnProcessMax:            "Exceeded the maximum number of simultaneous connections used by one process.",
	httpDescriptionReqUrl:                    "Invalid URL.",
	httpDescriptionReqConnMsgPort:            "Invalid CONNECT send port number.",
	httpDescriptionReqUnknownMethod:          "Unknown method.",
	httpDescriptionResHeader:                 "Invalid HTTP header.",
	httpDescriptionResNoNewline:              "No newline character in the HTTP header.",
	httpDescriptionResBodyBuf:                "HTTP body receive buffer error.",
	httpDescriptionResBodyBufShortage:        "The HTTP body receive buffer is too small.",
	httpDescriptionPostAddedAnother:          "Failure adding to POST data. (A different POST type already exists.)",
	httpDescriptionPostBoundary:              "Boundary cannot be set properly.",
	httpDescriptionPostSend:                  "POST request send failure.",
	httpDescriptionPostUnknownEncType:        "Unknown encoding type.",
	httpDescriptionPostNoData:                "Send data not set.",
	httpDescriptionSsl:                       "SSL error.",
	httpDescriptionSslCertExist:              "The SSL certificate is already set. (It must be deleted before re-registering.)",
	httpDescriptionSslNoCaCertStore:          "No such SSL CA certificate store is registered.",
	httpDescriptionSslNoClientCert:           "No such SSL client certificate is registered.",
	httpDescriptionSslCaCertStoreMax:         "The maximum number of simultaneously registerable SSL per-process CA certificate stores is already registered.",
	httpDescriptionSslClientCertMax:          "The maximum number of simultaneously registerable SSL per-process client certificates is already registered.",
	httpDescriptionSslFailToCreateCertStore:  "Failed to create SSL certificate store.",
	httpDescriptionSslFailToCreateClientCert: "Failed to create SSL client certificate.",
	httpDescriptionSslCrlExist:               "The SSL CRL is already set. (It must be deleted before re-registering.)",
	httpDescriptionSslNoCrlStore:             "No such SSL CRL store is registered.",
	httpDescriptionSslCrlStoreMax:            "The maximum number of simultaneously registerable SSL per-process CRL stores is already registered.",
	httpDescriptionSslFailToCreateCrlStore:   "Failed to create SSL CRL store.",
	httpDescriptionSocDns:                    "DNS name resolution failure.",
	httpDescriptionSocSend:                   "Socket data send failure.",
	httpDescriptionSocRecv:                   "Socket data receive failure.",
	httpDescriptionSocConn:                   "Socket connection failure.",
	httpDescriptionSocKeepAliveDisconnected:  "Keep-alive communications have been disconnected from the server side.",
	httpDescriptionGetProxySetting:           "Failed to get the proxy value set by the device.",
}

func (d httpDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", httpDescriptionToString[d], d, httpDescriptionDescription[d])
}
