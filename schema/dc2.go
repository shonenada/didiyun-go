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
	Eip        Dc2Eip   `json:"eip"`
	Spec       Dc2Spec  `json:"spec"`
	Vpc        Dc2Vpc   `json:"vpc"`
	Ebs        []Dc2Ebs `json:"ebs"`
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
	Uuid       string   `json:"vpcUuid"`
	Tags       []string `json:"vpcTags"`
	CreateTime int64    `json:"createTime"`
	UpdateTime int64    `json:"updateTime"`
}

type Dc2EbsSpec struct {
	Name        string `json:"name"`
	Model       string `json:"model"`
	Size        int64  `json:"size"`
	DiskType    string `json:"diskType"`
	MaxDiskSize int64  `json:"maxDiskSize"`
	MinDiskSize int64  `json:"minDiskSize"`
}

type Dc2Ebs struct {
	Uuid       string `json:"ebsUuid"`
	Name       string `json:"name"`
	Attr       string `json:"attr"`
	Region     Region `json:"region"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}

type Dc2Eip struct {
	Id         string   `json:"eipId"`
	Uuid       string   `json:"eipUuid"`
	Ip         string   `json:"ip"`
	State      string   `json:"state"`
	Status     string   `json:"status"`
	Tags       []string `json:"eipTags"`
	Charge     Charge   `json:"charge"`
	Spec       EipSpec  `json:"spec"`
	Region     Region   `json:"region"`
	CreateTime string   `json:"createTime"`
	UpdateTime string   `json:"updateTime"`
}
