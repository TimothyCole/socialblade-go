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
