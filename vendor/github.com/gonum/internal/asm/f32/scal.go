// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package f32

func ScalUnitary(alpha float32, x []float32) {
	for i := range x {
		x[i] *= alpha
	}
}

func ScalUnitaryTo(dst []float32, alpha float32, x []float32) {
	for i, v := range x {
		dst[i] = alpha * v
	}
}

// incX must be positive.
func ScalInc(alpha float32, x []float32, n, incX uintptr) {
	var ix uintptr
	for i := 0; i < int(n); i++ {
		x[ix] *= alpha
		ix += incX
	}
}

// incDst and incX must be positive.
func ScalIncTo(dst []float32, incDst uintptr, alpha float32, x []float32, n, incX uintptr) {
	var idst, ix uintptr
	for i := 0; i < int(n); i++ {
		dst[idst] = alpha * x[ix]
		ix += incX
		idst += incDst
	}
}
