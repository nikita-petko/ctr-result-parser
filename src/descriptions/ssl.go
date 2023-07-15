package descriptions

import "fmt"

type sslDescription int32

const (
	sslDescriptionNone sslDescription = iota
	sslDescriptionFailed
	sslDescriptionWantRead
	sslDescriptionWantWrite
	sslDescriptionSysCall sslDescription = iota + 1
	sslDescriptionZeroReturn
	sslDescriptionWantConnect
	sslDescriptionSslId
	sslDescriptionVerifyCommonName
	sslDescriptionVerifyRootCa
	sslDescriptionVerifyChain
	sslDescriptionVerifyDate
	sslDescriptionGetServerCert
	sslDescriptionVerifyRevokedCert sslDescription = iota + 3
	sslDescriptionState
	sslDescriptionRandom sslDescription = iota + 4
	sslDescriptionVerifyCert
	sslDescriptionCertBufAlreadySet
	sslDescriptionNotAssignServer sslDescription = iota + 50 - 18
	sslDescriptionAlreadyAssignServer
	sslDescriptionIpcSession
	sslDescriptionConnProcessMax
	sslDescriptionFailToCreateCertStore
	sslDescriptionFailToCreateCrlStore
	sslDescriptionFailToCreateClientCert
	sslDescriptionInvalidParam
	sslDescriptionClientProcessMax
	sslDescriptionIpcSessionMax
	sslDescriptionInternalCert
	sslDescriptionInternalCrl
)

var sslDescriptionToString = map[sslDescription]string{
	sslDescriptionNone:                   "None",
	sslDescriptionFailed:                 "Failed",
	sslDescriptionWantRead:               "WantRead",
	sslDescriptionWantWrite:              "WantWrite",
	sslDescriptionSysCall:                "SysCall",
	sslDescriptionZeroReturn:             "ZeroReturn",
	sslDescriptionWantConnect:            "WantConnect",
	sslDescriptionSslId:                  "SslId",
	sslDescriptionVerifyCommonName:       "VerifyCommonName",
	sslDescriptionVerifyRootCa:           "VerifyRootCa",
	sslDescriptionVerifyChain:            "VerifyChain",
	sslDescriptionVerifyDate:             "VerifyDate",
	sslDescriptionGetServerCert:          "GetServerCert",
	sslDescriptionVerifyRevokedCert:      "VerifyRevokedCert",
	sslDescriptionState:                  "State",
	sslDescriptionRandom:                 "Random",
	sslDescriptionVerifyCert:             "VerifyCert",
	sslDescriptionCertBufAlreadySet:      "CertBufAlreadySet",
	sslDescriptionNotAssignServer:        "NotAssignServer",
	sslDescriptionAlreadyAssignServer:    "AlreadyAssignServer",
	sslDescriptionIpcSession:             "IpcSession",
	sslDescriptionConnProcessMax:         "ConnProcessMax",
	sslDescriptionFailToCreateCertStore:  "FailToCreateCertStore",
	sslDescriptionFailToCreateCrlStore:   "FailToCreateCrlStore",
	sslDescriptionFailToCreateClientCert: "FailToCreateClientCert",
	sslDescriptionInvalidParam:           "InvalidParam",
	sslDescriptionClientProcessMax:       "ClientProcessMax",
	sslDescriptionIpcSessionMax:          "IpcSessionMax",
	sslDescriptionInternalCert:           "InternalCert",
	sslDescriptionInternalCrl:            "InternalCrl",
}

var sslDescriptionDescription = map[sslDescription]string{
	sslDescriptionNone:                   "No error.",
	sslDescriptionFailed:                 "Error due to SSL protocol failure. (If verification of the client certificate fails on the server side, etc.)",
	sslDescriptionWantRead:               "Processing of the Read function is not completed when using an asynchronous socket. (Please try again.)",
	sslDescriptionWantWrite:              "Processing of the Write function is not completed when using an asynchronous socket. (Please try again.)",
	sslDescriptionSysCall:                "An internal system function returned an unexpected error.",
	sslDescriptionZeroReturn:             "Zero was returned at an unexpected timing when Socket Read/Write was executed internally.",
	sslDescriptionWantConnect:            "SSL handshake (DoHandshake) not completed when using asynchronous socket. (Please try again.)",
	sslDescriptionSslId:                  "Internal error (bad SSLID)",
	sslDescriptionVerifyCommonName:       "Server authentication error. The CommonName of the server certificate does not match the host name of the destination server specified in AssignServer().",
	sslDescriptionVerifyRootCa:           "Server authentication error. The Root CA certificate of the server certificate does not match the certificate configured for Connection.",
	sslDescriptionVerifyChain:            "Server authentication error. The certificate chain of the server certificate is invalid.",
	sslDescriptionVerifyDate:             "Server authentication error. Server certificate is out of date.",
	sslDescriptionGetServerCert:          "Failed to store certificate data in buffer (Occurs when certificate size is larger than buffer in DoHandshake() with arguments)",
	sslDescriptionVerifyRevokedCert:      "Server authentication error. The server certificate was registered in the revocation list.",
	sslDescriptionState:                  "The status of the SSL library is incorrect (occurs when \"another library function was executed without initializing\").",
	sslDescriptionRandom:                 "Random number processing error.",
	sslDescriptionVerifyCert:             "Server certificate verification NG.",
	sslDescriptionCertBufAlreadySet:      "A buffer for storing the server certificate has already been set.",
	sslDescriptionNotAssignServer:        "Destination server not assigned",
	sslDescriptionAlreadyAssignServer:    "Destination server already assigned",
	sslDescriptionIpcSession:             "Bad IPC session",
	sslDescriptionConnProcessMax:         "Exceeded the maximum number of connections used by one process",
	sslDescriptionFailToCreateCertStore:  "Failed to create certificate store",
	sslDescriptionFailToCreateCrlStore:   "Failed to create CRL store",
	sslDescriptionFailToCreateClientCert: "Failed to create client certificate",
	sslDescriptionInvalidParam:           "Invalid parameter",
	sslDescriptionClientProcessMax:       "The number of clients that can be used simultaneously is already in use.",
	sslDescriptionIpcSessionMax:          "The number of IPC sessions that can be connected at the same time has already been connected (that is, both the number of clients and the number of connections have reached their maximum values).",
	sslDescriptionInternalCert:           "Failed to use built-in certificate",
	sslDescriptionInternalCrl:            "Failed to use built-in CRL",
}

func (d sslDescription) toString() string {
	return fmt.Sprintf("%s (%d) - %s", sslDescriptionToString[d], d, sslDescriptionDescription[d])
}
