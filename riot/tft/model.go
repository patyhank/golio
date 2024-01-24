package tft

type Match struct {
	Metadata struct {
		DataVersion  string   `json:"data_version"`
		MatchId      string   `json:"match_id"`
		Participants []string `json:"participants"`
	} `json:"metadata"`
	Info struct {
		GameDatetime   int64          `json:"game_datetime"`
		GameLength     float64        `json:"game_length"`
		GameVersion    string         `json:"game_version"`
		Participants   []Participants `json:"participants"`
		QueueId        int            `json:"queue_id"`
		TftGameType    string         `json:"tft_game_type"`
		TftSetCoreName string         `json:"tft_set_core_name"`
		TftSetNumber   int            `json:"tft_set_number"`
	} `json:"info"`
}

type Participants struct {
	Augments             []string  `json:"augments"`
	Companion            Companion `json:"companion"`
	GoldLeft             int       `json:"gold_left"`
	LastRound            int       `json:"last_round"`
	Level                int       `json:"level"`
	Placement            int       `json:"placement"`
	PlayersEliminated    int       `json:"players_eliminated"`
	Puuid                string    `json:"puuid"`
	TimeEliminated       float64   `json:"time_eliminated"`
	TotalDamageToPlayers int       `json:"total_damage_to_players"`
	Traits               []Trait   `json:"traits"`
	Units                []Unit    `json:"units"`
}

type Companion struct {
	ContentID string `json:"content_ID"`
	ItemID    int    `json:"item_ID"`
	SkinID    int    `json:"skin_ID"`
	Species   string `json:"species"`
}

type Trait struct {
	Name        string `json:"name"`
	NumUnits    int    `json:"num_units"`
	Style       int    `json:"style"`
	TierCurrent int    `json:"tier_current"`
	TierTotal   int    `json:"tier_total"`
}

type Unit struct {
	CharacterId string   `json:"character_id"`
	ItemNames   []string `json:"itemNames"`
	Name        string   `json:"name"`
	Rarity      int      `json:"rarity"`
	Tier        int      `json:"tier"`
}

// Summoner represents a summoner with several related IDs
type Summoner struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	PUUID         string `json:"puuid"`
	SummonerLevel int    `json:"summonerLevel"`
	RevisionDate  int    `json:"revisionDate"`
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
}
