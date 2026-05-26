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
	"sync"
)

// SnapshotWrapper implements the moss.Snapshot interface.
type SnapshotWrapper struct {
	m        sync.Mutex
	refCount uint64
	ss       Snapshot
	closer   io.Closer // Optional, may be nil.
}

// NewSnapshotWrapper creates a wrapper which provides ref-counting
// around a snapshot.  The snapshot (and an optional io.Closer) will
// be closed when the ref-count reaches zero.
func NewSnapshotWrapper(ss Snapshot, closer io.Closer) *SnapshotWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (w *SnapshotWrapper) addRef() *SnapshotWrapper { _ = "STUB: not implemented"; return nil }

func (w *SnapshotWrapper) decRef() (err error) { _ = "STUB: not implemented"; return nil }

// ChildCollectionNames returns an array of child collection name strings.
func (w *SnapshotWrapper) ChildCollectionNames() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ChildCollectionSnapshot returns a Snapshot on a given child
// collection by its name.
func (w *SnapshotWrapper) ChildCollectionSnapshot(childCollectionName string) (
	Snapshot, error) {
	_ = "STUB: not implemented"
	return *new(Snapshot), nil
}

// Close will decRef the underlying snapshot.
func (w *SnapshotWrapper) Close() (err error) {
	_ = "STUB: not implemented"

	// Get returns the key from the underlying snapshot.
	return nil
}

func (w *SnapshotWrapper) Get(key []byte, readOptions ReadOptions) (
	[]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StartIterator initiates a start iterator over the underlying snapshot.
func (w *SnapshotWrapper) StartIterator(
	startKeyInclusive, endKeyExclusive []byte,
	iteratorOptions IteratorOptions,
) (Iterator, error) {
	_ = "STUB: not implemented"
	return *new(Iterator), nil
}
