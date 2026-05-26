//  Copyright 2016-Present Couchbase, Inc.
//
//  Use of this software is governed by the Business Source License included
//  in the file licenses/BSL-Couchbase.txt.  As of the Change Date specified
//  in that file, in accordance with the Business Source License, use of this
//  software will be governed by the Apache License, Version 2.0, included in
//  the file licenses/APL2.txt.

//go:build safe
// +build safe

package moss

// Uint64SliceToByteSlice gives access to []uint64 as []byte
func Uint64SliceToByteSlice(in []uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ByteSliceToUint64Slice gives access to []byte as []uint64
func ByteSliceToUint64Slice(in []byte) ([]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// --------------------------------------------------------------

func endian() string {
	_ = "STUB: not implemented"
	// Need unsafe package to tell endian'ess.
	return ""
}
