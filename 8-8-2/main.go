package main

import "fmt"

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	if t == u.T {
		return u.Value
	}

	switch u.T {
	case Inch:
		return u.Value * 2.54
	case CM:
		return u.Value / 2.54
	}

	return u.Value
}

type DemensionsInch struct {
	length Unit
	width  Unit
	height Unit
}

func (d DemensionsInch) Length() Unit {
	return d.length
}

func (d DemensionsInch) Width() Unit {
	return d.width
}

func (d DemensionsInch) Height() Unit {
	return d.height
}

type DemensionsCM struct {
	length Unit
	width  Unit
	height Unit
}

func (d DemensionsCM) Length() Unit {
	return d.length
}

func (d DemensionsCM) Width() Unit {
	return d.width
}

func (d DemensionsCM) Height() Unit {
	return d.height
}

func GetLen(d Dimensions) float64 {
	return d.Length().Get(CM)
}

func main() {
	demensionsCM := DemensionsCM{
		length: Unit{Value: 5, T: CM},
		width:  Unit{Value: 3, T: CM},
		height: Unit{Value: 1, T: CM},
	}

	fmt.Println(GetLen(demensionsCM))
}
