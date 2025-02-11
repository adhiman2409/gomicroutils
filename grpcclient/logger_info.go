package grpcclient

type LogRequest struct {
	Log    string `json:"log"`
	Domain string `json:"domain"`
}

type LogResponse struct {
	Ok bool `json:"ok"`
}
