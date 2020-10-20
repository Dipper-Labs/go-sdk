package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Dipper-Labs/go-sdk/config"
	"github.com/Dipper-Labs/go-sdk/keys"
)

func TestNewKeyManager(t *testing.T) {
	if km, err := keys.NewKeyManager(config.KeyStoreFileAbsPath, config.KeyStorePasswd); err != nil {
		t.Fatal(err)
	} else {
		msg := []byte("hello world")
		signature, err := km.GetPrivKey().Sign(msg)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(km.GetAddr().String())

		assert.Equal(t, km.GetPrivKey().PubKey().VerifyBytes(msg, signature), true)
	}
}
