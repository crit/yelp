package business

type Region struct {
	Center struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"center"`
}

type Business struct {
	Id          string `json:"id"`
	Alias       string `json:"alias"`
	Name        string `json:"name"`
	ImageUrl    string `json:"image_url"`
	IsClosed    bool   `json:"is_closed"`
	Url         string `json:"url"`
	ReviewCount int    `json:"review_count"`
	Categories  []struct {
		Alias string `json:"alias"`
		Title string `json:"title"`
	} `json:"categories"`
	Rating      float64 `json:"rating"`
	Coordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
	Transactions []string `json:"transactions"`
	Price        string   `json:"price"`
	Location     struct {
		Address1       string      `json:"address1"`
		Address2       interface{} `json:"address2"`
		Address3       string      `json:"address3"`
		City           string      `json:"city"`
		ZipCode        string      `json:"zip_code"`
		Country        string      `json:"country"`
		State          string      `json:"state"`
		DisplayAddress []string    `json:"display_address"`
	} `json:"location"`
	Phone        string            `json:"phone"`
	DisplayPhone string            `json:"display_phone"`
	Distance     float64           `json:"distance"`
	Attributes   map[string]string `json:"attributes"`
}
