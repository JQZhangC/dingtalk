package response

type GetAttendanceVacationsRecords struct {
	Response
	Result GetAttendanceVacationsRecordsResult `json:"result"`
}

type GetAttendanceVacationsRecordsResult struct {
	//是否存在更多结果
	HasMore bool `json:"hasMore"`
	//假期消费记录列表
	LeaveRecords []LeaveRecords `json:"leaveRecords"`
}

type LeaveRecords struct {
	//员工userId
	UserId string `json:"userId"`
	//假期类型唯一标识
	LeaveCode string `json:"leaveCode"`
	//假期消费记录唯一标识
	RecordId string `json:"recordId"`
	//假期额度唯一标识
	QuotaId string `json:"quotaId"`
	//计算类型。
	//insert：新纪录
	//add：新增
	//delete：删除
	//update：更新
	//null（或者不返回该字段）：请假消耗
	CalType string `json:"calType"`
	//额度有效期开始时间或请假开始时间，毫秒级时间戳
	StartTime int64 `json:"startTime"`
	//额度有效期结束时间或请假结束时间，毫秒级时间戳
	EndTime int64 `json:"endTime"`
	//显示单位。day：天 hour：小时
	LeaveViewUnit string `json:"leaveViewUnit"`
	//原因
	LeaveReason string `json:"leaveReason"`
	//假期记录类型。
	//leave：请假
	//update：更新配额
	//modify_quota:初始化余额或者更新余额
	LeaveRecordType string `json:"leaveRecordType"`
	//请假状态。
	//init：请假申请中
	//success：请假并已通过
	//refuse：请假但被被拒
	//abort：请假撤销
	//revoke：请假已通过但是撤销了请假并已同意
	LeaveStatus string `json:"leaveStatus"`
	//以天计算的消费额度。
	RecordNumPerDay int64 `json:"recordNumPerDay"`
	//以小时计算的消费额度
	RecordNumPerHour int64 `json:"recordNumPerHour"`
	//记录创建时间，毫秒级时间戳
	GmtCreate int64 `json:"gmtCreate"`
	//记录修改时间，毫秒级时间戳
	GmtModified int64 `json:"gmtModified"`
	//记录的操作人Id
	OpUserId string `json:"opUserId"`
}
