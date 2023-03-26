package api

type Comment struct {
	ID       int    // id комментария
	NewsID   int    // id новости
	ParentID int    // id родительского комментария
	Msg      string // тело комментария
	PubTime  int64  // время публикации
}
