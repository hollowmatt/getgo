package main

import (
	"fmt"
)

// Define constants for measurement systems and units - go has no enums
const (
	Metric   MeasurementSystem = "Metric"
	Imperial MeasurementSystem = "Imperial"
)

const (
	Gram     Unit = "gram"
	Kilogram Unit = "kilogram"
	Pound    Unit = "pound"
	Ounce    Unit = "ounce"
	Litre    Unit = "litre"
	Gallon   Unit = "gallon"
	Cup      Unit = "cup"
	Pint     Unit = "pint"
	Quart    Unit = "quart"
	Tsp      Unit = "teaspoon"
	Tbl      Unit = "tablespoon"
)

type MeasurementSystem string
type Unit string
type Magnitude float32

type Measurement struct {
	Unit      Unit
	Magnitude Magnitude
}

func main() {
	fmt.Println("Chapter 5 - Types")
	meas := Measurement{Unit: Kilogram, Magnitude: 2.5}
	fmt.Printf("Measurement: %.2f %s\n", meas.Magnitude, meas.Unit)
}
