package post

import "github.com/cilloparch/cillop/i18np"

type RuleFunc func(e *Entity, value interface{}) *i18np.Error
