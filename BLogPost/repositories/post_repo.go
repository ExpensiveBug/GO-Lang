package repositories

// sql queries

import (
	"BLogPost/models"
	"uuid"

	"github.com/jmoiron/sqlx"
)

type PostRepository struct {
	DB *sqlx.DB
}

func GetPosts(db *sqlx.DB) ([]models.Post, error) {
	query := `SELECT id, title, content, category, author, created_at FROM posts ORDER BY created_at DESC; `
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var p models.Post
		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.Category, &p.Author, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func GetPostByID(db *sqlx.DB, id uuid.UUID) (models.Post, error) {
	query := ` SELECT id, title, content, category, author, created_at FROM posts WHERE id = $1; `
	var post models.Post
	err := db.Get(&post, query, id)
	if err != nil {
		return post, err //database error
	}
	return post, nil
}

func (repo *PostRepository) CreatePost(post models.Post) error {
	query := ` INSERT INTO posts (title, category, author, content) VALUES ($1, $2, $3, $4) `
	_, err := repo.DB.Exec(query, post.Title, post.Category, post.Author, post.Content)
	return err
}

func UpdatePost(db *sqlx.DB, id uuid.UUID, title string, category string, content string) error {
	query := ` UPDATE posts SET title = $1, category = $2, content = $3 WHERE id = $4 `
	_, err := db.Exec(query, title, category, content, id)
	if err != nil {
		return err
	}
	return nil
}

func DeletePost(db *sqlx.DB, id uuid.UUID) error {
	query := ` DELETE FROM posts WHERE id = $1; `
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
