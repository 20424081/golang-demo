package Controllers

type Resp struct {
		Count     int64       `json:"count,omitempty"`
		Result    interface{} `json:"result,omitempty"`
		Error     interface{} `json:"error,omitempty"`
		Message   string	  `json:"message"`
}

type Filter struct{
	Limit     int         `query:"limit"`
	Offset    int         `query:"offset"`
	Search    string      `query:"search"`
}