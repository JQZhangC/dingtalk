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
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
	"net/http"
	"net/url"
)

// GetAttendanceGroups 批量获取考勤组详情
func (ding *DingTalk) GetAttendanceGroups(offset, size int) (rsp response.GetAttendanceGroup, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceGroupsKey, nil,
		request.NewGetAttendanceGroup(offset, size), &rsp)
}

// GetAttendanceUserGroup 获取用户考勤组
func (ding *DingTalk) GetAttendanceUserGroup(userId string) (rsp response.GetAttendanceUserGroup, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceUserGroupKey, nil,
		request.NewGetAttendanceUserGroup(userId), &rsp)
}

// GetAttendanceGroupMinimalism 批量获取考勤组摘要
func (ding *DingTalk) GetAttendanceGroupMinimalism(userId string, cursor int) (rsp response.GetAttendanceGroupMinimalism,
	err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceGroupMinimalismKey, nil,
		request.NewGetAttendanceGroupMinimalism(userId, cursor), &rsp)
}

// GetAttendanceGroupDetail 获取考勤组详情
func (ding *DingTalk) GetAttendanceGroupDetail(userId string, groupId int) (rsp response.GetAttendanceGroupDetail,
	err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceGroupDetailKey, nil,
		request.NewGetAttendanceGroupDetail(userId, groupId), &rsp)
}

// SearchAttendanceGroup 搜索考勤组摘要
func (ding *DingTalk) SearchAttendanceGroup(userId, groupName string) (rsp response.SearchAttendanceGroup, err error) {
	return rsp, ding.Request(http.MethodPost, constant.SearchAttendanceGroupKey, nil,
		request.NewSearchAttendanceGroup(userId, groupName), &rsp)
}

// CreateAttendanceGroup 创建考勤组
func (ding *DingTalk) CreateAttendanceGroup(res *request.CreateAttendanceGroup) (rsp response.CreateAttendanceGroup, err error) {
	return rsp, ding.Request(http.MethodPost, constant.CreateAttendanceGroupKey, nil, res, &rsp)
}

// TODO
// UpdateAttendanceGroup 更新考勤组
func (ding *DingTalk) UpdateAttendanceGroup(res *request.UpdateAttendanceGroup) (rsp response.UpdateAttendanceGroup, err error) {
	return rsp, ding.Request(http.MethodPost, constant.UpdateAttendanceGroupKey, nil, res, &rsp)
}

// DeleteAttendanceGroup 删除考勤组
func (ding *DingTalk) DeleteAttendanceGroup(userId, groupKey string) (rsp response.DeleteAttendanceGroup, err error) {
	res := request.NewDeleteAttendanceGroup(userId, groupKey)
	return rsp, ding.Request(http.MethodPost, constant.DeleteAttendanceGroupKey, nil, res, &rsp)
}

// IdToKeyAttendanceGroup groupId转换为groupKey
func (ding *DingTalk) IdToKeyAttendanceGroup(userId string, group_id int64) (rsp response.IdToKeyAttendanceGroup, err error) {
	res := request.NewIdToKeyAttendanceGroup(userId, group_id)
	return rsp, ding.Request(http.MethodPost, constant.IdToKeyAttendanceGroupKey, nil, res, &rsp)
}

// CreateAttendanceShift 创建班次
func (ding *DingTalk) CreateAttendanceShift(res *request.CreateAttendanceShift) (rsp response.CreateAttendanceShift, err error) {
	return rsp, ding.Request(http.MethodPost, constant.CreateAttendanceShiftKey, nil, res, &rsp)
}

// DeleteAttendanceShift 删除班次
func (ding *DingTalk) DeleteAttendanceShift(userId string, shiftId int64) (rsp response.Response, err error) {
	res := request.NewDeleteAttendanceShift(userId, shiftId)
	return rsp, ding.Request(http.MethodPost, constant.DeleteAttendanceShiftKey, nil, res, &rsp)
}

// // UpdateAttendanceShiftPunches 修改打卡时段设置
func (ding *DingTalk) UpdateAttendanceShiftPunches(res *request.UpdateAttendanceShiftPunches) (rsp response.Response, err error) {
	return rsp, ding.Request(http.MethodPost, constant.UpdateAttendanceShiftPunchesKey, nil, res, &rsp)
}

// GetAttendanceShiftList 获取班次摘要信息
func (ding *DingTalk) GetAttendanceShiftList(userId string, cursor int64) (rsp response.GetAttendanceShiftList, err error) {
	res := request.NewGetAttendanceShiftList(userId, cursor)
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceShiftListKey, nil, res, &rsp)
}

// GetAttendanceShiftDetail 获取班次详情
func (ding *DingTalk) GetAttendanceShiftDetail(userId string, shiftId int64) (rsp response.GetAttendanceShiftDetail, err error) {
	res := request.NewGetAttendanceShiftDetail(userId, shiftId)
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceShiftDetailKey, nil, res, &rsp)
}

// GetAttendanceScheduleList 查询成员排班信息
func (ding *DingTalk) GetAttendanceScheduleList(opUserId, userId string, dateTime int64) (rsp response.GetAttendanceScheduleList, err error) {
	res := request.NewGetAttendanceScheduleList(opUserId, userId, dateTime)
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceScheduleDayListKey, nil, res, &rsp)
}

// GetAttendanceScheduleUsersList 批量查询人员排班信息
func (ding *DingTalk) GetAttendanceScheduleUsersList(opUserId, userIds string, fromDateTime, toDateTime int64) (rsp response.GetAttendanceScheduleUsersList, err error) {
	res := request.NewGetAttendanceScheduleShift(opUserId, userIds, fromDateTime, toDateTime)
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceScheduleUsersListKey, nil, res, &rsp)
}

// UpdateAttendanceSchedule 排班制考勤组排班
func (ding *DingTalk) UpdateAttendanceSchedule(res *request.UpdateAttendanceSchedule) (rsp response.Response, err error) {
	return rsp, ding.Request(http.MethodPost, constant.UpdateAttendanceScheduleKey, nil, res, &rsp)
}

// GetAttendanceScheduleShift 批量查询成员排班概要信息
func (ding *DingTalk) GetAttendanceScheduleShift(res *request.GetAttendanceScheduleShift) (rsp response.GetAttendanceScheduleShift, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceScheduleShiftKey, nil, res, &rsp)
}

// GetAttendanceListSchedule 查询企业考勤排班详情
func (ding *DingTalk) GetAttendanceListSchedule(workDate string, offset, size int) (rsp response.GetAttendanceListSchedule, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceListScheduleKey, nil,
		request.GetAttendanceListSchedule{workDate, offset, size}, &rsp)
}

func (ding *DingTalk) CreateLeavesTypes(opUserId string, res *request.CreateAttendanceLeavesTypes) (rsp response.CreateAttendanceLeavesTypes, err error) {
	query := url.Values{}
	query.Set("opUserId", opUserId)
	return rsp, ding.Request(http.MethodPost, constant.CreateLeavesTypes, query, res, &rsp)
}

func (ding *DingTalk) UpdateLeavesTypes(opUserId string, res *request.UpdateAttendanceLeavesTypes) (rsp response.UpdateAttendanceLeavesTypes, err error) {
	query := url.Values{}
	query.Set("opUserId", opUserId)
	return rsp, ding.Request(http.MethodPut, constant.UpdateLeavesTypes, query, res, &rsp)
}

func (ding *DingTalk) DeleteAttendanceVacationType(opUserId, leaveCode string) (rsp response.UpdateAttendanceLeavesTypes, err error) {
	type Request struct {
		OpUserId  string `json:"op_userid"`
		LeaveCode string `json:"leave_code"`
	}

	res := &Request{
		OpUserId:  opUserId,
		LeaveCode: leaveCode,
	}
	return rsp, ding.Request(http.MethodPost, constant.DeleteAttendanceVacationType, nil, res, &rsp)
}

func (ding *DingTalk) InitAttendanceVacationQuota(res *request.InitAttendanceVacationQuota) (rsp response.AttendanceVacationQuotaInit, err error) {
	return rsp, ding.Request(http.MethodPost, constant.InitAttendanceVacationQuota, nil, res, &rsp)
}

func (ding *DingTalk) UpdateAttendanceVacationQuota(res *request.UpdateAttendanceVacationQuota) (rsp response.AttendanceVacationQuotaInit, err error) {
	return rsp, ding.Request(http.MethodPost, constant.UpdateAttendanceVacationQuota, nil, res, &rsp)
}

func (ding *DingTalk) GetAttendanceVacationTypeList(opUserId, vacationSource string) (rsp response.GetAttendanceVacationTypeList, err error) {
	type Request struct {
		OpUserId       string `json:"op_userid"`
		VacationSource string `json:"vacation_source"`
	}

	res := &Request{
		OpUserId:       opUserId,
		VacationSource: vacationSource,
	}
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceVacationTypeList, nil, res, &rsp)
}

func (ding *DingTalk) GetAttendanceVacationQuotaList(res *request.GetAttendanceVacationQuotaList) (rsp response.GetAttendanceVacationQuotaList, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceVacationQuotaList, nil, res, &rsp)
}

func (ding *DingTalk) GetAttendanceVacationsRecords(res *request.GetAttendanceVacationsRecords) (rsp response.GetAttendanceVacationsRecords, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceVacationsRecords, nil, res, &rsp)
}

func (ding *DingTalk) GetAttColumns() (rsp response.GetAttColumns, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttColumns, nil, nil, &rsp)
}

func (ding *DingTalk) GetAttColumnValues(res *request.GetAttColumnValues) (rsp response.GetAttColumnValues, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttColumnValues, nil, res, &rsp)
}
