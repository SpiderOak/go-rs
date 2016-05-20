# Go-RS

This is a Go interface around the general byte-oriented Reed-Solomon
routines found in Phil Karn's [FEC library](http://www.ka9q.net/code/fec/).
It is licensed under the LGPL v2.1.

## API

`func NewRS(symsize, gfpoly, fcr, prim, nroots, pad int) (*rsEncoder, error)`

Creates and returns a new Reed-Solomon encoder with the given parameters.

  - `symsize` - The symbol size, in bits. Must be 8 or fewer.
  - `gfpoly` - The extended Galois field generator polynomial coefficients,
    with the 0th coefficient in the low order bit. The polynomial *must* be
    primitive; if not, the call will fail with an error.
  - `fcr` - In index form, the first consecutive root of the Reed-Solomon codeg
    generator polynomial.
  - `prim` - In index form, the primitive element in the Galois field used to
    generate the Reed-Solomon code generator polynomial.
  - `nroots` - The number of roots in the Reed-Solomon code generator
    polynomial.  This equals the number of parity symbols per code block.
  - `pad` - The number of leading symbols in the codeword that are implicitly
    padded to zero in a shortened code block.

The resulting Reed-Solomon code has parameters (N, K), where N = 2^symsize -
pad - 1 and K = N - nroots.

`func NewRSSimple(dataSize, paritySize int) (*rsEncoder, error)`

Creates and returns a new Reed-Solomon encoder with the given data and parity
size. The sum of data and parity sizes must be less than 255.  The other
parameters default to the following: `symsize: 8, gfpoly: 0x11d, fcr: 0, prim: 2`

`func (self rsEncoder) Encode(data []byte) []byte`

Encodes the data slice, which must have as many bytes as the data size (or K).
Returns a slice with data and parity (of size N).

`func (self rsEncoder) Decode(data []byte) ([]byte, error)`

Decodes the data slice, which contains data and parity (of size N).  Returns
the corrected data, or error.

`func (self rsEncoder) Close()`

Frees resources used by the encoder.  DO NOT USE THE OBJECT AFTER CALLING
`Close()` ON IT.
