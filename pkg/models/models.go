package models

//Port structure info
type Port struct {
	Name        string    `json:"name" bson:"Name,omitempty"`
	City        string    `json:"city" bson:"City,omitempty"`
	Country     string    `json:"country" bson:"Country,omitempty"`
	Alias       []string  `json:"alias" bson:"Alias,omitempty"`
	Regions     []string  `json:"regions" bson:"Regions,omitempty"`
	Coordinates []float64 `json:"coordinates" bson:"Coordinates,omitempty"`
	Province    string    `json:"province" bson:"Province,omitempty"`
	Timezone    string    `json:"timezone" bson:"Timezone,omitempty"`
	Unlocs      []string  `json:"unlocs" bson:"Unlocs,omitempty"`
	Code        string    `json:"code" bson:"Code,omitempty"`
}

//PortDetails structure info
type PortDetails struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	Port Port   `json:"port" bson:"Port,omitempty"`
}
