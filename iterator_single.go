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

// An iteratorSingle implements the Iterator interface, and is an edge
// case optimization when there's only a single segment to iterate and
// there's no lower-level iterator.  In contrast to the main iterator
// implementation, iteratorSingle doesn't have any heap operations.
type iteratorSingle struct {
	s  *segment
	sc SegmentCursor

	op uint64
	k  []byte
	v  []byte

	closer io.Closer

	options *CollectionOptions

	iteratorOptions IteratorOptions
}

// Close must be invoked to release resources.
func (iter *iteratorSingle) Close() error { _ = "STUB: not implemented"; return nil }

func (iter *iteratorSingle) InitCloser(closer io.Closer) error {
	_ = "STUB: not implemented"
	return nil
}

// Next returns ErrIteratorDone if the iterator is done.
func (iter *iteratorSingle) Next() error { _ = "STUB: not implemented"; return nil }

// we DO want to return ErrIteratorDone here

func (iter *iteratorSingle) SeekTo(seekToKey []byte) error { _ = "STUB: not implemented"; return nil }

// Try a loop of naive Next()'s for several attempts.

// we DO want to return ErrIteratorDone here

// Current returns ErrIteratorDone if the iterator is done.
// Otherwise, Current() returns the current key and val, which should
// be treated as immutable or read-only.  The key and val bytes will
// remain available until the next call to Next() or Close().
func (iter *iteratorSingle) Current() ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// CurrentEx is a more advanced form of Current() that returns more
// metadata.  It returns ErrIteratorDone if the iterator is done.
// Otherwise, the current operation, key, val are returned.
func (iter *iteratorSingle) CurrentEx() (
	entryEx EntryEx, key, val []byte, err error) {
	_ = "STUB: not implemented"
	return *new(EntryEx), nil, nil, nil
}
