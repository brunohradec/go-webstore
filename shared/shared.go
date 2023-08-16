package shared

import (
	"github.com/brunohradec/go-webstore/initializers"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Env *initializers.Env
