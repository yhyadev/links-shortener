package database

import "github.com/kamva/mgm/v3"

type Link struct {
	mgm.DefaultModel `bson:",inline"`
	Slug             string `json:"slug"`
	Redirect         string `json:"redirect"`
}
