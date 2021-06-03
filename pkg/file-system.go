package pkg

import (
	"fmt"
	"io/fs"
	"personal-blog/pkg/blog"
)

func New(postsDir fs.FS) (blog.Posts, error) {
	dir, err := fs.ReadDir(postsDir, ".")
	if err != nil {
		return nil, fmt.Errorf("cannot read the posts directory: %s", err)
	}
	return getPosts(postsDir, dir)
}

func getPosts(postsDir fs.FS, dir []fs.DirEntry) ([]blog.Post, error) {
	var posts []blog.Post
	for _, file := range dir {
		post, err := newPostFromFile(postsDir, file.Name())
		if err != nil {
			return nil, fmt.Errorf("cannot create a new post: %s", err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func newPostFromFile(postsDir fs.FS, fileName string) (blog.Post, error) {
	f, err := postsDir.Open(fileName)
	defer f.Close()

	if err != nil {
		return blog.Post{}, fmt.Errorf("cannot open the file %s: %s", fileName, err)
	}

	return blog.NewPost(f)
}
