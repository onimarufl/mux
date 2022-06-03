package testservice

type Grocery struct {
	Name     string `json: "name"`
	Quantity int    `json: "quantity"`
}

type Request struct {
	Name     string `json: "name"`
	Quantity int    `json: "quantity"`
}
