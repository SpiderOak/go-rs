/* Library tests
 *
 * Copyright 2016 SpiderOak, Inc.
 * May be used under the terms of the GNU Lesser General Public License
 * (LGPL)
 */
package rs

import (
    "bytes"
    "testing"
)

var refData []byte = []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 48, 49, 30, 234, 31, 196, 33, 191, 177}

func TestRS(t *testing.T) {
    c, err := NewRS(8, 0x11d, 0, 2, 7, 236)
    if err != nil {
        t.Fatal(err)
    }
    t.Log(c)

    data := []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 48, 49}
    codedData := c.Encode(data)
    if bytes.Compare(codedData, refData) != 0 {
        t.Fatal("Encoded data is wrong")
    }
    t.Log(codedData)

    decodedData, err := c.Decode(codedData)
    if err != nil {
        t.Fatal(err)
    }
    if bytes.Compare(decodedData, data) != 0 {
        t.Fatal("Decoded data is wrong")
    }
    t.Log(decodedData)
}

func TestRSSimple(t *testing.T) {
    data := []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 48, 49}

    c, err := NewRSSimple(12, 7)
    if err != nil {
        t.Fatal(err)
    }

    codedData := c.Encode(data)
    if bytes.Compare(codedData, refData) != 0 {
        t.Fatal("Encoded data is wrong")
    }
    t.Log(codedData)

    decodedData, err := c.Decode(codedData)
    if err != nil {
        t.Fatal(err)
    }
    if bytes.Compare(decodedData, data) != 0 {
        t.Fatal("Decoded data is wrong")
    }
    t.Log(decodedData)
}
