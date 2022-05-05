package recommendation

type R12n struct {
	UserId     int64   `json:"user_id"`
	ActivityId int64   `json:"activity_id"`
	Rank       float32 `json:"rank"`
	Shown      bool    `json:"shown"`
}
