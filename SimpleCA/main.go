package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	go server()
	client()
}

func client() {

	cert, err := tls.LoadX509KeyPair("cert/ca.pem", "cert/ca.key")
	if err != nil {
		log.Fatal(err)
	}

	caCert, err := ioutil.ReadFile("cert/ca.pem")

	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}
	_, err = client.Get("https://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

}

func server() {
	server := http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello, World!")
		}),
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	log.Println("Server started")
	server.ListenAndServeTLS("cert/ca.pem", "cert/ca.key")

}
