//  Copyright 2016-Present Couchbase, Inc.
//
//  Use of this software is governed by the Business Source License included
//  in the file licenses/BSL-Couchbase.txt.  As of the Change Date specified
//  in that file, in accordance with the Business Source License, use of this
//  software will be governed by the Apache License, Version 2.0, included in
//  the file licenses/APL2.txt.

//go:build gofuzz
// +build gofuzz

package moss

import (
	"sort"

	"github.com/mschoch/smat"
)

// TODO: Test pre-allocated batches and AllocSet/Del/Merge().

var smatDebug = false

var smatCompactionSync = true

var smatCompactionConcern = CompactionAllow

// ------------------------------------------------

func smatLog(prefix, format string, args ...interface{}) { _ = "STUB: not implemented"; return }

// fuzz test using state machine driven by byte stream.
func Fuzz(data []byte) int { _ = "STUB: not implemented"; return 0 }

type smatContext struct {
	tmpDir string

	coll       Collection // Initialized in setupFunc().
	collMirror mirrorColl

	mo MergeOperator

	curBatch    int
	curSnapshot int
	curIterator int
	curKey      int

	batches      []Batch
	batchMirrors []map[string]smatBatchOp // Mirrors the entries in batches.

	snapshots       []Snapshot
	snapshotMirrors []*mirrorColl

	iterators       []Iterator
	iteratorMirrors []*mirrorIter

	keys []string

	actions int
}

type smatBatchOp struct {
	op, v string
}

type mirrorColl struct { // Used to validate coll and snapshot entries.
	kvs  map[string]string
	keys []string // Will be nil unless this is a snapshot.
}

type mirrorIter struct { // Used to validate iterator entries.
	pos int
	ss  *mirrorColl
}

// ------------------------------------------------------------------

var actionMap = smat.ActionMap{
	smat.ActionID('.'): action("      +batch", delta(func(c *smatContext) { c.curBatch++ })),
	smat.ActionID(','): action("      -batch", delta(func(c *smatContext) { c.curBatch-- })),
	smat.ActionID('{'): action("      +snapshot", delta(func(c *smatContext) { c.curSnapshot++ })),
	smat.ActionID('}'): action("      -snapshot", delta(func(c *smatContext) { c.curSnapshot-- })),
	smat.ActionID('['): action("      +itr", delta(func(c *smatContext) { c.curIterator++ })),
	smat.ActionID(']'): action("      -itr", delta(func(c *smatContext) { c.curIterator-- })),
	smat.ActionID(':'): action("      +key", delta(func(c *smatContext) { c.curKey++ })),
	smat.ActionID(';'): action("      -key", delta(func(c *smatContext) { c.curKey-- })),
	smat.ActionID('s'): action("    set", opSetFunc),
	smat.ActionID('d'): action("    del", opDelFunc),
	smat.ActionID('m'): action("    merge", opMergeFunc),
	smat.ActionID('g'): action("    get", opGetFunc),
	smat.ActionID('B'): action("  batchCreate", batchCreateFunc),
	smat.ActionID('b'): action("  batchExecute", batchExecuteFunc),
	smat.ActionID('H'): action("  snapshotCreate", snapshotCreateFunc),
	smat.ActionID('h'): action("  snapshotClose", snapshotCloseFunc),
	smat.ActionID('I'): action("  itrCreate", iteratorCreateFunc),
	smat.ActionID('i'): action("  itrClose", iteratorCloseFunc),
	smat.ActionID('>'): action("  itrNext", iteratorNextFunc),
	smat.ActionID('K'): action("  keyRegister", keyRegisterFunc),
	smat.ActionID('k'): action("  keyUnregister", keyUnregisterFunc),
	smat.ActionID('$'): action("CLOSE-REOPEN", closeReopenFunc),
}

var runningPercentActions []smat.PercentAction

func init() {
	var ids []int
	for actionId := range actionMap {
		ids = append(ids, int(actionId))
	}
	sort.Ints(ids)

	pct := 100 / len(actionMap)
	for _, actionId := range ids {
		runningPercentActions = append(runningPercentActions,
			smat.PercentAction{Percent: pct, Action: smat.ActionID(actionId)})
	}

	actionMap[smat.ActionID('S')] = action("SETUP", setupFunc)
	actionMap[smat.ActionID('T')] = action("TEARDOWN", teardownFunc)
}

// We only have one state: running.
func running(next byte) smat.ActionID { _ = "STUB: not implemented"; return *new(smat.ActionID) }

// Creates an action func based on a callback, used for moving the curXxxx properties.
func delta(cb func(c *smatContext)) func(ctx smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"
	return nil
}

func action(name string, f func(ctx smat.Context) (smat.State, error)) func(ctx smat.Context) (smat.State, error) {
	_ = "STUB: not implemented"
	return nil
}

// ------------------------------------------------------------------

var prefix = "                          "

func setupFunc(ctx smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"
	return *new(smat.State), nil
}

func teardownFunc(ctx smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"
	return *new(smat.State), nil
}

// ------------------------------------------------------------------

func opSetFunc(ctx smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"
	return *new(smat.State), nil
}

func opDelFunc(ctx smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"
	return *new(smat.State), nil
}

func opMergeFunc(ctx smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"
	return *new(smat.State), nil
}

func opGetFunc(ctx smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"
	return *new(smat.State), nil
}

func batchCreateFunc(ctx smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"
	return *new(smat.State), nil
}

func batchExecuteFunc(ctx smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"
	return *new(smat.State), nil
}

func snapshotCreateFunc(ctx smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"
	return *new(smat.State), nil
}

func snapshotCloseFunc(ctx smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"
	return *new(smat.State), nil
}

// Close any child iterators.

func iteratorCreateFunc(ctx smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"
	return *new(smat.State), nil
}

func iteratorCloseFunc(ctx smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"
	return *new(smat.State), nil
}

func iteratorNextFunc(ctx smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"
	return *new(smat.State), nil
}

func keyRegisterFunc(ctx smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"
	return *new(smat.State), nil
}

func keyUnregisterFunc(ctx smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"
	return *new(smat.State), nil
}

// ------------------------------------------------------

func closeReopenFunc(ctx smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"
	return *new(smat.State), nil
}

// Wait until dirty ops are drained.

func closeParts(c *smatContext) (err error) { _ = "STUB: not implemented"; return nil }

// ------------------------------------------------------

func (c *smatContext) getCurKey() string { _ = "STUB: not implemented"; return "" }

func (c *smatContext) getCurBatch() (Batch, map[string]smatBatchOp, error) {
	_ = "STUB: not implemented"
	return *new(Batch), nil, nil
}

func (c *smatContext) getCurSnapshot() (Snapshot, *mirrorColl, error) {
	_ = "STUB: not implemented"
	return *new(Snapshot), nil, nil
}

// ------------------------------------------------------------------

func (mc *mirrorColl) snapshot() *mirrorColl { _ = "STUB: not implemented"; return nil }
