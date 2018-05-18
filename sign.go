package bankid

type SignRequirement struct {
	CardReader             string   `json:"cardReader,omitempty"`
	CertificatePolices     []string `json:"certificatePolices,omitempty"`
	IssuerCn               string   `json:"issuerCn,omitempty"`
	AutoStartTokenRequired bool     `json:"autoStartTokenRequired,omitempty"`
	AllowFingerprint       bool     `json:"allowFingerprint,omitempty"`
}

type SignRequest struct {
	EndUserIP      string `json:"endUserIp"`
	PersonalNumber string `json:"personalNumber,omitempty"`
	// UserVisibleData needs to be base64 encoded
	UserVisibleData    string           `json:"userVisibleData"`
	UserNonVisibleData string           `json:"userNonVisibleData,omitempty"`
	Requirement        *SignRequirement `json:"requirement,omitempty"`
}

type SignResponse struct {
	AutoStartToken string `json:"autoStartToken"`
	OrderRef       string `json:"orderRef"`
}

func (c *Conn) Sign(req *SignRequest) (res *SignResponse, err error) {
	res = &SignResponse{}
	err = c.postJSON("/rp/v5/sign", req, res)
	return
}
