// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elliptic

import (
	"testing"
)

func TestIssue52075(t *testing.T) {
	Gx, Gy := P256().Params().Gx, P256().Params().Gy
	scalar := make([]byte, 33)
	scalar[32] = 1
	x, y := P256().ScalarBaseMult(scalar)
	if x.Cmp(Gx) != 0 || y.Cmp(Gy) != 0 {
		t.Errorf("unexpected output (%v,%v)", x, y)
	}
	x, y = P256().ScalarMult(Gx, Gy, scalar)
	if x.Cmp(Gx) != 0 || y.Cmp(Gy) != 0 {
		t.Errorf("unexpected output (%v,%v)", x, y)
	}
}
