package main

import "testing"

func TestUnpack(t *testing.T) {
	check := func(t testing.TB, got, want string) {
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}
		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}

	}
	t.Run("digits only", func(t *testing.T) {
		got, err := unpack("4532")
		want := ""
		check(t, got, want)
		assertError(t, err, "incorrect string")

	})
	t.Run("letters only", func(t *testing.T) {
		got, err := unpack("abschtwejt")
		if err != nil {
			t.Error("letters only testcase has an error")
		}
		want := "abschtwejt"
		check(t, got, want)
	})
	t.Run("empty string", func(t *testing.T) {
		got, err := unpack("")
		if err != nil {
			t.Error("empty string testcase has an error")
		}
		want := ""
		check(t, got, want)
	})
	t.Run("incorrect string: 'sdfhfs4g3h46'", func(t *testing.T) {
		got, err := unpack("sdfhfs4g3h46")
		want := ""
		check(t, got, want)
		assertError(t, err, "incorrect string")
	})
	t.Run("example 1: a4bc2d5e", func(t *testing.T) {
		got, err := unpack("a4bc2d5e")
		if err != nil {
			t.Error("example 1 testcase has an error")
		}
		want := "aaaabccddddde"
		check(t, got, want)
	})
	t.Run("example 2: abcd", func(t *testing.T) {
		got, err := unpack("abcd")
		if err != nil {
			t.Error("example 2 testcase has an error")
		}
		want := "abcd"
		check(t, got, want)
	})
	t.Run("example 3: 45", func(t *testing.T) {
		got, err := unpack("45")
		want := ""
		check(t, got, want)
		assertError(t, err, "incorrect string")
	})
}
