package model

import "time"

type ContentData struct {
	Data []struct {
		UUID      string `json:"uuid"`
		Type      string `json:"type"`
		TokenID   string `json:"token_id"`
		IP        string `json:"ip"`
		Hostname  string `json:"hostname"`
		Method    string `json:"method"`
		UserAgent string `json:"user_agent"`
		Content   struct {
			ContactID     string    `json:"contact_id"`
			FirstName     string    `json:"first_name"`
			LastName      string    `json:"last_name"`
			FullName      string    `json:"full_name"`
			Email         string    `json:"email"`
			Phone         string    `json:"phone"`
			Tags          string    `json:"tags"`
			State         string    `json:"state"`
			Country       string    `json:"country"`
			Timezone      string    `json:"timezone"`
			DateCreated   time.Time `json:"date_created"`
			PostalCode    string    `json:"postal_code"`
			DateOfBirth   time.Time `json:"date_of_birth"`
			ContactSource string    `json:"contact_source"`
			FullAddress   string    `json:"full_address"`
			ContactType   string    `json:"contact_type"`
			Gclid         any       `json:"gclid"`
			Location      struct {
				Name        string `json:"name"`
				Country     string `json:"country"`
				FullAddress string `json:"fullAddress"`
				ID          string `json:"id"`
			} `json:"location"`
			User struct {
				FirstName string `json:"firstName"`
				LastName  string `json:"lastName"`
				Email     string `json:"email"`
				Phone     string `json:"phone"`
			} `json:"user"`
			Workflow struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"workflow"`
			Contact struct {
				AttributionSource struct {
					Gclid         any    `json:"gclid"`
					UserAgent     string `json:"userAgent"`
					SessionSource string `json:"sessionSource"`
					UtmTerm       any    `json:"utmTerm"`
					URL           string `json:"url"`
					FbEventID     string `json:"fbEventId"`
					UtmContent    any    `json:"utmContent"`
					Referrer      any    `json:"referrer"`
					UtmMedium     any    `json:"utmMedium"`
					IP            string `json:"ip"`
					Msclkid       any    `json:"msclkid"`
					UtmSource     any    `json:"utmSource"`
				} `json:"attributionSource"`
				LastAttributionSource struct {
					UserAgent     string `json:"userAgent"`
					URL           string `json:"url"`
					IP            string `json:"ip"`
					Msclkid       any    `json:"msclkid"`
					UtmSource     any    `json:"utmSource"`
					SessionSource string `json:"sessionSource"`
					Referrer      any    `json:"referrer"`
					FbEventID     string `json:"fbEventId"`
					UtmContent    any    `json:"utmContent"`
					Gclid         any    `json:"gclid"`
					UtmTerm       any    `json:"utmTerm"`
					UtmMedium     any    `json:"utmMedium"`
				} `json:"lastAttributionSource"`
			} `json:"contact"`
			AttributionSource struct {
			} `json:"attributionSource"`
			CustomData struct {
			} `json:"customData"`
		} `json:"content"`
		Query   any `json:"query"`
		Headers struct {
			Connection    []string `json:"connection"`
			Host          []string `json:"host"`
			ContentLength []string `json:"content-length"`
			UserAgent     []string `json:"user-agent"`
			ContentType   []string `json:"content-type"`
			Accept        []string `json:"accept"`
		} `json:"headers"`
		URL                string `json:"url"`
		Size               int    `json:"size"`
		Files              []any  `json:"files"`
		CreatedAt          string `json:"created_at"`
		UpdatedAt          string `json:"updated_at"`
		Sorting            int64  `json:"sorting"`
		CustomActionOutput []any  `json:"custom_action_output"`
	} `json:"data"`
	Total       int  `json:"total"`
	PerPage     int  `json:"per_page"`
	CurrentPage int  `json:"current_page"`
	IsLastPage  bool `json:"is_last_page"`
	From        int  `json:"from"`
	To          int  `json:"to"`
}
