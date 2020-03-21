package printer

import (
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	m := make(map[string]interface{})
	m["Test"] = "that"
	m["Number"] = 100
	m["Date"] = time.Now()
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
	m := make(map[string]interface{})
	m["Test"] = "that"
	e := Blueprint{
		Data:   m,
		Colour: Red,
		Title:  "my table",
	}
	Print(e)
	// Output:
	//[101;30m MY TABLE [0m[101;30m      [0m
	//[107;30m Test     [0m[107;30m that [0m
}
