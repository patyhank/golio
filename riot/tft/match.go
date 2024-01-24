package tft

import (
	"fmt"
	"github.com/KnutZuidema/golio/api"

	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/internal"
)

// MatchClient provides methods for the match endpoints of the VALORANT API.
type MatchClient struct {
	c *internal.Client
}

// GetMatchByID returns information about a match using match id
func (cc *MatchClient) GetMatchByID(matchID string) (*Match, error) {
	logger := cc.logger().WithField("method", "GetMatchByID")
	url := endpointGetMatch
	var match *Match
	if err := cc.c.GetInto(fmt.Sprintf(url, matchID), &match); err != nil {
		logger.Debug(err)
		fmt.Println(err)
		return nil, err
	}
	return match, nil
}

// List returns  a list of match ids by puuid
func (cc *MatchClient) List(puuid string, start, count int) (
	[]string, error,
) {
	logger := cc.logger().WithField("method", "List")
	c := *cc.c                                         // copy client
	c.Region = api.Region(api.RegionToRoute[c.Region]) // Match v5 uses a route instead of a region
	var matches []string
	endpoint := fmt.Sprintf(endpointGetMatchIDs, puuid, start, count)

	if err := c.GetInto(endpoint, &matches); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return matches, nil
}

func (cc *MatchClient) logger() log.FieldLogger {
	return cc.c.Logger().WithField("category", "match")
}
