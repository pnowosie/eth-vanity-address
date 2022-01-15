compile:

```
# M1 Mac:
GOOS=darwin GOARCH=arm64 go build

# Intel Mac:
go build
```

run:

```
./simple-eth-vanity-address -prefix 0x123 -suffix 456
```
