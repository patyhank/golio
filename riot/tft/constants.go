package tft

const (
	endpointBase                    = "/tft"
	endpointMatchBase               = endpointBase + "/match/v1"
	endpointGetMatchIDsBase         = endpointMatchBase + "/matches/by-puuid"
	endpointGetMatchIDs             = endpointGetMatchIDsBase + "/%s/ids?start=%d&count=%d"
	endpointGetMatch                = endpointMatchBase + "/matches/%s"
	endpointSummonerBase            = endpointBase + "/summoner/v1"
	endpointGetSummonerBySummonerID = endpointSummonerBase + "/summoners/%s"
	endpointGetSummonerBy           = endpointSummonerBase + "/summoners/by-%s/%s"
)
