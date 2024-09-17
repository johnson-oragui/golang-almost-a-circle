package models

import (
	"testing"
)


func TestNewBaseWithoutId(t *testing.T) {
	nbObjects = 0

	b1 := NewBase()

	if b1.Id != 1 {
		t.Logf("Expected Id to be 1, but got %d", b1.Id)
		t.Fail()
	}

	b2 := NewBase()

	if b2.Id != 2 {
		t.Errorf("Expected Id to be 2, but got %d", b2.Id)
	}
}


func TestNewBaseWithId(t *testing.T) {
	nbObjects = 0

	b3 := NewBase(5)

	if b3.Id != 5 {
		t.Logf("Expected Id to be 5, but got %d", b3.Id)
		t.Fail()
	}

	b4 := NewBase(10)

	if b4.Id != 10 {
		t.Errorf("Expected Id to be 10, but got %d", b4.Id)
	}
}

func TestNbObjectsWithoutIdValue(t *testing.T) {
	nbObjects = 0

	b5 := NewBase()

	if nbObjects != 1 {
		t.Logf("Expected nbObjects to be 1, but got %d", nbObjects)
		t.Logf("%d", b5)
		t.Fail()
	}

	b6 := NewBase()

	if nbObjects != 2 {
		t.Logf("Expected nbObjects to be 2, but got %d", nbObjects)
		t.Logf("%d", b6)
		t.Fail()
	}
}

func TestNbObjectsWithIdValue(t *testing.T) {
	nbObjects = 0

	b7 := NewBase(1)

	if nbObjects != 0 {
		t.Logf("Expected nbObjects to be 0, but got %d", nbObjects)
		t.Logf("%d", b7)
		t.Fail()
	}

	b8 := NewBase()

	if nbObjects != 1 {
		t.Logf("Expected nbObjects to be 1, but got %d", nbObjects)
		t.Logf("%d", b8)
		t.Fail()
	}
}

func TestNewBaseMixed(t *testing.T) {
	nbObjects = 0

	b9 := NewBase(9)

	if b9.Id != 9 {
		t.Logf("Expected Id to be 9, but got %d", b9.Id)
		t.Fail()
	}

	b10 := NewBase()

	if b10.Id != 1 {
		t.Errorf("Expected Id to be 1, but got %d", b10.Id)
	}

	b11 := NewBase(11)

	if b11.Id != 11 {
		t.Logf("Expected Id to be 11, but got %d", b11.Id)
		t.Fail()
	}

	b12 := NewBase()

	if b12.Id != 2 {
		t.Logf("Expected Id to be 2, but got %d", b12.Id)
		t.Fail()
	}
}
