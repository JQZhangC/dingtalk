package response

type GetAttendanceList2 struct {
	Response
	// 游标位置
	Cursor int `json:"cursor"`

	// 是否有更多数据
	HasMore bool `json:"has_more"`

	RecordResult []map[string]interface{} `json:"recordresult"`
}

type GetAttendanceList struct {
	Response
	// 游标位置
	Cursor int `json:"cursor"`

	// 是否有更多数据
	HasMore bool `json:"has_more"`

	RecordResult []GetAttendanceRecordResult `json:"recordresult"`
}

type GetAttendanceRecordResult struct {
	SourceType     string `json:"sourceType"`     //ATM：考勤机打卡（指纹/人脸打卡)
	BaseCheckTime  int64  `json:"baseCheckTime"`  //计算迟到和早退，基准时间
	UserCheckTime  int64  `json:"userCheckTime"`  //实际打卡时间, 用户打卡时间的毫秒数
	ProcInstId     string `json:"procInstId"`     //关联的审批实例ID，当该字段非空时，表示打卡记录与请假、加班等审批有关
	ApproveId      string `json:"approveId"`      //关联的审批ID，当该字段非空时，表示打卡记录与请假、加班等审批有关
	LocationResult string `json:"locationResult"` //位置结果
	TimeResult     string `json:"timeResult"`     //打卡结果
	CheckType      string `json:"checkType"`      //考勤类型
	UserId         string `json:"userId"`         //打卡人的userId
	WorkDate       int64  `json:"workDate"`       //工作日
	RecordId       int64  `json:"recordId"`       //打卡记录ID
	PlanId         int64  `json:"planId"`         //排班ID
	GroupId        int64  `json:"groupId"`        //考勤组ID
	Id             int64  `json:"id"`             //唯一标识ID
}
