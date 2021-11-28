package bag

import (
	"errors"
	"testing"
)

func TestEmpty(t *testing.T) {
	b := make(Bag)

	_, err := Get[string](b, "foo")
	if !errors.Is(err, errNotFound) {
		t.Fatalf("expected not found error, saw %v", err)
	}
}

func TestAdd(t *testing.T) {
	b := make(Bag)

	if !Add(b, "foo", "bar") {
		t.Fatalf("weird add failure")
	}

	foo, err := Get[string](b, "foo")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if foo != "bar" {
		t.Fatalf("unexpected value: %v", foo)
	}

	if Add(b, "foo", "again") {
		t.Fatalf("weird add success")
	}

	_, err = Get[int](b, "foo")
	if !errors.Is(err, errTypeError) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRef(t *testing.T) {
	b := make(Bag)

	name := "Jordan"
	if !Ref(b, "name", &name) {
		t.Fatal("ref failed")
	}

	_, err := Get[*string](b, "name")
	if !errors.Is(err, errTypeError) {
		t.Fatal("retrieving pointer for ref did not fail")
	}

	_, err = Get[int](b, "name")
	if !errors.Is(err, errTypeError) {
		t.Fatal("retrieving value of differing type for ref did not fail")
	}

	readName, err := Get[string](b, "name")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if readName != "Jordan" {
		t.Fatalf("unexpected value: %v", readName)
	}

	name = "Jordan Orelli"
	if readName != "Jordan" {
		t.Fatalf("unexpected value: %v", readName)
	}
	readName, err = Get[string](b, "name")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if readName != "Jordan Orelli" {
		t.Fatalf("unexpected value: %v", readName)
	}

	fn := func(s *string) {
		*s = "mute"
	}
	fn(&name)

	readName, err = Get[string](b, "name")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if readName != "mute" {
		t.Fatalf("unexpected value: %v", readName)
	}
}
