//  Copyright 2016-Present Couchbase, Inc.
//
//  Use of this software is governed by the Business Source License included
//  in the file licenses/BSL-Couchbase.txt.  As of the Change Date specified
//  in that file, in accordance with the Business Source License, use of this
//  software will be governed by the Apache License, Version 2.0, included in
//  the file licenses/APL2.txt.

package moss

// calcTargetTopLevel() heuristically computes a new top level that
// the segmentStack should be merged to.
func (ss *segmentStack) calcTargetTopLevel() int { _ = "STUB: not implemented"; return 0 }

// ------------------------------------------------------

// merge() returns a new segmentStack, merging all the segments that
// are at the given newTopLevel and higher.
func (ss *segmentStack) merge(mergeAll bool, base *segmentStack) (
	*segmentStack, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// If we have not been asked to merge all segments,
// then heuristically calc a newTopLevel.

// ----------------------------------------------------
// First, rough estimate the bytes neeeded.

// ----------------------------------------------------
// Next, use an iterator for the actual merge.

// ---------------------------------------------------
// Recursively merge all the child segmentStacks with the base
// stack, dropping any deleted collections present in base but not
// in me.

// The base segment stack carries a child collection
// which was subsequently recreated.

// The dirtyBase's old segmentStacks will be closed by
// the collection merger after successful merge.

func (ss *segmentStack) mergeInto(minSegmentLevel, maxSegmentHeight int,
	dest SegmentMutator, base *segmentStack, includeDeletions, optimizeTail bool,
	cancelCh chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// NO-OP.

// When only 1 cursor remains, copy the remains of the
// last segment more directly instead of Next()'ing
// through the iterator.

// TODO: the merge operator implementation is currently
// inefficient and not lazy enough right now.
