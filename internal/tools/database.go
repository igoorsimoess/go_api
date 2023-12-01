package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	AuthToken string
	Username string
}

type CoinDetails struct {
	Coins int64
	Username string
}


// using a interface for easily swapping as long as we define the methods inside it
type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

// returns the above interface
func NewDatabase() (*DatabaseInterface, error){

	var database DatabaseInterface = &mockDB{}
	var err error = database.SetupDatabase()

	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &database, nil
}
