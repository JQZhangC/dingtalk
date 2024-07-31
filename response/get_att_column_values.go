package response

type GetAttColumnValues struct {
	Response
	Result GetAttColumnValuesResult `json:"result"`
}

type GetAttColumnValuesResult struct {
	ColumnVals []ColumnValForTopVo `json:"column_vals"`
}

type ColumnValForTopVo struct {
	ColumnVals []ColumnDayAndVal `json:"column_vals"` //列值数据
	ColumnVo   ColumnForTopVo    `json:"column_vo"`   //列信息
	FixedValue string            `json:"fixed_value"` //固定值
}

type ColumnDayAndVal struct {
	Date  string `json:"date"`  // 日期
	Value string `json:"value"` // 列值
}
type ColumnForTopVo struct {
	Id int `json:"id"` //报表列ID
}
