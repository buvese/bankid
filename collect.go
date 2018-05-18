package bankid

type CollectStatus string

const (
	Pending  CollectHintCode = "pending"
	Complete                 = "complete"
)

type CollectHintCode string

const (
	// Pending
	OutstandingTransaction CollectHintCode = "outstandingTransaction"
	NoClient                               = "noClient"
	Started                                = "started"
	UserSign                               = "userSign"
	// Failed
	ExpiredTransaction = "expiredTransaction"
	CertificateErr     = "certificateErr"
	UserCancel         = "userCancel"
	Cancelled          = "cancelled"
	StartFailed        = "startFailed"
)

type User struct {
	PersonalNumber string `json:"personalNumber"`
	Name           string `json:"name"`
	GiveName       string `json:"giveName"`
	Surname        string `json:"surname"`
}

type Device struct {
	IpAddress string `json:"ipAddress"`
}

type Cert struct {
	NotBefore string `json:"notBefore"`
	NotAfter  string `json:"notAfter"`
}

type CollectCompletionData struct {
	User         User   `json:"user"`
	Device       Device `json:"device"`
	Cert         Cert   `json:"cert"`
	Signature    string `json:"signature"`
	OcspResponse string `json:"ocspResponse"`
}

type CollectRequest struct {
	OrderRef string `json:"orderRef"`
}

type CollectResponse struct {
	OrderRef       string                 `json:"orderRef"`
	Status         CollectStatus          `json:"status"`
	HintCode       CollectHintCode        `json:"hintCode"`
	CompletionData *CollectCompletionData `json:"completionData,omitempty"`
}

func (c *Conn) Collect(orderRef string) (res *CollectResponse, err error) {
	req := &CollectRequest{OrderRef: orderRef}
	res = &CollectResponse{}
	err = c.postJSON("/rp/v5/collect", req, res)

	return
}
