// Copyright 2020 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"


)

func TestError_NotFound(t *testing.T) {
	tests := []struct {
		err    error
		expVal bool
	}{
		{err: os.ErrNotExist, expVal: true},
		{err: os.ErrClosed, expVal: false},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expVal, IsNotFound(NewError(test.err)))
		})
	}
}
