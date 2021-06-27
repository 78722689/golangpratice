package https

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

// addr, :8443
func smain(addr string) {

	caCert, err := ioutil.ReadFile(ROOT_CA_FILE)
	if err != nil {
		log.Fatal(err)
	}

	caCertPool, _ := x509.SystemCertPool()
	if caCertPool == nil {
		//log.Println("system cert pool is empty")
		caCertPool = x509.NewCertPool()
	}

	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal("Append cert to pool failed")
	}

	cfg := &tls.Config{
		ServerName: "localhost",
		ClientAuth: tls.RequireAnyClientCert,
		ClientCAs:  caCertPool,
		MinVersion: tls.VersionTLS12, // TLS1.2 at least.
	}

	srv := &http.Server{
		Addr:      addr,
		TLSConfig: cfg,
	}
	http.Handle("/hellozz", new(helloZZHandler))

	if err := srv.ListenAndServeTLS(SERVER_CA_FILE, SERVER_KEY_FILE); err != nil {
		log.Fatal(err)
	}
}
