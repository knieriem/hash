This module provides implementations of Go's hash.Hash interface
for calculating 8- and 16-bit CRCs.

Initially, the crc16, and crc8 implementations were derived
from [an early version of Go's crc32.go][crc32.go] by substituting some values only;
now these implementations are thin wrappers around [crcutil]
to allow for more flexibility regarding polynomial representations,
and CRC models.

Implemented models:

	crc4.ITU     0x3
	crc8.DOWCRC  0x8C   / 0x31    (e.g. Dallas/Maxim)
	crc16.IBMCRC 0xA001 / 0x8005  (e.g. Modbus)

[crc32.go]: https://github.com/golang/go/blob/9feddd0bae188825e01771d182b84e47b159aa30/src/pkg/hash/crc32/crc32.go
[crcutil]: https://pkg.go.dev/github.com/knieriem/crcutil
