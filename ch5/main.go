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
	Description    string
	IngredientList IngredientList
}

type Recipe struct {
	Title       string
	Description string
	Yield       int
	Ingredients IngredientList
	Steps       []Step
}

func main() {
	fmt.Println("Chapter 5 - Types")
	ingredients := IngredientList{
		{Ingredient: "Flour", Measurement: Measurement{Unit: Gram, Magnitude: 200}},
		{Ingredient: "Sugar", Measurement: Measurement{Unit: Gram, Magnitude: 100}},
		{Ingredient: "Butter", Measurement: Measurement{Unit: Gram, Magnitude: 50}},
	}
	steps := []Step{
		{
			Description: "Mix dry ingredients",
			IngredientList: IngredientList{
				{Ingredient: "Flour", Measurement: Measurement{Unit: Gram, Magnitude: 200}},
				{Ingredient: "Sugar", Measurement: Measurement{Unit: Gram, Magnitude: 100}},
			},
		},
		{
			Description: "Add butter and mix",
			IngredientList: IngredientList{
				{Ingredient: "Butter", Measurement: Measurement{Unit: Gram, Magnitude: 50}},
			},
		},
	}
	recipe := Recipe{
		Title:       "Simple Cake",
		Description: "A simple cake recipe",
		Yield:       1,
		Ingredients: ingredients,
		Steps:       steps,
	}
	fmt.Printf("Recipe: %s\nDescription: %s\nYield: %d\n", recipe.Title, recipe.Description, recipe.Yield)
	fmt.Println("Ingredients:")
	for _, im := range recipe.Ingredients {
		fmt.Printf("- %s: %.2f %s\n", im.Ingredient, im.Measurement.Magnitude, im.Measurement.Unit)
	}
	fmt.Println("Steps:")
	for i, step := range recipe.Steps {
		fmt.Printf("%d. %s\n", i+1, step.Description)
		for _, im := range step.IngredientList {
			fmt.Printf("   - %s: %.2f %s\n", im.Ingredient, im.Measurement.Magnitude, im.Measurement.Unit)
		}
	}
}
