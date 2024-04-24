package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Classification int

func (Classification) Good() Classification {
	return 1
}
func (Classification) Defected() Classification {
	return 2
}
// Ceranuc Image Model !
type Ceramic struct {
	uadmin.Model
	Name string
	Classification Classification `uadmin:"display_name:Classification"`
	Image string `uadmin:"image"`
	Date time.Time `uadmin:"display_name:Date Uploaded"`
}
