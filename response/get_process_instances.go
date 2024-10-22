package response

type GetProcessInstances struct {
	Response
	Result GetProcessInstancesResult `json:"result"`
}

type GetProcessInstancesResult struct {
	Title               string                `json:"title"`               //审批实例标题
	FinishTime          string                `json:"finishTime"`          //结束时间
	OriginatorUserId    string                `json:"originatorUserId"`    //发起人的userId
	Status              string                `json:"status"`              //审批状态
	Result              string                `json:"result"`              //审批结果
	FormComponentValues []FormComponentValues `json:"formComponentValues"` //用户自定义业务参数透出
	CreateTime          string                `json:"createTime"`          //创建时间
	CcUserIds           []string              `json:"ccUserIds"`           //参与人userId
	Tasks               []Task                `json:"tasks"`               //task
}

type FormComponentValues struct {
	ComponentType string `json:"componentType"`
	Name          string `json:"name"`
	BizAlias      string `json:"bizAlias"`
	Id            string `json:"id"`
	Value         string `json:"value"`
	ExtValue      string `json:"extValue"`
}

type Task struct {
	Result string `json:"result"`
	Status string `json:"status"`
	UserId string `json:"userId"`
}
