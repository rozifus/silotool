# Silotools

Tooling for interacting with [SiLos](https://arx.org/)

Ported to golang from [kong-org/silo-usb-nfc](https://github.com/kong-org/silo-usb-nfc)

[Kong Discord](https://discord.com/invite/dypeg4JfTX)


## Compatability

This tool requires a system with the `pcsclite` driver and a compatible smart card reader.

The following card readers have been tested as working:

- ACS ACR122U-A9

## Usage

(go commands require go compiler)

Run from directory

```bash
$ go run ./silotools/cmd/* {commands}
```

Compile and run

```bash
$ go build -o silotools ./silotools/cmd/*
```

```bash
$ ./silotools {commands}
```

### Commands

#### Silo

Read silo core data

```bash
$ silotools silo read
```

Test silo signing (random address and block)

```bash
$ silotools silo test
```
