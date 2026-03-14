package types

type Role struct{
	Contracts []string `json:"contracts"`
}

type Policy struct{
	Roles map[string]Role `json:"roles"`
}


