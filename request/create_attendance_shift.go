package request

import "encoding/json"

type CreateAttendanceShift struct {
	// 操作人的userid
	OpUserId string `json:"op_user_id" validate:"required"`

	// 考勤组相关信息
	AttendanceShift *AttendanceShift `json:"shift" validate:"required"`
}

type AttendanceShift struct {
	// 班次owner
	Owner string `json:"owner,omitempty"`
	// 班次组名
	ClassGroupName string `json:"class_group_name,omitempty"`
	// 企业的corpId
	CorpId string `json:"corp_id,omitempty"`
	// 班次名称
	Name string `json:"name" validate:"required"`
	// 班次id
	Id int64 `json:"id,omitempty"`
	// 卡段
	Sections []AttendanceShiftSection `json:"sections" validate:"required"`
	// 设置
	Setting *AttendanceShiftSetting `json:"setting"`
	// 高级排班绑定服务id
	ServiceId int64 `json:"service_id,omitempty"`
}

type AttendanceShiftSection struct {
	// 打卡信息，需要设定一对Time包含 OnDuty：上班 OffDuty：下班
	Times []AttendanceShiftTime `json:"times" validate:"required"`
}

type AttendanceShiftTime struct {
	// 打卡类型：OnDuty：上班 OffDuty：下班
	CheckType string `json:"check_type,omitempty"`
	// 是否跨天：0：不跨天 1：跨天
	Across int `json:"across"`
	// 允许的最晚打卡时间，单位分钟（-1表示不限制）。
	EndMin int `json:"end_min,omitempty"`
	// 打卡时间
	CheckTime string `json:"check_time,omitempty"`
	// 是否免打卡：false：需打卡 true：免打卡
	FreeCheck bool `json:"free_check,omitempty"`
	// 允许的最早提前打卡时间，分钟为单位
	BeginMin int `json:"begin_min,omitempty"`
}

type AttendanceShiftSetting struct {
	// 休息开始
	RestBeginTime *AttendanceShiftTime `json:"rest_begin_time,omitempty"`
	// 班次ID
	ClassId *int `json:"class_id,omitempty"`
	// 是否弹性 true：弹性 false：非弹性
	IsFlexible bool `json:"is_flexible,omitempty"`
	// 企业的corpId
	CorpId string `json:"corp_id,omitempty"`
	// 是否删除。 N：是 Y：否
	IsDeleted string `json:"is_deleted,omitempty"`
	// 休息结束
	RestEndTime *AttendanceShiftTime `json:"rest_end_time,omitempty"`
	// 严重早退/迟到的时长，单位分钟
	SeriousLateMinutes int `json:"serious_late_minutes,omitempty"`
	// 旷工早退/迟到的时长，单位分钟
	AbsenteeismLateMinutes int `json:"absenteeism_late_minutes,omitempty"`
	// 班次设置扩展字段, json 字符串
	Extras string `json:"extras,omitempty"`
	// 班次tags
	Tags string `json:"tags,omitempty"`
}

type CreateAttendanceShiftBuilder struct {
	shift *CreateAttendanceShift
}

func NewCreateAttendanceShift(opUserId, name string, sections []AttendanceShiftSection) *CreateAttendanceShiftBuilder {
	return &CreateAttendanceShiftBuilder{shift: &CreateAttendanceShift{
		OpUserId:        opUserId,
		AttendanceShift: &AttendanceShift{Name: name, Sections: sections, Setting: &AttendanceShiftSetting{}},
	}}
}

func (cas *CreateAttendanceShiftBuilder) SetOwner(owner string) *CreateAttendanceShiftBuilder {
	cas.shift.AttendanceShift.Owner = owner
	return cas
}

func (cas *CreateAttendanceShiftBuilder) SetClassGroupName(classGroupName string) *CreateAttendanceShiftBuilder {
	cas.shift.AttendanceShift.ClassGroupName = classGroupName
	return cas
}

func (cas *CreateAttendanceShiftBuilder) SetCorpId(corpId string) *CreateAttendanceShiftBuilder {
	cas.shift.AttendanceShift.CorpId = corpId
	return cas
}

func (cas *CreateAttendanceShiftBuilder) SetId(id int64) *CreateAttendanceShiftBuilder {
	cas.shift.AttendanceShift.Id = id
	return cas
}

func (cas *CreateAttendanceShiftBuilder) SetServiceId(serviceId int64) *CreateAttendanceShiftBuilder {
	cas.shift.AttendanceShift.ServiceId = serviceId
	return cas
}

func (cas *CreateAttendanceShiftBuilder) SetSetting(classId, seriousLateMinutes, absenteeismLateMinutes int, isFlexible bool, corpId, extras, tags, isDeleted string) *CreateAttendanceShiftBuilder {
	setting := cas.shift.AttendanceShift.Setting
	setting.ClassId = &classId
	setting.IsFlexible = isFlexible
	setting.CorpId = corpId
	setting.IsDeleted = isDeleted
	setting.SeriousLateMinutes = seriousLateMinutes
	setting.AbsenteeismLateMinutes = absenteeismLateMinutes
	setting.Extras = extras
	setting.Tags = tags
	return cas
}

func (cas *CreateAttendanceShiftBuilder) SetSettingRestBeginTime(checkType, checkTime string, across, endMin, beginMin int, freeCheck bool) *CreateAttendanceShiftBuilder {
	st := cas.shift.AttendanceShift.Setting
	st.RestBeginTime = &AttendanceShiftTime{
		CheckType: checkType,
		Across:    across,
		EndMin:    endMin,
		CheckTime: checkTime,
		FreeCheck: freeCheck,
		BeginMin:  beginMin,
	}
	return cas
}

func (cas *CreateAttendanceShiftBuilder) SetSettingRestEndTime(checkType, checkTime string, across, endMin, beginMin int, freeCheck bool) *CreateAttendanceShiftBuilder {
	st := cas.shift.AttendanceShift.Setting
	st.RestEndTime = &AttendanceShiftTime{
		CheckType: checkType,
		Across:    across,
		EndMin:    endMin,
		CheckTime: checkTime,
		FreeCheck: freeCheck,
		BeginMin:  beginMin,
	}
	return cas
}

func (cas *CreateAttendanceShiftBuilder) Build() *CreateAttendanceShift {
	return cas.shift
}

func (cas *CreateAttendanceShift) String() string {
	if str, err := json.Marshal(cas); err != nil {
		return ""
	} else {
		return string(str)
	}
}
