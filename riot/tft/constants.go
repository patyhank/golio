package tft

const (
	endpointBase             = "/tft"
	endpointStatusBase       = endpointBase + "/status/v3"
	endpointGetStatus        = endpointStatusBase + "/shard-data"
	endpointMatchBase        = endpointBase + "/match/v5"
	endpointGetMatchIDsBase  = endpointMatchBase + "/matches/by-puuid"
	endpointGetMatchIDs      = endpointGetMatchIDsBase + "/%s/ids?start=%d&count=%d"
	endpointGetMatch         = endpointMatchBase + "/matches/%s"
	endpointGetMatchTimeline = endpointMatchBase + "/matches/%s/timeline"
)
