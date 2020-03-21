# Printer

This package takes as input a structured Blueprint and prints it according to the chosen style.

## Usage

### Installation

To install the package just require it with go get:

```go
go get github.com/soldatigabriele/printer
```

### Example

```go
package main

import (
    "github.com/soldatigabriele/printer"
)

func main() {
    // Create the data that will populate the table
    data := make(map[string]interface{})
    data["Test"] = "that"
    data["Number"] = 100
    data["Date"] = time.Now()

    bp := printer.Blueprint{
        Data:   data,
        Colour: printer.Red, // Choose the color
        Title:  "my table", // Set a title
    }
    printer.Print(bp)
}
```

Data is a map of `interface{}`, so anything printable can be passed to the printer.

### Output

Available colours are `printer.Red`, `printer.Blue`, `printer.Green`.
![Output](/images/output.png)
