package rso

const (
	endpointBase            = "/riot"
	endpointAccountBase     = endpointBase + "/account/v1/accounts"
	endpointAccountByPUUID  = endpointAccountBase + "/by-puuid/%s"
	endpointAccountByRiotID = endpointAccountBase + "/%s/%s"
)
