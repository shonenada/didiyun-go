package dc2

improt (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"

    . "github.com/shonenada/didiyun-go/schema"
)

type ListDC2Request struct {
    // 地域 ID
    RegionId string `json:"regionId"`

    // 可用区 ID，希望查询的可用区，不传则表示查询此地域下的所有可用区
    ZoneId string `json:"zondId"`

    // 查询 DC2 列表起始 index，从 0 开始
    Start int `json:"start"`

    // 查询 DC2 列表元素数量
    Limit int `json:"limit"`

    // 是否简化输出
    Simplify bool `json:"simplify"`

    // 查询 DC2 列表筛选条件
    condition DC2Condition `json:"condition"`
}

type DC2Response struct {
    // 此 DC2 正在进行的任务，若无任务则没有此字段
    Job Job `json:"job"`

    // DC2 唯一标识
    DC2UUID string `json:"dc2Uuid"`

    // DC2 名称
    Name string `json:"name"`

    // DC2 创建时间
    CreateTime int64 `json:"createTime"`

    // DC2 更新时间
    UpdateTime int64 `json:"updateTime"`

    // DC2 内网 IP
    IP string `json:"ip"`

    // DC2 的 tags
    Tags []string `json:"tags"`

    // DC2 状态
    Status string `json:"status"`

    // DC2 操作系统发行版及版本号
    OSType string `json:"osType"`

    // 与 DC2 关联的 EIP 信息，没有 EIP 则没有该字段
    EIP EIP `json:"eip"`

    // 与 DC2 关联的 EBS 信息，没有 EBS 则没有该字段，如果是通用型 DC2，则必有这个字段，且根盘信息包含在内
    EBS []EBS `json:"ebs"`

    // region 信息
    Region Region `json:"region"`
}

type ListDC2Response struct {
	Errno     int      `json:"errno"`
	Errmsg    string   `json:"errmsg"`
	RequestId string   `json:"requestId"`
	Data      []DC2Response `json:"data"`
}

type ListDC2RequestBuilder struct {

}

func (b *ListDC2RequestBuilder) SetRegionId(regionId string) {
    b.RegionId = regionId
}

func (c *Client) ListDC2 () ListDC2Response {
}
