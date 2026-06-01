package dto

// StoreSettingsRequest is the DTO for updating store settings (all at once)
type StoreSettingsRequest struct {
	// General
	StoreName        string `json:"store_name" form:"store_name"`
	StoreDescription string `json:"store_description" form:"store_description"`
	StoreLogoURL     string `json:"store_logo_url" form:"store_logo_url"`
	StoreFaviconURL  string `json:"store_favicon_url" form:"store_favicon_url"`
	StoreBannerURL   string `json:"store_banner_url" form:"store_banner_url"`

	// Theme & Layout
	PrimaryColor    string `json:"primary_color" form:"primary_color"`
	SecondaryColor  string `json:"secondary_color" form:"secondary_color"`
	AccentColor     string `json:"accent_color" form:"accent_color"`
	FontFamily      string `json:"font_family" form:"font_family"`
	LayoutStyle     string `json:"layout_style" form:"layout_style"` // grid, list
	ProductsPerPage string `json:"products_per_page" form:"products_per_page"`

	// SEO
	MetaTitle       string `json:"meta_title" form:"meta_title"`
	MetaDescription string `json:"meta_description" form:"meta_description"`
	MetaKeywords    string `json:"meta_keywords" form:"meta_keywords"`
	GoogleAnalytics string `json:"google_analytics" form:"google_analytics"`

	// Social Media
	SocialFacebook  string `json:"social_facebook" form:"social_facebook"`
	SocialLine      string `json:"social_line" form:"social_line"`
	SocialInstagram string `json:"social_instagram" form:"social_instagram"`
	SocialTwitter   string `json:"social_twitter" form:"social_twitter"`
	SocialYoutube   string `json:"social_youtube" form:"social_youtube"`
	SocialTiktok    string `json:"social_tiktok" form:"social_tiktok"`

	// Contact
	ContactPhone   string `json:"contact_phone" form:"contact_phone"`
	ContactEmail   string `json:"contact_email" form:"contact_email"`
	ContactAddress string `json:"contact_address" form:"contact_address"`
	ContactMapURL  string `json:"contact_map_url" form:"contact_map_url"`
	BusinessHours  string `json:"business_hours" form:"business_hours"`

	// Shipping
	ShippingEnabled     string `json:"shipping_enabled" form:"shipping_enabled"`
	ShippingFreeMinimum string `json:"shipping_free_minimum" form:"shipping_free_minimum"`
	ShippingFlatRate    string `json:"shipping_flat_rate" form:"shipping_flat_rate"`
	ShippingNote        string `json:"shipping_note" form:"shipping_note"`

	// Payment
	PaymentCOD         string `json:"payment_cod" form:"payment_cod"`
	PaymentTransfer    string `json:"payment_transfer" form:"payment_transfer"`
	PaymentBankName    string `json:"payment_bank_name" form:"payment_bank_name"`
	PaymentBankAccount string `json:"payment_bank_account" form:"payment_bank_account"`
	PaymentPromptpay   string `json:"payment_promptpay" form:"payment_promptpay"`
	PaymentNote        string `json:"payment_note" form:"payment_note"`

	// Notification
	NotifyOrderEmail  string `json:"notify_order_email" form:"notify_order_email"`
	NotifyLineToken   string `json:"notify_line_token" form:"notify_line_token"`
	NotifyLowStock    string `json:"notify_low_stock" form:"notify_low_stock"`
	LowStockThreshold string `json:"low_stock_threshold" form:"low_stock_threshold"`

	// Footer
	FooterText      string `json:"footer_text" form:"footer_text"`
	FooterAbout     string `json:"footer_about" form:"footer_about"`
	FooterCopyright string `json:"footer_copyright" form:"footer_copyright"`
}

// StoreSettingsResponse is the DTO for returning all store settings
type StoreSettingsResponse struct {
	// General
	StoreName        string `json:"store_name"`
	StoreDescription string `json:"store_description"`
	StoreLogoURL     string `json:"store_logo_url"`
	StoreFaviconURL  string `json:"store_favicon_url"`
	StoreBannerURL   string `json:"store_banner_url"`

	// Theme & Layout
	PrimaryColor    string `json:"primary_color"`
	SecondaryColor  string `json:"secondary_color"`
	AccentColor     string `json:"accent_color"`
	FontFamily      string `json:"font_family"`
	LayoutStyle     string `json:"layout_style"`
	ProductsPerPage string `json:"products_per_page"`

	// SEO
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	MetaKeywords    string `json:"meta_keywords"`
	GoogleAnalytics string `json:"google_analytics"`

	// Social Media
	SocialFacebook  string `json:"social_facebook"`
	SocialLine      string `json:"social_line"`
	SocialInstagram string `json:"social_instagram"`
	SocialTwitter   string `json:"social_twitter"`
	SocialYoutube   string `json:"social_youtube"`
	SocialTiktok    string `json:"social_tiktok"`

	// Contact
	ContactPhone   string `json:"contact_phone"`
	ContactEmail   string `json:"contact_email"`
	ContactAddress string `json:"contact_address"`
	ContactMapURL  string `json:"contact_map_url"`
	BusinessHours  string `json:"business_hours"`

	// Shipping
	ShippingEnabled     string `json:"shipping_enabled"`
	ShippingFreeMinimum string `json:"shipping_free_minimum"`
	ShippingFlatRate    string `json:"shipping_flat_rate"`
	ShippingNote        string `json:"shipping_note"`

	// Payment
	PaymentCOD         string `json:"payment_cod"`
	PaymentTransfer    string `json:"payment_transfer"`
	PaymentBankName    string `json:"payment_bank_name"`
	PaymentBankAccount string `json:"payment_bank_account"`
	PaymentPromptpay   string `json:"payment_promptpay"`
	PaymentNote        string `json:"payment_note"`

	// Notification
	NotifyOrderEmail  string `json:"notify_order_email"`
	NotifyLineToken   string `json:"notify_line_token"`
	NotifyLowStock    string `json:"notify_low_stock"`
	LowStockThreshold string `json:"low_stock_threshold"`

	// Footer
	FooterText      string `json:"footer_text"`
	FooterAbout     string `json:"footer_about"`
	FooterCopyright string `json:"footer_copyright"`
}

// ToMap converts the request to a key-value map
func (r *StoreSettingsRequest) ToMap() map[string]string {
	return map[string]string{
		// General
		"store_name":        r.StoreName,
		"store_description": r.StoreDescription,
		"store_logo_url":    r.StoreLogoURL,
		"store_favicon_url": r.StoreFaviconURL,
		"store_banner_url":  r.StoreBannerURL,
		// Theme
		"primary_color":     r.PrimaryColor,
		"secondary_color":   r.SecondaryColor,
		"accent_color":      r.AccentColor,
		"font_family":       r.FontFamily,
		"layout_style":      r.LayoutStyle,
		"products_per_page": r.ProductsPerPage,
		// SEO
		"meta_title":       r.MetaTitle,
		"meta_description": r.MetaDescription,
		"meta_keywords":    r.MetaKeywords,
		"google_analytics": r.GoogleAnalytics,
		// Social
		"social_facebook":  r.SocialFacebook,
		"social_line":      r.SocialLine,
		"social_instagram": r.SocialInstagram,
		"social_twitter":   r.SocialTwitter,
		"social_youtube":   r.SocialYoutube,
		"social_tiktok":    r.SocialTiktok,
		// Contact
		"contact_phone":   r.ContactPhone,
		"contact_email":   r.ContactEmail,
		"contact_address": r.ContactAddress,
		"contact_map_url": r.ContactMapURL,
		"business_hours":  r.BusinessHours,
		// Shipping
		"shipping_enabled":      r.ShippingEnabled,
		"shipping_free_minimum": r.ShippingFreeMinimum,
		"shipping_flat_rate":    r.ShippingFlatRate,
		"shipping_note":         r.ShippingNote,
		// Payment
		"payment_cod":          r.PaymentCOD,
		"payment_transfer":     r.PaymentTransfer,
		"payment_bank_name":    r.PaymentBankName,
		"payment_bank_account": r.PaymentBankAccount,
		"payment_promptpay":    r.PaymentPromptpay,
		"payment_note":         r.PaymentNote,
		// Notification
		"notify_order_email":  r.NotifyOrderEmail,
		"notify_line_token":   r.NotifyLineToken,
		"notify_low_stock":    r.NotifyLowStock,
		"low_stock_threshold": r.LowStockThreshold,
		// Footer
		"footer_text":      r.FooterText,
		"footer_about":     r.FooterAbout,
		"footer_copyright": r.FooterCopyright,
	}
}

// StoreSettingsFromMap creates a response from a key-value map
func StoreSettingsFromMap(m map[string]string) StoreSettingsResponse {
	return StoreSettingsResponse{
		StoreName:           m["store_name"],
		StoreDescription:    m["store_description"],
		StoreLogoURL:        m["store_logo_url"],
		StoreFaviconURL:     m["store_favicon_url"],
		StoreBannerURL:      m["store_banner_url"],
		PrimaryColor:        m["primary_color"],
		SecondaryColor:      m["secondary_color"],
		AccentColor:         m["accent_color"],
		FontFamily:          m["font_family"],
		LayoutStyle:         m["layout_style"],
		ProductsPerPage:     m["products_per_page"],
		MetaTitle:           m["meta_title"],
		MetaDescription:     m["meta_description"],
		MetaKeywords:        m["meta_keywords"],
		GoogleAnalytics:     m["google_analytics"],
		SocialFacebook:      m["social_facebook"],
		SocialLine:          m["social_line"],
		SocialInstagram:     m["social_instagram"],
		SocialTwitter:       m["social_twitter"],
		SocialYoutube:       m["social_youtube"],
		SocialTiktok:        m["social_tiktok"],
		ContactPhone:        m["contact_phone"],
		ContactEmail:        m["contact_email"],
		ContactAddress:      m["contact_address"],
		ContactMapURL:       m["contact_map_url"],
		BusinessHours:       m["business_hours"],
		ShippingEnabled:     m["shipping_enabled"],
		ShippingFreeMinimum: m["shipping_free_minimum"],
		ShippingFlatRate:    m["shipping_flat_rate"],
		ShippingNote:        m["shipping_note"],
		PaymentCOD:          m["payment_cod"],
		PaymentTransfer:     m["payment_transfer"],
		PaymentBankName:     m["payment_bank_name"],
		PaymentBankAccount:  m["payment_bank_account"],
		PaymentPromptpay:    m["payment_promptpay"],
		PaymentNote:         m["payment_note"],
		NotifyOrderEmail:    m["notify_order_email"],
		NotifyLineToken:     m["notify_line_token"],
		NotifyLowStock:      m["notify_low_stock"],
		LowStockThreshold:   m["low_stock_threshold"],
		FooterText:          m["footer_text"],
		FooterAbout:         m["footer_about"],
		FooterCopyright:     m["footer_copyright"],
	}
}
