package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func demo02() {

	key := `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAnI30fe5eqBZsI/7e6xLHx3jVMFK34RS9FI5Tp3T3hwldoCSW
BMo3Y+gEnoc5F7yEF/rL1QNe/370ikpeVtAnUAEf4QMsqL4s2CTGciG0mmKMFuVR
TYWkWKa26EWooF02Od6Lk1Zfvl0pSa+FpuSeWUmUXkzoTg2NIeVb9thKL8ZEcUDv
KEBVxn3UujZYSeGrOU7LqvzgZ1Hf8mFSurUgI895Ptl5heR/GP9I3sdlEPk9Bxti
IyJkahk/SvswZcNU1PP9UHtkM+8Rz1+qV015Wr4SlzCti1iqlv6tDnPOAaa27iWO
xv8FedOHrj/iy2CWs5PVwLCQonYYTgZdp6DvNQIDAQABAoIBABNQ2HepL3ihu5n9
Wle6q3eLCxK5QgK9Yk7k+KFSpy+8EhcuLGxO3uKuv9Mnb/3jvpCO2xNfJRt0+IsP
ffBMIm2Wle5XS+1N+Y55ixqN2hCqVAqHJK7h7FMbzrl9zr5qgcRGFwaIw6tjJzL2
OxFKZ9cwYuP3rTnkJiOmfuBH829NX7nCiybYgTWYIaHn9zTFxsfUilitr3LVrKf0
fnbrKO/UwCgUJtdxUux798QJMkY87hNFhMc+taED5x0MtGD8xFVT/OpHP1fjA69P
SfEuRwdPAE6gRusWysEOYl5J3O7Z6ZJ2XjOoAmJO0eolXk2g+EgQrjOJ1nEagx3Z
Ai6pz0kCgYEAz3HVjmpXQ68JzhpuOdLSyFtiRMeT+pREuUNbpS+xGdoMn7Vq/CCk
/+mzRz1e/aQDELQ4ND694xvRIgsNu1rkZ9zzdMm6Dlb/VeesFFBcsV5oO9CevAgc
jIJaNGkzSsh80G7a1+oklInmZutgsmNYsj38uANl0/GeXlDyAvOOAv8CgYEAwTLD
TLBe+DDdeh4Nm8ApLYhmqg4mN+7gVk/A5lvrBJxGrlOp6uMvVrrTDv2iNGh9utlj
+BwhO3kp6ChxHbwg7QdxG3tqeSjIHBOtd77gT8FiUzW9hYKMKHhFYkXCV4KmkA3z
wX/k0wJAMiS5Dn+8Oop0iQ8JM8O+tVrZ+RNOccsCgYBkzwTguHy/v4gcSGvYXcNg
6kfO9M3Mr9DV1w3qMEr4LSaFVjwvWg3U71TYAdLvn4x9yZamjPbqLtJqoMSDj5//
eTg2QPHRliVvMa5GEOd6qk4tAyygY7C793yui5EJfnRpNfXTh01PgdQdn7TV+8FW
nRj6s2wAmyJC40M4e/M/MQKBgEkrpaYE4mMTdr5KZZMIi0/wahmb9EOY2CQ4XnGg
fHkpiud5TlBkrMgnR8pUqzs4G+5JaUBNvxRcRk/kCTtexgWfnOnwqxmqJWfk1/0T
MLdkIY4sO1Q2ZkjpjDW0/+7Oz7CdaCVv/8UfBBs5p1DjdfGiAJLsO+r8Bt1+xDE+
a6ZTAoGBAI/JnDSwIkc8Fgywaf4ohhDT3lyloNjXebEZmqdssirPkJJfRMTT9nGc
NNFZgKmnq1iEn1BJKV4G+ah5FX3D1SNzgnNduFgkCHHzYaWsJEPupb14BhxkD8tR
qz6hetldGocG6VtilK0tDakR+U+suTYJOlQ48NuamorhnDDpHkDN
-----END RSA PRIVATE KEY-----`

	pem := `-----BEGIN CERTIFICATE-----
MIICaDCCAe4CFFX3ACVuVZ5LGTnb/eGOe72dwiVaMAoGCCqGSM49BAMCMCYxJDAi
BgNVBAoTG0FwcGxpYW5jZSBSb290IENBIEF1dGhvcml0eTAeFw0yMDExMDIwODA4
MjJaFw0yMTExMDIwODA4MjJaMFwxCzAJBgNVBAYTAmNuMQswCQYDVQQIDAJzaDEL
MAkGA1UEBwwCc2gxEjAQBgNVBAoMCWNlcnR1c25ldDERMA8GA1UECwwId2lyZWxl
c3MxDDAKBgNVBAMMA0E6QjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
AJyN9H3uXqgWbCP+3usSx8d41TBSt+EUvRSOU6d094cJXaAklgTKN2PoBJ6HORe8
hBf6y9UDXv9+9IpKXlbQJ1ABH+EDLKi+LNgkxnIhtJpijBblUU2FpFimtuhFqKBd
Njnei5NWX75dKUmvhabknllJlF5M6E4NjSHlW/bYSi/GRHFA7yhAVcZ91Lo2WEnh
qzlOy6r84GdR3/JhUrq1ICPPeT7ZeYXkfxj/SN7HZRD5PQcbYiMiZGoZP0r7MGXD
VNTz/VB7ZDPvEc9fqldNeVq+EpcwrYtYqpb+rQ5zzgGmtu4ljsb/BXnTh64/4stg
lrOT1cCwkKJ2GE4GXaeg7zUCAwEAATAKBggqhkjOPQQDAgNoADBlAjEA29lULM/o
G8xtg+AVqZmR/jmsYSE/aTdhs3zfyiL6/iP1AqN68sgSLMvbfof+n3K+AjAE2l4V
FdWaX5Y9MOtx6y1YhIYG+z2eBA1NCpZBnzU9grgT8Vka2cxWoWdw+sQx7ac=
-----END CERTIFICATE-----
`

	// 目前用不上的 ca
	//pool := x509.NewCertPool()
	//caCrt, err := ioutil.ReadFile(cert.K8sAgentCrtPath)
	//if err != nil {
	//	log.Fatal("read ca.crt file error:", err.Error())
	//}
	//pool.AppendCertsFromPEM(caCrt)
	//cliCrt, err := tls.LoadX509KeyPair(cert.ClientCrt, cert.ClientKey)

	cliCrt, err := tls.X509KeyPair([]byte(pem), []byte(key))
	if err != nil {
		log.Fatalln("LoadX509KeyPair error:", err.Error())
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			//RootCAs:      pool,
			InsecureSkipVerify: true,
			Certificates:       []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://bilibili.com/")
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
