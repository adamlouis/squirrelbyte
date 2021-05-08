package documentserver

func NewAPIHandler() APIHandler {
	return &hdl{}
}

type hdl struct{}
