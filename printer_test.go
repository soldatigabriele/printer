package printer

import (
	"testing"
	"time"

	"github.com/elliotchance/orderedmap"
)

func TestPrint(t *testing.T) {
	m := orderedmap.NewOrderedMap()

	m.Set("Test", "that")
	m.Set("Number", 100)
	m.Set("Date", time.Now())
	e := Blueprint{
		Data:   m,
		Colour: Red,
		Title:  "my table",
	}
	err := Print(e)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
}

func Example() {
	m := orderedmap.NewOrderedMap()
	m.Set("Test", "that")
	m.Set("Number", 1000)
	e := Blueprint{
		Data:   m,
		Colour: Red,
		Title:  "my table",
	}
	Print(e)
	// Output:
	//[101;30m MY TABLE [0m[101;30m      [0m
	//[107;30m Test     [0m[107;30m that [0m
	//[47;30m Number   [0m[47;30m 1000 [0m
}
