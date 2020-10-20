package keys

import (
	"fmt"
	"io/ioutil"

	"github.com/btcsuite/btcd/btcec"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/crypto"

	"github.com/Dipper-Labs/Dipper-Protocol/app/v0/auth"
	"github.com/Dipper-Labs/Dipper-Protocol/crypto/keys/mintkey"
	sdk "github.com/Dipper-Labs/Dipper-Protocol/types"
	"github.com/Dipper-Labs/go-sdk/types"
)

type KeyManager interface {
	Sign(msg types.StdSignMsg) ([]byte, error)
	SignBytes(msg []byte) ([]byte, error)
	GetPrivKey() crypto.PrivKey
	GetAddr() sdk.AccAddress

	GetUCPubKey() (UCPubKey []byte, err error)
}

type keyManager struct {
	privKey crypto.PrivKey
	addr    sdk.AccAddress
}

func NewKeyManager(file string, auth string) (KeyManager, error) {
	k := keyManager{}
	err := k.ImportKeystore(file, auth)
	return &k, err
}

func (k *keyManager) Sign(msg types.StdSignMsg) ([]byte, error) {
	sig, err := k.makeSignature(msg)
	if err != nil {
		return nil, err
	}

	newTx := auth.NewStdTx(msg.Msgs, msg.Fee, []auth.StdSignature{sig}, msg.Memo)
	bz, err := types.Cdc.MarshalBinaryLengthPrefixed(newTx)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

func (k *keyManager) SignBytes(msg []byte) ([]byte, error) {
	return k.privKey.Sign(msg)
}

func (k *keyManager) GetPrivKey() crypto.PrivKey {
	return k.privKey
}

func (k *keyManager) GetAddr() sdk.AccAddress {
	return k.addr
}

func (k *keyManager) GetUCPubKey() (UCPubKey []byte, err error) {
	pubkey, err := btcec.ParsePubKey(k.GetPrivKey().PubKey().Bytes()[5:], btcec.S256())
	if err != nil {
		return nil, err
	}

	return pubkey.SerializeUncompressed(), nil
}

func (k *keyManager) makeSignature(msg types.StdSignMsg) (sig auth.StdSignature, err error) {
	if err != nil {
		return
	}
	sigBytes, err := k.privKey.Sign(msg.Bytes())
	if err != nil {
		return
	}
	return auth.StdSignature{
		PubKey:    k.privKey.PubKey(),
		Signature: sigBytes,
	}, nil
}

func (k *keyManager) ImportKeystore(keystoreFile string, passphrase string) error {
	if passphrase == "" {
		return fmt.Errorf("Password is missing ")
	}

	armor, err := ioutil.ReadFile(keystoreFile)
	if err != nil {
		return err
	}

	privKey, err := mintkey.UnarmorDecryptPrivKey(string(armor), passphrase)
	if err != nil {
		return errors.Wrap(err, "couldn't import private key")
	}

	addr := sdk.AccAddress(privKey.PubKey().Address())
	k.addr = addr
	k.privKey = privKey
	return nil
}
