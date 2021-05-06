A technical approach on building this blog from scratch 
2021-May-03
golangPic2.png
golang
-----

I am writing a more technical counterpart to my [previous blog post](https://www.riyadattani.com/blog/My%20journey%20to%20writing%20my%20blog%20in%20Go%20). I am going to explore my technical plan of achieving my MVP (Minimum Viable Product) and how my code base evolved.

### What is my MVP?

1. A home page which has a list of my blogs in descending order by date
2. The blogs are links that open up on a separate page

When starting a program, I think it's particularly intimidating to start writing that first bit of code when faced with a
blank canvas. In this situation, I find it useful to broadly think about _how_ I am going to build the system. There is
a lot of discussion around waterfall development vs agile development which I am not going to dwell on now, but
developers generally strive to be agile. I think this illustration best describes the difference between the two and why we should aim to be agile:

<img src="../css/images/building-software.jpg" alt="Building software - agile" />

> The key to agile is incremental development

To build the "skateboard" of my site, I used the **Steel thread** concept to figure out where to start. A Steel thread
is some important, minimal thread of functionality that runs throughout the system. What is the minimal amount of work
you can do to deliver value to the user of your site? Here is a sequence diagram of what I defined as my Steel thread:

<img src="../css/images/sequence-diagram-steel-thread.png" alt="Steel thread diagram" />

In respect to _my own_ experience of adding blogs to the website, my ultimate goal was to use markdown files to write
the blogs that I can simply commit to the Github repository for this site. Markdown files are perfect as I can reap the
benefits of basic formatting, and I can easily embed videos or images.

Here are the steps I took to get started:

1. The "Hello world" version
    - This is the simplest way to get the show on the road. I used this example: https://gobyexample.com/http-servers,
      which steered me to create a http server and print out "Hello world" on a page.
    - I extracted out the server into a separate package and used http://www.gorillatoolkit.org/pkg/mux to handle my
      routing. At this point, I had a server up and something was on the page - a great start. My next step was to
      render a list of links that represent blog posts.
2. Hard-coded blogs
    - This was where the fun began. Instead of printing text on the page, I wanted to render some html with an unordered
      list. I used the built-in `html/template` to point to a html file that had a list of hard-coded blog titles.
    - I pushed the hard-coded titles from the html file into an `InMemoryRespository`. Here is where the list of blog
      posts would live. I passed the list from this repository to the handler for the homepage in order to map over
      the posts in the html template. My `Post` type only required a title and the content:

```go
type Post struct {
  Title   string
  Content string
}
```

For a user to complete their journey, they should be able to click on the blog post link and view the full contents of
that particular blog. I iteratively added a route `/blogs/{blogPostTitle}`and a html template to render the contents of
the post.

Hooray, the Steel thread is complete. The proof of concept is now a reality, and the site had an end-to-end user journey.
The scaffolding of the software is shaping up, and the next step was to shift the hard-coded blog posts from
the `InMemoryRepository` to markdown files.

This is how I did it:

- I created a `blog-posts` folder to store markdown files.
- The main gist is to map over each file in the `blog-posts` folder and create a `Post` out of them. I had some fun
  challenges here. How do I read a markdown file? How do I transform it into html?
- I used `ioutil.ReadFile` to read the file.
- I transformed the content into html using https://github.com/russross/blackfriday.

```go
func readPost(title string) ([]byte, error) {
	body, err := ioutil.ReadFile(fmt.Sprintf("../web/blogs/%s", title))
	if err != nil {
		return nil, err
	}

	output := blackfriday.Run(body)
	return output, nil
}
```

- The `content` in the `Post` struct is no longer a `string` but of type `template.HTML`.

This was a great milestone to achieve however, I had not actually built my MVP yet. I had two problems:

1. The titles were still hardcoded on the homepage
2. I want to order the blogs by date

I had to think of a solution to specify the key attributes of a `Post`. I decided to split the
markdown file by metadata, and the content of the blog post. Here is an example of a typical blog post:

<img src="../css/images/markdown-example.png" alt="Markdown example" />

This was a tricky challenge because I had to figure out a way to read the metadata line by line, and after the dash
characters `-----`, I wanted to assign the rest of the file as the content. TDD (Test Driven Development) to my
rescue. _Properly_ test-driving this part of the code helped me break down the problem into smaller chucks. I used the
built-in package `bufio` to scan the file. Here is a code snippet of my test and solution:

Test:

```go
func TestBlog(t *testing.T) {
   t.Run("it should split the markdown file into the metadata and the content", func(t *testing.T) {
      byteArray := []byte(markdownDoc)
      title, body, date, _ := blog.CreatePost(byteArray)

      expectedTitle := `This is the title of the first blog post`
      expectedDate := `2021-05-05`
      expectedBody := `This is the first sentence of the post. This is the second sentence of the post.`

      if string(body) != expectedBody {
         t.Errorf("got %q, want %q", body, expectedBody)
      }

      if title != expectedTitle {
         t.Errorf("got %q, want %q", title, expectedTitle)
      }

      if date != expectedDate {
         t.Errorf("got %q, want %q", date, expectedDate)
      }
   })
}
```

Solution:
```go
type Post struct {
   Title   string
   Content template.HTML
   Date    time.Time
}

func NewPost(fileName string) (Post, error) {
   fileContent, err := ioutil.ReadFile(fmt.Sprintf("../../cmd/web/posts/%s", fileName))
   if err != nil {
      return Post{}, err
   }

   title, body, date, err := CreatePost(fileContent)
   if err != nil {
      return Post{}, err
   }

   content := blackfriday.Run(body)

   const shortForm = "2006-Jan-02"
   parsedDate, err := time.Parse(shortForm, date)
   if err != nil {
      return Post{}, err
   }

   return Post{
      Title:   title,
      Content: template.HTML(content),
      Date:    parsedDate,
   }, nil
}

func CreatePost(fileContent []byte) (title string, body []byte, date string, err error) {
   r := bytes.NewReader(fileContent)

   metaData := getMetaData(r)
   title = metaData[0]
   date = metaData[1]

   body = getContentBody(fileContent)

   return title, body, date, nil
}

func getMetaData(r io.Reader) []string {
   metaData := make([]string, 0)
   scanner := bufio.NewScanner(r)
   scanner.Split(bufio.ScanLines)

   for scanner.Scan() {
      line := scanner.Text()
      if line == "-----" {
         break
      }
      metaData = append(metaData, line)
   }

   return metaData
}

func getContentBody(byteArray []byte) []byte {
   content := bytes.Split(byteArray, []byte("-----\n"))[1]
   return content
}
```


To finally accomplish the MVP, I want the blogs to be listed in descending order by date. I chose to make my life easier by writing
the date in a particular format in the markdown file so that I can transform that `string` into a`time.Time` type using
the built-in `time` package. As a`time.Time` type, I could sort the posts in descending order in
the `InMemoryRepository`. It was a pleasant surprise when I discovered that I could format the date to look more readable in the html
template:

```html
<time>
   {{.Date.Format "Jan 02, 2006"}}
</time>
```

The MVP is complete! Building this in an iterative way gave me the opportunity to break down the scary big problems into
digestible small problems. Soon after, I added tags and a picture to the metadata (and to the `Post`) with ease.

### Future features

- An RSS feed to enable users to subscribe to the site.
- Dark mode.

### To refactor 

- When creating a post, `getMetaData` should return structured data `MetaData` rather than `string[]`. 
- Use `io.Reader` effectively instead of bytes. 
