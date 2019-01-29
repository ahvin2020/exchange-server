package model

import (
	"github.com/gocraft/dbr"
)

var GenderModel Gender

type Gender struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

func (GenderModel Gender) GetAll(dbSess *dbr.Session) ([]*Gender, error) {
	if cacheData.genderList == nil {
		cacheData.genderList = []*Gender{
			{Id: "M", Name: "Male"}, 
			{Id: "F", Name: "Female"}, 
			{Id: "O", Name: "Others"},
		}
	}

	return cacheData.genderList, nil
}
