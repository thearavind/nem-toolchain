<p align="center">
  <img alt="Nem Toolchain Logo" src="assets/logo.png" height="128" />
  <h3 align="center">nem-toolchain</h3>
  <p align="center">Command line toolchain for <a href=https://nem.io>Nem blockchain</a>.</p>
  <p align="center">
    <a href="https://gitter.im/nem-toolchain/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge"><img alt="Join the chat at https://gitter.im/nem-toolchain/Lobby" src="https://img.shields.io/gitter/room/badges/shields.svg?style=flat-square"></a>
    <a href="https://github.com/r8d8/nem-toolchain/releases/latest"><img alt="Release" src="https://img.shields.io/github/release/r8d8/nem-toolchain.svg?style=flat-square"></a>
    <a href="https://circleci.com/gh/r8d8/nem-toolchain"><img alt="CircleCI" src="https://img.shields.io/circleci/project/github/r8d8/nem-toolchain.svg?style=flat-square"></a>
    <a href="https://travis-ci.org/r8d8/nem-toolchain"><img alt="Travis" src="https://img.shields.io/travis/r8d8/nem-toolchain.svg?style=flat-square"></a>
    <a href="https://codecov.io/gh/r8d8/nem-toolchain"><img alt="Coverage Status" src="https://img.shields.io/codecov/c/github/r8d8/nem-toolchain/master.svg?style=flat-square"></a>
    <a href="http://godoc.org/github.com/r8d8/nem-toolchain"><img alt="Go Doc" src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"></a>
    <a href="https://goreportcard.com/report/github.com/r8d8/nem-toolchain"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/r8d8/nem-toolchain?style=flat-square"></a>
    <a href="LICENSE"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square"></a>
  </p>
</p>

---

## Prerequisites

* [Go 1.8+](http://golang.org/doc/install)

## How to install

To install, simply run:

```shell
go get github.com/r8d8/nem-toolchain/cmd/nem
```

Make sure your `PATH` includes the `$GOPATH/bin` directory so your commands can be easily used:

```bash
export PATH=$PATH:$GOPATH/bin
```

## Usage examples

### New accounts

How to create a new account for testnet:

```
> nem --chain testnet account generate
Address: TBYLAB-4RILJG-ZAUH65-YVTRDO-VH52P4-ZDPDO3-SWHS
Public key: bda2ddf2cdb65267aad0032435c5669f4391f02c681aed62b06762ddda97f1c1
Private key: 7d77192b3cc6f3770ffae4f7a85f3c40bca1256abfa1e59a3cf57cefd260af80
```

Instead of option you can specify the desired network via environment variables:

```
> NEM_CHAIN=testnet nem account generate
...
```

### Vanity addresses

How to find vanity address with predefined prefix `TCGQQK` for testnet:

```
> nem --chain testnet account vanity TCGQQK
Address: TCGQQK-N5HED6-6OQ67Z-2F7GGW-Z66DWV-BFJUW6-F5WC
Public key: c342dbf7cdd3096c4c3910c511a57049e62847dd5030c7e644bc855acc1fd626
Private key: 4e017065d62f10223b989ff3f75a845fbe3df73d6c0e6d67cc4c59bea3213002
```

If you would like to search for multiple prefixes at the same time and without digits at all:

```
> nem --chain testnet account vanity --no-digits TABC TACB TBAC TBCA TCAB TCBA
Address: TACBLF-CJBFVE-TPTUIP-HBIRVI-PHQKKD-OSJMOF-KGNU
Public key: 8a68fdf463b4531f409369ffe368f9d78eb5e0b713459b767fbb4c4bfd148667
Private key: ae4e943300554508d52c863329a53e40787e994e7c2733d54c378fb88421d387
```

Important notes:

1. Mainnet addresses start with `N`, Mijin - with `M`, Testnet - with `T`.
1. Second symbols are `A`, `B`, `C`, or `D` only, so for mainnet you won't find addresses that start with `NE` or `N4`.
1. The digits `0`, `1`, `8` and `9` are not part of [Base32 encoding](https://en.wikipedia.org/wiki/Base32) and therefore will not appear in any address. 

## Other topics

* [How to contribute](CONTRIBUTING.md)
* [How to get support](SUPPORT.md)

## Thanks to

* [JetBrains](https://www.jetbrains.com) for [IntelliJ IDEA Ultimate](https://www.jetbrains.com/idea) free open-source license.

## Licence

[MIT](LICENSE)
