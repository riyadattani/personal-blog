package acceptance_tests

import (
	is2 "github.com/matryer/is"
	"personal-blog/pkg/blog"
	"personal-blog/pkg/teshelpers"
	"testing"
)

type BlogDriver interface {
	Publish(post blog.Post) error
	ReadPost(urlTitle string) (exists bool, err error)
}

func BlogAcceptanceCriteria(t *testing.T, blog BlogDriver) {
	t.Run("Can publish a post, find by title and read the content", func(t *testing.T) {
		is := is2.New(t)
		post := teshelpers.NewPostBuilder().Build()

		err := blog.Publish(post)
		is.NoErr(err)

		exists, err := blog.ReadPost(post.URLTitle)
		is.NoErr(err)
		is.True(exists)
	})
}
