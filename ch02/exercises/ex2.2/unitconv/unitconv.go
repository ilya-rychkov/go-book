package unitconv

import "fmt"

type Foot float64
type Metre float64
type Pound float64
type Kilogram float64

func (f Foot) String() string     { return fmt.Sprintf("%g Фут", f) }
func (m Metre) String() string    { return fmt.Sprintf("%g Метров", m) }
func (p Pound) String() string    { return fmt.Sprintf("%g Фунтов", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%g Кг", k) }
