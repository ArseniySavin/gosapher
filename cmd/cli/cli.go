package cli

import "gosapher/pkg/sapher"

type Secret struct {
	Str sapher.Sapher
}

func Encode(str string) string {
	secret := Secret{Str: sapher.Sapher(str)}

	return secret.Str.Encode()
}


func Decode(str string) string {
	secret := Secret{Str: sapher.Sapher(str)}

	return secret.Str.Decode()
}