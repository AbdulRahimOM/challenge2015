package data

type Person struct {
	// Name       string      `json:"name"`
	// URL        string      `json:"url"`
	// Type       string      `json:"type"`
	MovieRoles []MovieRole `json:"movies"`
}

type MovieRole struct {
	// Name string `json:"name"`
	URL string `json:"url"`
	// Role string `json:"role"`
}

type Movie struct {
	// Name string     `json:"name"`
	// URL  string     `json:"url"`
	// Type string     `json:"type"`
	Cast []CastCrew `json:"cast"`
	Crew []CastCrew `json:"crew"`
}

type CastCrew struct {
	// Name string `json:"name"`
	URL string `json:"url"`
	// Role string `json:"role"`
}
