package types

type Package struct {
	Name string `json:"name"`
}

type ServiceResponse struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
}
