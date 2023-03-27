package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lekht/news-comments-service/config"
	"github.com/lekht/news-comments-service/pkg/storage"
)

type Postgres struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, cfg *config.PG) (*Postgres, error) {
	var connstr string = fmt.Sprintf("postgres://%s:%s@%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.DB)
	dbpool, err := pgxpool.New(ctx, connstr)
	if err != nil {
		return nil, err
	}
	db := Postgres{
		pool: dbpool,
	}
	return &db, nil
}

func (p *Postgres) Close() {
	if p.pool != nil {
		p.pool.Close()
	}
}

func (p *Postgres) CommentsByNewsID(id int) ([]*storage.Comment, error) {
	var comments []*storage.Comment
	rows, err := p.pool.Query(context.Background(), `
		SELECT 
			id,
			news_id,
			parent_id,
			msg,
			pubTime
		FROM comments.messages
		WHERE news_id = $1
		ORDER BY pubTime DESC;
		`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment storage.Comment

		err = rows.Scan(&comment.ID, &comment.NewsID, &comment.ParentID, &comment.Msg, &comment.PubTime)
		if err != nil {
			return nil, err
		}

		comments = append(comments, &comment)
	}

	if err = rows.Err(); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return comments, nil

}

func (p *Postgres) AddComment(c *storage.Comment) error {
	_, err := p.pool.Exec(context.Background(), `
		INSERT INTO comments.messages (news_id, parent_id, msg, pubTime)
		VALUES ($1, $2, $3, $4);`,
		c.NewsID,
		c.ParentID,
		c.Msg,
		c.PubTime,
	)
	if err != nil {
		return err
	}

	return nil
}
