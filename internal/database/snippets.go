package database

import (
	"context"

	"github.com/aadi-1024/identikit-backend/internal/models"
)

func (d *Database) GetAllSnippets(ctx context.Context, conds ...string) ([]models.Snippet, error) {
	data := make([]models.Snippet, 0)

	res := d.conn.WithContext(ctx).Model(&models.Snippet{})
	for i := 0; i < len(conds); {
		res.Where(conds[i], conds[i+1])
		i += 2
	}
	res.Find(&data)
	return data, res.Error
}

func (d *Database) CreateSnippet(ctx context.Context, snippet models.Snippet) error {
	return d.conn.Create(&snippet).Error
}
