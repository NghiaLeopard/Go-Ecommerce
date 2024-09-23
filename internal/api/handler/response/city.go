package IResponse

import "time"

type City struct {
	Id       int64  `json:"_id"`
	Name     string `json:"name"`
	CreateAt time.Time
}
