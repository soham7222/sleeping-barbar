package sleepingbarber

import (
	"fmt"

	"time"
)

type Client struct {
	Human
}

func (c *Client) EnterShop(shop *Shop) {
	for i := 0; i < 3; i++ {
		if i > 0 {
			fmt.Printf(">> %s is back\n", c.GetName())
		}
		select {
		case shop.WaitingRoom <- c:
			fmt.Printf(":) %s found a seat\n", c.GetName())
			select {
			case shop.Barber.WakeMe <- c:
				fmt.Printf("waking up barber %s\n", c.GetName())
			default:
			}
			return

		default:
			fmt.Printf("... %s will be back later ...\n", c.GetName())
			time.Sleep(time.Millisecond * 100)
		}
	}

	fmt.Printf(":( %s didn't found a seat, get out...\n", c.GetName())
}

func NewClient(name string) *Client {
	c := new(Client)
	c.name = name
	return c
}
