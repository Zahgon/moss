//  Copyright 2017-Present Couchbase, Inc.
//
//  Use of this software is governed by the Business Source License included
//  in the file licenses/BSL-Couchbase.txt.  As of the Change Date specified
//  in that file, in accordance with the Business Source License, use of this
//  software will be governed by the Apache License, Version 2.0, included in
//  the file licenses/APL2.txt.

package moss

type segmentKeysIndex struct {
	// Number of keys that can be indexed.
	numIndexableKeys int

	// Keys that have been added so far.
	numKeys int

	// Size in bytes of all the indexed keys.
	numKeyBytes int

	// In-memory byte array of keys.
	data []byte

	// Start offsets of keys in the data array.
	offsets []uint32

	// Number of skips over keys in the segment kvs to arrive at the
	// next adjacent key in the data array.
	hop int

	// Total number of keys in the source segment.
	srcKeyCount int
}

// newSegmentKeysIndex preallocates the data/offsets arrays
// based on a calculated hop.
func newSegmentKeysIndex(quota int, srcKeyCount int,
	keyAvgSize int) *segmentKeysIndex {
	_ = "STUB: not implemented"
	return nil
}

/* 4 for the offset */

// Adds a qualified entry to the index. Returns true if space
// still available, false otherwise.
func (s *segmentKeysIndex) add(keyIdx int, key []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// All keys that can be indexed already have been,
// return false indicating that there's no room for
// anymore.

// No room for any more keys.

// Key does not satisfy the hop condition.

// Fetches the range of offsets between which the key exists,
// if present at all. The returned leftPos and rightPos can
// directly be used as the left and right extreme cursors
// while binary searching over the source segment.
func (s *segmentKeysIndex) lookup(key []byte) (leftPos int, rightPos int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// The index either wasn't used or isn't of any use.

// If key smaller than the first key, return early.

// If key larger than last key, return early.

// Direct hit.
