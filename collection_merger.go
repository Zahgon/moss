//  Copyright 2016-Present Couchbase, Inc.
//
//  Use of this software is governed by the Business Source License included
//  in the file licenses/BSL-Couchbase.txt.  As of the Change Date specified
//  in that file, in accordance with the Business Source License, use of this
//  software will be governed by the Apache License, Version 2.0, included in
//  the file licenses/APL2.txt.

package moss

// NotifyMerger sends a message (optionally synchronously) to the merger
// to run another cycle.  Providing a kind of "mergeAll" forces a full
// merge and can be useful for applications that are no longer
// performing mutations and that want to optimize for retrievals.
func (m *collection) NotifyMerger(kind string, synchronous bool) error {
	_ = "STUB: not implemented"
	return nil
}

// ------------------------------------------------------

// runMerger() implements the background merger task.
func (m *collection) runMerger() { _ = "STUB: not implemented"; return }

// Way to disable merger.

// ---------------------------------------------
// Notify ping'ers from the previous loop.

// ---------------------------------------------
// Wait for new stackDirtyTop entries and/or pings.

// ---------------------------------------------
// Atomically ingest stackDirtyTop into stackDirtyMid.

// m.stackDirtyMid takes 1 refs, and
// stackDirtyMid takes 1 refs.

// While waiting for persistence, might as well do
// a full merge to optimize reads.

// Awake writers waiting for space in stackDirtyTop.

// The collection level lock needs to be acquired.

// ---------------------------------------------
// Merge multiple stackDirtyMid layers.

// ---------------------------------------------
// Notify persister.

// ---------------------------------------------

// TODO: Concurrent merging of disjoint slices of stackDirtyMid
// instead of the current, single-threaded merger?
//
// TODO: A busy merger means no feeding of the persister?
//
// TODO: Delay merger until lots of deletion tombstones?
//
// TODO: The base layer is likely the largest, so instead of heap
// merging the base layer entries, treat the base layer with
// special case to binary search to find better start points?
//
// TODO: Dynamically calc'ed soft max dirty top height, for
// read-heavy (favor lower) versus write-heavy (favor higher)
// situations?

// ------------------------------------------------------

func (m *collection) idleMergerWaker() { _ = "STUB: not implemented"; return }

// -1 will disable idle compactions.

// Helps run idle merger IFF new data has come in.

// Merger is indeed asleep.
// Nap only while merger naps.
// New data.

// mergerWaitForWork() is a helper method that blocks until there's
// either pings or incoming segments (from ExecuteBatch()) of work for
// the merger.
func (m *collection) mergerWaitForWork(pings []ping) (
	stopped, mergeAll bool, pingsOut []ping) {
	_ = "STUB: not implemented"
	return false, false, nil
}

// NO-OP.

// ------------------------------------------------------

// mergerMain() is a helper method that performs the merging work on
// the stackDirtyMid and swaps the merged result into the collection.
func (m *collection) mergerMain(stackDirtyMid, stackDirtyBase *segmentStack,
	mergeAll bool) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

// Do this only for idle-compactions.

// Allow an empty stackDirtyMid to kick persistence.

// ------------------------------------------------------

// mergerNotifyPersister() is a helper method that notifies the
// optional persister goroutine that there's a dirty segment stack
// that needs persistence.
func (m *collection) mergerNotifyPersister() { _ = "STUB: not implemented"; return }

// NO-OP.
