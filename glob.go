package main

import "strings"

type matcher interface {
	Match(string) bool
}

type globMatcher struct {
	parts         []string
	caseSensitive bool
	anyMatch      bool
	exactMatch    bool
	leadingGlob   bool
	trailingGlob  bool
}

func newGlob(pattern string) *globMatcher {
	m := new(globMatcher)
	if pattern == "*" {
		m.anyMatch = true
		return m
	}
	m.caseSensitive = pattern != strings.ToLower(pattern)
	if strings.HasPrefix(pattern, "*") {
		m.leadingGlob = true
		pattern = pattern[1:]
	}
	if strings.HasSuffix(pattern, "*") {
		m.trailingGlob = true
		pattern = pattern[:len(pattern)-1]
	}
	m.parts = strings.Split(pattern, "*")
	m.exactMatch = len(m.parts) == 1 && !m.leadingGlob && !m.trailingGlob
	return m
}

func (m *globMatcher) Match(s string) bool {
	if m.anyMatch {
		return true
	}
	if !m.caseSensitive {
		s = strings.ToLower(s)
	}
	if m.exactMatch {
		return s == m.parts[0]
	}

	parts := m.parts
	if !m.leadingGlob {
		if strings.HasPrefix(s, parts[0]) {
			parts = parts[1:]
		} else {
			return false
		}
	}
	if !m.trailingGlob {
		if strings.HasSuffix(s, parts[len(parts)-1]) {
			parts = parts[:len(parts)-1]
		} else {
			return false
		}
	}

	for _, part := range parts {
		n := strings.Index(s, part)
		if n < 0 {
			return false
		}
		s = s[n+len(part):]
	}
	return true
}
