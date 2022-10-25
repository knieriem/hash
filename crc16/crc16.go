// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crc16

import (
	"github.com/knieriem/crcutil/crc16"
	"github.com/knieriem/hash"
)

// The size of a CRC-16 checksum in bytes.
const Size = 2

// Predefined models.
var (
	Modbus = crc16.Modbus
	IBMCRC = crc16.Modbus
)

type Model = crc16.Model

// MakeTable used to create a Table constructed from a specified polynomial.
//
// Deprecated: Now it just returns the provided Model to keep the package's API compatible.
func MakeTable(m *Model) *Model {
	return m
}

// New creates a new hash.Hash16 computing the CRC-16 checksum
// using the polynomial represented by the Model.
func New(m *Model) hash.Hash16 {
	return &digest{Inst: m.New()}
}

// digest represents the partial evaluation of a checksum.
type digest struct {
	*crc16.Inst
}

func (d *digest) Size() int { return Size }

func (d *digest) BlockSize() int { return 1 }

func (d *digest) Sum16() uint16 {
	return d.Inst.Sum()
}

func (d *digest) Sum(in []byte) []byte {
	return d.Inst.AppendSum(in)
}

// Checksum returns the CRC-16 checksum of data
// using the polynomial, and parameters represented by the Model.
func Checksum(data []byte, m *Model) uint16 {
	return m.Checksum(data)
}
