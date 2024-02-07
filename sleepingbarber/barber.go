package sleepingbarber

import (
	"fmt"
	"time"
)

type Barber struct {
	Human
	WakeMe chan *Client
}

func (b *Barber) ManageShop(shop *Shop) {
	for {
		select {
		case c := <-shop.WaitingRoom:
			fmt.Printf("%s cuts %s's hair\n", b.GetName(), c.GetName())
			time.Sleep(time.Millisecond * 6)
			fmt.Printf("%s finished %s\n", b.GetName(), c.GetName())

		default:
			fmt.Printf("%s want to sleep... zzZZZzzz\n", b.GetName())
			c := <-b.WakeMe
			fmt.Printf("%s waked %s\n", c.GetName(), b.GetName())
		}
	}
}

func NewBarber(name string) *Barber {
	b := &Barber{}
	b.name = name
	b.WakeMe = make(chan *Client)
	return b
}
