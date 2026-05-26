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

// Stats returns a map of stats.
func (s *Store) Stats() (map[string]interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// Histograms returns a snapshot of the histograms for this store.
func (s *Store) Histograms() ghistogram.Histograms {
	_ = "STUB: not implemented"
	return *new(ghistogram.Histograms)
}
