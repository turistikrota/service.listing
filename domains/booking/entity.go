package booking

import "github.com/cilloparch/cillop/i18np"

type People struct {
	Adult int `json:"adult"`
	Kid   int `json:"kid"`
	Baby  int `json:"baby"`
}

type ValidationError struct {
	Field   string  `json:"field"`
	Message string  `json:"message"`
	Params  i18np.P `json:"params"`
}
