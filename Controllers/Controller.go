package Controllers

type Resp struct {
		Count     int         `json:"count,omitempty"`
		Result    interface{} `json:"result,omitempty"`
		Error     interface{} `json:"error,omitempty"`
		Message   string	  `json:"message"`
}