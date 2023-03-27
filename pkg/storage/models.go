package storage

type Comment struct {
	ID       int    `json:"id"`
	NewsID   int    `json:"news_id"`
	ParentID int    `json:"parent_id"`
	Msg      string `json:"msg"`
	PubTime  int64  `json:"pub_time"`
}
