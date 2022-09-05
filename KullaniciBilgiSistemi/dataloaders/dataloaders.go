package dataloaders

import (
	"encoding/json"

	. "github.com/dogukan9/myProject/models"
	util "github.com/dogukan9/myProject/utils"
)

func LoadUsers() []User {

	bytes, _ := util.ReadFile("../json/users.json")
	var data []User
	json.Unmarshal([]byte(bytes), &data)

	return data
}

func LoadInterest() []Interest {

	bytes, _ := util.ReadFile("../json/interests.json")
	var data []Interest
	json.Unmarshal([]byte(bytes), &data)

	return data
}

func LoadInterestMapping() []InterestMapping {

	bytes, _ := util.ReadFile("../json/userInterestMapping.json")
	var data []InterestMapping
	json.Unmarshal([]byte(bytes), &data)

	return data
}
