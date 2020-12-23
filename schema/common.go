package schema

type Period struct {
	AutoRenewCnt int    `json:"autoRenewCnt"` // 资源到期前自动续费数目（包月即为月数）
	AutoSwitch   bool   `json:"autoSwitch"`   // 付费方式，先付费为prePaid，后付费为postPaid
	PayType      string `json:"payType"`      // 付费方式，先付费为prePaid，后付费为postPaid
}

type ChargeInfo struct {
	Period Period `json:"period"` // 付费时长信息
}

type Charge struct {
	ChargeInfo    ChargeInfo `json:"chargeInfo"`    // 计费信息
	CostThisMonth float64    `json:"costThisMonth"` // 本月消费
	EndTime       int64      `json:"endTime"`       // 包月到期时间（单位：毫秒）
}

// 区域信息
type Region struct {
	AreaName string `json:"areaName"` // 区域名称
	Id       string `json:"id"`       // region id
	Name     string `json:"name"`     // region name
	Zone     []Zone `json:"zone"`     // zone信息
}

// 可用区信息
type Zone struct {
	Id   string `json:"id"`   // zone id
	Name string `json:"name"` // zone name
}

// 任务进度
type Job struct {
	Done         bool    `json:"done"`         // 任务是否已经执行结束
	Progress     float64 `json:"progress"`     // 任务执行进度，为0-100之间的一个百分比值
	ResourceUuid string  `json:"ResourceUuid"` // 当前正在操作的资源Uuid
	Result       string  `json:"result"`       // 任务执行失败情况下，返回失败原因
	Success      bool    `json:"success"`      // 任务执行的结果是否成功
	Type         string  `json:"type"`         // 任务类型，描述这个任务在做什么事情
	Uuid         string  `json:"jobUuid"`      // 任务唯一标识
}

// DC2 规格
type Dc2Spec struct {
	CpuNum                int    `json:"cpuNum"` // CPU 数
	CpuSpeed              int64  `json:"CpuSpeed"`
	OfferingUuid          string `json:"offeringUuid"`
	DiskSize              int64  `json:"diskSize"` // 根盘大小
	GpuDeviceOfferingUuid string `json:"gpuDeviceOfferingUuid"`
	GpcNum                int    `json:"gpuNum"`     // GCP 数
	MemorySize            int64  `json:"memorySize"` // 内存大小
	Name                  string `json:"name"`
	Model                 string `json:"model"` // DC2型号
	RootDiskOfferingUuid  string `json:"rootDiskOfferingUuid"`
	Type                  string `json:"type"`
	TypeName              string `json:"typeName"`
}

// EIP规格
type EipSpec struct {
	Bandwidth         int64  `json:"babandwidth"` // 带宽
	CreateDate        string `json:"createDate"`
	ChargeType        string `json:"chargeType"` // 计费类型,`bandwidth` 为按带宽计费;`flow` 为按流量计费
	Description       string `json:"description"`
	InboundBandwidth  int64  `json:"inboundBandwidth"`
	LastOpDate        string `json:"lastOpDate"`
	Name              string `json:"name"`
	OfferingUuid      string `json:"offeringUuid"`
	OutboundBandwidth int64  `json:"outboundBandwidth"`
	PeerOfferingUuid  string `json:"peerOfferingUuid"`
	State             string `json:"state"`
	Type              string `json:"type"`
	Uuid              string `json:"uuid"`
}

type BEIP struct {
	BeipUuid   string `json:"beipUuid"`
	Ip         string `json:"ip"`
	CreateTime int    `json:"createTime"`
	UpdateTime int    `json:"updateTime"`
}

type VPC struct {
	VpcUuid    string `json:"vpcUuid"`
	Name       string `json:"name"`
	IsDefault  bool   `json:"isDefault"`
	createTime int    `json:"CreateTime"`
	updateTime int    `json:"UpdateTime"`
}

type Flow struct {
	In  float32 `json:"float32"`
	Out float32 `json:"float32"`
}

type SlbSpec struct {
	OfferingUuid string `json:"offeringUuid"`
}
