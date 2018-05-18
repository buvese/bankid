package bankid

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

type Conn struct {
	serverName string
	c          http.Client
}

// New will do all you want!!
func New(addr string, srvCrt string, clntCrt tls.Certificate, srvName string) (*Conn, error) {
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(srvCrt))
	if !ok {
		return nil, fmt.Errorf("Failed to parse server certificate")
	}
	/*
		conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%s", addr, "443"), &tls.Config{
			Certificates: []tls.Certificate{clntCrt},
			RootCAs:      roots,
			ServerName:   addr,
		})
		if err != nil {
			return nil, fmt.Errorf("Failed to connect: " + err.Error())
		}
		conn.Write([]byte("POST /rp/v5/sign HTTP/1.1\nContent-Type: application/json\nHost: appapi2.test.bankid.com\nContent-Length: 129\n\n{\"personalNumber\":\"195504166249\",\"endUserIp\":\"194.168.2.25\",\"userVisibleData\":\"IFRoaXMgaXMgYSBzYW1wbGUgdGV4dCB0byBiZSBzaWduZWQ=\"}"))
		io.Copy(os.Stdout, conn)
		conn.Close()*/
	// 129 {"personalNumber":"195504166249","endUserIp":"194.168.2.25","userVisibleData":"IFRoaXMgaXMgYSBzYW1wbGUgdGV4dCB0byBiZSBzaWduZWQ="}
	tran := &http.Transport{TLSClientConfig: &tls.Config{
		Certificates: []tls.Certificate{clntCrt},
		RootCAs:      roots,
		ServerName:   addr,
	}}
	/*
		tran.DialTLS = func(network, addr2 string) (net.Conn, error) {
			conn, err := tls.Dial(network, addr2, &tls.Config{
				Certificates: []tls.Certificate{clntCrt},
				RootCAs:      roots,
				ServerName:   addr,
			})
			if err != nil {
				return conn, err
			}
			return &wrappedConn{mw: io.MultiWriter(conn, os.Stdout), conn: conn}, nil
		}*/
	return &Conn{serverName: addr, c: http.Client{Transport: tran}}, nil
}

func (c *Conn) postJSON(url string, req, res interface{}) error {
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}
	var resp *http.Response
	resp, err = c.c.Post(fmt.Sprintf("https://%s%s", c.serverName, url), "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	fmt.Printf("Status code: %s\n", resp.Status)
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(res)
}

type wrappedConn struct {
	mw   io.Writer
	conn net.Conn
}

func (w *wrappedConn) Read(b []byte) (n int, err error) {
	n, err = w.conn.Read(b)
	os.Stdout.Write(b)
	return
}
func (w *wrappedConn) Write(b []byte) (n int, err error)  { return w.mw.Write(b) }
func (w *wrappedConn) Close() error                       { return w.conn.Close() }
func (w *wrappedConn) LocalAddr() net.Addr                { return w.conn.LocalAddr() }
func (w *wrappedConn) RemoteAddr() net.Addr               { return w.conn.RemoteAddr() }
func (w *wrappedConn) SetDeadline(t time.Time) error      { return w.conn.SetDeadline(t) }
func (w *wrappedConn) SetReadDeadline(t time.Time) error  { return w.conn.SetReadDeadline(t) }
func (w *wrappedConn) SetWriteDeadline(t time.Time) error { return w.conn.SetWriteDeadline(t) }
