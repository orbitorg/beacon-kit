// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.
package core

import (
	"fmt"

	"github.com/berachain/beacon-kit/mod/errors"
)

var (
	// ErrBlockSlotTooLow is returned when the block slot is too low.
	ErrBlockSlotTooLow = errors.New("block slot too low")

	// ErrBeaconStateOutOfSync is returned when the state is either too far
	// behind
	// or too far ahead of the head and we must abort the state transition.
	ErrBeaconStateOutOfSync = errors.New("state is out of sync with head")

	// ErrSlotMismatch is returned when the slot in a block header does not
	// match the expected value.
	ErrSlotMismatch = errors.New("slot mismatch")

	// ErrParentRootMismatch is returned when the parent root in an execution
	// payload does not match the expected value.
	ErrParentRootMismatch = errors.New("parent root mismatch")

	// ErrParentPayloadHashMismatch is returned when the parent hash of an
	// execution payload does not match the expected value.
	ErrParentPayloadHashMismatch = errors.New("payload parent hash mismatch")

	// ErrRandaoMixMismatch is returned when the randao mix in an execution
	// payload does not match the expected value.
	ErrRandaoMixMismatch = errors.New("randao mix mismatch")

	// ErrExceedsBlockDepositLimit is returned when the block exceeds the
	// deposit limit.
	ErrExceedsBlockDepositLimit = errors.New("block exceeds deposit limit")

	// ErrRewardsLengthMismatch is returned when the length of the rewards
	// in a block does not match the expected value.
	ErrRewardsLengthMismatch = errors.New("rewards length mismatch")

	// ErrPenaltiesLengthMismatch is returned when the length of the penalties
	// in a block does not match the expected value.
	ErrPenaltiesLengthMismatch = errors.New("penalties length mismatch")

	// ErrSlashedProposer is returned when a block is processed in which
	// the proposer is slashed.
	ErrSlashedProposer = errors.New(
		"attempted to process a block with a slashed proposer")

	// ErrStateRootMismatch is returned when the state root in a block header
	// does not match the expected value.
	ErrStateRootMismatch = errors.New("state root mismatch")

	// ErrInvalidSignature is returned when the signature is invalid.
	ErrInvalidSignature = errors.New("invalid signature")

	// ErrXorInvalid is returned when the XOR operation is invalid.
	ErrXorInvalid = errors.New("xor invalid")

	// ErrOutdatedSlashingCount is returned when the count
	// of total slashing is not up to date.
	//nolint: lll
	ErrOutdatedSlashingCount = errors.New("count of total slashing is not up to date")
)

// ErrTooManyWithdrawal is returned
// when the number of withdrawals exceeds the limit.
type TooManyWithdrawalError struct {
	Expected uint64
	Actual   uint64
}

func (e TooManyWithdrawalError) Error() string {
	return fmt.Sprintf("too many withdrawals, expected %d, got %d",
		e.Expected, e.Actual)
}

// ErrExceedsBlockBlobLimit is returned
// when the number of blobs in a block exceeds the limit.
type ExceedsBlockBlobLimitError struct {
	Expected uint64
	Actual   int
}

func (e ExceedsBlockBlobLimitError) Error() string {
	return fmt.Sprintf("block exceeds blob limit, expected: %d, got: %d",
		e.Expected, e.Actual)
}

// ErrMismatchWithdrawalCount is returned when the count
// of the withdrawals in a block does not match the expected value.
type MismatchWithdrawalCountError struct {
	Expected uint64
	Actual   uint64
}

func (e MismatchWithdrawalCountError) Error() string {
	return fmt.Sprintf("withdrawal count mismatch, expected %d, got %d",
		e.Expected, e.Actual)
}

// ErrMismatchWithdrawal is returned when
// the withdrawal is not as expected.
type MismatchWithdrawalError struct {
	Expected string
	Actual   string
	Index    int
}

func (e MismatchWithdrawalError) Error() string {
	return fmt.Sprintf("withdrawals do not match expected %s, got %s at index %d",
		e.Expected, e.Actual, e.Index)
}
