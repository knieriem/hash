// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crc8

import (
	"sync"
	"te/hash"
)

// The size of a CRC-8 checksum in bytes.
const Size = 1

// Predefined polynomials.
const (
	// CRC-8-Dallas/Maxim, reversed polynomial of 0x31
	DOWCRC = 0x8C
)

// Table is a 256-word table representing the polynomial for efficient processing.
type Table [256]uint8

var dowcrcTable *Table
var dowcrcOnce sync.Once

func dowcrcInit() {
	dowcrcTable = makeTable(DOWCRC)
}

// MakeTable returns the Table constructed from the specified polynomial.
func MakeTable(poly uint8) *Table {
	switch poly {
	case DOWCRC:
		dowcrcOnce.Do(dowcrcInit)
		return dowcrcTable
	}
	return makeTable(poly)
}

// makeTable returns the Table constructed from the specified polynomial.
func makeTable(poly uint8) *Table {
	t := new(Table)
	for i := 0; i < 256; i++ {
		crc := uint8(i)
		for j := 0; j < 8; j++ {
			if crc&1 == 1 {
				crc = (crc >> 1) ^ poly
			} else {
				crc >>= 1
			}
		}
		t[i] = crc
	}
	return t
}

// digest represents the partial evaluation of a checksum.
type digest struct {
	crc uint8
	tab *Table
}

// New creates a new hash.Hash8 computing the CRC-8 checksum
// using the polynomial represented by the Table.
func New(tab *Table) hash.Hash8 { return &digest{0, tab} }

func (d *digest) Size() int { return Size }

func (d *digest) BlockSize() int { return 1 }

func (d *digest) Reset() { d.crc = 0 }

func update(crc uint8, tab *Table, p []byte) uint8 {
	for _, v := range p {
		crc = tab[byte(crc)^v]
	}
	return crc
}

// Update returns the result of adding the bytes in p to the crc.
func Update(crc uint8, tab *Table, p []byte) uint8 {
	return update(crc, tab, p)
}

func (d *digest) Write(p []byte) (n int, err error) {
	d.crc = Update(d.crc, d.tab, p)
	return len(p), nil
}

func (d *digest) Sum8() uint8 {
	return d.crc
}

func (d *digest) Sum(in []byte) []byte {
	return append(in, d.Sum8())
}

// Checksum returns the CRC-8 checksum of data
// using the polynomial represented by the Table.
func Checksum(data []byte, tab *Table) (crc uint8) {
	crc = Update(0, tab, data)
	return crc
}
