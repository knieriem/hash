// Package crc4 implements hash.Hash8 for 4-bit CRCs
//
// Deprecated: Instead of crc4, package crc8 should be used
// with polys from github.com/knieriem/crcutil/poly4.
package crc4

import (
	"github.com/knieriem/hash"
)

// The size of a CRC-4 checksum in bytes.
const Size = 1

// Predefined polynomials.
const (
	ITU = 0x3
)

// Table is a 16-word table representing the polynomial for efficient processing.
type Table struct {
	Data            [16]uint8
	HighNibbleFirst bool
}

// MakeTable returns the Table constructed from the specified polynomial.
func MakeTable(poly uint8) *Table {
	t := new(Table)
	t.HighNibbleFirst = true
	for i := 0; i < 16; i++ {
		crc := uint8(i)
		for j := 0; j < 4; j++ {
			if crc&8 == 8 {
				crc = (crc << 1) ^ poly
			} else {
				crc <<= 1
			}
		}
		t.Data[i] = crc & 0xF
	}
	return t
}

// MakeTableReversed returns the Table constructed from the reversed
// form of the specified polynomial.
func MakeTableReversed(poly uint8) *Table {
	poly = reverseNibbles[poly]
	t := new(Table)
	for i := 0; i < 16; i++ {
		crc := uint8(i)
		for j := 0; j < 4; j++ {
			if crc&1 == 1 {
				crc = (crc >> 1) ^ poly
			} else {
				crc >>= 1
			}
		}
		t.Data[i] = crc
	}
	return t
}

var reverseNibbles = [16]uint8{
	0, 0x8, 0x4, 0xC,
	2, 0xA, 6, 0xe,
	1, 9, 5, 0xD,
	3, 0xB, 7, 0xF,
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

func update(crc uint8, t *Table, p []byte) uint8 {
	tab := t.Data
	if t.HighNibbleFirst {
		for _, v := range p {
			crc = tab[crc^((v>>4)&0xF)]
			crc = tab[crc^(v&0xF)]
		}
		return crc
	}
	for _, v := range p {
		crc = tab[crc^(v&0xF)]
		crc = tab[crc^((v>>4)&0xF)]
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

// Checksum returns the CRC-4 checksum of data
// using the polynomial represented by the Table.
func Checksum(data []byte, tab *Table) (crc uint8) {
	crc = Update(0, tab, data)
	return crc
}
