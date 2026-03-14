package types

type Request struct{
	contract string
}

func (r *Request) SetContract(contract string) {
	r.contract = contract
}

func (r *Request) GetContract() string {
	return r.contract
}