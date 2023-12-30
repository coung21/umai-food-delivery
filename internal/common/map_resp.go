package common

type GoongGeocodeResponse struct {
	Results []struct {
		AddressComponents []struct {
			LongName  string `json:"long_name"`
			ShortName string `json:"short_name"`
		} `json:"address_components"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			Boundary struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"boundary"`
		} `json:"geometry"`
		PlaceID   string `json:"place_id"`
		Reference string `json:"reference"`
		PlusCode  struct {
			GlobalCode   string `json:"global_code"`
			CompoundCode string `json:"compound_code"`
		} `json:"plus_code"`
		Compound struct {
			District string `json:"district"`
			Comune   string `json:"comune"`
			Province string `json:"province"`
		} `json:"compound"`
		Types   []string `json:"types"`
		Name    string   `json:"name"`
		Address string   `json:"address"`
	} `json:"results"`
	Status string `json:"status"`
}
