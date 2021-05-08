# GORM

GORM V2 moved to https://github.com/go-gorm/gorm

GORM V1 Doc https://v1.gorm.io/

Edit

type Model struct {
ID int64 `gorm:"primary_key"`
CreatedAt int64
UpdatedAt int64
DeletedAt \*int64 `sql:"index"`
}
