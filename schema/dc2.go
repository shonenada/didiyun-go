package schema

type Dc2 struct {
	Charge     Charge   `json:"charge"`
	CreateTime int64    `json:"createTime"`
	Ebs        []Ebs    `json:"ebs"`
	Eip        Dc2Eip   `json:"eip"`
	Id         string   `json:"dc2Id"`
	ImageUuid  string   `json:"imgUuid"`
	Ip         string   `json:"ip"`
	Name       string   `json:"name"`
	Job        Job      `json:"job"`
	OSType     string   `json:"osType"`
	Platform   string   `json:"platform"`
	Region     Region   `json:"region"`
	Spec       Dc2Spec  `json:"spec"`
	Status     string   `json:"status"`
	Tags       []string `json:"dc2Tags"`
	UpdateTime int64    `json:"updateTime"`
	Uuid       string   `json:"dc2Uuid"`
	Vpc        Dc2Vpc   `json:"vpc"`
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
	Attr       string     `json:"attr"`
	CreateTime int64      `json:"createTime"`
	Dc2        Dc2        `json:"dc2"`
	EbsTags    []string   `json:"ebsTags"`
	EbsUuid    string     `json:"ebsUuid"`
	Job        Job        `json:"job"`
	Name       string     `json:"name"`
	Region     Region     `json:"region"`
	Size       int64      `json:"size"`
	Spec       Dc2EbsSpec `json:"spec"`
	UpdateTime int64      `json:"updateTime"`
}

type Dc2Eip struct {
	Charge     Charge   `json:"charge"`
	CreateTime string   `json:"createTime"`
	EipId      string   `json:"eipId"`
	EipTags    []string `json:"eipTags"`
	Ip         string   `json:"ip"`
	Region     Region   `json:"region"`
	Spec       EipSpec  `json:"spec"`
	State      string   `json:"state"`
	Status     string   `json:"status"`
	UpdateTime string   `json:"updateTime"`
	Uuid       string   `json:"eipUuid"`
}
