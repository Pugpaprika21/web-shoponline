package dto

// StoreSettingsRequest is the DTO for updating store settings (all at once)
type StoreSettingsRequest struct {
	StoreName       string `json:"store_name" form:"store_name"`
	StoreLogoURL    string `json:"store_logo_url" form:"store_logo_url"`
	StoreBannerURL  string `json:"store_banner_url" form:"store_banner_url"`
	PrimaryColor    string `json:"primary_color" form:"primary_color"`
	SecondaryColor  string `json:"secondary_color" form:"secondary_color"`
	FooterText      string `json:"footer_text" form:"footer_text"`
	SocialFacebook  string `json:"social_facebook" form:"social_facebook"`
	SocialLine      string `json:"social_line" form:"social_line"`
	SocialInstagram string `json:"social_instagram" form:"social_instagram"`
	ContactPhone    string `json:"contact_phone" form:"contact_phone"`
	ContactEmail    string `json:"contact_email" form:"contact_email"`
	ContactAddress  string `json:"contact_address" form:"contact_address"`
}

// StoreSettingsResponse is the DTO for returning all store settings
type StoreSettingsResponse struct {
	StoreName       string `json:"store_name"`
	StoreLogoURL    string `json:"store_logo_url"`
	StoreBannerURL  string `json:"store_banner_url"`
	PrimaryColor    string `json:"primary_color"`
	SecondaryColor  string `json:"secondary_color"`
	FooterText      string `json:"footer_text"`
	SocialFacebook  string `json:"social_facebook"`
	SocialLine      string `json:"social_line"`
	SocialInstagram string `json:"social_instagram"`
	ContactPhone    string `json:"contact_phone"`
	ContactEmail    string `json:"contact_email"`
	ContactAddress  string `json:"contact_address"`
}

// ToMap converts the request to a key-value map
func (r *StoreSettingsRequest) ToMap() map[string]string {
	return map[string]string{
		"store_name":       r.StoreName,
		"store_logo_url":   r.StoreLogoURL,
		"store_banner_url": r.StoreBannerURL,
		"primary_color":    r.PrimaryColor,
		"secondary_color":  r.SecondaryColor,
		"footer_text":      r.FooterText,
		"social_facebook":  r.SocialFacebook,
		"social_line":      r.SocialLine,
		"social_instagram": r.SocialInstagram,
		"contact_phone":    r.ContactPhone,
		"contact_email":    r.ContactEmail,
		"contact_address":  r.ContactAddress,
	}
}

// StoreSettingsFromMap creates a response from a key-value map
func StoreSettingsFromMap(m map[string]string) StoreSettingsResponse {
	return StoreSettingsResponse{
		StoreName:       m["store_name"],
		StoreLogoURL:    m["store_logo_url"],
		StoreBannerURL:  m["store_banner_url"],
		PrimaryColor:    m["primary_color"],
		SecondaryColor:  m["secondary_color"],
		FooterText:      m["footer_text"],
		SocialFacebook:  m["social_facebook"],
		SocialLine:      m["social_line"],
		SocialInstagram: m["social_instagram"],
		ContactPhone:    m["contact_phone"],
		ContactEmail:    m["contact_email"],
		ContactAddress:  m["contact_address"],
	}
}
