package pie

//go:generate pie cars.* carPointers.*
type cars []car
type carPointers []*car

type car struct {
	Name, Color string
}

func (c *car) Equals(c2 *car) bool { _ = "STUB: not implemented"; return false }

func (c *car) String() string { _ = "STUB: not implemented"; return "" }
