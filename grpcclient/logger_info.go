package grpcclient

type LogRequest struct {
	Log string `json:"log"`
}

type LogResponse struct {
	Ok bool `json:"ok"`
}
