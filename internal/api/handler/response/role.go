package IResponse

type Role struct {
	Id         int64    `json:"_id"`
	Name       string   `json:"name"`
	Permission []string `json:"permissions"`
}
