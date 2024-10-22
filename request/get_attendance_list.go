package request

type GetAttendanceList struct {
	WorkDateFrom string   `json:"workDateFrom"` //查询考勤打卡记录的起始工作日。格式为yyyy-MM-dd HH:mm:ss，HH:mm:ss可以使用00:00:00，将返回此日期从0点到24点的结果
	WorkDateTo   string   `json:"workDateTo"`   //查询考勤打卡记录的结束工作日。格式为“yyyy-MM-dd HH:mm:ss”，HH:mm:ss可以使用00:00:00，将返回此日期从0点到24点的结果。
	UserIdList   []string `json:"userIdList"`   //员工在企业内的userId列表，最大值50
	Offset       int64    `json:"offset"`       //表示获取考勤数据的起始点。第一次传0，如果还有多余数据，下次获取传的offset值为之前的offset+limit，0、1、2...依次递增
	Limit        int64    `json:"limit"`        //表示获取考勤数据的条数，最大值50
	IsI18n       bool     `json:"isI18n"`       //是否为海外企业使用
}
