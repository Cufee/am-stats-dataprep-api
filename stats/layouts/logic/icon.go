package logic

import (
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

type Icon struct {
	GetName  func(Values) string      `json:"-"`
	GetStyle func(Values) style.Style `json:"-"`
}
