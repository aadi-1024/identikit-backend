package database

import (
	"context"

	"github.com/aadi-1024/identikit-backend/internal/models"
)

func (d *Database) GetAllSnippets(ctx context.Context, conds ...string) ([]models.Snippet, error) {
	ctx, cancel := d.context(ctx)
	defer cancel()

	data := make([]models.Snippet, 0)

	res := d.conn.WithContext(ctx).Model(&models.Snippet{})
	for i := 0; i < len(conds); {
		res.Where(conds[i], conds[i+1])
		i += 2
	}
	res.Find(&data)
	return data, res.Error
}

func (d *Database) GetSnippetById(ctx context.Context, id string) (models.Snippet, error) {
	ctx, cancel := d.context(ctx)
	defer cancel()

	ret := models.Snippet{}
	err := d.conn.WithContext(ctx).Table("snippets").Find(&ret).Where("id = ?", id).Error
	return ret, err
}

func (d *Database) UpdateSnippet(ctx context.Context, snippet models.Snippet) error {
	ctx, cancel := d.context(ctx)
	defer cancel()

	return d.conn.WithContext(ctx).Model(&snippet).UpdateColumns(snippet).Error
}

func (d *Database) CreateSnippet(ctx context.Context, snippet models.Snippet) error {
	ctx, cancel := d.context(ctx)
	defer cancel()

	return d.conn.WithContext(ctx).Create(&snippet).Error
}
