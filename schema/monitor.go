package schema

const DSTYPE_GAUGE = "GAUGE"
const DSTYPE_COUNTER = "COUNTER"

type Counter struct {
	MonitorTags string `json:"monitorTags"`
	Step        int    `json:"step"`
	DStype      string `json:"dstype"`
}
