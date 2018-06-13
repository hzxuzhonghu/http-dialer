package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/hzxuzhonghu/http-dialer"
)

func main() {
	dialFunc := func(ctx context.Context, network, address string) (net.Conn, error) {
		conn, err := (&net.Dialer{Timeout: 30 * time.Second, KeepAlive: 30 * time.Second}).DialContext(ctx, network, address)
		log.Printf("created a connection with address: %s", conn.LocalAddr().String())
		return conn, err
	}

	dialer := http_dialer.NewDialer(dialFunc)

	transport := &http.Transport{
		DialContext: dialer.DialContext,
	}

	client := http.Client{
		Transport: transport,
	}

	res, err := client.Get("https://www.github.com")
	if err != nil {
		log.Printf("visit www.github.com failed: %v", err)
		return
	}
	io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()

	dialer.CloseAll() // without this, the underlying connection will be reused.

	res, err = client.Get("https://www.github.com")
	if err != nil {
		log.Printf("visit www.github.com failed: %v", err)
		return
	}
	io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()
	select {}
}
