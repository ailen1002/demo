package backend

type PortlistReq struct {
}

type PortlistRes struct {
	Lists []string `json:"lists"`
}
