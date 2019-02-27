package schema

type Period struct {
	// 付费方式，先付费为prePaid，后付费为postPaid
	PayType string `json:"payType"`

	// 资源到期前自动续费数目（包月即为月数）
	AutoRenewCnt int `json:"autoRenewCnt"`

	// 付费方式，先付费为prePaid，后付费为postPaid
	AutoSwitch bool `json:"autoSwitch"`
}

type ChargeInfo struct {
	// 付费时长信息
	Period Period `json:"period"`
}

type Charge struct {
	// 计费信息
	ChargeInfo ChargeInfo `json:"chargeInfo"`

	// 包月到期时间（单位：毫秒）
	EndTime int64 `json:"endTime"`

	// 本月消费
	CostThisMonth float64 `json:"costThisMonth"`
}

// 区域信息
type Region struct {
	// 区域名称
	AreaName string `json:"areaName"`

	// region id
	Id string `json:"id"`

	// region name
	Name string `json:"name"`

	// zone信息
	Zone []Zone `json:"zone"`
}

// 可用区信息
type Zone struct {
	// zone id
	Id string `json:"id"`

	// zone name
	Name string `json:"name"`
}

// 任务进度
type Job struct {
	// 任务唯一标识
	Uuid string `json:"jobUuid"`

	// 当前正在操作的资源Uuid
	ResourceUuid string `json:"ResourceUuid"`

	// 任务类型，描述这个任务在做什么事情
	Type string `json:"type"`

	// 任务执行进度，为0-100之间的一个百分比值
	Progress float64 `json:"progress"`

	// 任务是否已经执行结束
	Done bool `json:"done"`

	// 任务执行的结果是否成功
	Success bool `json:"success"`

	// 任务执行失败情况下，返回失败原因
	Result string `json:"result"`
}

// DC2 规格
type Dc2Spec struct {
	CpuNum                int    `json:"cpuNum"` // CPU 数
	CpuSpeed              int64  `json:"CpuSpeed"`
	OfferingUuid          string `json:"offeringUuid"`
	GpcNum                int    `json:"gpuNum"` // GCP 数
	GpuDeviceOfferingUuid string `json:"gpuDeviceOfferingUuid"`
	MemorySize            int64  `json:"memorySize"` // 内存大小
	DiskSize              int64  `json:"diskSize"`   // 根盘大小
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
	PeerOfferingUuid  string `json:"peerOfferingUuid"`
	OutboundBandwidth int64  `json:"outboundBandwidth"`
	State             string `json:"state"`
	Type              string `json:"type"`
	Uuid              string `json:"uuid"`
}
