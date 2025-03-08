package models

import "gorm.io/gorm"

// HINT: Higher level component!
// FIXME: should not be aware of json schema!
// HINT: models and repositories doesn't have any logic!
// HINT FOR FAR: it's an entity
// HINT: entities have ID and should get mutated!
// HINT: DDD
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}
