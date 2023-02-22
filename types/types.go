package types

type ResolveColumns struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

type ResolveTables struct {
	Name    string           `json:"name"`
	Comment string           `json:"comment"`
	Columns []ResolveColumns `json:"columns"`
}

type ResolveResp struct {
	Data []ResolveTables `json:"data"`
}
