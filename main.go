package main

import (
	"fmt"
	"time"

	barbershop "sleeping-barbar/sleepingbarber"
)

func main() {
	barber := barbershop.NewBarber("John")
	shop := barbershop.NewShop(barber, 4)

	go barber.ManageShop(shop)
	time.Sleep(time.Second * 2)

	clients := []string{"Alain"}
	clients = append(clients, "Bernard")
	clients = append(clients, "Claude")
	clients = append(clients, "Dan")
	clients = append(clients, "Eric")
	clients = append(clients, "Franck")
	clients = append(clients, "Gérard")
	clients = append(clients, "Henri")
	clients = append(clients, "Ignasse")

	// create clients
	for _, c := range clients {
		client := barbershop.NewClient(c)
		// let them enter the shop, but use "go"
		// to create other without waiting the end
		// of the "entrance"
		go client.EnterShop(shop)

	}

	// press a key to finish
	fmt.Scanln()

}
