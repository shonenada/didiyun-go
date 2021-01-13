package schema

type Dc2 struct {
	Id         string   `json:"dc2Id"`
	Uuid       string   `json:"dc2Uuid"`
	Name       string   `json:"name"`
	Ip         string   `json:"ip"`
	OSType     string   `json:"osType"`
	Platform   string   `json:"platform"`
	Status     string   `json:"status"`
	ImageUuid  string   `json:"imgUuid"`
	Tags       []string `json:"dc2Tags"`
	Charge     Charge   `json:"charge"`
	Ebs        []Dc2Ebs `json:"ebs"`
	Eip        Dc2Eip   `json:"eip"`
	Spec       Dc2Spec  `json:"spec"`
	Vpc        Dc2Vpc   `json:"vpc"`
	Region     Region   `json:"region"`
	Job        Job      `json:"job"`
	CreateTime int64    `json:"createTime"`
	UpdateTime int64    `json:"updateTime"`
}

type Dc2Spec struct {
	CpuNum                int    `json:"cpuNum"`
	CpuSpeed              int64  `json:"CpuSpeed"`
	OfferingUuid          string `json:"offeringUuid"`
	DiskSize              int64  `json:"diskSize"`
	GpuDeviceOfferingUuid string `json:"gpuDeviceOfferingUuid"`
	GpcNum                int    `json:"gpuNum"`
	MemorySize            int64  `json:"memorySize"`
	Name                  string `json:"name"`
	Model                 string `json:"model"`
	RootDiskOfferingUuid  string `json:"rootDiskOfferingUuid"`
	Type                  string `json:"type"`
	TypeName              string `json:"typeName"`
}

type Dc2Vpc struct {
	Id         string   `json:"vpcId"`
	Tags       []string `json:"vpcTags"`
	Uuid       string   `json:"vpcUuid"`
	CreateTime int64    `json:"createTime"`
	UpdateTime int64    `json:"updateTime"`
}

type Dc2EbsSpec struct {
	DiskType    string `json:"diskType"`
	MaxDiskSize int64  `json:"maxDiskSize"`
	MinDiskSize int64  `json:"minDiskSize"`
	Model       string `json:"model"`
	Name        string `json:"name"`
	Size        int64  `json:"size"`
}

type Dc2Ebs struct {
	Name       string `json:"name"`
	Uuid       string `json:"ebsUuid"`
	Attr       string `json:"attr"`
	Region     Region `json:"region"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}

type Dc2Eip struct {
	EipId      string   `json:"eipId"`
	EipTags    []string `json:"eipTags"`
	Ip         string   `json:"ip"`
	State      string   `json:"state"`
	Status     string   `json:"status"`
	Uuid       string   `json:"eipUuid"`
	Charge     Charge   `json:"charge"`
	Region     Region   `json:"region"`
	Spec       EipSpec  `json:"spec"`
	CreateTime string   `json:"createTime"`
	UpdateTime string   `json:"updateTime"`
}
