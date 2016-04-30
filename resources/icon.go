package resources

import "time"

// Icon type of Icon
type Icon struct {
	*Resource
	URL   string `json:",omitempty"`
	Name  string `json:",omitempty"`
	Scope string `json:",omitempty"` //TODO scopeってenumかも？
	*EAvailability
	CreatedAt  time.Time `json:",omitempty"`
	ModifiedAt time.Time `json:",omitempty"`
	Tags       []string  `json:",omitempty"`
}
