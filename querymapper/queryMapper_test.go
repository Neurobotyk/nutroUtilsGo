package querymapper

import (
	"fmt"
	"os"
	"strings"
	"testing"

	internalutils "github.com/Neurobotyk/nutroUtilsGo"
	"github.com/stretchr/testify/assert"
)

func prepData() (unprocData map[string][]string, reqfieldsl, istFields []string, err error) {
	unprocData = nil
	err = internalutils.SetEnvir("../test_config.txt")
	if err != nil {
		err = fmt.Errorf("Setting environment error: %s", err)
		return
	}

	unprocData = map[string][]string{
		"name": []string{"hi"},
		"pssw": []string{os.Getenv("myPass")},
	}
	reqfields := strings.Split(os.Getenv("reqFields"), ",")
	for i, item := range reqfields {
		unprocData[item] = []string{fmt.Sprintf("value%v", i)}
	}
	listFields := strings.Split(os.Getenv("listFields"), ",")
	for i, item := range listFields {
		unprocData[item] = []string{fmt.Sprintf("value%v,2value%v,3value%v", i, i, i)}
	}
	return
}

func TestPsswValidation(t *testing.T) {
	unprocData, _, _, err := prepData()
	if err != nil {
		t.Errorf("%s", err)
	}
	customErrMsg := "Not enought pylons"
	mymap := NewQueryMap(unprocData, os.Getenv("reqFields"), os.Getenv("listFields"))
	// good
	mymap.ValidatePassword("pssw", customErrMsg, []string{os.Getenv("myPass")})
	assert.Equal(t, mymap.err, nil, "error = %v", mymap.err)

	// wrong key
	mymap.ValidatePassword("NoSuchKey", customErrMsg, []string{os.Getenv("myPass")})
	assert.Equal(t, mymap.err.Error(), "Missing field NoSuchKey required", "wrong key test failed with erroor = %v", mymap.err)

	// wrong password
	mymap.ValidatePassword("pssw", customErrMsg, []string{"WRONG"})
	assert.Equal(t, mymap.err.Error(), customErrMsg, "testing wrong  = %v", mymap.err)
}

func TestMapFields(t *testing.T) {
	unprocData, reqfields, _, err := prepData()
	if err != nil {
		t.Errorf("%s", err)
	}
	mymap := NewQueryMap(unprocData, os.Getenv("reqFields"), os.Getenv("listFields"))

	mymap.MapFields(reqfields...)
	assert.Equal(t, mymap.err, nil, "error = %v", mymap.err)
	for _, field := range reqfields {
		assert.Equal(t, mymap.mappedData[field], unprocData[field], "error = %v", mymap.err)
	}

	mymap.MapFields("noSuchField")
	assert.Equal(t, mymap.err.Error(), "Missing field noSuchField required", "error = there is no such field err value =  %v", mymap.err)
	mymap.err = nil

	// mymap.GetDataList()
	// if mymap.err != nil {
	// 	t.Errorf("error = %v", mymap.err)
	// }
}
