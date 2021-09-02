package http_api

import "time"

type ServerConfig struct {
	CSSDir           string
	HTMLDir          string
	Port             string
	HTTPReadTimeout  time.Duration
	HTTPWriteTimeout time.Duration
	PostsDir         string
	EventsDir        string
}

func (c ServerConfig) TCPAddress() string {
	return ":" + c.Port
}
