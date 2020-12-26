package localca

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)


func Init() {
	// First, define a CA struct
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2021),
		Subject: pkix.Name{
			Organization: []string{"Meyer Household, Ltd."},
			Country: []string{""},
			Province: []string{""},
			Locality: []string{""},
			StreetAddress: []string{""},
			PostalCode: []string{""},
		},
		NotBefore: time.Now(),
		NotAfter: time.Now().AddDate(10, 0, 0),
		IsCA: true,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage: x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}
	fmt.Printf("cert struct: %+v\n", ca)

	// Next, get a basic keypair
	caPrivateKey, shit := rsa.GenerateKey(rand.Reader, 4096)
	if shit != nil {
		fmt.Println(shit.Error())
	}
	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &caPrivateKey.PublicKey, caPrivateKey)
	if err != nil { fmt.Println(err.Error())}

	caPEM := new(bytes.Buffer)
	pem.Encode(caPEM, &pem.Block{
		Type: "CERTIFICATE",
		Bytes: caBytes,
	})

	caPrivateKeyPEM := new(bytes.Buffer)
	pem.Encode(caPrivateKeyPEM, &pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caPrivateKey),
	})

	certFile, err := os.Create("certfile.crt")
	if err != nil {fmt.Println(err.Error())}
	caPEM.WriteTo(certFile)

	keyFile, err := os.Create("certfile.key")
	if err != nil {fmt.Println(err.Error())}
	caPrivateKeyPEM.WriteTo(keyFile)
}
