package infrastructure

import (
	"api/entity"
	"database/sql"
)

type BlogPostgresRepository struct {
	DB *sql.DB
}

func (r *BlogPostgresRepository) GetAll() ([]entity.Blog, error) {
	query := "SELECT id, title, content, author, created_at FROM blogs"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []entity.Blog
	for rows.Next() {
		var blog entity.Blog
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.CreatedAt); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}
	return blogs, nil
}

func (r *BlogPostgresRepository) GetByID(id int) (*entity.Blog, error) {
	query := "SELECT id, title, content, author, created_at FROM blogs WHERE id = $1"
	row := r.DB.QueryRow(query, id)

	var blog entity.Blog
	if err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.CreatedAt); err != nil {
		return nil, err
	}
	return &blog, nil
}

func (r *BlogPostgresRepository) Create(blog entity.Blog) error {
	query := "INSERT INTO blogs (title, content, author, created_at) VALUES ($1, $2, $3, $4)"
	_, err := r.DB.Exec(query, blog.Title, blog.Content, blog.Author, blog.CreatedAt)
	return err
}

func (r *BlogPostgresRepository) Delete(id int) error {
	query := "DELETE FROM blogs WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}
