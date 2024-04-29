package models

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/uadmin/uadmin"
)

type Classification int

// func (Classification) Good() Classification {
// 	return 1
// }
// func (Classification) Defected() Classification {
// 	return 2
// }

const (
	Good =  iota + 1
	Defected
)

// Ceramic Image Model !
type Ceramic struct {
	uadmin.Model
	Name string
	ClassificationNum Classification `uadmin:"list_exclude"`
	Classification string `uadmin:"display_name:Classification"`
	Image string `uadmin:"image"`
	Date time.Time `uadmin:"display_name:Date Uploaded"`
}

func (i *Ceramic) Save() {
	images := Ceramic{}
	baseCount := uadmin.Count(&images, "id > 0")
	recentCount := baseCount + 1

	// randomNum := rand.Intn(10000)

	// newName := "Plate" + strconv.Itoa(recentCount) + "-" + strconv.Itoa(randomNum)
	newName := "Plate" + strconv.Itoa(recentCount) 

	i.Name = newName

	// Generate a random number between 1 and 2 (inclusive)
	randomInt := rand.Intn(2) + 1

	// Set classification based on the random number
	if randomInt == 1 {
		i.ClassificationNum = Good
		i.Classification = "Good"
	} else {
		i.ClassificationNum = Defected
		i.Classification = "Defected"
	}

	i.Date = time.Now()

	// Save the model
	uadmin.Save(i)

}
