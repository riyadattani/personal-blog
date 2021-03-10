My journey to writing my blog in Go as junior developer 
2021-03-10

-----

I built a blog from scratch, without a blog engine, using Golang.

### First try

I looked at other desirable blogs and gathered a list of my favourite feature and thought I would figure out a way to replicate that. I wanted very cool features like an animated background and an interactive tool such as a built in game. These features sounded fabulous but in reality, they were a nice-to-have rather than necessary. In fact, the more I thought about what my blog should look like, the more I ventured further away from the purpose of the blog: to share articles. In other words, I envisioned a shiny gold standard website rather than a MVP (Minimum Viable Product).

I started building the "Hello World" version of my site and I cheated a little bit by not really writing any tests. Not very long after, I frantically searched for tutorials to get me these cool features. I was directed to multiple plugins and various other tutorials... I eventually got frustrated. Naturally, I didn't want to open the project anymore because I didn't really know what to do and where to start. The tutorials and articles I read felt overwhelming and not perfectly suitable to what I am trying to build. I told myself some unworthy excuses and left the blog alone for a couple of weeks.

> I needed to go back to the basics and start simple.

On a fine Hackday, I was experimenting with Hotwire and Golang with my colleagues. We uniquely decided to build a simple todo list app with basic CRUD functionality. Thanks to the fast, non-hassle set up to get a website up and running in Golang, once the bare bones were wired up, we iteratively fleshed out all the routes and basically built the app in a day. The simplicity inspired me to start building a blog again. The process felt easy, simple and doable. I wanted to restart building my blog using the same strategy and foundation.

### Second try

During my second attempt at writing my blog I reflected on my previous approach and process and learned from my mistakes. Here were my learnings:

1. Remember the _purpose_ of what I am trying to build. As a programmer, we should focus on WHAT, WHY and HOW we are going to solve a specific problem.
    - **What**: I want to build a blog where I can share articles of things that I am interested in and I am learning
    - **Why**: I want to build a network and connect with other people interested in the same thing, share knowledge that other people might find useful and practice communicating through writing effectively.
    - **How**: I made a decision to not use a blog engine and to build the blog from scratch in Golang in an _iterative_ way. I want to make the MVP first and let it evolve in the future.

2. Decide on the _MVP_: I wrote down 2 features which I think I need for my blog to be functional
    - A home page which has a list of my blogs in descending order by date
    - The blogs are links that open up on a separate page

3. Write _tests_ or you will regret it in the future.
    - Go through [Learn Go With Tests](https://github.com/quii/learn-go-with-tests) to learn how to test the code _correctly_ in Golang
    - What would be included in my acceptance test? Do I need one?

4. Have a _process_ and develop it to be more systematic.
    - **Github Issues**:  I used these to keep a backlog of tasks, bugs and general thoughts that need doing especially when they crop up in the middle of something else.
    - **Continuous Integration**:
        - I tried to apply some processes from work to make it as easy as possible to continuously integrate my code to the live site. I created a pipeline that builds, tests and deploys the code at every time I push to master. I always work on master.
        - I took this opportunity to learn more about docker. I used docker to build my code in a container and this is used by Heroku when deploying my site. 