package models

type Guest struct {
	Adults   int  `json:"adults"`
	Children *int `json:"children"`
}

type LanguageRequest struct {
	Language string `json:"language"`
}

type HotelInfoRequest struct {
	ID       string `json:"id"`
	Language string `json:"language"`
}

type SearchMulticompleteRequest struct {
	Query    string `json:"query"`
	Language string `json:"language"`
}

type SearchRequest struct {
	RegionId    int     `json:"region_id"`
	Checkin     string  `json:"checkin"`
	Checkout    string  `json:"checkout"`
	Guests      []Guest `json:"guests"`
	Currency    *string `json:"currency,omitempty"`
	Residency   *string `json:"residency,omitempty"`
	HotelsLimit *int    `json:"hotels_limit,omitempty"`
	Timeout     *int    `json:"timeout,omitempty"`
	Language    *string `json:"language,omitempty"`
}

type SearchGeoRequest struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Radius    int     `json:"radius"`
	Checkin   string  `json:"checkin"`
	Checkout  string  `json:"checkout"`
	Guests    []Guest `json:"guests"`
	Currency  *string `json:"currency,omitempty"`
	Residency *string `json:"residency,omitempty"`
	Timeout   *int    `json:"timeout,omitempty"`
	Language  *string `json:"language,omitempty"`
}

type SearchHotelsRequest struct {
	IDs       []string `json:"ids"`
	Checkin   string   `json:"checkin"`
	Checkout  string   `json:"checkout"`
	Guests    []Guest  `json:"guests"`
	Currency  *string  `json:"currency,omitempty"`
	Residency *string  `json:"residency,omitempty"`
	Timeout   *int     `json:"timeout,omitempty"`
	Language  *string  `json:"language,omitempty"`
}

type SearchHotelPage struct {
	ID        string  `json:"id"`
	Checkin   string  `json:"checkin"`
	Checkout  string  `json:"checkout"`
	Guests    []Guest `json:"guests"`
	Currency  *string `json:"currency,omitempty"`
	Residency *string `json:"residency,omitempty"`
	Upsells   *[]struct {
		EarlyCheckin *[]struct {
			Time *string `json:"time,omitempty"`
		} `json:"early_checkin,omitempty"`
		LateCheckout *[]struct {
			Time *string `json:"time,omitempty"`
		} `json:"late_checkout,omitempty"`
		MultipleECLC *bool `json:"multiple_eclc"`
		OnlyECLC     *bool `json:"only_eclc"`
	} `json:"upsells,omitempty"`
	Timeout  *int    `json:"timeout,omitempty"`
	Language *string `json:"language,omitempty"`
}
