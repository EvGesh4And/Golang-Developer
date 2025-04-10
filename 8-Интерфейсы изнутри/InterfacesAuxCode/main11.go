package main

type BaseStorage interface {
	Close()
}

type UsersStorage struct{}

func (UsersStorage) Close() {}

type TicketsStorage struct{}

func (TicketsStorage) Close()      {}
func (TicketsStorage) GetTickets() {}

func main() {
	var s BaseStorage

	s = UsersStorage{}
	s = TicketsStorage{}
	_ = s
}
