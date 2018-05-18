package bankid

import (
	"fmt"
	"testing"
)

/*
Namn: 			Elly Brodin
Gata: 			Malsjö kvarn
Postort: 		371 33 KARLSKRONA
Telefon: 		0455-1553789
Mail: 			elly.brodin@dodgit.com
Ålder: 			63 år
Födelsedatum: 	1955-04-16
Personnummer: 	550416-6249
IP-adress: 		85.229.174.166
PIN-kod: 		8303
Lösenord: 		r^{DiSNn
Skostorlek: 	43
*/

func TestConnect(t *testing.T) {
	cert, err := ClientTestCert()
	if err != nil {
		t.Error("Failed to get client test certificate", err)
	}
	conn, err := New(TestUri, TestCert, cert, "")
	if err != nil {
		t.Error(err)
	}
	signReq := &SignRequest{EndUserIP: "85.229.174.166", PersonalNumber: "195504166249", UserVisibleData: "IFRoaXMgaXMgYSBzYW1wbGUgdGV4dCB0byBiZSBzaWduZWQ="}
	signRes, err := conn.Sign(signReq)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Token: %s, OrderRef: %s\n", signRes.AutoStartToken, signRes.OrderRef)
	colResp, err := conn.Collect(signRes.OrderRef)
	fmt.Printf("OrderRef: %s, Status: %s, HintCode: %s, CompletionData: %s\n", colResp.OrderRef, colResp.Status, colResp.HintCode, colResp.CompletionData)
}
