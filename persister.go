//  Copyright 2016-Present Couchbase, Inc.
//
//  Use of this software is governed by the Business Source License included
//  in the file licenses/BSL-Couchbase.txt.  As of the Change Date specified
//  in that file, in accordance with the Business Source License, use of this
//  software will be governed by the Apache License, Version 2.0, included in
//  the file licenses/APL2.txt.

package moss

// runPersister() implements the persister task.
func (m *collection) runPersister() { _ = "STUB: not implemented"; return }

// There's a concurrency scenario where imagine that
// persistence takes a long time.  Also, imagine that
// there are no more incoming batches (so, stackDirtyTop
// is empty).
//
// That allows the merger to complete a merging cycle (so,
// stackDirtyMid is non-empty with unpersisted data) and
// the merger is now just waiting for either more incoming
// batches or waiting to be awoken.
//
// So, we notify/awake the merger here so that it can feed
// stackDirtyMid down to the persister as stackDirtyBase.
// Merger is indeed asleep.

// ---------------------------------------------

// TODO: More advanced eviction of stackClean.
// TODO: Timer based eviction of stackClean?
// TODO: Randomized eviction?
// TODO: Merging of stackClean to 1 level?
// TODO: WaitForMerger() also considers stackClean?
// TODO: Track popular Get() keys?
// TODO: Track shadowing during merges for writes.
