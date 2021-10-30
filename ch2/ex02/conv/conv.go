package conv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k) - 273.15 }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FeetToMeter converts a Feet to Meter.
func FeetToMeter(c Feet) Meter { return Meter(c * 0.3048) }

// MeterToFeet converts a Meter to Feet
func MeterToFeet(m Meter) Feet { return Feet(m / 0.3048) }

// PoundToKilogram converts a Pound to Kilogram
func PoundToKilogram(p Pound) Kilogram { return Kilogram(p * 0.45359237) }

// KilogramToPound converts a Kilogram to Pound
func KilogramToPound(k Kilogram) Pound { return Pound(k / 0.45359237) }
