// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crc8

import (
	"github.com/knieriem/crcutil/crc8"
	"github.com/knieriem/hash"
)

// The size of a CRC-8 checksum in bytes.
const Size = 1

// Predefined polynomials.
var (
	DOWCRC = crc8.DOW
)

type Model = crc8.Model

// MakeTable used to create a Table constructed from a specified polynomial.
//
// Deprecated: Now it just returns the provided Model to keep the package's API compatible.
func MakeTable(m *Model) *Model {
	return m
}

// New creates a new hash.Hash8 computing the CRC-8 checksum
// using the polynomial represented by the Model.
func New(m *Model) hash.Hash8 {
	return &digest{Inst: m.New()}
}

// digest represents the partial evaluation of a checksum.
type digest struct {
	*crc8.Inst
}

func (d *digest) Size() int { return Size }

func (d *digest) BlockSize() int { return 1 }

func (d *digest) Sum8() uint8 {
	return d.Inst.Sum()
}

func (d *digest) Sum(in []byte) []byte {
	return d.Inst.AppendSum(in)
}

// Checksum returns the CRC-8 checksum of data
// using the polynomial, and parameters represented by the Model.
func Checksum(data []byte, m *Model) uint8 {
	return m.Checksum(data)
}
