package oauthserver

func NewAPIHandler() APIHandler {
	return &hdl{}
}

type hdl struct{}
