package main

import (
	"bytes"
	"encoding/base64"
	"github.com/tjfoc/gmsm/sm2"
	"log"
	"math/big"
	"os"
)

var (
	default_uid = []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38}
)

func Sign(body string) (string, error) {
	cwd, _ := os.Getwd()
	PriKeyPath := cwd + string(os.PathSeparator) + "sm2PriKeyPkcs8.pem"

	priKey, e := sm2.ReadPrivateKeyFromPem(PriKeyPath, nil)
	if e != nil {
		log.Println("priKeyPem read failed, error: ", e)
		return "", e
	}

	r, s, err := sm2.Sm2Sign(priKey, []byte(body), default_uid)
	if err != nil {
		log.Println("priKey sign error: ", err)
		return "", err
	}

	//Buffer是一个实现了读写方法的可变大小的字节缓冲
	var buffer bytes.Buffer
	buffer.Write(r.Bytes())
	buffer.Write(s.Bytes())

	signature := base64.StdEncoding.EncodeToString(buffer.Bytes())
	log.Println("priKey signature base64: ", signature)
	return signature, nil
}

func Verify(body, signature string) {
	cwd, _ := os.Getwd()
	PubKeyPath := cwd + string(os.PathSeparator) + "sm2PubKey.pem"

	pubKey, e := sm2.ReadPublicKeyFromPem(PubKeyPath, nil)

	if e != nil {
		log.Println("pubKeyPem read failed, error: ", e)
	}

	d64, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		log.Println("base64 decode error: ", err)
	}

	l := len(d64)
	br := d64[:l/2]
	bs := d64[l/2:]

	var ri, si big.Int
	r := ri.SetBytes(br)
	s := si.SetBytes(bs)
	v := sm2.Sm2Verify(pubKey, []byte(body), default_uid, r, s)
	log.Printf("pubKey verified: %v\n", v)
}

func main() {
	body := `{"name":"mike","gender":"male"}`
	signature, _ := Sign(body)
	Verify(body, signature)
}
