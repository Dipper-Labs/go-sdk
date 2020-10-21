# dip-chain Go SDK


- **client**: provide httpClient, LcdClient, RpcClient and TxClient for query or send transaction on dippernetwork
- **keys**: implement KeyManage to manage private key and accounts
- **types**: common types
- **util**: define constant and common functions

# Install

## Requirement

Go version above 1.13

## Use go mod(recommend)

Add "github.com/Dipper-Labs/go-sdk" dependency into your go.mod file.

```go
require (
	github.com/Dipper-Labs/go-sdk latest
)
```

# Usage

## Key Manager

Before start using API, you should have some accounts which have udippernetwork tokens. then exporting keysotre file by dipcli tool through command below:
```cassandraql
dipcli keys export <account_name>

Enter passphrase to decrypt your key:
enter your passphrase of <account_name> account
Enter passphrase to encrypt the exported key:
enter passphrase to encrypt the keystore file which can be used to import keystore to sdk

e.g.:
dipcli keys export alice

```

When you have a keystore file and corresponding passphrase, you should construct a Key Manager to help sign the transaction msg or verify signature. Key Manager is an Identity Manger to define who you are in the dippernetwork

We provide follow construct functions to generate Key Manager(other keyManager will coming soon):

```go
NewKeyManager(keystoreFile string, passphrase string) (KeyManager, error)
```

Examples:

for keyStore:

```go
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
```

## Init Client

```go
import (
	"github.com/Dipper-Labs/go-sdk/client"
)

client, err := client.NewClient("/home/ubuntu/sdk.yaml")
```

Note:
- `baseUrl`: should be lcd endpoint if you want to use liteClient
- `nodeUrl`: should be dippernetwork node endpoint, format is `tcp://host:port`
- `networkType`: `alphanet` or `mainnet`(mainnet will come later)

after you init dipclient, it include follow clients which you can use:

- `lcdClient`: lcd client for dippernetwork
- `rpcClient`: query dippernetwork info by rpc
- `txClient`: send transaction on dippernetwork

