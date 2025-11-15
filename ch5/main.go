package main

import (
	"fmt"
)

// Define constants for measurement units - go has no enums
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
type Ingredient string
type Measurement struct {
	Unit      Unit
	Magnitude Magnitude
}

// embedded types in a type.
type MetricWeight struct{ Measurement }
type ImperialWeight struct{ Measurement }
type MetricVolume struct{ Measurement }
type ImperialVolume struct{ Measurement }

// complex types
type IngredientMeasurement struct {
	Ingredient  Ingredient
	Measurement Measurement
}

type IngredientList []IngredientMeasurement
type Step struct {
	Description string
	Ingredient
}

func main() {
	fmt.Println("Chapter 5 - Types")
	mw := MetricWeight{Measurement{Unit: Gram, Magnitude: 500}}
	fmt.Printf("Metric Weight: %.2f %s\n", mw.Magnitude, mw.Unit)
	mv := MetricVolume{Measurement{Unit: Litre, Magnitude: 1.5}}
	fmt.Printf("Metric Volume: %.2f %s\n", mv.Magnitude, mv.Unit)
}
