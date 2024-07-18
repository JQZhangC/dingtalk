/*
 * Copyright 2020 zhaoyunxing.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dingtalk

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/zhaoyunxing92/dingtalk/v2/constant/attendance"
	"github.com/zhaoyunxing92/dingtalk/v2/request"

	"github.com/stretchr/testify/assert"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

func TestDingTalk_GetAttendanceGroups(t *testing.T) {
	groups, err := client.GetAttendanceGroups(0, 10)

	assert.Nil(t, err)
	assert.NotNil(t, groups)
}

func TestGetAttendanceGroups_Json(t *testing.T) {
	str := `{
    "errcode": 0,
    "result": {
        "groups": [
            {
                "classes_list": [
                    "周一、二、三、四、五 A 09:30-18:30 ",
                    "周六、日 休息"
                ],
                "default_class_id": 686275261,
                "disable_check_when_rest": false,
                "disable_check_without_schedule": false,
                "enable_emp_select_class": true,
                "group_id": 792440025,
                "group_name": "技术考勤",
                "is_default": false,
                "manager_list": [
                    "manager164"
                ],
                "member_count": 1,
                "owner_user_id": "manager164",
                "selected_class": [
                    {
                        "class_id": 686275261,
                        "class_name": "A",
                        "sections": [
                            {
                                "times": [
                                    {
                                        "across": 0,
                                        "check_time": "1970-01-01 09:30:00",
                                        "check_type": "OnDuty"
                                    },
                                    {
                                        "across": 0,
                                        "check_time": "1970-01-01 18:30:00",
                                        "check_type": "OffDuty"
                                    }
                                ]
                            }
                        ],
                        "setting": {
                            "absenteeism_late_minutes": -1,
                            "class_setting_id": 599750584,
                            "is_off_duty_free_check": "N",
                            "permit_late_minutes": -1,
                            "serious_late_minutes": -1,
                            "work_time_minutes": 540
                        }
                    }
                ],
                "type": "FIXED",
                "work_day_list": [
                    "0",
                    "686275261",
                    "686275261",
                    "686275261",
                    "686275261",
                    "686275261",
                    "0"
                ]
            }
        ],
        "has_more": false
    },
    "request_id": "5wut35udt0h9"
}`
	group := &response.GetAttendanceGroup{}

	err := json.Unmarshal([]byte(str), group)

	assert.Nil(t, err)
	assert.Equal(t, group.RequestId, "5wut35udt0h9")
	assert.Equal(t, group.HasMore, false)
	assert.Equal(t, group.Code, 0)

	assert.Equal(t, group.Groups[0].GroupId, 792440025)
	assert.Equal(t, group.Groups[0].GroupName, "技术考勤")
	assert.Equal(t, group.Groups[0].Default, false)
	assert.Equal(t, group.Groups[0].EnableEmpSelectClass, true)
	assert.Equal(t, group.Groups[0].DisableCheckWhenRest, false)
	assert.Equal(t, group.Groups[0].DefaultClassId, 686275261)
	assert.Equal(t, group.Groups[0].DisableCheckWithoutSchedule, false)
	assert.Equal(t, group.Groups[0].MemberCount, 1)
	assert.Equal(t, group.Groups[0].ManagerList[0], "manager164")
	assert.Equal(t, group.Groups[0].OwnerUserId, "manager164")
	assert.Equal(t, group.Groups[0].Type, "FIXED")

	assert.Equal(t, len(group.Groups[0].WorkDayList), 7)

	assert.Equal(t, group.Groups[0].SelectedClass[0].ClassId, 686275261)
	assert.Equal(t, group.Groups[0].SelectedClass[0].ClassName, "A")
	assert.Equal(t, group.Groups[0].SelectedClass[0].Setting.ClassSettingId, 599750584)
	assert.Equal(t, group.Groups[0].SelectedClass[0].Setting.AbsenteeismLateMinutes, -1)
	assert.Equal(t, group.Groups[0].SelectedClass[0].Setting.IsOffDutyFreeCheck, "N")
	assert.Equal(t, group.Groups[0].SelectedClass[0].Setting.PermitLateMinutes, -1)
	assert.Equal(t, group.Groups[0].SelectedClass[0].Setting.SeriousLateMinutes, -1)
	assert.Equal(t, group.Groups[0].SelectedClass[0].Setting.WorkTimeMinutes, 540)

	assert.Equal(t, group.Groups[0].SelectedClass[0].Sections[0].Times[0].CheckTime, "1970-01-01 09:30:00")
	assert.Equal(t, group.Groups[0].SelectedClass[0].Sections[0].Times[0].CheckType, "OnDuty")
	assert.Equal(t, group.Groups[0].SelectedClass[0].Sections[0].Times[0].Across, 0)

	assert.Equal(t, group.Groups[0].SelectedClass[0].Sections[0].Times[1].CheckTime, "1970-01-01 18:30:00")
	assert.Equal(t, group.Groups[0].SelectedClass[0].Sections[0].Times[1].CheckType, "OffDuty")
	assert.Equal(t, group.Groups[0].SelectedClass[0].Sections[0].Times[1].Across, 0)
}

func TestGetAttendanceUserGroup(t *testing.T) {
	t.Skip()
	group, err := client.GetAttendanceUserGroup("manager164")

	assert.Nil(t, err)
	assert.NotNil(t, group)
	assert.Equal(t, group.AttendanceUserGroup.Type, "FIXED")
	assert.Equal(t, group.AttendanceUserGroup.Name, "技术考勤")
	assert.Equal(t, group.AttendanceUserGroup.GroupId, 792440025)
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Name, "考勤班次")
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].ClassId, 686275261)

	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Setting.RestEndTime.Across, 0)
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Setting.RestEndTime.CheckType, "OffDuty")
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Setting.RestEndTime.CheckTime, "1970-01-01 13:00:00")

	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Setting.RestBeginTime.Across, 0)
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Setting.RestBeginTime.CheckType, "OnDuty")
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Setting.RestBeginTime.CheckTime, "1970-01-01 12:00:00")

	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[0].Across, 0)
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[0].BeginMin, 1)
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[0].EndMin, 1)
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[0].CheckType, "OnDuty")
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[0].CheckTime, "1970-01-01 09:30:00")

	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[1].Across, 0)
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[1].BeginMin, 1)
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[1].EndMin, 1)
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[1].CheckType, "OffDuty")
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[1].CheckTime, "1970-01-01 18:30:00")
}

func TestDingTalk_GetAttendanceGroupMinimalism(t *testing.T) {
	minimalism, err := client.GetAttendanceGroupMinimalism("manager164", 0)

	assert.Nil(t, err)
	assert.Equal(t, minimalism.MinimalismGroup.Attendance[0].Id, 792440025)
	assert.Equal(t, minimalism.MinimalismGroup.Attendance[0].Name, "技术考勤")
}

func TestDingTalk_GetAttendanceGroupDetail(t *testing.T) {
	detail, err := client.GetAttendanceGroupDetail("manager164", 792440025)

	assert.Nil(t, err)
	assert.Equal(t, detail.AttendanceGroup.Id, 792440025)
	assert.Equal(t, detail.AttendanceGroup.Name, "技术考勤")
	assert.Equal(t, detail.AttendanceGroup.AddressList[0], "绿城未来park")
	assert.Equal(t, detail.AttendanceGroup.MemberCount, 3)
	assert.Equal(t, detail.AttendanceGroup.OwnerUserId, "manager164")
	assert.Equal(t, detail.AttendanceGroup.ShiftIds[0], 686275261)
	assert.Equal(t, detail.AttendanceGroup.Type, "FIXED")
	assert.Equal(t, detail.AttendanceGroup.Wifis[0], "wifi")
	assert.Equal(t, len(detail.AttendanceGroup.WorkDayList), 7)
}

func TestDingTalk_SearchAttendanceGroup(t *testing.T) {
	group, err := client.SearchAttendanceGroup("manager164", "技术")

	assert.Nil(t, err)
	assert.Equal(t, len(group.AttendanceGroups), 1)
}

func TestDingTalk_CreateAttendanceGroup_TURN(t *testing.T) {
	//t.Skip()
	group := request.NewCreateAttendanceGroup("manager164",
		"创建排班制考勤组",
		attendance.TURN,
		[]request.AttendanceMember{request.NewAttendanceMember("manager164", "")}).
		Build()

	res, err := client.CreateAttendanceGroup(group)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.NotNil(t, res.AttendanceGroup.Name, "创建排班制考勤组")
}

func TestDingTalk_UpdateAttendanceGroup_TURN(t *testing.T) {
	t.Skip()
	group := request.NewCreateAttendanceGroup("manager164",
		"创建排班制考勤组",
		attendance.TURN,
		[]request.AttendanceMember{request.NewAttendanceMember("manager164", "")}).
		SetDefaultClassId(1197110225).
		Build()

	res, err := client.CreateAttendanceGroup(group)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.NotNil(t, res.AttendanceGroup.Name, "创建排班制考勤组")
}

func TestDingTalk_DeleteAttendanceGroup(t *testing.T) {
	//t.Skip()
	res, err := client.DeleteAttendanceGroup("manager164", "EBA6891DAA04268FB3BFB43E5DC1910D")
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.Result, "EBA6891DAA04268FB3BFB43E5DC1910D")
}

func TestDingTalk_IdToKeyAttendanceGroup(t *testing.T) {
	res, err := client.IdToKeyAttendanceGroup("manager164", 1210435013)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.Result, "EBA6891DAA04268FB3BFB43E5DC1910D")
}
func TestDingTalk_CreateAttendanceShift(t *testing.T) {
	uid := "020628151124053635"
	name := "早午班13"
	shift := request.NewCreateAttendanceShift(uid,
		name, []request.AttendanceShiftSection{
			request.AttendanceShiftSection{
				Times: []request.AttendanceShiftTime{
					request.AttendanceShiftTime{
						FreeCheck: true,
						BeginMin:  0,
						CheckType: "OffDuty",
						Across:    0,
						EndMin:    0,
						CheckTime: "2024-06-07 12:30:00",
					},
					request.AttendanceShiftTime{
						FreeCheck: true,
						BeginMin:  0,
						CheckType: "OnDuty",
						Across:    0,
						EndMin:    0,
						CheckTime: "2024-06-07 15:50:00",
					},
				},
			},
		}).
		SetOwner(uid).
		Build()

	res, err := client.CreateAttendanceShift(shift)
	assert.Nil(t, err)
	assert.Equal(t, res.AttendanceShift.Name, name)
}

func TestDingTalk_DeleteAttendanceShift(t *testing.T) {
	res, err := client.DeleteAttendanceShift("manager164", 1212315041)
	assert.Nil(t, err)
	fmt.Println(res)
}

func TestDingTalk_UpdateAttendanceShiftPunches(t *testing.T) {
	// IMPORTANT 修改成功以后节点id会发生变化
	punches := []request.AttendanceShiftPunches{
		{
			Id:        937800001,
			FreeCheck: true,
		},
		{
			Id:        937800001,
			FreeCheck: true,
		},
	}

	shiftPunches := request.NewUpdateAttendanceShiftPunches("manager164", 1219930037, punches)
	res, err := client.UpdateAttendanceShiftPunches(shiftPunches)
	assert.Nil(t, err)
	assert.Equal(t, res.Success, true)
}

func TestDingTalk_GetAttendanceShiftList(t *testing.T) {
	res, err := client.GetAttendanceShiftList("manager164", 0)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetAttendanceShiftDetail(t *testing.T) {
	res, err := client.GetAttendanceShiftDetail("manager164", 1219930037)
	fmt.Println(res)
	assert.Nil(t, err)
}

func TestDingTalk_GetAttendanceListSchedule(t *testing.T) {
	t.Skip()
	res, err := client.GetAttendanceListSchedule("2024-06-07", 1, 10)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.AttendanceListSchedule.Schedules[0].UserID, "020628151124053635")
	assert.Equal(t, res.AttendanceListSchedule.Schedules[0].ClassID, 1213915286)
	assert.Equal(t, res.AttendanceListSchedule.Schedules[0].GroupID, 1197110225)
}

func TestDingTalk_GetAttendanceScheduleList(t *testing.T) {
	list, err := client.GetAttendanceScheduleList("manager164", "020628151124053635", 1717658407000)

	assert.Nil(t, err)
	assert.NotNil(t, list)
}

func TestDingTalk_GetAttendanceScheduleUsersList(t *testing.T) {
	list, err := client.GetAttendanceScheduleUsersList("manager164", "020628151124053635,1315165650938287,285566344124169124,09285247331246086", 1717430400000, 1717516800000)

	assert.Nil(t, err)
	assert.NotNil(t, list)
}

func TestDingTalk_UpdateAttendanceSchedule(t *testing.T) {
	schedule := request.NewAttendanceSchedule(1213915286, 1717689600000, false, "020628151124053635")
	req := request.NewUpdateAttendanceSchedule("manager164", 1197110225, []request.AttendanceSchedule{schedule})
	res, err := client.UpdateAttendanceSchedule(req)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetAttendanceScheduleShift(t *testing.T) {

	req := request.NewGetAttendanceScheduleShift("manager164", "020628151124053635,1315165650938287,285566344124169124,09285247331246086", 1717430400000, 1717506800000)
	res, err := client.GetAttendanceScheduleShift(req)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}
