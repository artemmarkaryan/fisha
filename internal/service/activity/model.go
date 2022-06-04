package activity

type Activity struct {
	Id        int64   `db:"id"`
	Name      string  `db:"name"`
	CreatedAt string  `db:"created_at"`
	UpdatedAt string  `db:"updated_at"`
	Address   string  `db:"address"`
	Lon       float64 `db:"lon"`
	Lat       float64 `db:"lat"`
	Meta      []byte  `db:"meta"`
}

type Metadata struct {
	Id         string               `json:"id"`
	Name       string               `json:"name"`
	URL        string               `json:"url"`
	Address    string               `json:"address"`
	Hours      MetadataHours        `json:"Hours"`
	Phones     []MetadataPhones     `json:"Phones"`
	Categories []MetadataCategories `json:"Categories"`
}

type MetadataHours struct {
	Text           string `json:"text"`
	Availabilities []struct {
		Monday    bool `json:"Monday,omitempty"`
		Tuesday   bool `json:"Tuesday,omitempty"`
		Thursday  bool `json:"Thursday,omitempty"`
		Intervals []struct {
			To   string `json:"to"`
			From string `json:"from"`
		} `json:"Intervals"`
		Wednesday bool `json:"Wednesday,omitempty"`
		Friday    bool `json:"Friday,omitempty"`
		Saturday  bool `json:"Saturday,omitempty"`
		Sunday    bool `json:"Sunday,omitempty"`
	} `json:"Availabilities"`
}

type MetadataPhones struct {
	Type      string `json:"type"`
	Formatted string `json:"formatted"`
}

type MetadataCategories struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}
