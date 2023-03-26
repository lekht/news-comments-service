package storage

type Comment struct {
	ID       int
	NewsID   int // id новости
	ParentID int // id родительского комментария
	Msg      string
	PubTime  int64
}
