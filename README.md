This module provides implementations of Go's hash.Hash interface.
Since the code has been derived from Go's crc32.go initially by substituting some values only, this project is using the same license as Go; only the implementation of the 4-bit CRC differs slightly more from the original.

	crc4.ITU     0x3
	crc8.DOWCRC  0x8C   / 0x31    (e.g. Dallas/Maxim)
	crc8.CCITT   0xE0   / 0x07
	crc16.IBMCRC 0xA001 / 0x8005  (e.g. Modbus)
