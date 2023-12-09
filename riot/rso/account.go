package rso

import (
	"fmt"
	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/internal"
)

// AccountClient provides methods for the account endpoints of the Riot Accounts RSO API.
type AccountClient struct {
	c *internal.Client
}

// GetByPUUID retrieves a Riot RSO account by its puuid
func (a *AccountClient) GetByPUUID(puuid string) (*Account, error) {
	logger := a.c.Logger().WithField("method", "GetByPUUID")

	endpoint := fmt.Sprintf(endpointAccountByPUUID, puuid)
	var account *Account
	if err := a.c.GetInto(endpoint, &account); err != nil {
		logger.Debug(err)
		return nil, err
	}

	return account, nil
}

// GetByRiotID retrieves a Riot RSO account by its gameName and tagLine
func (a *AccountClient) GetByRiotID(gameName string, tagLine string) (*Account, error) {
	logger := a.c.Logger().WithField("method", "GetByRiotID")

	endpoint := fmt.Sprintf(endpointAccountByRiotID, gameName, tagLine)
	var account *Account
	if err := a.c.GetInto(endpoint, &account); err != nil {
		logger.Debug(err)
		return nil, err
	}

	return nil, nil
}

func (a *AccountClient) logger() log.FieldLogger {
	return a.c.Logger().WithField("category", "account")
}
