package weightconv

import "fmt"

type Pound float64
type Gram float64

func (p Pound) String() string { return fmt.Sprintf("%g [lb]", p) }
func (g Gram) String() string  { return fmt.Sprintf("%g [kg]", g/1000) }

func PToG(p Pound) Gram { return Gram((p / 2.2046) * 1000) }
func GToP(g Gram) Pound { return Pound((g / 1000) * 2.2046) }
