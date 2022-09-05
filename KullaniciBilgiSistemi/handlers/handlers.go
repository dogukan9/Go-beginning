package handlers

import (
	"encoding/json"
	"net/http"

	. "github.com/dogukan9/myProject/dataloaders"
	. "github.com/dogukan9/myProject/models"
)

func Run() {

	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(res http.ResponseWriter, req *http.Request) {

	page := Page{ID: 7, Name: "Kullanıcılar", Description: "Kullanıcı Listesi", URI: "/users"}

	users := LoadUsers()
	interests := LoadInterest()

	interestsMapping := LoadInterestMapping()

	var newUsers []User

	for _, user := range users {
		for _, interestMapping := range interestsMapping {
			if user.ID == interestMapping.UserId {
				for _, interest := range interests {
					if interestMapping.InterestId == interest.ID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)

	}

	viewModel := UserViewModel{Page: page, Users: newUsers}
	data, _ := json.Marshal(viewModel)
	res.Write([]byte(data))
}
