package opt

import (
	"testing"
	"unicode/utf8"
)

func TestNone(t *testing.T) {
	s := None[string]()
	if s.ok {
		t.Error("should not be ok")
	}
	if _, ok := s.Open(); ok {
		t.Error("should not be ok")
	}

	s2 := NoneOf("poop")
	if s2.ok {
		t.Error("should not be ok")
	}
	if _, ok := s2.Open(); ok {
		t.Error("should not be ok")
	}
}

func TestSome(t *testing.T) {
	s := Some("poop")
	if !s.ok {
		t.Error("should be ok")
	}

	v, ok := s.Open()
	if !ok {
		t.Fatal("should be ok")
	}

	if v != "poop" {
		t.Error("should be poop")
	}
}

func TestBind(t *testing.T) {
	count := Bind(utf8.RuneCountInString)

	t.Run("bind some", func(t *testing.T) {
		m := count(Some("poop"))

		n, ok := m.Open()
		if !ok {
			t.Fatal("should be ok")
		}
		if n != 4 {
			t.Errorf("wanted 4 but got %d instead", n)
		}
	})

	t.Run("bind non", func(t *testing.T) {
		m := count(None[string]())

		n, ok := m.Open()
		if ok {
			t.Fatal("should not be ok")
		}
		if n != 0 {
			t.Errorf("wanted 0 but got %d instead", n)
		}
	})
}

func TestNew(t *testing.T) {
	count := Bind(utf8.RuneCountInString)

	getCity := func(name string) (string, bool) {
		if name == "Jordan" {
			return "Chicago", true
		}
		return "", false
	}

	if n, ok := count(New(getCity("Jordan"))).Open(); !ok {
		t.Errorf("should be ok")
	} else {
		if n != 7 {
			t.Errorf("what")
		}
	}

	if n, ok := count(New(getCity("poopface jones"))).Open(); ok {
		t.Errorf("should not be ok")
	} else {
		if n != 0 {
			t.Errorf("why do we have this value %d", n)
		}
	}
}
