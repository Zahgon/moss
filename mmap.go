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

	"github.com/blevesearch/mmap-go"
)

// mmapRef provides a ref-counting wrapper around a mmap handle.
type mmapRef struct {
	fref *FileRef
	mm   mmap.MMap
	buf  []byte
	m    sync.Mutex // Protects the fields that follow.
	refs int
	ext  interface{} // Extra user/associated data.
}

func (r *mmapRef) AddRef() *mmapRef { _ = "STUB: not implemented"; return nil }

func (r *mmapRef) DecRef() error { _ = "STUB: not implemented"; return nil }

func (r *mmapRef) Close() error { _ = "STUB: not implemented"; return nil }

func (r *mmapRef) SetExt(v interface{}) { _ = "STUB: not implemented"; return }

func (r *mmapRef) GetExt() (v interface{}) { _ = "STUB: not implemented"; return nil }
