package test

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	btcsecp256k1 "github.com/btcsuite/btcd/btcec"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	ethsecp256k1 "github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"

	"github.com/Dipper-Labs/Dipper-Protocol/hexutil"
	sdk "github.com/Dipper-Labs/Dipper-Protocol/types"
	"github.com/Dipper-Labs/go-sdk/keys"
)

const (
	uncompressedPubKey = "0487e7a605af50b0e57838bc8508fe80f74dfd8710f92a2c165e10b407b9385b57968620bbd71b7888915e9fa861e3e47b38aa49f029886277404ad5b82771c2e4"
	compressedPubKey   = "0287e7a605af50b0e57838bc8508fe80f74dfd8710f92a2c165e10b407b9385b57"
	address            = "1c0311d33691aa5bf659fe7ae8276cce19b304b5"
	addressBech32      = "dip1rsp3r5ekjx49hajeleawsfmvecvmxp94hdur8m"
)

func Test_Bech32AddrToHexAddr(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32(addressBech32)
	require.Nil(t, err)
	require.Equal(t, address, fmt.Sprintf("%x", addr.Bytes()))
}

func Test_UNCompressedPubKey2CompressedPubKey(t *testing.T) {
	pubKeyBytes, err := keys.UNCompressedPubKey2CompressedPubKey(uncompressedPubKey)
	require.True(t, err == nil)
	require.Equal(t, compressedPubKey, fmt.Sprintf("%x", pubKeyBytes))
}

func Test_CompressedPubKey2UNCompressedPubKey(t *testing.T) {
	pubKeyBytes, err := keys.CompressedPubKey2UNCompressedPubKey(compressedPubKey)
	require.True(t, err == nil)
	require.Equal(t, uncompressedPubKey, fmt.Sprintf("%x", pubKeyBytes))
}

func Test_UNCompressedPubKey2Address(t *testing.T) {
	addr, err := keys.UNCompressedPubKey2Address(uncompressedPubKey)
	require.True(t, err == nil)
	require.Equal(t, address, strings.ToLower(addr.String()))
}

func Test_UNCompressedPubKey2AddressBech32(t *testing.T) {
	addr, err := keys.UNCompressedPubKey2AddressBech32(uncompressedPubKey)
	require.True(t, err == nil)
	require.Equal(t, addressBech32, addr)
}

func Test_PubKeyHexString2AddressBech32(t *testing.T) {
	addr, err := keys.PubKeyHexString2AddressBech32(compressedPubKey)
	require.True(t, err == nil)
	require.Equal(t, addressBech32, addr)
}

func Test_PubKey2AddressBech32(t *testing.T) {
	pubKey, err := hex.DecodeString(compressedPubKey)
	require.True(t, err == nil)

	var pk secp256k1.PubKeySecp256k1
	copy(pk[:], pubKey)

	addrBech32, err := keys.PubKey2AddressBech32(pk)
	require.True(t, err == nil)
	require.Equal(t, addressBech32, addrBech32)
}

func Test_test(t *testing.T) {
	h1 := ethcrypto.Keccak256([]byte("abc"))
	h2 := crypto.Sha256([]byte("abc"))

	t.Log(fmt.Sprintf("\n%x\n%x", h1, h2))
}

func Test_1(t *testing.T) {
	hash := crypto.Sha256([]byte("abdadfasdfadfcd"))
	hash1, err := hexutil.Decode("0xce0677bb30baa8cf067c88db9811f4333d131bf8bcf12fe7065d211dce971008")
	t.Log(fmt.Sprintf("%x", hash1))
	t.Log(fmt.Sprintf("%x", hash))

	pri, err := btcsecp256k1.NewPrivateKey(btcsecp256k1.S256())
	t.Log(fmt.Sprintf("%x", pri))
	t.Log(fmt.Sprintf("%x", pri.D))
	t.Log(err)
	sig, err := pri.Sign(hash)
	t.Log(fmt.Sprintf("%x", sig.Serialize()))
	t.Log(err)
	sig1, err := btcsecp256k1.SignCompact(btcsecp256k1.S256(), pri, hash, true)
	t.Log(fmt.Sprintf("%x", sig1))
	sig2, err := btcsecp256k1.SignCompact(btcsecp256k1.S256(), pri, hash, false)
	t.Log(fmt.Sprintf("%x", sig2))
	t.Log(fmt.Sprintf("%d:%d:%d\n", len(sig.Serialize()), len(sig1), len(sig2)))

	x, err := hexutil.Decode("2a6e636831787735396864307a74677a35366c6d307534716567336335767330726b6b72727775656e7072")
	t.Log(fmt.Sprintf("%s\n", string(x)))
	addr, err := sdk.AccAddressFromBech32("dip1xw59hd0ztgz56lm0u4qeg3c5vs0rkkrrwuenpr")

	t.Log(hexutil.Encode(addr.Bytes()))
}

func Test_2(t *testing.T) {
	hash, _ := hexutil.Decode("0xd86cc96fd39ec373298bf7958f9a67ae334c922e9d83678ebf440bee460c3f15")

	curve := btcsecp256k1.S256()

	pk, _ := btcsecp256k1.NewPrivateKey(curve)
	t.Log(fmt.Sprintf("pk %x", pk))

	pubkey := pk.PubKey()
	t.Log(fmt.Sprintf("pubkey: %x", pubkey.SerializeCompressed()))

	t.Log(fmt.Sprintf("uncompressed pubkey: %x", pubkey.SerializeUncompressed()))
	addr, _ := keys.UNCompressedPubKey2Address(fmt.Sprintf("%x", pubkey.SerializeUncompressed()))
	t.Log(fmt.Sprintf("addr = %v", addr))

	sig, _ := btcsecp256k1.SignCompact(curve, pk, hash, false)
	t.Log(fmt.Sprintf("sig: %x", sig))

	rpubkey, _, _ := btcsecp256k1.RecoverCompact(btcsecp256k1.S256(), sig, hash)
	t.Log(fmt.Sprintf("recovered pubkey: %x", rpubkey))
}

func Test_3(t *testing.T) {
	curve := btcsecp256k1.S256()

	pk, _ := btcsecp256k1.NewPrivateKey(curve)
	t.Log(fmt.Sprintf("pk %x", pk))

	pubkey := pk.PubKey()
	t.Log(fmt.Sprintf("pubkey: %x", pubkey.SerializeCompressed()))

	t.Log(fmt.Sprintf("uncompressed pubkey: %x", pubkey.SerializeUncompressed()))
	addr, _ := keys.UNCompressedPubKey2Address(fmt.Sprintf("%x", pubkey.SerializeUncompressed()))
	t.Log(fmt.Sprintf("addr = %v", addr))

	d := fmt.Sprintf("%x%x%x%02x%016x", pubkey.SerializeCompressed()[1:], pubkey.SerializeCompressed()[1:], addr.Bytes(), 1, 100)
	dhex, _ := hexutil.Decode(d)
	t.Log(fmt.Sprintf("d: %x", dhex))

	hash := sha256.Sum256(dhex)
	t.Log(fmt.Sprintf("hash: %x", hash))
	sig, _ := btcsecp256k1.SignCompact(curve, pk, hash[:], false)
	t.Log(fmt.Sprintf("sig: %x", sig))

	rpubkey, _, _ := btcsecp256k1.RecoverCompact(btcsecp256k1.S256(), sig, hash[:])
	t.Log(fmt.Sprintf("recovered pubkey: %x", rpubkey))
}

func Test_4(t *testing.T) {
	curve := btcsecp256k1.S256()

	pk, _ := btcsecp256k1.NewPrivateKey(curve)
	t.Log(fmt.Sprintf("pk %x", pk))

	pubkey := pk.PubKey()
	t.Log(fmt.Sprintf("pubkey: %x", pubkey.SerializeCompressed()))

	t.Log(fmt.Sprintf("uncompressed pubkey: %x", pubkey.SerializeUncompressed()))
	addr, _ := keys.UNCompressedPubKey2Address(fmt.Sprintf("%x", pubkey.SerializeUncompressed()))
	t.Log(fmt.Sprintf("addr = %v", addr))

	fk := fmt.Sprintf("%x", pubkey.SerializeCompressed())
	tk := fmt.Sprintf("%x", pubkey.SerializeCompressed())
	d := fmt.Sprintf("%s%s%x%02x%016x", hexutil.Encode([]byte(fk)), hexutil.Encode([]byte(tk)), addr.Bytes(), 1, 100)
	dhex, _ := hexutil.Decode(d)
	t.Log(fmt.Sprintf("d: %x", dhex))

	hash := sha256.Sum256(dhex)
	t.Log(fmt.Sprintf("hash: %x", hash))
	sig, _ := btcsecp256k1.SignCompact(curve, pk, hash[:], false)
	t.Log(fmt.Sprintf("sig: %x", sig))

	rpubkey, _, _ := btcsecp256k1.RecoverCompact(btcsecp256k1.S256(), sig, hash[:])
	t.Log(fmt.Sprintf("recovered pubkey: %x", rpubkey))
}

func Test_Ecrecover(t *testing.T) {
	hash, err := hexutil.Decode("d2a93c9b60dc0861011ea1510c838dc3af99f15cc9b88060ae29bf2dbc3026a4")
	require.True(t, err == nil)

	sig, err := hexutil.Decode("3c8e6518da47c6e9dd2172de39971558882783017d8fdb3e105d608170751ea734887cd5367bd8222723094775a43e8b3cafa9170e00c6dd97ffafc4e52a88821c")
	require.True(t, err == nil)

	sig[64] = sig[64] - 27

	pubkey, err := ethsecp256k1.RecoverPubkey(hash, sig)
	require.True(t, err == nil)
	t.Log(fmt.Sprintf("pubkey = %x", pubkey))
}
