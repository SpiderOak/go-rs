/* Go interface to init_rs_char, free_rs_char, encode_rs_char, and
 * decode_rs_char.
 *
 * Copyright 2016 SpiderOak, Inc.
 * May be used under the terms of the GNU Lesser General Public License
 * (LGPL)
 */
package rs

/*
#include "rs.h"
*/
import "C"

import (
    "errors"
    "math"
    "unsafe"
)

type rsEncoder struct {
    encoder unsafe.Pointer
    dataSize int
    paritySize int
}

func NewRS(symsize, gfpoly, fcr, prim, nroots, pad int) (*rsEncoder, error) {
    var newEncoder unsafe.Pointer

    newEncoder = C.init_rs_char(C.int(symsize), C.int(gfpoly), C.int(fcr), C.int(prim), C.int(nroots), C.int(pad))
    if newEncoder != nil {
        return &rsEncoder{
            newEncoder,
            int(math.Exp2(float64(symsize))) - pad - 1 - nroots,
            nroots,
        }, nil
    }
    return nil, errors.New("Could not create RS encoder")
}

func NewRSSimple(dataSize, paritySize int) (*rsEncoder, error) {
    pad := 255 - dataSize - paritySize
    return NewRS(8, 0x11d, 0, 2, paritySize, pad)
}

func (self rsEncoder) Encode(data []byte) []byte {
    parity := make([]byte, self.paritySize)
    C.encode_rs_char(self.encoder, (*C.uchar)(&data[0]), (*C.uchar)(&parity[0]))
    ret := make([]byte, self.dataSize + self.paritySize)
    copy(ret, data)
    copy(ret[self.dataSize:], parity)

    return ret
}

func (self rsEncoder) Decode(data []byte) ([]byte, error) {
    r := C.decode_rs_char(self.encoder, (*C.uchar)(&data[0]), nil, 0)
    if r == -1 {
        return nil, errors.New("Decoding failed")
    }
    return data[0:self.dataSize], nil
}

func (self rsEncoder) Close() {
    C.free_rs_char(self.encoder)
}
