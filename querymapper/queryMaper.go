package querymapper

import (
	"fmt"
	"strings"
)

// QueryMap maps required fields of unprocessed query data to the form of organized map
type QueryMap struct {
	err         error
	mappedData  map[string]interface{}
	unprocessed map[string][]string
	reqFields   string
	listFields  string
}

// NewQueryMap builds new query mapping obiect
func NewQueryMap(unproc map[string][]string, reqfields, listfields string) QueryMap {
	mymap := QueryMap{}
	mymap.err = nil
	mymap.unprocessed = unproc
	mymap.mappedData = make(map[string]interface{})
	mymap.reqFields = reqfields
	mymap.listFields = listfields
	return mymap
}

// ValidatePassword enables password or token validation (what should be in querry under psswKey value key)
// and returning customised unauthenthicated message.
func (mymap *QueryMap) ValidatePassword(psswKey, errMessage string, passwords []string) *QueryMap {
	var pssw string
	switch len(mymap.unprocessed[psswKey]) {
	case 0:
		mymap.err = fmt.Errorf("Missing field %s required", psswKey)
		return mymap
	case 1:
		pssw = mymap.unprocessed[psswKey][0]
	default:
		mymap.err = fmt.Errorf("Data error when accessing %s", psswKey)
		return mymap
	}
	for _, password := range passwords {
		if pssw == password {
			return mymap
		}
	}
	mymap.err = fmt.Errorf("%s", errMessage)
	return mymap
}

// ? why MapFields breaks when I return string instead of interface{}
func returnSingle(value string) interface{} {
	return value
}

// ? why MapListFields breaks when I return []string instead of interface{}
func returnList(value string) interface{} {
	return strings.Split(value, ",")
}

func (mymap *QueryMap) mapData(myfunc func(string) interface{}, fieldNames ...string) *QueryMap {
	if mymap.err != nil {
		return mymap
	}
	for _, fieldName := range fieldNames {
		switch len(mymap.unprocessed[fieldName]) {
		case 0:
			mymap.err = fmt.Errorf("Missing field %s required", fieldName)
			return mymap
		case 1:
			mymap.mappedData[fieldName] = myfunc(mymap.unprocessed[fieldName][0])
		default:
			mymap.err = fmt.Errorf("Data error when accessing %s", fieldName)
			return mymap
		}
	}
	return mymap
}

// MapFields allows ensuring that required fields do exist in the query and organise them into map
func (mymap *QueryMap) MapFields(fieldNames ...string) *QueryMap {
	return mymap.mapData(returnSingle, fieldNames...)
}

// MapListFields allows ensuring that required fields do exist in the query and subtract their values in a list form into map
func (mymap *QueryMap) MapListFields(fieldNames ...string) *QueryMap {
	return mymap.mapData(returnList, fieldNames...)
}
