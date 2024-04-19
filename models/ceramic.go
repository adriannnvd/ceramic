package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Name int

func (Name) Good() Name {
	return 1
}
func (Name) Defected() Name {
	return 2
}
// Ceranuc Image Model !
type Ceramic struct {
	uadmin.Model
	Name Name `uadmin:"display_name:Classification"`
	Image string `uadmin:"image"`
	Date time.Time `uadmin:"display_name:Date Uploaded"`
}
