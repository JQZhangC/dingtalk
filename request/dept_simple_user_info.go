package request

//DeptSimpleUserInfo 获取部门用户基础信息
type DeptSimpleUserInfo struct {
	//部门ID，根部门ID为1
	DeptId int `json:"dept_id" validate:"required,min=1"`

	//分页查询的游标，最开始传0，后续传返回参数中的next_cursor值
	Cursor int `json:"cursor" validate:"omitempty,min=0"`

	//分页长度，最大值100。
	Size int `json:"size" validate:"required,max=100"`

	//部门成员的排序规则：
	//
	//entry_asc：代表按照进入部门的时间升序。
	//
	//entry_desc：代表按照进入部门的时间降序。
	//
	//modify_asc：代表按照部门信息修改时间升序。
	//
	//modify_desc：代表按照部门信息修改时间降序。
	//
	//custom：代表用户定义(未定义时按照拼音)排序。
	//
	//默认值：custom。
	OrderField string `json:"order_field,omitempty" validate:"omitempty,oneof=entry_asc entry_desc modify_asc modify_desc"`

	//是否返回访问受限的员工
	ContainAccessLimit bool `json:"contain_access_limit,omitempty"`

	//通讯录语言，默认zh_CN。如果是英文，请传入en_US。
	Language string `json:"language,omitempty" validate:"omitempty,oneof=zh_CN en_US"`
}

type deptSimpleUserInfoBuilder struct {
	du *DeptSimpleUserInfo
}

func NewDeptSimpleUserInfo(deptId, cursor, size int) *deptSimpleUserInfoBuilder {
	return &deptSimpleUserInfoBuilder{du: &DeptSimpleUserInfo{DeptId: deptId, Cursor: cursor, Size: size}}
}

//SetOrderField 部门成员的排序规则：
//
//entry_asc：代表按照进入部门的时间升序。
//
//entry_desc：代表按照进入部门的时间降序。
//
//modify_asc：代表按照部门信息修改时间升序。
//
//modify_desc：代表按照部门信息修改时间降序。
//
//custom：代表用户定义(未定义时按照拼音)排序。
//
//默认值：custom。
func (dub *deptSimpleUserInfoBuilder) SetOrderField(orderField string) *deptSimpleUserInfoBuilder {
	dub.du.OrderField = orderField
	return dub
}

//SetContainAccessLimit 是否返回访问受限的员工
func (dub *deptSimpleUserInfoBuilder) SetContainAccessLimit(containAccessLimit bool) *deptSimpleUserInfoBuilder {
	dub.du.ContainAccessLimit = containAccessLimit
	return dub
}

//SetLanguage 通讯录语言，默认zh_CN。如果是英文，请传入en_US。
func (dub *deptSimpleUserInfoBuilder) SetLanguage(language string) *deptSimpleUserInfoBuilder {
	dub.du.Language = language
	return dub
}

func (dub *deptSimpleUserInfoBuilder) Build() *DeptSimpleUserInfo {
	return dub.du
}