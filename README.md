simple and fast ethereum vanity address generator.

it only uses one external dependency: [ethereum/go-ethereum](https://github.com/ethereum/go-ethereum) -- the official go implementation of Ethereum.


**install & build:**

first, [install go](https://go.dev/doc/install), and then clone+build this program:
```sh
git clone https://github.com/6/simple-eth-vanity-address.git

cd simple-eth-vanity-address

# Most computers:
go build

# M1 Mac:
GOOS=darwin GOARCH=arm64 go build
```

**generating vanity address:**

```
./simple-eth-vanity-address -prefix 0xABC -suffix DEF

...some time passes...

Address: 0xABCaa219d2Ce67B09A4e5071c21a4A2B2b921DEF
Public key: ...
Private key: ...
```
