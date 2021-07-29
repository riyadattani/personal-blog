package http_api

import "time"

type ServerConfig struct {
	CSSDir           string
	HTMLDir          string
	Port             string
	HTTPReadTimeout  time.Duration
	HTTPWriteTimeout time.Duration
}

func (c ServerConfig) TCPAddress() string {
	return ":" + c.Port
}
