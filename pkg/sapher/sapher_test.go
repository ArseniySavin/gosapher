package sapher

import "testing"

type cfg struct {
	Secret Sapher
}

func Test_Encode(t *testing.T) {
	cfg := cfg{
		Secret: "Secret for mars trip",
	}

	data := cfg.Secret.Encode()

	t.Log(data)

	if data != "U2VjcmV0IGZvciBtYXJzIHRyaXA=" {
		t.Fatal("Secret does not equal data")
	}
}

func Test_Decode(t *testing.T) {
	cfg := cfg{
		Secret: "U2VjcmV0IGZvciBtYXJzIHRyaXA=",
	}

	data := cfg.Secret.Decode()

	t.Log(data)

	if data != "Secret for mars trip" {
		t.Fatal("Secret does not decode")
	}
}

func Test_Encode_to_Decode(t *testing.T) {
	cfg := cfg{
		Secret: "Secret for mars trip",
	}

	data := cfg.Secret.Encode()

	t.Log(data)

	if data != "U2VjcmV0IGZvciBtYXJzIHRyaXA=" {
		t.Fatal("Secret does not equal data")
	}

	cfg.Secret = Sapher(data)

	data = cfg.Secret.Decode()

	t.Log(data)

	if data != "Secret for mars trip" {
		t.Fatal("Secret does not decode")
	}
}
