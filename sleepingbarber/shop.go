package sleepingbarber

type Shop struct {
	Barber      *Barber
	WaitingRoom chan *Client
}

func NewShop(barber *Barber, seats int) *Shop {
	shop := new(Shop)
	shop.Barber = barber
	shop.WaitingRoom = make(chan *Client, seats)
	return shop
}
