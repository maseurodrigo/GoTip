package Structs

// JSON basic structure
type JSONData struct {
	Key       string `json:"APIKey"`
	Trigger   string `json:"Trigger"`
	CallStart string `json:"AICall_Start"`
	CallEnd   string `json:"AICall_End"`
}
