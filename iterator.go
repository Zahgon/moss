//  Copyright 2016-Present Couchbase, Inc.
//
//  Use of this software is governed by the Business Source License included
//  in the file licenses/BSL-Couchbase.txt.  As of the Change Date specified
//  in that file, in accordance with the Business Source License, use of this
//  software will be governed by the Apache License, Version 2.0, included in
//  the file licenses/APL2.txt.

package moss

import (
	"io"
)

// DefaultNaiveSeekToMaxTries is the max number of attempts a forward
// iterator.SeekTo() will loop using simple Next()'s before giving up
// and starting a binary search for a given, forward seekToKey.
var DefaultNaiveSeekToMaxTries = 100

// An iterator tracks a min-heap "scan-line" of cursors through a
// segmentStack.  Iterator implements the sort.Interface and
// heap.Interface on its cursors.
type iterator struct {
	ss *segmentStack

	cursors []*cursor // The len(cursors) <= len(ss.a) (+1 when lowerLevelIter).

	startKeyInclusive []byte
	endKeyExclusive   []byte

	prefixLen int

	lowerLevelIter Iterator // May be nil.

	closer io.Closer

	iteratorOptions IteratorOptions
}

// A cursor rerpresents a logical entry position inside a segment in a
// segmentStack.  An ssIndex < 0 and pos < 0 mean that the op/k/v came
// from the lowerLevelIter.
type cursor struct {
	ssIndex int // Index into Iterator.ss.a.
	sc      SegmentCursor

	op uint64
	k  []byte
	v  []byte
}

// StartIterator returns a new iterator on the given segmentStack.
//
// On success, the returned Iterator will be positioned so that
// Iterator.Current() will either provide the first entry in the
// iteration range or ErrIteratorDone.
//
// A startKeyInclusive of nil means the logical "bottom-most" possible
// key and an endKeyExclusive of nil means the logical "top-most"
// possible key.
//
// StartIterator can optionally include deletion operations in the
// enumeration via the IteratorOptions.IncludeDeletions flag.
//
// StartIterator can skip lower segments, via the
// IteratorOptions.MinSegmentLevel parameter.  For example, to ignore
// the lowest, 0th segment, use MinSegmentLevel of 1.
func (ss *segmentStack) StartIterator(
	startKeyInclusive, endKeyExclusive []byte,
	iteratorOptions IteratorOptions) (Iterator, error) {
	_ = "STUB: not implemented"
	return *new(Iterator), nil
}

// startIterator() returns a new iterator on the given segmentStack.
//
// On success, the returned Iterator will be positioned so that
// Iterator.Current() will either provide the first entry in the
// iteration range or ErrIteratorDone.
//
// A startKeyInclusive of nil means the logical "bottom-most" possible
// key and an endKeyExclusive of nil means the logical "top-most"
// possible key.
//
// startIterator() can optionally include deletion operations in the
// enumeration via the IteratorOptions.IncludeDeletions flag.
//
// startIterator() can skip lower segments, via the
// IteratorOptions.MinSegmentLevel parameter.  For example, to ignore
// the lowest, 0th segment, use MinSegmentLevel of 1.
func (ss *segmentStack) startIterator(
	startKeyInclusive, endKeyExclusive []byte,
	iteratorOptions IteratorOptions) (*iterator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ----------------------------------------------
// Add cursors for our allowed segments.

// ----------------------------------------------
// Add cursor for the lower level, if wanted.

// ----------------------------------------------
// Heap-ify the cursors.

// Close must be invoked to release resources.
func (iter *iterator) Close() error { _ = "STUB: not implemented"; return nil }

func (iter *iterator) InitCloser(closer io.Closer) error { _ = "STUB: not implemented"; return nil }

// Next returns ErrIteratorDone if the iterator is done.
func (iter *iterator) Next() error { _ = "STUB: not implemented"; return nil }

func iteratorBytesEqual(a, b []byte) bool { _ = "STUB: not implemented"; return false }

// Optimization to compare right-hand-side of keys first.

func (iter *iterator) SeekTo(seekToKey []byte) error { _ = "STUB: not implemented"; return nil }

// Try a loop of naive Next()'s for several attempts.

// The seekToKey is before our current position, or we gave up on
// the naiveSeekTo(), so start a brand new iterator to replace our
// current iterator, bounded by the startKeyInclusive.
//

// Clone current iterator before overwriting it.

func naiveSeekTo(iter Iterator, seekToKey []byte, maxTries int) error {
	_ = "STUB: not implemented"
	return nil
}

// Current returns ErrIteratorDone if the iterator is done.
// Otherwise, Current() returns the current key and val, which should
// be treated as immutable or read-only.  The key and val bytes will
// remain available until the next call to Next() or Close().
func (iter *iterator) Current() ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// CurrentEx is a more advanced form of Current() that returns more
// metadata.  It is used when IteratorOptions.IncludeDeletions is
// true.  It returns ErrIteratorDone if the iterator is done.
// Otherwise, the current operation, key, val are returned.
func (iter *iterator) CurrentEx() (
	entryEx EntryEx, key, val []byte, err error) {
	_ = "STUB: not implemented"
	return *new(EntryEx), nil, nil, nil
}

func (iter *iterator) Len() int { _ = "STUB: not implemented"; return 0 }

func (iter *iterator) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (iter *iterator) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (iter *iterator) Push(x interface{}) {
	_ = "STUB: not implemented"
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	return
}

func (iter *iterator) Pop() interface{} { _ = "STUB: not implemented"; return nil }

// --------------------------------------------

// The optimize method tries to optimize an iterator.  For example,
// when there's only a single segment, then the heap can be avoided by
// using a simpler, faster iteratorSingle implementation.
func (iter *iterator) optimize() (Iterator, error) {
	_ = "STUB: not implemented"
	return *new(Iterator), nil
}

// Optimization to return lowerLevelIter directly.

// --------------------------------------------

// sharedPrefixLen returns the length of the prefix shared by a and b,
// which can might be 0 length.
func sharedPrefixLen(a, b []byte) int { _ = "STUB: not implemented"; return 0 }
