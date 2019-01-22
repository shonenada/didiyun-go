package dc2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/shonenada/didiyun-go/schema"
)

type CreateDc2Request struct {
	RegionId                string     `json:"regionId"`     // 地域 ID
	ZoneId                  string     `json:"zoneId"`       // 可用区 ID
	AutoContinue            bool       `json:"autoContinue"` // 是否设置 DC2 自动续费
	PayPeriod               int        `json:"payPeriod"`    // 购买包月时长，单位为月，不传或传 0 表示后付费
	Count                   int        `json:"count"`        // 批量购买参数，不传默认购买一台 DC2，不能超过 20
	CouponId                string     `json:"couponId"`     // 本次创建使用的优惠券 ID
	SubnetUuid              string     `json:"subnetUuid"`   // 在此指定子网下创建 DC2。如果未传，则会在目标地域的默认 VPC 下对应可用区的默认子网中创建 DC2
	Dc2Model                string     `json:"dc2Model"`     // 要创建的dc2型号
	ImgUuid                 string     `json:"imgUuid"`      // 使用何镜像创建 DC2，与 SnapUuid 二选一
	SnapUuid                string     `json:"snapUuid"`     // 使用何快照创建DC2，与imgUuid二选一
	PubKeyUuids             []string   `json:"pubKeyUuids"`  // 使用公钥 Uuid 列表进行 DC2 创建，与 password 参数二选一。
	Password                string     `json:"password"`
	RootDiskSize            int        `json:"rootDiskSize"`
	RootDiskType            string     `json:"rootDiskType"`
	Dc2Tags                 []string   `json:"dc2Tags"`
	Name                    string     `json:"name"`
	ProSecurityAgentEnabled bool       `json:"proSecurityAgentEnabled"` // 是否同时安装主机安全 Agent 专业版
	MonitoringAgentEnabled  bool       `json:"monitoringAgentEnabled"`  // 是否同时安装监控 Agent
	SgUuids                 []string   `json:"sgUuids"`                 // 安全组列表
	Eip                     EipInput   `json:"eip"`
	Ebs                     []EbsInput `json:"ebs"`
}

type EipInput struct {
	BandWidth      int      `json:"bandWidth"`      // 带宽，为 1-100 之间的值
	ChargeWithFlow bool     `json:"chargeWithFlow"` // 是否要按流量计费 （为按流量计费时，默认为后付费）
	EipTags        []string `json:"eipTags"`
}

type EbsInput struct {
	Count    int      `json:"count"`
	Name     string   `json:"name"`
	Size     int64    `json:"size"`
	DiskType string   `json:"diskType"`
	SnapUuid string   `json:"snapUuid"`
	EbsTags  []string `json:"ebsTags"`
}

type CreateDc2Response struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `jon:"job"`
}

type CreateDc2RequestBuilder struct {
	regionId                string
	zoneId                  string
	autoContinue            bool
	payPeriod               int
	count                   int
	couponId                string
	subnetUuid              string
	model                   string
	imgUuid                 string
	snapUuid                string
	pubKeyUuids             []string
	password                string
	rootDiskType            string
	dc2Tags                 []string
	proSecurityAgentEnabled bool
	monitoringAgentEnabled  bool
	sgUuids                 []string

	eipBandWidth      int
	eipChargeWithFlow bool
	eipTags           []string

	ebsCount    int
	ebsName     string
	ebsSize     int64
	ebsDiskType string
	ebsSnapUuid string
	ebsTags     []string
}

func (b *CreateDc2RequestBuilder) SetRegionId(regionId string) {
	b.regionId = regionId
}

func (b *CreateDc2RequestBuilder) SetZoneId(zoneId string) {
	b.zoneId = zoneId
}

func (b *CreateDc2RequestBuilder) SetAutoContinue(isAuto bool) {
	b.autoContinue = isAuto
}

func (b *CreateDc2RequestBuilder) SetPayPeriod(payPeriod int) {
	b.payPeriod = payPeriod
}

func (b *CreateDc2RequestBuilder) SetCount(count int) {
	b.count = count
}

func (b *CreateDc2RequestBuilder) SetCouponId(couponId string) {
	b.couponId = couponId
}

func (b *CreateDc2RequestBuilder) SetSubnetUuid(subnetUuid string) {
	b.subnetUuid = subnetUuid
}

func (b *CreateDc2RequestBuilder) SetDc2Model(model string) {
	b.model = model
}

func (b *CreateDc2RequestBuilder) SetImgUuid(uuid string) {
	b.imgUuid = uuid
}

func (b *CreateDc2RequestBuilder) SetSnapUuid(snapUuid string) {
	b.snapUuid = snapUuid
}

func (b *CreateDc2RequestBuilder) SetPubKeyUuids(pubKeyUuids []string) {
	b.pubKeyUuids = pubKeyUuids
}

func (b *CreateDc2RequestBuilder) SetPassword(password string) {
	b.password = password
}

func (b *CreateDc2RequestBuilder) SetRootDiskType(rootDiskType string) {
	b.rootDiskType = rootDiskType
}

func (b *CreateDc2RequestBuilder) SetDc2Tags(tags []string) {
	b.dc2Tags = tags
}

func (b *CreateDc2RequestBuilder) InstallProSecurityAgent() {
	b.proSecurityAgentEnabled = true
}

func (b *CreateDc2RequestBuilder) InstallMonitoringAgentEnabled() {
	b.monitoringAgentEnabled = true
}

func (b *CreateDc2RequestBuilder) SetSgUuids(sgUuids []string) {
	b.sgUuids = sgUuids
}

func (b *CreateDc2RequestBuilder) SetBandWidth(bw int) {
	b.eipBandWidth = bw
}

func (b *CreateDc2RequestBuilder) SetChargeWithFlow(withFlow bool) {
	b.eipChargeWithFlow = withFlow
}

func (b *CreateDc2RequestBuilder) SetEipTags(tags []string) {
	b.eipTags = tags
}

func (b *CreateDc2RequestBuilder) SetEbsCount(count int) {
	b.ebsCount = count
}

func (b *CreateDc2RequestBuilder) SetEbsName(name string) {
	b.ebsName = name
}

func (b *CreateDc2RequestBuilder) SetEbsSize(size int64) {
	b.ebsSize = size
}

func (b *CreateDc2RequestBuilder) SetDiskType(dtype string) {
	b.ebsDiskType = dtype
}

func (b *CreateDc2RequestBuilder) SetEbsSnapUuid(uuid string) {
	b.ebsSnapUuid = uuid
}

func (b *CreateDc2RequestBuilder) SetEbsTags(tags []string) {
	b.ebsTags = tags
}

func (b *CreateDc2RequestBuilder) Build() CreateDc2Request {
	return CreateDc2Request{
		RegionId:                b.regionId,
		ZoneId:                  b.zoneId,
		AutoContinue:            b.autoContinue,
		PayPeriod:               b.payPeriod,
		Count:                   b.count,
		CouponId:                b.couponId,
		SubnetUuid:              b.subnetUuid,
		Dc2Model:                b.model,
		ImgUuid:                 b.imgUuid,
		SnapUuid:                b.snapUuid,
		PubKeyUuids:             b.pubKeyUuids,
		Password:                b.password,
		RootDiskType:            b.rootDiskType,
		Dc2Tags:                 b.dc2Tags,
		ProSecurityAgentEnabled: b.proSecurityAgentEnabled,
		MonitoringAgentEnabled:  b.monitoringAgentEnabled,
		SgUuids:                 b.sgUuids,

		Eip: EipInput{
			BandWidth:      b.eipBandWidth,
			ChargeWithFlow: b.eipChargeWithFlow,
			EipTags:        b.eipTags,
		},

		Ebs: []EbsInput{
			{
				Count:    b.ebsCount,
				Name:     b.ebsName,
				Size:     b.ebsSize,
				DiskType: b.ebsDiskType,
				SnapUuid: b.snapUuid,
				EbsTags:  b.ebsTags,
			},
		},
	}
}

func (c *Client) CreateDc2(reqBody *CreateDc2Request) (*Job, error) {
	httpClient := &http.Client{}
	reqStr, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}

	req, err := http.NewRequest("POST", CREATE_DC2_URL, bytes.NewBuffer([]byte(reqStr)))
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}

	ret := CreateDc2Response{}
	json.Unmarshal([]byte(body), &ret)

	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}

	return &ret.Data[0], nil
}
