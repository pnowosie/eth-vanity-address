# eth vanity address
simple and fast ethereum vanity address generator forked with :heart: from [6/simple-eth-vanity-address](https://github.com/6/simple-eth-vanity-address)


## Usage

### Flags
- `-prefix` e.g. `0xABC` - prefix pattern preceded with 0x
- `-suffix` e.g. `DEF` - suffix pattern
- `-ignore-case` - for case-insensitive match
- `-password` - if provided found keys with be saved in keyfile encrypted with password. It follows Ethereum's V3 keystore schema, can be inspected with `ethkey` tool.

At least either **prefix** or **suffix** have to be provided.

### Example

```bash
./eth-vanity-address \
  -prefix 0xabc \
  -suffix def \
  -ignore-case 2>> vanity.log
```
it will use available number of CPUs to spin worker goroutines. Progress (number of generated keys) is reported by separate goroutine in 15-minutes (hardcoded) intervals. It's recommended to redirect output to file as shown above.

**Sample output file**
```json
2023/07/26 18:44:09 Generating address with 12 workers, prefix=0x123456, suffix=

2023/07/26 18:44:39 Total keys checked: 1,281,799
2023/07/26 18:45:09 Total keys checked: 2,313,285
...
2023/07/26 18:47:04 Worker 10 found address:

Address    : 0xAbCaA219d2Ce67B09A4e5071c21a4A2B2b92fdeF
Public key : *****
Private key: *****

...
2023/07/26 18:48:09 Total keys checked: 8,552,048
2023/07/26 18:48:31 Received interrupt signal. BYE!
```

## Instalation

### Dependencies

- Golang version >= `1.20`


```bash
git clone https://github.com/pnowosie/eth-vanity-address.git
cd eth-vanity-address
make
```



**:warning: This software has not been audited. use at your own risk!**
