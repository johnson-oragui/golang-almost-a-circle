package models

import "testing"


func TestRectangleWithoutValues(t *testing.T) {
	nbObjects = 0
	rec1 := NewRectangle()

	if rec1.Width != 0 {
		t.Logf("Expected rec1.Width to be 0, but got %d", rec1.Width)
		t.Fail()
	}

	if rec1.Height != 0 {
		t.Logf("Expected rec1.Height to be 0, but got %d", rec1.Height)
		t.Fail()
	}

	if rec1.X != 0 {
		t.Logf("Expected rec1.X to be 0, but got %d", rec1.X)
		t.Fail()
	}

	if rec1.Y != 0 {
		t.Logf("Expected rec1.Y to be 0, but got %d", rec1.Y)
		t.Fail()
	}

	if rec1.Id != 1 {
		t.Logf("Expected rec1.Id to be 0, but got %d", rec1.Id)
		t.Fail()
	}
}

func TestRectangleWithValues(t *testing.T) {
	nbObjects = 0
	rec2 := NewRectangle(10, 5)

	rec3 := NewRectangle(0, 0, 2, 3, 5)

	if rec2.Width != 10 {
		t.Errorf("Expected rec2.Width to be 10, but got %d", rec2.Width)
	}
	if rec2.Height != 5 {
		t.Errorf("Expected rec2.Height to be 5, but got %d", rec2.Height)
	}
	if rec2.Id != 1 {
		t.Errorf("Expected rec2.Id to be 5, but got %d", rec2.Id)
	}

	if rec3.Width != 0 {
		t.Errorf("Expected rec3.Width to be 0, but got %d", rec3.Width)
	}
	if rec3.Height != 0 {
		t.Errorf("Expected rec3.Height to be 0, but got %d", rec3.Height)
	}
	if rec3.X != 2 {
		t.Errorf("Expected rec3.Width to be 2, but got %d", rec3.X)
	}
	if rec3.Y != 3 {
		t.Errorf("Expected rec3.Y to be 3, but got %d", rec3.Y)
	}
	if rec3.Id != 5 {
		t.Errorf("Expected rec3.Id to be 5, but got %d", rec3.Id)
	}

}
