package ports

import "personal-blog/pkg/twitter"

type TwitterGateway interface {
	GetUserTimeline() (twitter.APIResponse, error)
}
