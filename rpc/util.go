package rpc

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
)

func clientTLSLoad(pemFile, keyFile string) (credentials.TransportCredentials, error) {

	certPool, cert, err := certLoad(pemFile, keyFile)
	if err != nil {
		return nil, err
	}

	tlsCred := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
		RootCAs:      certPool, // 注意client端使用RootCAs，server端使用ClientCAs
	})

	return tlsCred, nil
}

func serverTLSLoad(pemFile, keyFile string) (credentials.TransportCredentials, error) {
	certPool, cert, err := certLoad(pemFile, keyFile)
	if err != nil {
		return nil, err
	}

	tlsCred := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool, // 注意client端使用RootCAs，server端使用ClientCAs
	})

	return tlsCred, nil
}

func certLoad(pemFile, keyFile string) (*x509.CertPool, tls.Certificate, error) {
	ca, err := ioutil.ReadFile("./cert/rootca.crt")
	if err != nil {
		log.Println(err)
		return nil, tls.Certificate{}, err
	}

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Println("Unable to load CA")
		return nil, tls.Certificate{}, grpc.Errorf(codes.Unauthenticated, "Unable to load CA")
	}

	cert, err := tls.LoadX509KeyPair(pemFile, keyFile)
	if err != nil {
		log.Println(err)
		return nil, tls.Certificate{}, err
	}

	return certPool, cert, nil
}
