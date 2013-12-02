// DO NOT EDIT THIS FILE BY HAND. This file was generated using
// translate_fdb_options.go, part of the fdb-go repository, and a copy of the
// fdb.options file (installed as part of the FoundationDB client, typically
// found as /usr/include/foundationdb/fdb.options).

// To regenerate this file, from the top level of an fdb-go repository checkout,
// run:
// $ go run _util/translate_fdb_options.go < /usr/include/foundationdb/fdb.options > fdb/generated.go

package fdb

import (
	"bytes"
	"encoding/binary"
)

func int64ToBytes(i int64) ([]byte, error) {
	buf := new(bytes.Buffer)
	if e := binary.Write(buf, binary.LittleEndian, i); e != nil {
		return nil, e
	}
	return buf.Bytes(), nil
}

// Enables trace output to a file in a directory of the clients choosing
//
// Parameter: path to output directory (or NULL for current working directory)
func (o NetworkOptions) SetTraceEnable(param string) error {
	return o.setOpt(30, []byte(param))
}

// Set internal tuning or debugging knobs
//
// Parameter: knob_name=knob_value
func (o NetworkOptions) SetKnob(param string) error {
	return o.setOpt(40, []byte(param))
}

// Set the size of the client location cache. Raising this value can boost performance in very large databases where clients access data in a near-random pattern. Defaults to 100000.
//
// Parameter: Max location cache entries
func (o DatabaseOptions) SetLocationCacheSize(param int64) error {
	b, e := int64ToBytes(param)
	if e != nil {
		return e
	}
	return o.setOpt(10, b)
}

// Set the maximum number of watches allowed to be outstanding on a database connection. Increasing this number could result in increased resource usage. Reducing this number will not cancel any outstanding watches. Defaults to 10000 and cannot be larger than 1000000.
//
// Parameter: Max outstanding watches
func (o DatabaseOptions) SetMaxWatches(param int64) error {
	b, e := int64ToBytes(param)
	if e != nil {
		return e
	}
	return o.setOpt(20, b)
}

// Specify the machine ID that was passed to fdbserver processes running on the same machine as this client, for better location-aware load balancing.
//
// Parameter: Hexadecimal ID
func (o DatabaseOptions) SetMachineId(param string) error {
	return o.setOpt(21, []byte(param))
}

// Specify the datacenter ID that was passed to fdbserver processes running in the same datacenter as this client, for better location-aware load balancing.
//
// Parameter: Hexadecimal ID
func (o DatabaseOptions) SetDatacenterId(param string) error {
	return o.setOpt(22, []byte(param))
}

// The transaction, if not self-conflicting, may be committed a second time after commit succeeds, in the event of a fault
func (o TransactionOptions) SetCausalWriteRisky() error {
	return o.setOpt(10, nil)
}

// The read version will be committed, and usually will be the latest committed, but might not be the latest committed in the event of a fault or partition
func (o TransactionOptions) SetCausalReadRisky() error {
	return o.setOpt(20, nil)
}

func (o TransactionOptions) SetCausalReadDisable() error {
	return o.setOpt(21, nil)
}

// The next write performed on this transaction will not generate a write conflict range. As a result, other transactions which read the key(s) being modified by the next write will not conflict with this transaction. Care needs to be taken when using this option on a transaction that is shared between multiple threads. When setting this option, write conflict ranges will be disabled on the next write operation, regardless of what thread it is on.
func (o TransactionOptions) SetNextWriteNoWriteConflictRange() error {
	return o.setOpt(30, nil)
}

func (o TransactionOptions) SetCheckWritesEnable() error {
	return o.setOpt(50, nil)
}

// Reads performed by a transaction will not see any prior mutations that occured in that transaction, instead seeing the value which was in the database at the transaction's read version. This option may provide a small performance benefit for the client, but also disables a number of client-side optimizations which are beneficial for transactions which tend to read and write the same keys within a single transaction.
func (o TransactionOptions) SetReadYourWritesDisable() error {
	return o.setOpt(51, nil)
}

// Disables read-ahead caching for range reads. Under normal operation, a transaction will read extra rows from the database into cache if range reads are used to page through a series of data one row at a time (i.e. if a range read with a one row limit is followed by another one row range read starting immediately after the result of the first).
func (o TransactionOptions) SetReadAheadDisable() error {
	return o.setOpt(52, nil)
}

func (o TransactionOptions) SetDurabilityDatacenter() error {
	return o.setOpt(110, nil)
}

func (o TransactionOptions) SetDurabilityRisky() error {
	return o.setOpt(120, nil)
}

func (o TransactionOptions) SetDurabilityDevNullIsWebScale() error {
	return o.setOpt(130, nil)
}

// Specifies that this transaction should be treated as highest priority and that lower priority transactions should block behind this one. Use is discouraged outside of low-level tools
func (o TransactionOptions) SetPrioritySystemImmediate() error {
	return o.setOpt(200, nil)
}

// Specifies that this transaction should be treated as low priority and that default priority transactions should be processed first. Useful for doing batch work simultaneously with latency-sensitive work
func (o TransactionOptions) SetPriorityBatch() error {
	return o.setOpt(201, nil)
}

// This is a write-only transaction which sets the initial configuration
func (o TransactionOptions) SetInitializeNewDatabase() error {
	return o.setOpt(300, nil)
}

// Allows this transaction to read and modify system keys (those that start with the byte 0xFF)
func (o TransactionOptions) SetAccessSystemKeys() error {
	return o.setOpt(301, nil)
}

func (o TransactionOptions) SetDebugDump() error {
	return o.setOpt(400, nil)
}

// Set a timeout in milliseconds which, when elapsed, will cause the transaction automatically to be cancelled. Valid parameter values are ``[0, INT_MAX]``. If set to 0, will disable all timeouts. All pending and any future uses of the transaction will throw an exception. The transaction can be used again after it is reset.
//
// Parameter: value in milliseconds of timeout
func (o TransactionOptions) SetTimeout(param int64) error {
	b, e := int64ToBytes(param)
	if e != nil {
		return e
	}
	return o.setOpt(500, b)
}

// Set a maximum number of retries after which additional calls to onError will throw the most recently seen error code. Valid parameter values are ``[-1, INT_MAX]``. If set to -1, will disable the retry limit.
//
// Parameter: number of times to retry
func (o TransactionOptions) SetRetryLimit(param int64) error {
	b, e := int64ToBytes(param)
	if e != nil {
		return e
	}
	return o.setOpt(501, b)
}

type StreamingMode int
const (

    // Client intends to consume the entire range and would like it all
    // transferred as early as possible.
	StreamingModeWantAll StreamingMode = -1

    // The default. The client doesn't know how much of the range it is likely
    // to used and wants different performance concerns to be balanced. Only a
    // small portion of data is transferred to the client initially (in order to
    // minimize costs if the client doesn't read the entire range), and as the
    // caller iterates over more items in the range larger batches will be
    // transferred in order to minimize latency.
	StreamingModeIterator StreamingMode = 0

    // Infrequently used. The client has passed a specific row limit and wants
    // that many rows delivered in a single batch. Because of iterator operation
    // in client drivers make request batches transparent to the user, consider
    // ``WANT_ALL`` StreamingMode instead. A row limit must be specified if this
    // mode is used.
	StreamingModeExact StreamingMode = 1

    // Infrequently used. Transfer data in batches small enough to not be much
    // more expensive than reading individual rows, to minimize cost if
    // iteration stops early.
	StreamingModeSmall StreamingMode = 2

    // Infrequently used. Transfer data in batches sized in between small and
    // large.
	StreamingModeMedium StreamingMode = 3

    // Infrequently used. Transfer data in batches large enough to be, in a
    // high-concurrency environment, nearly as efficient as possible. If the
    // client stops iteration early, some disk and network bandwidth may be
    // wasted. The batch size may still be too small to allow a single client to
    // get high throughput from the database, so if that is what you need
    // consider the SERIAL StreamingMode.
	StreamingModeLarge StreamingMode = 4

    // Transfer data in batches large enough that an individual client can get
    // reasonable read bandwidth from the database. If the client stops
    // iteration early, considerable disk and network bandwidth may be wasted.
	StreamingModeSerial StreamingMode = 5
)

// Add performs an addition of little-endian integers. If the existing value in the database is not present or shorter than ``param``, it is first extended to the length of ``param`` with zero bytes.  If ``param`` is shorter than the existing value in the database, the existing value is truncated to match the length of ``param``. The integers to be added must be stored in a little-endian representation.  They can be signed in two's complement representation or unsigned. You can add to an integer at a known offset in the value by prepending the appropriate number of zero bytes to ``param`` and padding with zero bytes to match the length of the value. However, this offset technique requires that you know the addition will not cause the integer field within the value to overflow.
func (t Transaction) Add(key KeyConvertible, param []byte) {
	t.atomicOp(key.FDBKey(), param, 2)
}

// BitAnd performs a bitwise ``and`` operation.  If the existing value in the database is not present or shorter than ``param``, it is first extended to the length of ``param`` with zero bytes.  If ``param`` is shorter than the existing value in the database, the existing value is truncated to match the length of ``param``.
func (t Transaction) BitAnd(key KeyConvertible, param []byte) {
	t.atomicOp(key.FDBKey(), param, 6)
}

// BitOr performs a bitwise ``or`` operation.  If the existing value in the database is not present or shorter than ``param``, it is first extended to the length of ``param`` with zero bytes.  If ``param`` is shorter than the existing value in the database, the existing value is truncated to match the length of ``param``.
func (t Transaction) BitOr(key KeyConvertible, param []byte) {
	t.atomicOp(key.FDBKey(), param, 7)
}

// BitXor performs a bitwise ``xor`` operation.  If the existing value in the database is not present or shorter than ``param``, it is first extended to the length of ``param`` with zero bytes.  If ``param`` is shorter than the existing value in the database, the existing value is truncated to match the length of ``param``.
func (t Transaction) BitXor(key KeyConvertible, param []byte) {
	t.atomicOp(key.FDBKey(), param, 8)
}

type conflictRangeType int
const (

    // Used to add a read conflict range
	conflictRangeTypeRead conflictRangeType = 0

    // Used to add a write conflict range
	conflictRangeTypeWrite conflictRangeType = 1
)
