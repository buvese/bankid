package bankid

type AuthenticateRequirement struct {
	CardReader             string   `json:"cardReader,omitempty"`
	CertificatePolices     []string `json:"certificatePolices,omitempty"`
	IssuerCn               string   `json:"issuerCn,omitempty"`
	AutoStartTokenRequired bool     `json:"autoStartTokenRequired,omitempty"`
	AllowFingerprint       bool     `json:"allowFingerprint,omitempty"`
}

type AuthenticateRequest struct {
	EndUserIP          string                   `json:"endUserIp"`
	PersonalNumber     string                   `json:"personalNumber,omitempty"`
	UserVisibleData    string                   `json:"userVisibleData"`
	UserNonVisibleData string                   `json:"userNonVisibleData,omitempty"`
	Requirement        *AuthenticateRequirement `json:"requirement,omitempty"`
}

type AuthenticateResponse struct {
	AutoStartToken string `json:"autoStartToken"`
	OrderRef       string `json:"orderRef"`
}

func (c *Conn) Authenticate(req *AuthenticateRequest) (res *AuthenticateResponse, err error) {
	res = &AuthenticateResponse{}
	err = c.postJSON("/rp/v5/auth", req, res)
	return
}
