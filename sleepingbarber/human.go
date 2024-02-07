package sleepingbarber

type Human struct {
	name string
}

func (h *Human) GetName() string {
	return h.name
}
