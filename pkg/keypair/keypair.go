// Copyright 2017 The nem-toolchain project authors. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package keypair responses for account's private/public crypto keys.
package keypair

import (
	"encoding/base32"

	"strings"

	"regexp"

	"fmt"

	"github.com/r8d8/nem-toolchain/pkg/core"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

// Address length
const ADDRESS_LENGTH = 40

// Address is a readable string representation for a public key.
type Address string

// KeyPair is a private/public crypto key pair.
type KeyPair struct {
	Private []byte
	Public  []byte
}

// Gen generates a new private/public key pair using entropy from crypto rand.
func Gen() KeyPair {
	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic("assert: ed25519 generate key function internal error")
	}
	return KeyPair{priv[:32], pub}
}

// Address converts a key pair into corresponding address string representation.
func (pair KeyPair) Address(chain core.Chain) Address {
	h := sha3.SumKeccak256(pair.Public)
	r := ripemd160.New()
	_, err := r.Write(h[:])
	if err != nil {
		panic("assert: Ripemd160 hash function internal error")
	}
	b := append([]byte{chain.Id}, r.Sum(nil)...)
	h = sha3.SumKeccak256(b)
	a := append(b, h[:4]...)
	return Address(base32.StdEncoding.EncodeToString(a))
}

// PrettyString returns pretty formatted address with separators ('-').
func (addr Address) PrettyString() (string, error) {
	if len(addr) != ADDRESS_LENGTH {
		return "", fmt.Errorf(
			"invalid address length. Expected %v, but received %v", ADDRESS_LENGTH, len(addr))
	}
	ps := regexp.MustCompile(".{6}").FindAllString(string(addr), -1)
	ps = append(ps, string(addr)[ADDRESS_LENGTH-4:])
	return strings.Join(ps, "-"), nil
}

func (addr Address) String() string {
	return string(addr)
}
