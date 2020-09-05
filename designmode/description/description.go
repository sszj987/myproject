package description

const (
	coffeeCost = 3
	milkCost = 2
)

type drinkItf interface {
	 getCost() int
}

type drink struct {
	cost int
	drinkItf
}

func (d *drink) getCost() int {
	return d.cost
}

type coffee struct {
	itf drinkItf
}

func (c *coffee) getCost() int {
	return c.itf.getCost() + coffeeCost
}

type milk struct {
	itf drinkItf
}

func (m *milk) getCost() int {
	return m.itf.getCost() + milkCost
}