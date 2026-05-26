//  Copyright 2016-Present Couchbase, Inc.
//
//  Use of this software is governed by the Business Source License included
//  in the file licenses/BSL-Couchbase.txt.  As of the Change Date specified
//  in that file, in accordance with the Business Source License, use of this
//  software will be governed by the Apache License, Version 2.0, included in
//  the file licenses/APL2.txt.

package moss

import (
	"github.com/couchbase/ghistogram"
)

// Stats returns stats for this collection.
func (m *collection) Stats() (*CollectionStats, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *collection) Histograms() ghistogram.Histograms {
	_ = "STUB: not implemented"
	return *new(ghistogram.Histograms)
}

// statsSegmentsLOCKED retrieves stats related to segments.
func (m *collection) statsSegmentsLOCKED(rv *CollectionStats) { _ = "STUB: not implemented"; return }

// AtomicCopyTo copies stats from s to r (from source to result).
func (s *CollectionStats) AtomicCopyTo(r *CollectionStats) { _ = "STUB: not implemented"; return }
