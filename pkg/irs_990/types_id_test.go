// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package irs_990

import (
	"testing"

	"github.com/moov-io/1120x/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestFixedLengthIdsRejectPartialMatches(t *testing.T) {
	require.Error(t, EINType("1234567890").Validate())
	require.NoError(t, EINType("123456789").Validate())

	require.Error(t, SSNType("1234567890").Validate())
	require.NoError(t, SSNType("123456789").Validate())

	require.Error(t, EFINType("1234567").Validate())
	require.NoError(t, EFINType("123456").Validate())

	require.Error(t, ETINType("123456").Validate())
	require.NoError(t, ETINType("12345").Validate())

	require.Error(t, PINType("123456").Validate())
	require.NoError(t, PINType("12345").Validate())

	require.Error(t, PTINType("P123456789").Validate())
	require.NoError(t, PTINType("P12345678").Validate())

	require.Error(t, STINType("S123456789").Validate())
	require.NoError(t, STINType("S12345678").Validate())

	require.Error(t, RoutingTransitNumberType("0210000219").Validate())
	require.NoError(t, RoutingTransitNumberType("021000021").Validate())

	require.Error(t, SubmissionIdType("1234567890123abcdefgX").Validate())
	require.NoError(t, SubmissionIdType("1234567890123abcdefg").Validate())

	require.Error(t, utils.SubmissionIdType("1234567890123abcdefgX").Validate())
	require.NoError(t, utils.SubmissionIdType("1234567890123abcdefg").Validate())
}
