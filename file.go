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
	"os"
	"sync"
)

// An InitCloser holds onto an io.Closer, and is used for chaining
// io.Closer's.  That is, we often want the closing of one resource to
// close related resources.
type InitCloser interface {
	InitCloser(io.Closer) error
}

// The File interface is implemented by os.File.  App specific
// implementations may add concurrency, caching, stats, fuzzing, etc.
type File interface {
	io.ReaderAt
	io.WriterAt
	io.Closer
	Stat() (os.FileInfo, error)
	Sync() error
	Truncate(size int64) error
}

// The OpenFile func signature is similar to os.OpenFile().
type OpenFile func(name string, flag int, perm os.FileMode) (File, error)

// FileRef provides a ref-counting wrapper around a File.
type FileRef struct {
	file File
	m    sync.Mutex // Protects the fields that follow.
	refs int

	beforeCloseCallbacks []func() // Optional callbacks invoked before final close.
	afterCloseCallbacks  []func() // Optional callbacks invoked after final close.
}

type ioResult struct {
	kind string // Kind of io attempted.
	want int    // Num bytes expected to be written or read.
	got  int    // Num bytes actually written or read.
	err  error
}

// --------------------------------------------------------

// OnBeforeClose registers event callback func's that are invoked before the
// file is closed.
func (r *FileRef) OnBeforeClose(cb func()) { _ = "STUB: not implemented"; return }

// OnAfterClose registers event callback func's that are invoked after the
// file is closed.
func (r *FileRef) OnAfterClose(cb func()) { _ = "STUB: not implemented"; return }

// AddRef increases the ref-count on the file ref.
func (r *FileRef) AddRef() File { _ = "STUB: not implemented"; return *new(File) }

// DecRef decreases the ref-count on the file ref, and closing the
// underlying file when the ref-count reaches zero.
func (r *FileRef) DecRef() (err error) { _ = "STUB: not implemented"; return nil }

// Close allows the FileRef to implement the io.Closer interface.  It actually
// just performs what should be the final DecRef() call which takes the
// reference count to 0.  Once 0, it allows the file to actually be closed.
func (r *FileRef) Close() error {
	_ = "STUB: not implemented"

	// FetchRefCount fetches the ref-count on the file ref.
	return nil
}

func (r *FileRef) FetchRefCount() int { _ = "STUB: not implemented"; return 0 }

// --------------------------------------------------------

// OsFile interface allows conversion from a File to an os.File.
type OsFile interface {
	OsFile() *os.File
}

// ToOsFile provides the underlying os.File for a File, if available.
func ToOsFile(f File) *os.File { _ = "STUB: not implemented"; return nil }

// --------------------------------------------------------

type bufferedSectionWriter struct {
	err error
	w   io.WriterAt
	beg int64 // Start position where we started writing in file.
	cur int64 // Current write-at position in file.
	max int64 // When > 0, max number of bytes we can write.
	buf []byte
	n   int

	stopCh chan struct{}
	doneCh chan struct{}
	reqCh  chan ioBuf
	resCh  chan ioBuf
}

type ioBuf struct {
	buf []byte
	pos int64
	err error
}

// newBufferedSectionWriter converts incoming Write() requests into
// buffered, asynchronous WriteAt()'s in a section of a file.
func newBufferedSectionWriter(w io.WriterAt, begPos, maxBytes int64,
	bufSize int, s statsReporter) *bufferedSectionWriter {
	_ = "STUB: not implemented"
	return nil
}

// Offset returns the byte offset into the file where the
// bufferedSectionWriter is currently logically positioned.
func (b *bufferedSectionWriter) Offset() int64 { _ = "STUB: not implemented"; return 0 }

// Written returns the logical number of bytes written to this
// bufferedSectionWriter; or, the sum of bytes to Write() calls.
func (b *bufferedSectionWriter) Written() int64 { _ = "STUB: not implemented"; return 0 }

func (b *bufferedSectionWriter) Write(p []byte) (nn int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Would go over b.max.

func (b *bufferedSectionWriter) Flush() error { _ = "STUB: not implemented"; return nil }

func (b *bufferedSectionWriter) Stop() error { _ = "STUB: not implemented"; return nil }
