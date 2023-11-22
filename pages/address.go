package pages

import "webpage/components"

type Address struct {
	components.Page
}

func NewAddress() (*Address, error) {
	address := &Address{}

	label := components.NewLabel()
	label.Id.Set("id!!")
	label.Name.Set("name!!")

	address.Components = append(address.Components, components.NewLabel())
	address.Components = append(address.Components, components.NewInput())
}
