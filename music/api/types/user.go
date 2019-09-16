package types

type CellLoginRequest struct {
	Phone         string `json:"phone"`
	Password      string `json:"password"`
	RememberLogin string `json:"rememberLogin"`
}

type IsRegisterRequest struct {
	Cellphone string `json:"cellphone"`
}
