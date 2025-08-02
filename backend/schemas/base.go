package schemas

type Message struct {
	Status string `json:"messsage"`
}

type Error struct {
	Error string `json:"message"`
}