package jobserver

func NewAPIHandler() APIHandler {
	return &hdl{}
}

type hdl struct{}
