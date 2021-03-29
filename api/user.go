package api

import (
	"dungeon-comrade/daos/factory"
	"dungeon-comrade/models"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

/**
*  Register new user
**/
func OnCreateUser(userInput map[string]interface{}) {
	var user models.CreateUserRequest
	err := mapstructure.Decode(userInput, &user)
	if err != nil {
		fmt.Printf("Invalid input to 'onCreateUser' function\n")
		return
	}

	userDao := factory.UserFactoryDao("cassandra")

	createErr := userDao.Create(&user)
	if createErr != nil {
		fmt.Printf("Error: %s", err)
	}

	// TODO: send creation success back to client / empty data model or default datamodel for a new user
}

/**
* Login and fetch initial datamodel .
**/
func OnLogin(userInput map[string]interface{}) {
	var user models.CreateUserRequest
	err := mapstructure.Decode(userInput, &user)

	if err != nil {
		fmt.Printf("Invalid input to 'OnLogin' function\n")
		return
	}

	userDao := factory.UserFactoryDao("cassandra")
	fetchedUser, fetchUserErr := userDao.GetByEmail(user.Email)
	if fetchUserErr != nil {
		fmt.Printf("Fetch User Error: %s\n", fetchUserErr)
		// TODO: handle login error - return back to client...
		return
	}

	fmt.Printf("LOGGED IN: %s\n", fetchedUser.Email)

	// TODO: Fetch initial state from cassandra - user datamodel
}
