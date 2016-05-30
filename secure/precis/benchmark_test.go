// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.7

package precis

import "testing"

var benchData = []struct{ name, str string }{
	{"ASCII", "Malvolio"},
	{"Arabic", "دبي"},
	{"Hangul", "동일조건변경허락"},
}

func doBench(b *testing.B, f func(b *testing.B, p *Profile, s string)) {
	for _, tc := range testCases {
		for _, d := range benchData {
			b.Run(tc.name+"/"+d.name, func(b *testing.B) {
				f(b, tc.p, d.str)
			})
		}
	}
}

func BenchmarkString(b *testing.B) {
	doBench(b, func(b *testing.B, p *Profile, s string) {
		for i := 0; i < b.N; i++ {
			p.String(s)
		}
	})
}

func BenchmarkBytes(b *testing.B) {
	doBench(b, func(b *testing.B, p *Profile, s string) {
		src := []byte(s)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			p.Bytes(src)
		}
	})
}

func BenchmarkTransform(b *testing.B) {
	doBench(b, func(b *testing.B, p *Profile, s string) {
		src := []byte(s)
		dst := make([]byte, 2*len(s))
		t := p.NewTransformer()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			t.Transform(dst, src, true)
		}
	})
}