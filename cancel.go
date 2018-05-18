package bankid

type CancelRequest struct {
	OrderRef string `json:"orderRef"`
}

type CancelResponse struct {
	ErrorCode string `json:"errorCode"`
	Details   string `json:"details"`
}

func (c *Conn) Cancel(orderRef string) (res *CancelResponse, err error) {
	req := &CancelRequest{OrderRef: orderRef}
	res = &CancelResponse{}
	err = c.postJSON("/rp/v5/cancel", req, res)

	return
}
