package gorm

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
//    type User struct {
//      gorm.Model
//    }
type Model struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt int64
	UpdatedAt int64
	DeletedAt *int64 `gorm:"DEFAULT:0" sql:"index"`
}
