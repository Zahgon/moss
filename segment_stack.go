//  Copyright 2016-Present Couchbase, Inc.
//
//  Use of this software is governed by the Business Source License included
//  in the file licenses/BSL-Couchbase.txt.  As of the Change Date specified
//  in that file, in accordance with the Business Source License, use of this
//  software will be governed by the Apache License, Version 2.0, included in
//  the file licenses/APL2.txt.

package moss

import (
	"sync"
)

// A segmentStack is a stack of segments, where higher (later) entries
// in the stack have higher precedence, and should "shadow" any
// entries of the same key from lower in the stack.  A segmentStack
// implements the Snapshot interface.
type segmentStack struct {
	options *CollectionOptions
	stats   *CollectionStats

	a []Segment

	m sync.Mutex // Protects the fields the follow.

	refs int

	lowerLevelSnapshot *SnapshotWrapper

	// incarNum represents this segmentStack's unique incarnation number assigned
	// when the child collection was created. 0 for top-level collection.
	incarNum uint64

	// childSegStacks recursively store child collection segmentStacks.
	childSegStacks map[string]*segmentStack
}

func (ss *segmentStack) addRef() { _ = "STUB: not implemented"; return }

func (ss *segmentStack) decRef() { _ = "STUB: not implemented"; return }

// Only update stats if snapshot is on collection.

// ------------------------------------------------------

// Close releases associated resources.
func (ss *segmentStack) Close() error { _ = "STUB: not implemented"; return nil }

// ------------------------------------------------------

// Get retrieves a val from a segmentStack.
func (ss *segmentStack) Get(key []byte, readOptions ReadOptions) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get() retrieves a val from a segmentStack, but only considers
// segments at or below the segStart level.  The optional base
// segmentStack, when non-nil, is used instead of the
// lowerLevelSnapshot, as a form of controllable chaining.
func (ss *segmentStack) get(key []byte, segStart int, base *segmentStack,
	readOptions ReadOptions) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: else add a special return error indicating cache-miss!

// ------------------------------------------------------

// getMerged() retrieves a lower level val for a given key and returns
// a merged val, based on the configured merge operator.
func (ss *segmentStack) getMerged(key, val []byte, segStart int,
	base *segmentStack, readOptions ReadOptions) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ------------------------------------------------------

func (ss *segmentStack) ensureSorted(minSeg, maxSeg int) { _ = "STUB: not implemented"; return }

// Two phases allows for more concurrent sorting.

// ------------------------------------------------------

// SegmentStackStats represents the stats for a segmentStack.
type SegmentStackStats struct {
	CurOps      uint64
	CurBytes    uint64 // Counts key-val bytes only, not metadata.
	CurSegments uint64
}

// AddTo adds the values from this SegmentStackStats to the dest
// SegmentStackStats.
func (sss *SegmentStackStats) AddTo(dest *SegmentStackStats) { _ = "STUB: not implemented"; return }

// Stats returns the stats for this segment stack.
func (ss *segmentStack) Stats() *SegmentStackStats { _ = "STUB: not implemented"; return nil }

// ChildCollectionNames returns an array of child collection name strings.
func (ss *segmentStack) ChildCollectionNames() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ChildCollectionSnapshot returns a Snapshot on a given child
// collection by its name.
func (ss *segmentStack) ChildCollectionSnapshot(childCollectionName string) (
	Snapshot, error) {
	_ = "STUB: not implemented"
	return *new(Snapshot), nil
}

// ensureFullySorted recursively ensures that all child segmentStacks
// are sorted from 0 to end.
func (ss *segmentStack) ensureFullySorted() { _ = "STUB: not implemented"; return }

func (ss *segmentStack) isEmpty() bool { _ = "STUB: not implemented"; return false }
