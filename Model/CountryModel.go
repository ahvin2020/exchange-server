package model

import (
	"github.com/gocraft/dbr"
)

var CountryModel Country

type Country struct {
	Id             int16  `json:"Id"`
	Name           string `json:"Name"`
	IsoAlpha2      string `json:"IsoAlpha2,omitempty"`
	IsoAlpha3      string `json:"IsoAlpha3,omitempty"`
	IsoNumeric     int16  `json:"IsoNumeric,omitempty"`
	CurrencyCode   string `json:"CurrencyCode,omitempty"`
	CurrencyName   string `json:"CurrencyName,omitempty"`
	CurrencySymbol string `json:"CurrencySymbol,omitempty"`
}

func (CountryModel Country) GetAll(dbSess *dbr.Session) ([]*Country, error) {
	var err error = nil
	
	if cacheData.countryList == nil {
		cacheData.countryList = []*Country{}
		_, err = dbSess.SelectBySql("SELECT id, name FROM countries ORDER BY name ASC;").LoadStructs(&cacheData.countryList)
	}

	return cacheData.countryList, err
}
