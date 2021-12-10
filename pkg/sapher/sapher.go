package sapher

import (
	"encoding/base64"
	"log"
)

type Sapher string

func (r Sapher) Decode() string {
	decRes, err := base64.StdEncoding.DecodeString(string(r))

	if err != nil {
		log.Fatalf("Can not decode password | %s", err.Error())
	}

	return string(decRes)
}

func (r Sapher) Encode() string {

	data := []byte(r)

	if data == nil {
		log.Fatalf("Can not encode secret!")
	}

	return base64.StdEncoding.EncodeToString(data)

}
