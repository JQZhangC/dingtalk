package request

import "time"

type GetUpdateData struct {
	UserId   string    `json:"userid"`    //用户的userid
	WorkDate time.Time `json:"work_date"` //查询日期
}
