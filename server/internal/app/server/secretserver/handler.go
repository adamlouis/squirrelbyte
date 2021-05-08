package secretserver

func NewAPIHandler() APIHandler {
	return &hdl{}
}

type hdl struct{}
