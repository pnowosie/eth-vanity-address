# eth vanity address
simple and fast ethereum vanity address generator forked with :heart: from [6/simple-eth-vanity-address](https://github.com/6/simple-eth-vanity-address)


## Usage

### Flags
- `-prefix` e.g. `0xABC` - prefix pattern preceded with 0x
- `-suffix` e.g. `DEF` - suffix pattern
- `-ignore-case` - for case-insensitive match
- `-password` - if provided found keys will be saved in keyfile encrypted with password. It follows [Web3 Secret Storage Definition V3](https://github.com/ethereum/wiki/wiki/Web3-Secret-Storage-Definition) keystore schema, can be inspected with `ethkey` tool. Alternatively one can set `ethVA_PASSWORD` environment variable.

At least either **prefix** or **suffix** have to be provided.

### Example

```bash
./eth-vanity-address \
  -prefix 0xabc \
  -suffix def \
  -ignore-case 2>> vanity.log
```

- it will use available number of CPUs to spin worker goroutines. 
- Progress (number of generated keys) is reported by separate goroutine in 15-minutes (hardcoded) intervals. 
- It won't stop searching for key once found. 
- It's recommended to redirect output to file as shown above, because stdout might be less reliable (e.g. on [Replit](https://replit.com))


<details>
  <summary><b>Sample output file</b></summary>

```txt
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

</details>

## Installation

### RECOMMENDED: Use release

**GH Cli** isn't hard dependency, you can `curl` instead, but why make your life harder?
**Make sure:** pattern `-p` matches correct target architecture. Run `gh release view --repo pnowosie/eth-vanity-address` when uncertain.

```bash
gh release download -p "*darwin_amd64*" --repo pnowosie/eth-vanity-address \
  && tar -xvf eth-vanity-address*.tar.gz \
  && mv eth-vanity-address $GOBIN/
```

### Build from sources

This approach is recommended only if there is no release matching target architecture.

**Dependencies:**
- Golang version >= `1.20`
- `GOBIN` env variable is set and added to `PATH``

```bash
go install github.com/pnowosie/eth-vanity-address@latest
```

## How it is different from fork?

Most significant changes are:
- doesn't stop after first found,
- saves key in standard keystore file (see `password` option),
- removed `concurrency` option - uses all available cores instead,
- progress (number of checked keys) is reported periodically,
- subjectively less bloated output.

For more details, run:
```bash
git log --pretty=oneline --abbrev-commit 1a4f194..@
```


**:warning: This software has not been audited. use at your own risk!**
