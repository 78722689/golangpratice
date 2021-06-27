package https

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// TLSClient
type TLSClient struct {
	tlsClient  *http.Client
	serverAddr string // localhost:8443
	data       chan []string
	result     chan Result
}

// Data data to send
type Result struct {
	Data []bool
	err  error
}

var HTTPSClient *TLSClient

func NewClient(data []string) ([]bool, error) {
	var err error
	if HTTPSClient == nil {
		if HTTPSClient, err = NewTLSClient("localhost:8443"); err != nil {
			return nil, err
		}
		go HTTPSClient.Start()
	}

	HTTPSClient.data <- data
	result := <-HTTPSClient.result

	return result.Data, result.err
	//return HTTPSClient.Send(data)
}

// NewTLSClient
func NewTLSClient(serverAddr string) (*TLSClient, error) {

	tls := TLSClient{serverAddr: serverAddr,
		data:   make(chan []string),
		result: make(chan Result)}

	var err error
	if tls.tlsClient, err = tls.initHTTPSClient(); err != nil {
		return nil, err
	}

	//go tls.start(ctx)

	return &tls, nil
}

func (client *TLSClient) initHTTPSClient() (*http.Client, error) {
	cert, err := tls.LoadX509KeyPair(CLIENT_CERT_FILE, CLIENT_KEY_FILE)
	if err != nil {
		return nil, err
	}

	caCert, err := ioutil.ReadFile(ROOT_CA_FILE)
	if err != nil {
		return nil, err
	}
	caCertPool, _ := x509.SystemCertPool()
	if caCertPool == nil {
		caCertPool = x509.NewCertPool()
	}
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		return nil, errors.New("no cert append")
	}

	t := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
			RootCAs:      caCertPool,
		},
	}

	return &http.Client{Transport: t, Timeout: 10 * time.Second}, nil
}

func (client *TLSClient) Start() {
	for {
		select {
		case data := <-client.data:
			value, error := client.Send(data)
			result := Result{value, error}
			client.result <- result
		}
	}
}

func (client *TLSClient) Send(data []string /*ctx context.Context*/) ([]bool, error) {
	if client.tlsClient == nil {
		return nil, errors.New("Not created TLS client")
	}

	json_data, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	host := "https://" + client.serverAddr + "/hellozz"
	resp, err := client.tlsClient.Post(host, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}

	var res []bool
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func (client *TLSClient) Get(path string) ([]string, error) {
	host := "https://" + client.serverAddr + path
	resp, err := client.tlsClient.Get(host)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var result []string
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	//log.Printf("%s, %v\n", result, resp.Status)
	return result, nil
}
