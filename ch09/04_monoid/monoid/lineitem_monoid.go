package monoid

type LineItem struct {
	Quantity  int
	Price     int
	ListPrice int
}

type LineItemMonoid interface {
	Zero() []int
	Append(i ...int) LineItemMonoid
	Reduce() int
}

type lineItemContainer struct {
	LineItems []LineItem
}

func WrapLineItem(lineItems []LineItem) lineItemContainer {
	return lineItemContainer{LineItems: lineItems}
}

func (box lineItemContainer) Zero() []LineItem {
	return nil
}

func (box lineItemContainer) Append(lineItems ...LineItem) lineItemContainer {
	box.LineItems = append(box.LineItems, lineItems...)
	return box
}

func (box lineItemContainer) Reduce() LineItem {
	totalQuantity := 0
	totalPrice := 0
	totalListPrice := 0
	for _, item := range box.LineItems {
		totalQuantity += item.Quantity
		totalPrice += item.Price
		totalListPrice += item.ListPrice
	}
	return LineItem{
		Quantity:  totalQuantity,
		Price:     totalPrice,
		ListPrice: totalListPrice,
	}
}
