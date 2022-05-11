package Controllers

type Resp struct {
		Count     int         `json:"count"`
		Result    interface{} `json:"result,omitempty"`
		Error     interface{} `json:"error,omitempty"`
}