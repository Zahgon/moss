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

// MergeOperatorStringAppend implements a simple merger that appends
// strings.  It was originally built for testing and sample purposes.
type MergeOperatorStringAppend struct {
	Sep        string // The separator string between operands.
	m          sync.Mutex
	numFull    int
	numPartial int
}

// Name returns the name of this merge operator implemenation
func (mo *MergeOperatorStringAppend) Name() string { _ = "STUB: not implemented"; return "" }

// FullMerge performs the full merge of a string append operation
func (mo *MergeOperatorStringAppend) FullMerge(key, existingValue []byte,
	operands [][]byte) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// PartialMerge performs the partial merge of a string append operation
func (mo *MergeOperatorStringAppend) PartialMerge(key,
	leftOperand, rightOperand []byte) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
