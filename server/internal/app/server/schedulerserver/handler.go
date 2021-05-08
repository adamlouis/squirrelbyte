package schedulerserver

func NewAPIHandler() APIHandler {
	return &hdl{}
}

type hdl struct{}
