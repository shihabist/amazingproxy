package repository

import (
	"encoding/json"
	"io/ioutil"

	"proxy/model"
)

const (
	fileName = "allowedList.json"
)

type fileRepo struct{}

func NewFileRepository() ProxyRepository {
	return &fileRepo{}
}

func (*fileRepo) GetAllowedUriList() ([]model.AllowedUri, error) {
	allowedUriList := loadAllowedListFile().AllowedList
	return allowedUriList, nil
}
func (*fileRepo) GetRulesList() (map[string]string, error) {
	rulesList := loadAllowedListFile().Rules

	rules := make(map[string]string)
	rules["id"] = rulesList.ID
	rules["num"] = rulesList.Num

	return rules, nil
}

func loadAllowedListFile() model.AllowedListData {
	// Read config file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	// json data
	var allowedListData model.AllowedListData

	err = json.Unmarshal(data, &allowedListData)
	if err != nil {
		panic(err)
	}

	return allowedListData
}
