package dc2

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type CreateRequest struct {
	RegionId                  string      `json:"regionId"`               // 地域 ID
	ZoneId                    string      `json:"zoneId,omitempty"`       // 可用区 ID
	IsAutoContinue            bool        `json:"autoContinue,omitempty"` // 是否设置 DC2 自动续费
	PayPeriod                 int         `json:"payPeriod,omitempty"`    // 购买包月时长，单位为月，不传或传 0 表示后付费
	Count                     int         `json:"count,omitempty"`        // 批量购买参数，不传默认购买一台 DC2，不能超过 20
	CouponId                  string      `json:"couponId,omitempty"`     // 本次创建使用的优惠券 ID
	SubnetUuid                string      `json:"subnetUuid,omitempty"`   // 在此指定子网下创建 DC2。如果未传，则会在目标地域的默认 VPC 下对应可用区的默认子网中创建 DC2
	Model                     string      `json:"dc2Model"`               // 要创建的 dc2 型号
	ImgUuid                   string      `json:"imgUuid"`                // 使用何镜像创建 DC2，与 SnapUuid 二选一
	SnapUuid                  string      `json:"snapUuid"`               // 使用何快照创建 DC2，与 imgUuid 二选一
	PubKeyUuids               []string    `json:"pubKeyUuids"`            // 使用公钥 Uuid 列表进行 DC2 创建，与 password 参数二选一。
	Password                  string      `json:"password"`
	RootDiskSize              int         `json:"rootDiskSize"`
	RootDiskType              string      `json:"rootDiskType"`
	Tags                      []string    `json:"dc2Tags,omitempty"`
	Name                      string      `json:"name,omitempty"`
	IsProSecurityAgentEnabled bool        `json:"proSecurityAgentEnabled,omitempty"` // 是否同时安装主机安全 Agent 专业版
	IsMonitoringAgentEnabled  bool        `json:"monitoringAgentEnabled,omitempty"`  // 是否同时安装监控 Agent
	SgUuids                   []string    `json:"sgUuids,omitempty"`                 // 安全组列表
	Eip                       EipParams   `json:"eip,omitempty"`
	Ebs                       []EbsParams `json:"ebs,omitempty"`
}

type EipParams struct {
	BandWidth        int      `json:"bandWidth"`                // 带宽，为 1-100 之间的值
	IsChargeWithFlow bool     `json:"chargeWithFlow,omitempty"` // 是否要按流量计费 （为按流量计费时，默认为后付费）
	Tags             []string `json:"eipTags,omitempty"`
}

type EbsParams struct {
	Count    int      `json:"count,omitempty"`
	Name     string   `json:"name,omitempty"`
	Size     int64    `json:"size"`
	DiskType string   `json:"diskType"`
	SnapUuid string   `json:"snapUuid,omitempty"`
	Tags     []string `json:"ebsTags,omitempty"`
}

type CreateResponse struct {
	Errno     int          `json:"errno"`
	Errmsg    string       `json:"errmsg"`
	RequestId string       `json:"requestId"`
	Data      []schema.Job `json:"data"`
}

type CreateRequestBuilder struct {
	regionId                  string
	zoneId                    string
	isAutoContinue            bool
	payPeriod                 int
	count                     int
	couponId                  string
	subnetUuid                string
	model                     string
	imgUuid                   string
	snapUuid                  string
	pubKeyUuids               []string
	password                  string
	rootDiskType              string
	dc2Tags                   []string
	isProSecurityAgentEnabled bool
	isMonitoringAgentEnabled  bool
	sgUuids                   []string

	eipBandWidth        int
	eipIsChargeWithFlow bool
	eipTags             []string

	ebsCount    int
	ebsName     string
	ebsSize     int64
	ebsDiskType string
	ebsSnapUuid string
	ebsTags     []string
}

func (b *CreateRequestBuilder) SetRegionId(regionId string) {
	b.regionId = regionId
}

func (b *CreateRequestBuilder) SetZoneId(zoneId string) {
	b.zoneId = zoneId
}

func (b *CreateRequestBuilder) SetIsAutoContinue(isAuto bool) {
	b.isAutoContinue = isAuto
}

func (b *CreateRequestBuilder) SetPayPeriod(payPeriod int) {
	b.payPeriod = payPeriod
}

func (b *CreateRequestBuilder) SetCount(count int) {
	b.count = count
}

func (b *CreateRequestBuilder) SetCouponId(couponId string) {
	b.couponId = couponId
}

func (b *CreateRequestBuilder) SetSubnetUuid(subnetUuid string) {
	b.subnetUuid = subnetUuid
}

func (b *CreateRequestBuilder) SetDc2Model(model string) {
	b.model = model
}

func (b *CreateRequestBuilder) SetImgUuid(uuid string) {
	b.imgUuid = uuid
}

func (b *CreateRequestBuilder) SetSnapUuid(snapUuid string) {
	b.snapUuid = snapUuid
}

func (b *CreateRequestBuilder) SetPubKeyUuids(pubKeyUuids []string) {
	b.pubKeyUuids = pubKeyUuids
}

func (b *CreateRequestBuilder) SetPassword(password string) {
	b.password = password
}

func (b *CreateRequestBuilder) SetRootDiskType(rootDiskType string) {
	b.rootDiskType = rootDiskType
}

func (b *CreateRequestBuilder) SetDc2Tags(tags []string) {
	b.dc2Tags = tags
}

func (b *CreateRequestBuilder) InstallProSecurityAgent() {
	b.isProSecurityAgentEnabled = true
}

func (b *CreateRequestBuilder) InstallMonitoringAgent() {
	b.isMonitoringAgentEnabled = true
}

func (b *CreateRequestBuilder) SetSgUuids(sgUuids []string) {
	b.sgUuids = sgUuids
}

func (b *CreateRequestBuilder) SetBandWidth(bw int) {
	b.eipBandWidth = bw
}

func (b *CreateRequestBuilder) SetIsChargeWithFlow(withFlow bool) {
	b.eipIsChargeWithFlow = withFlow
}

func (b *CreateRequestBuilder) SetEipTags(tags []string) {
	b.eipTags = tags
}

func (b *CreateRequestBuilder) SetEbsCount(count int) {
	b.ebsCount = count
}

func (b *CreateRequestBuilder) SetEbsName(name string) {
	b.ebsName = name
}

func (b *CreateRequestBuilder) SetEbsSize(size int64) {
	b.ebsSize = size
}

func (b *CreateRequestBuilder) SetDiskType(dtype string) {
	b.ebsDiskType = dtype
}

func (b *CreateRequestBuilder) SetEbsSnapUuid(uuid string) {
	b.ebsSnapUuid = uuid
}

func (b *CreateRequestBuilder) SetEbsTags(tags []string) {
	b.ebsTags = tags
}

func (b *CreateRequestBuilder) Build() CreateRequest {
	return CreateRequest{
		RegionId:                  b.regionId,
		ZoneId:                    b.zoneId,
		IsAutoContinue:            b.isAutoContinue,
		PayPeriod:                 b.payPeriod,
		Count:                     b.count,
		CouponId:                  b.couponId,
		SubnetUuid:                b.subnetUuid,
		Model:                     b.model,
		ImgUuid:                   b.imgUuid,
		SnapUuid:                  b.snapUuid,
		PubKeyUuids:               b.pubKeyUuids,
		Password:                  b.password,
		RootDiskType:              b.rootDiskType,
		Tags:                      b.dc2Tags,
		IsProSecurityAgentEnabled: b.isProSecurityAgentEnabled,
		IsMonitoringAgentEnabled:  b.isMonitoringAgentEnabled,
		SgUuids:                   b.sgUuids,

		Eip: EipParams{
			BandWidth:        b.eipBandWidth,
			IsChargeWithFlow: b.eipIsChargeWithFlow,
			Tags:             b.eipTags,
		},

		Ebs: []EbsParams{
			{
				Count:    b.ebsCount,
				Name:     b.ebsName,
				Size:     b.ebsSize,
				DiskType: b.ebsDiskType,
				SnapUuid: b.snapUuid,
				Tags:     b.ebsTags,
			},
		},
	}
}

func (c *Client) Create(request *CreateRequest) (*schema.Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.CREATE_DC2_URL, data)
	ret := CreateResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
