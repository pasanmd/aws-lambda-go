// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertJsonsEqual asserts two JSON files are semantically equal
// (ignores white-space and attribute order)
func AssertJsonsEqual(t *testing.T, expectedJSON []byte, actualJSON []byte) {
	assert.JSONEq(t, string(expectedJSON), string(actualJSON))
}
