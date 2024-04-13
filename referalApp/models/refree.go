package models

import "github.com/sahilwahi7/referalApp/database"

type ConcreteRefree struct {
	RefreeDB *map[string]database.Refree //key is refreeid and value is struct
}

type Refree interface {
	SaveRefree()
	FetchRefree(name string) *database.Refree
}

// TO-DO this will be from auth serivce, will implement after auth service
func (r *ConcreteRefree) SaveRefree() {}

func (r *ConcreteRefree) FetchRefree(name string) *database.Refree {
	var refree *database.Refree
	for _, v := range *r.RefreeDB {
		//k--> refree Id
		//v --> refree struct

		if name == v.RefreeAuth.Name { //fetching refree from refree array
			refree = &v
		}
	}
	return refree
}
