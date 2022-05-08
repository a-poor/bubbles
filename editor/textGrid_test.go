package editor

import (
	"fmt"
	"testing"
)

func TestNewTextGrid(t *testing.T) {
	_ = newTextGrid()
}

func TestTextGridFromString(t *testing.T) {
	d := "hello world"
	tg := textGridFromString(d)
	s := tg.String()
	if d != s {
		t.Errorf("Expected %q, got %q", d, s)
	}
}

func TestTextGrid_Length(t *testing.T) {
	cases := []struct {
		text   string
		expect int
	}{
		{
			text:   "hello world",
			expect: 1,
		},
		{
			text:   "",
			expect: 1,
		},
		{
			text: `hello
			...
			world`,
			expect: 3,
		},
	}
	for i, c := range cases {
		name := fmt.Sprintf("case-%d", i)
		t.Run(name, func(t *testing.T) {
			tg := textGridFromString(c.text)
			n := tg.length()
			if c.expect != n {
				t.Errorf("For case %q, expected %d, got %d", c.text, c.expect, n)
			}
		})
	}
}

func TestTextGrid_WidthAt(t *testing.T) {
	cases := []struct {
		text   string
		at     int
		expect int
	}{
		{
			text:   "hello world",
			at:     0,
			expect: 11,
		},
		{
			text:   "",
			at:     0,
			expect: 0,
		},
		{
			text: `hello
			...
			world`,
			at:     0,
			expect: 5,
		},
		{
			text: `hello

			world`,
			at:     1,
			expect: 0,
		},
		{
			text: `hello
			...
			world`,
			at:     2,
			expect: 8,
		},
	}
	for i, c := range cases {
		name := fmt.Sprintf("case-%d", i)
		t.Run(name, func(t *testing.T) {
			tg := textGridFromString(c.text)
			n := tg.widthAt(c.at)
			if c.expect != n {
				t.Errorf("For case %q, expected %d at %d, got %d", c.text, c.expect, c.at, n)
			}
		})
	}
}
