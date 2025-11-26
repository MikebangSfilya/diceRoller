package parser

type PersonRequest struct {
	Name string `json:"name"`
	Dext int8   `json:"dext"`
	Wits int8   `json:"wits"`
}

type PersonResult struct {
	Name string `json:"name"`
	Sum  int8   `json:"sum"`
}
