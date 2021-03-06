package twitter

type PlaceType struct {
	Name string `json:"name"`
	Code int    `json:"code"`
}

type TrendLocation struct {
	Name        string    `json:"name"`
	CountryCode string    `json:"countryCode"`
	URL         string    `json:"url"`
	Woeid       int       `json:"woeid"`
	PlaceType   PlaceType `json:"placeType"`
	Parentid    int       `json:"parentid"`
	Country     string    `json:"country"`
}

type TimeZone struct {
	Name       string `json:"name"`
	UtcOffset  int    `json:"utc_offset"`
	TzinfoName string `json:"tzinfo_name"`
}

type SleepTime struct {
	Enabled   bool        `json:"enabled"`
	EndTime   interface{} `json:"end_time"`
	StartTime interface{} `json:"start_time"`
}

// Account hold information about account
type Account struct {
	TimeZone                  TimeZone        `json:"time_zone"`
	Protected                 bool            `json:"protected"`
	ScreenName                string          `json:"screen_name"`
	AlwaysUseHTTPS            bool            `json:"always_use_https"`
	UseCookiePersonalization  bool            `json:"use_cookie_personalization"`
	SleepTime                 SleepTime       `json:"sleep_time"`
	GeoEnabled                bool            `json:"geo_enabled"`
	Language                  string          `json:"language"`
	DiscoverableByEmail       bool            `json:"discoverable_by_email"`
	DiscoverableByMobilePhone bool            `json:"discoverable_by_mobile_phone"`
	DisplaySensitiveMedia     bool            `json:"display_sensitive_media"`
	AllowContributorRequest   string          `json:"allow_contributor_request"`
	AllowDmsFrom              string          `json:"allow_dms_from"`
	AllowDmGroupsFrom         string          `json:"allow_dm_groups_from"`
	SmartMute                 bool            `json:"smart_mute"`
	TrendLocation             []TrendLocation `json:"trend_location"`
}

// User hold information about user
type User struct {
	Name            string `json:"name"`
	ScreenName      string `json:"screen_name"`
	FollowersCount  int    `json:"followers_count"`
	ProfileImageURL string `json:"profile_image_url"`
}

type Place struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
}

type Entities struct {
	HashTags []struct {
		Indices [2]int `json:"indices"`
		Text    string `json:"text"`
	}
	UserMentions []struct {
		Indices    [2]int `json:"indices"`
		ScreenName string `json:"screen_name"`
	} `json:"user_mentions"`
	Urls []struct {
		Indices [2]int `json:"indices"`
		URL     string `json:"url"`
	} `json:"urls"`
}

// Tweet hold information about tweet
type Tweet struct {
	ID        string   `json:"id_str"`
	Text      string   `json:"text"`
	FullText  string   `json:"full_text,omitempty"`
	Source    string   `json:"source"`
	CreatedAt string   `json:"created_at"`
	User      User     `json:"user"`
	Place     *Place   `json:"place"`
	Entities  Entities `json:"entities"`

	InReplyToStatusID string `json:"in_reply_to_status_id_str"`

	RetweetCount  int  `json:"retweet_count"`
	Retweeted     bool `json:"retweeted"`
	FavoriteCount int  `json:"favorite_count"`
	Favorited     bool `json:"favorited"`
}

// SearchMetadata hold information about search metadata
type SearchMetadata struct {
	CompletedIn float64 `json:"completed_in"`
	MaxID       int64   `json:"max_id"`
	MaxIDStr    string  `json:"max_id_str"`
	NextResults string  `json:"next_results"`
	Query       string  `json:"query"`
	RefreshURL  string  `json:"refresh_url"`
	Count       int     `json:"count"`
	SinceID     int     `json:"since_id"`
	SinceIDStr  string  `json:"since_id_str"`
}

type UploadMediaResponse struct {
	MediaID          int64  `json:"media_id"`
	MediaIDString    string `json:"media_id_string"`
	Size             int    `json:"size"`
	ExpiresAfterSecs int    `json:"expires_after_secs"`
	Image            struct {
		ImageType string `json:"image_type"`
		W         int    `json:"w"`
		H         int    `json:"h"`
	} `json:"image"`
}

type SearchTweetsResponse struct {
	Statuses       []Tweet        `json:"statuses"`
	SearchMetadata SearchMetadata `json:"search_metadata"`
}
