package conv

import "fmt"

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k) + AbsoluteZeroC }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }

// FeetToMeter converts a Feet to Meter.
func FeetToMeter(f Feet) Meter { return Meter(f * 0.3048) }

// MeterToFeet converts a Meter to Feet
func MeterToFeet(m Meter) Feet { return Feet(m / 0.3048) }

// PoundToKilogram converts a Pound to Kilogram
func PoundToKilogram(p Pound) Kilogram { return Kilogram(p * 0.45359237) }

// KilogramToPound converts a Kilogram to Pound
func KilogramToPound(k Kilogram) Pound { return Pound(k / 0.45359237) }

func Temp(t float64) {
	f := Fahrenheit(t)
	c := Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, FToC(f), c, CToF(c))
}

func Length(l float64) {
	f := Feet(l)
	m := Meter(l)
	fmt.Printf("%s = %s, %s = %s\n",
		f, FeetToMeter(f), m, MeterToFeet(m))
}

func Weight(w float64) {
	p := Pound(w)
	k := Kilogram(w)
	fmt.Printf("%s = %s, %s = %s\n",
		p, PoundToKilogram(p), k, KilogramToPound(k))
}
