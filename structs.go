package socialblade

// Client is the authed data for the client
type Client struct {
	user  bool
	email string
	token string
	key   string
	api   string
}

// section is the allowed sections for Favourites
type section string

type tpAuth struct {
	Status struct {
		Response int    `json:"response"`
		Type     string `json:"type"`
		Error    string `json:"error"`
	} `json:"status"`
	Access struct {
		NAME string `json:"NAME,omitempty"`
		KEY  string `json:"KEY,omitempty"`
	} `json:"access,omitempty"`
}

type userAuth struct {
	Status struct {
		Response int    `json:"response"`
		Error    bool   `json:"error,omitempty"`
		Message  string `json:"message,omitempty"`
	} `json:"status"`
	ID struct {
		Result   int `json:"result"`
		Userid   int `json:"userid"`
		Settings struct {
			Currency      int    `json:"currency"`
			Ads           int    `json:"ads"`
			DisplayName   string `json:"display_name"`
			EmailProgress int    `json:"email_progress"`
		} `json:"settings"`
		SocialBladePartner bool   `json:"socialblade_partner"`
		Email              string `json:"email"`
		Token              string `json:"token"`
		Limits             struct {
			Hour struct {
				Current string `json:"current"`
				Limit   int    `json:"limit"`
			} `json:"hour"`
			Day struct {
				Current string `json:"current"`
				Limit   int    `json:"limit"`
			} `json:"day"`
			Success interface{} `json:"success"`
			Failure interface{} `json:"failure"`
		} `json:"limits"`
	} `json:"id"`
	Premium struct {
		Bronze struct {
			Active    bool   `json:"active"`
			Expires   string `json:"expires,omitempty"`
			ExpiresAt int    `json:"expires_at,omitempty"`
		} `json:"bronze"`
		Silver struct {
			Active    bool   `json:"active"`
			Expires   string `json:"expires,omitempty"`
			ExpiresAt int    `json:"expires_at,omitempty"`
		} `json:"silver"`
		Gold struct {
			Active    bool   `json:"active"`
			Expires   string `json:"expires,omitempty"`
			ExpiresAt int    `json:"expires_at,omitempty"`
		} `json:"gold"`
		Platinum struct {
			Active    bool   `json:"active"`
			Expires   string `json:"expires,omitempty"`
			ExpiresAt int    `json:"expires_at,omitempty"`
		} `json:"platinum"`
	} `json:"premium"`
}

// Favourite is the data of a user favourite
type Favourite struct {
	Username string `json:"username,omitempty"`
	Section  string `json:"section,omitempty"`
	Data     struct {
		Followers string `json:"followers,omitempty"`
		Tweets    string `json:"tweets,omitempty"`
		Avatar    string `json:"avatar,omitempty"`
		Banner    string `json:"banner,omitempty"`
		Name      string `json:"name,omitempty"`
	} `json:"data,omitempty"`
	Growth []struct {
		Followers string `json:"followers,omitempty"`
		Following string `json:"following,omitempty"`
		Tweets    string `json:"tweets,omitempty"`
		Favorites string `json:"favorites,omitempty"`
		Date      string `json:"date,omitempty"`
	} `json:"growth,omitempty"`
	Type string `json:"type,omitempty"`
}

type favouriteUpdate struct {
	Status    interface{} `json:"status"`
	Favourite bool        `json:"favorite,omitempty"`
}

// YouTubeStatsData is data for YouTube Stats
type YouTubeStatsData struct {
	Status struct {
		Response int         `json:"response"`
		Error    interface{} `json:"error"`
		Message  string      `json:"message,omitempty"`
	} `json:"status"`
	ID struct {
		Type      string `json:"type,omitempty"`
		ChannelID string `json:"channelid,omitempty"`
		Username  string `json:"username,omitempty"`
		CUsername string `json:"cusername,omitempty"`
	} `json:"id"`
	Data struct {
		Username      string `json:"username"`
		DisplayName   string `json:"displayname"`
		CreatedAt     string `json:"created_at"`
		Uploads       string `json:"uploads"`
		Subs          string `json:"subs"`
		Views         string `json:"views"`
		CountryCode   string `json:"country_code"`
		Country       string `json:"country"`
		ChannelType   string `json:"channeltype"`
		Avatar        string `json:"avatar"`
		Banner        string `json:"banner"`
		AvgDailySubs  string `json:"avgdailysubs"`
		AvgDailyViews string `json:"avgdailyviews"`
		Partner       string `json:"partner"`
		IsVerified    string `json:"isVerified"`
	} `json:"data,omitempty"`
	Rank struct {
		SBRank          string `json:"sbrank"`
		Rank            string `json:"rank"`
		ViewsRank       string `json:"viewsrank"`
		CountryRank     string `json:"countryrank"`
		ChannelTypeRank string `json:"channeltyperank"`
		Raw             struct {
			Sbrank          string `json:"sbrank"`
			Rank            string `json:"rank"`
			Viewsrank       string `json:"viewsrank"`
			Countryrank     string `json:"countryrank"`
			Channeltyperank string `json:"channeltyperank"`
		} `json:"raw"`
		GradeRaw string `json:"grade_raw"`
		Grade    string `json:"grade"`
	} `json:"rank,omitempty"`
	Charts struct {
		Subs struct {
			Subs14  string `json:"subs14"`
			Subs30  string `json:"subs30"`
			Subs60  string `json:"subs60"`
			Subs90  string `json:"subs90"`
			Subs180 string `json:"subs180"`
			Subs365 string `json:"subs365"`
		} `json:"subs"`
		Views struct {
			Views14  string `json:"views14"`
			Views30  string `json:"views30"`
			Views60  string `json:"views60"`
			Views90  string `json:"views90"`
			Views180 string `json:"views180"`
			Views365 string `json:"views365"`
		} `json:"views"`
		Growth struct {
			Subs  string `json:"subs"`
			Views string `json:"views"`
		} `json:"growth"`
	} `json:"charts,omitempty"`
	Social struct {
		Googleplus string `json:"googleplus,omitempty"`
		Facebook   string `json:"facebook,omitempty"`
		Twitter    string `json:"twitter,omitempty"`
		Instagram  string `json:"instagram,omitempty"`
	} `json:"social,omitempty"`
	DataDaily []struct {
		Date  string `json:"date"`
		Subs  string `json:"subs"`
		Views string `json:"views"`
	} `json:"data_daily,omitempty"`
}

// YouTubeTopList is used for youtube top lists
type YouTubeTopList struct {
	Status struct {
		Cache    string `json:"cache"`
		Type     string `json:"type"`
		Response int    `json:"response"`
	} `json:"status"`
	Result []struct {
		Username          string `json:"username"`
		Channelid         string `json:"channelid"`
		Cusername         string `json:"cusername"`
		Displayname       string `json:"displayname"`
		Subscribers       string `json:"subscribers"`
		Vidviews          string `json:"vidviews"`
		Uploads           string `json:"uploads"`
		CreatedAt         string `json:"created_at"`
		Channeltype       string `json:"channeltype"`
		Country           string `json:"country"`
		Avgdailyviews     string `json:"avgdailyviews"`
		Views30           string `json:"views30"`
		Views365          string `json:"views365"`
		Sbrank            string `json:"sbrank"`
		RankSubscribers   int    `json:"rank_subscribers"`
		RankViews         string `json:"rank_views"`
		EstimatedEarnings struct {
			Daily struct {
				Low  float32 `json:"low"`
				High float32 `json:"high"`
			} `json:"daily"`
			Monthly struct {
				Low  float32 `json:"low"`
				High float32 `json:"high"`
			} `json:"monthly"`
			Yearly struct {
				Low  float32 `json:"low"`
				High float32 `json:"high"`
			} `json:"yearly"`
		} `json:"estimated_earnings"`
	} `json:"result"`
}

// TwitchStatsData is data for YouTube Stats
type TwitchStatsData struct {
	Status struct {
		Response int         `json:"response"`
		Error    interface{} `json:"error"`
		Message  string      `json:"message"`
	} `json:"status"`
	ID struct {
		Results  int    `json:"results"`
		TwitchID int    `json:"twitch_id"`
		Username string `json:"username"`
		Mod      string `json:"mod"`
	} `json:"id"`
	Data struct {
		Username           string `json:"username"`
		Game               string `json:"game"`
		Status             string `json:"status"`
		Avatar             string `json:"avatar"`
		Banner             string `json:"banner"`
		Followers          int    `json:"followers"`
		Followersdaygain   int    `json:"followersdaygain"`
		Followersmonthgain int    `json:"followersmonthgain"`
		Avgdailyfollowers  int    `json:"avgdailyfollowers"`
		Avgdailyviews      int    `json:"avgdailyviews"`
		Activity           int    `json:"activity"`
		Views              int    `json:"views"`
		Viewsdaygain       int    `json:"viewsdaygain"`
		Viewsmonthgain     int    `json:"viewsmonthgain"`
		CreatedAt          string `json:"created_at"`
		IsVerified         int    `json:"isVerified"`
	} `json:"data"`
	Rank struct {
		Rank        interface{} `json:"rank"`
		Vidviewrank interface{} `json:"vidviewrank"`
		Grade       struct {
			Raw     interface{} `json:"raw"`
			Grade   string      `json:"grade"`
			Display string      `json:"display"`
		} `json:"grade"`
	} `json:"rank"`
	DataDaily []struct {
		Date      string `json:"date"`
		Followers int    `json:"followers"`
		Views     int    `json:"views"`
	} `json:"data_daily"`
}

// TwitchTopList is used for youtube top lists
type TwitchTopList struct {
	Status struct {
		Cache    string `json:"cache"`
		Type     string `json:"type"`
		Response int    `json:"response"`
	} `json:"status"`
	Result []struct {
		Username          string      `json:"username"`
		Displayname       interface{} `json:"displayname"`
		Followers         string      `json:"followers"`
		Views             string      `json:"views"`
		Avgdailyfollowers string      `json:"avgdailyfollowers"`
		Avgdailyviews     string      `json:"avgdailyviews"`
		CreatedAt         string      `json:"created_at"`
		Game              string      `json:"game"`
		Status            string      `json:"status"`
		// rank_followers and rank_channelviews currently have a glich where the filtered type (followers/views) will be an int and the other will be a string
		//  it's an interface so it to unmarshal correctly and you can still use it by check the type of it first
		RankFollowers    interface{} `json:"rank_followers"`
		RankChannelviews interface{} `json:"rank_channelviews"`
	} `json:"result"`
}

// TeamData is the struct for the Social Blade Team
type TeamData struct {
	Team []struct {
		Name    string   `json:"name"`
		Title   string   `json:"title,omitempty"`
		Image   string   `json:"image,omitempty"`
		Socials []string `json:"socials,omitempty"`
		Discord int64    `json:"discord"`
	} `json:"team"`
}
