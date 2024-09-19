package models

import (
	"testing"
)


func TestRectangleWithoutValues(t *testing.T) {
	nbObjects = 0
	rec1, err := NewRectangle()

	if rec1 != nil && err == nil {
		t.Logf("Expected rec1 to be nill, but got %v\n", rec1)
		t.Fail()
	}
}

func TestRectangleWithWidth(t *testing.T) {
	nbObjects = 0

	if rec2, err := NewRectangle(10); rec2 != nil && err == nil {
		t.Errorf("Expected rec2 to nill, but got %v\n", rec2)
	}
}

func TestRectangleWithIdthAndHeight(t *testing.T) {
	nbObjects = 0

	rec3, err := NewRectangle(10, 20)

	if err != nil {
		t.Errorf("Expected error to be nil, but got %v\n", err)
	}

	if rec3.X != 0 {
		t.Errorf("Expected rec3.X to be 0, but got %v\n", rec3.X)
	}
	if rec3.Y != 0 {
		t.Errorf("Expected rec3.Y to be 0, but got %v\n", rec3.Y)
	}
	if rec3.Id != 1 {
		t.Errorf("Expected rec3.Id to be 0, but got %v\n", rec3.Id)
	}
}

func TestRectangleWithNegativeX(t *testing.T) {
	nbObjects = 0

	rec4, err := NewRectangle(10, 20, -1)

	if rec4 != nil {
		t.Errorf("Expected rec4 to be nil, but got %v\n", rec4)
	}

	if err.Error() != "x must be >= 0" {
		t.Errorf("Expected err to be 'x must be >= 0', but got %v\n", err)
	}
}

func TestRectangleWithNegativeY(t *testing.T) {
	nbObjects = 0

	rec5, err := NewRectangle(10, 20, 1, -2)

	if rec5 != nil {
		t.Errorf("Expected rec5 to be nil, but got %v\n", rec5)
	}

	if err.Error() != "y must be >= 0" {
		t.Errorf("Expected err to be 'x must be >= 0', but got %v\n", err)
	}
}

func TestRectangleWithValues(t *testing.T) {
	nbObjects = 0
	rec6, _ := NewRectangle(10, 20, 1, 2, 10)
	rec7, _ := NewRectangle(102, 201, 5, 4)

	if rec6.Width != 10 {
		t.Errorf("Expected rec6.Width to be 10, but got %v\n", rec6.Width)
	}
	if rec6.Height != 20 {
		t.Errorf("Expected rec6.Height to be 20, but got %v\n", rec6.Height)
	}
	if rec6.X != 1 {
		t.Errorf("Expected rec6.X to be 1, but got %v\n", rec6.X)
	}
	if rec6.Y != 2 {
		t.Errorf("Expected rec6.Y to be 2, but got %v\n", rec6.Y)
	}
	if rec6.Id != 10 {
		t.Errorf("Expected rec6.Id to be 10, but got %v\n", rec6.Id)
	}


	if rec7.Width != 102 {
		t.Errorf("Expected rec7.Width to be 102, but got %v\n", rec7.Width)
	}
	if rec7.Height != 201 {
		t.Errorf("Expected rec7.Height to be 201, but got %v\n", rec7.Height)
	}
	if rec7.X != 5 {
		t.Errorf("Expected rec7.X to be 5, but got %v\n", rec7.X)
	}
	if rec7.Y != 4 {
		t.Errorf("Expected rec7.Y to be 4, but got %v\n", rec7.Y)
	}
	if rec7.Id != 1 {
		t.Errorf("Expected rec7.Id to be 1, but got %v\n", rec7.Id)
	}

}

func TestRectangleArea(t *testing.T) {
	rec8, _ := NewRectangle(3, 2)

	area := rec8.Area()

	if area != 6 {
		t.Errorf("Expected are to be 6, but got %d\n", area)
	}
}
