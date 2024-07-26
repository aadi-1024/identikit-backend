package models

type Snippet struct {
	Id            string `json:"id" gorm:"primaryKey;type:uuid"`
	Name          string `json:"name" validate:"required"`
	Description   string `json:"description"`
	Language      string `json:"language" validate:"required"`
	Prefix        JSONB  `json:"prefix" gorm:"type:jsonb" validate:"required"`
	Body          JSONB  `json:"body" gorm:"type:jsonb" validate:"required"`
	Documentation string `json:"documentation,omitempty" gorm:"default:null"`
	Security      string `json:"security,omitempty" gorm:"default:null"`
}
