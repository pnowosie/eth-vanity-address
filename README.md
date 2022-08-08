simple and fast ethereum vanity address generator. 

only uses a single external dependency: [ethereum/go-ethereum](https://github.com/ethereum/go-ethereum), the official go implementation of Ethereum.


**compile:**

```
# M1 Mac:
GOOS=darwin GOARCH=arm64 go build

# Intel Mac:
go build
```

**run:**

```
./simple-eth-vanity-address -prefix 0xABC -suffix DEF

...some time passes...
Address: 0xABCaa219d2Ce67B09A4e5071c21a4A2B2b921DEF
Public key: ...
Private key: ...
```
