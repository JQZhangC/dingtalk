package response

import "time"

type GetUpdateData struct {
	Response
	Result AtCheckInfoForOpenVo `json:"result"`
}

type AtCheckInfoForOpenVo struct {
	WorkDate             string                        `json:"work_date"`              //查询日期
	Userid               string                        `json:"userid"`                 //用户userId
	CorpId               string                        `json:"corpId"`                 //企业corpId
	ApproveList          []AtApproveForOpenVo          `json:"approve_list"`           //审批单列表
	AttendanceResultList []AtAttendanceResultForOpenVo `json:"attendance_result_list"` //打卡结果
	CheckRecordList      []AtAttendanceRecordForOpenVo `json:"check_record_list"`      //打卡详情
	ClassSettingInfo     AtClassSettingInfoForOpenVo   `json:"class_setting_info"`     //当前排班对应的休息时间段
}

type AtApproveForOpenVo struct {
	DurationUnit string    `json:"duration_unit"` //审批单的单位
	Duration     string    `json:"duration"`      //时长
	SubType      string    `json:"sub_type"`      //子类型名称
	TagName      string    `json:"tag_name"`      //审批单类型名称;请假,出差,外出,加班
	ProcInstId   string    `json:"procInst_id"`   //审批单ID
	BeginTime    time.Time `json:"begin_time"`    //审批单开始时间
	BizType      int       `json:"biz_type"`      //审批单类型;1：加班;2：出差/外出;3：请假
	EndTime      time.Time `json:"end_time"`      //审批单结束时间
	GmtFinished  time.Time `json:"gmt_finished"`  //审批单审批完成时间
}

type AtAttendanceResultForOpenVo struct {
	RecordId       int64     `json:"record_id"`       //打卡流水ID
	SourceType     string    `json:"source_type"`     //打卡来源。 ATM：考勤机 BEACON：IBeacon DING_ATM：钉钉考勤机 USER：用户打卡 BOSS：老板改签 APPROVE：审批系统 SYSTEM：考勤系统 AUTO_CHECK：自动打卡
	PlanCheckTime  time.Time `json:"plan_check_time"` //标准打卡时间
	ClassId        int       `json:"class_id"`        //班次ID
	LocationMethod string    `json:"location_method"` //定位方法
	LocationResult string    `json:"location_result"` //定位结果 Normal：范围内 Outside：范围外 NotSigned：未打卡
	OutsideRemark  string    `json:"outside_remark"`  //外勤备注
	PlanId         int64     `json:"plan_id"`         //排班ID
	UserAddress    string    `json:"user_address"`    //用户打卡地址
	GroupId        int       `json:"group_id"`        //考勤组ID
	UserCheckTime  time.Time `json:"user_check_time"` //用户打卡时间
	ProInstId      string    `json:"procInst_id"`     //审批单ID
	CheckType      string    `json:"check_type"`      //打卡类型 OnDuty：上班 OffDuty：下班
	TimeResult     string    `json:"time_result"`     //打卡的时间结果 Normal：正常 Early：早退 Late：迟到 SeriousLate：严重迟到 Absenteeism：旷工迟到 NotSigned：未打卡
}

type AtAttendanceRecordForOpenVo struct {
	RecordId          int64     `json:"record_id"`           //打卡记录ID
	SourceType        string    `json:"source_type"`         //打卡来源。 ATM：考勤机 BEACON：IBeacon DING_ATM：钉钉考勤机 USER：用户打卡 BOSS：老板改签 APPROVE：审批系统 SYSTEM：考勤系统 AUTO_CHECK：自动打卡
	UserAccuracy      string    `json:"user_accuracy"`       //用户定位精度
	ValidMatched      bool      `json:"valid_matched"`       //是否匹配打卡结果的流水。 true：匹配 false：不匹配
	UserCheckTime     time.Time `json:"user_check_time"`     //用户打卡时间
	UserLongitude     string    `json:"user_longitude"`      //打卡经度
	UserSsid          string    `json:"user_ssid"`           //wifi名称
	BaseAccuracy      string    `json:"base_accuracy"`       //基本定位精度
	UserMacAddr       string    `json:"user_mac_addr"`       //MAC地址
	UserLatitude      string    `json:"user_latitude"`       //打卡纬度
	BaseAddress       string    `json:"base_address"`        //打卡基础地址
	InvalidRecordMsg  string    `json:"invalid_record_msg"`  //打卡无效的原因
	InvalidRecordType string    `json:"invalid_record_type"` //打卡无效的类型
}

type AtClassSettingInfoForOpenVo struct {
	RestTimeVoList []AtRestTimeVo `json:"rest_time_vo_list"` //班次内休息信息
}

type AtRestTimeVo struct {
	RestBeginTime int `json:"rest_begin_time"` //休息开始时间
	RestEndTime   int `json:"rest_end_time"`   //休息结束时间
}
