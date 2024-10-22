package request

type GetUpdateData struct {
	UserId   string `json:"userid"`    //用户的userid
	WorkDate string `json:"work_date"` //查询日期
}
