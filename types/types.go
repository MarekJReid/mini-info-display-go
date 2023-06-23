package mytypes



type SidResponse struct {
	Message string `json:"message"`
}

type McidResponse struct {
	Message string `json:"message"`
}

type PostDataRequest struct {
	Data string `json:"data"`
}

type PostDataResponse struct {
	Message string `json:"message"`
}