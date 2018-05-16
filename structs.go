package socialblade

// Client is the authed data for the client
type Client struct {
	user  bool
	email string
	token string
	key   string
}

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
		Result   int    `json:"result"`
		Userid   string `json:"userid"`
		Settings struct {
			Currency      string `json:"currency"`
			Ads           string `json:"ads"`
			DisplayName   string `json:"display_name"`
			EmailProgress string `json:"email_progress"`
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
			Success string `json:"success"`
			Failure string `json:"failure"`
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

// Section is the allowed sections for Favourites
type section string

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
