package request

type GetAttColumnValues struct {
	UserId       string `json:"userid"`         // 用户的userid
	ColumnIdList string `json:"column_id_list"` // 报表列ID，多值用英文逗号分隔，最大长度20
	FromDate     string `json:"from_date"`      // 开始日期，格式为yyyy-MM-dd
	ToDate       string `json:"to_date"`        // 结束时间
}
