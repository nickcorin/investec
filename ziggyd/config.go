package ziggyd

// Config contains the various configuration settings for ziggy.
type Config struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	PrivacyMask  bool   `json:"privacyMask"`
}
