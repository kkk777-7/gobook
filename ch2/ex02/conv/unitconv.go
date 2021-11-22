package conv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

type Feet float64
type Meter float64

func (f Feet) String() string  { return fmt.Sprintf("%g feet", f) }
func (m Meter) String() string { return fmt.Sprintf("%g meters", m) }

type Pound float64
type Kilogram float64

func (p Pound) String() string    { return fmt.Sprintf("%g pounds", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%g kilograms", k) }
