//  Copyright 2016-Present Couchbase, Inc.
//
//  Use of this software is governed by the Business Source License included
//  in the file licenses/BSL-Couchbase.txt.  As of the Change Date specified
//  in that file, in accordance with the Business Source License, use of this
//  software will be governed by the Apache License, Version 2.0, included in
//  the file licenses/APL2.txt.

package moss

func (s *Store) compactMaybe(higher Snapshot,
	persistOptions StorePersistOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// First compute size of the incoming batch of data.

// Try leveled compaction to same file.

// Else append data to the end of the same file.

// Swallow the error internally.

// Fetch old file size.

// Fetch new file size.

// calcPartialCompactionStart - returns an index into the segment
// locations from which point it would be better to merge segments
// into a larger one.
//
// Return Values:
//
//	 0 => Full compaction into a new file.
//	>0 => Partially Compact all the segments starting at this
//	      return index, and append this one segment at the end of
//	      the file while retaining all segments before this
//	      starting segment.
//
//	false => Append data to the end of the file.
//	true  => Perform compaction.
//
// Example:                                    ||
//
//	                                        ||
//	 Say CompactionLevelMaxSegments = 2     ||
//	                                        ||
//	    ||                  ||              ||
//	 || ||               || || ??           ||
//	 || ||          ==>  || || ??     ==>   ||
//	 || ||               || || ??           ||
//	 || ||               || || ??           ||
//	 || ||  || ||  ##    || || ??           ||
//	 || ||  || ||  ##    || || ??           ||
//	----------------->  ----------->      -----> (new file)
//	 level1 level0 new  level1 hits max!  final file
func calcPartialCompactionStart(slocs SegmentLocs, newDataSize uint64,
	options *StoreOptions) (compStartIdx int, doCompact bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// Idle compaction => attempt full compaction.

// No segments => append to end of same file.

// Assume we need to append to end of file.
// sizeSoFar represents the estimated size of the partially compacted
// future segment if we were to start compacting from current segment.

// Incoming batch is assumed to be appended in L0.

// File is too fragmented for partial compaction.

func (s *Store) compact(footer *Footer, partialCompactStart int,
	higher Snapshot, persistOptions StorePersistOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Segments to compact (all segs when full compaction).
// Segments not compacted (nil when full compaction).

// No incoming data to persist and 1 or fewer footer segments.

// Safe as footer ref count is held positive.
// No incoming data & 1 or fewer footer segments.
// no need to perform compaction.

// Include deletions for partialCompactions.

// Prefix restore the footer's partialCompactStart.

// Owns the frefCompact ref-count.

func (s *Store) mergeSegStacks(footer *Footer, splicePoint int,
	higher *segmentStack) (rv, rvBase *segmentStack) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fast child collection recreation, must not merge
// segments from prior incarnation.

func (right *Footer) spliceFooter(left *Footer, splicePoint int) { _ = "STUB: not implemented"; return }

// Fast child collection recreation, ok to drop store footer's
// segments from prior incarnation.

func (s *Store) writeSegments(newSS, base *segmentStack,
	frefCompact *FileRef, fileCompact File,
	includeDeletes bool, syncAfterBytes int) (compactFooter *Footer, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: IMPORTANT: See MB-29664 - merge-operators, child
// collections, and partial/leveled compaction does not work
// correctly.  You need to use full compaction if you're using
// merge-operators with child collections.  The fix will be to
// compute and provide the right childSegStackBase to the
// recursive writeSegments() calls.
//

type compactWriter struct {
	file      File
	kvsWriter *bufferedSectionWriter
	bufWriter *bufferedSectionWriter

	// Bytes after which file's Sync() is to be invoked provided
	// sync is enabled. If <= 0, Sync() is not invoked.
	syncAfterBytes int

	// Bytes written since the last Sync().
	bytesSinceSync int

	totOperationSet   uint64
	totOperationDel   uint64
	totOperationMerge uint64
	totKeyByte        uint64
	totValByte        uint64
}

func (cw *compactWriter) Mutate(operation uint64, key, val []byte) error {
	_ = "STUB: not implemented"
	return nil
}

/* len(kvsBuf) */
