package constants

// Constant represents a constant
type Constant struct {
	Bytes   []byte     `json:"bytes"`
	Boolean *bool      `json:"bool"`
	String  *string    `json:"string"`
	Int     *int       `json:"int"`
	Uint    *uint      `json:"uint"`
	Float   *float64   `json:"float"`
	List    []Constant `json:"constant"`
}
