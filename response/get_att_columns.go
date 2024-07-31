package response

type GetAttColumns struct {
	Response
	Result GetAttColumnsResult `json:"result"`
}

type GetAttColumnsResult struct {
	Columns []GetAttColumnsColumns `json:"columns"`
}

type GetAttColumnsColumns struct {
	Id      int    `json:"id"`
	Type    int    `json:"type"`
	Name    string `json:"name"`
	Alias   string `json:"alias"`
	Status  int    `json:"status"`
	SubType int    `json:"sub_type"`
}
