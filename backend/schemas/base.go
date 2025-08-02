package schemas

type Message struct {
	Status string `json:"status"`
}

type Error struct {
	Error string `json:"error"`
}