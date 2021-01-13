package schema

type SSHKey struct {
	Uuid        string `json:"pubKeyUuid"`
	Name        string `json:"name"`
	Key         string `json:"key"`
	Fingerprint string `json:"fingerprint"`
	CreateTime  int64  `json:"createTime"`
}
