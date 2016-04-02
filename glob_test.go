package main

import (
	"regexp"
	"testing"
)

func TestGlob(t *testing.T) {
	for _, c := range []struct {
		pattern string
		target  string
		expect  bool
	}{
		{"", "", true},
		{"", "a", false},
		{"*", "a", true},
		{"*a", "a", true},
		{"*a", "aa", true},
		{"*a", "ab", false},
		{"a*", "a", true},
		{"a*", "aa", true},
		{"a*", "ab", true},
		{"a*", "ba", false},
		{"*a*", "a", true},
		{"*a*", "bab", true},
		{"*a*a*", "a", false},
		{"*a*a*", "aba", true},
		{"*a*a*", "babab", true},
		{"*a**a*", "babab", true},
		{"*a*A*", "babab", false},
		{"*a*A*", "babAb", true},
	} {
		m := newGlob(c.pattern)
		if m.Match(c.target) != c.expect {
			t.Errorf("Unexpected: %#v", c)
		}
	}
}

func BenchmarkLeadingGlob(b *testing.B) {
	m := newGlob("*bar")
	s := "foobar"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Match(s)
	}
}

func BenchmarkTrailingGlob(b *testing.B) {
	m := newGlob("foo*")
	s := "foobar"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Match(s)
	}
}

func BenchmarkAroundGlob(b *testing.B) {
	m := newGlob("*bar*")
	s := "foobarbaz"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Match(s)
	}
}

func BenchmarkLeadingRegexp(b *testing.B) {
	m := regexp.MustCompile(".*bar$")
	s := "foobar"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.MatchString(s)
	}
}

func BenchmarkTrailingRegexp(b *testing.B) {
	m := regexp.MustCompile("^foo.*")
	s := "foobar"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.MatchString(s)
	}
}

func BenchmarkAroundRegexp(b *testing.B) {
	m := regexp.MustCompile(".*bar.*")
	s := "foobarbaz"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.MatchString(s)
	}
}
