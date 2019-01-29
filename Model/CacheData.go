package model

var cacheData CacheModel

// this struct will cache all fixed data
type CacheModel struct {
	countryList []*Country
	genderList []*Gender
}
