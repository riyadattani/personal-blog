My journey to writing my blog in Go
2021-Apr-12
golangPic1.png
golang
-----

I built this blog from scratch, without a blog engine. I used Go to build it and some html + css to make it look pretty.
This is my first personal project in Go, in fact, its my first personal project on a live
website. [Tweet me](https://twitter.com/DattaniRiya) if you have any feedback whether its kind, mean or funny! I also wrote a technical counterpart to this blog - you can find it [here](https://www.riyadattani.com/blog/A%20technical%20approach%20on%20building%20this%20blog%20from%20scratch%20).

So I made the decision to make a blog in Go because it was different language to what I normally use at work (Node) and
when I was introduced to it, it felt fun and simple.

### First attempt

Once the decision was made, the question was where do I begin? I surfed the web and looked for other desirable blogs. I
gathered a list of my favourite features and thought I would figure out a way to replicate them. I wanted very cool
features like an animated background and an interactive interface such as a built in game. These features sounded
fabulous but in reality, they were a nice-to-have rather than necessary. Frankly, the more I thought about what my blog
should look like, the more I ventured further away from the purpose of the blog: to share articles. "Agile" developers
would say I envisioned a shiny gold standard website rather than a MVP (Minimum Viable Product).

I started building the `Hello World` version of my site and I cheated a little by not writing any tests. Not very long after, I frantically searched for tutorials to get me these fancy features. I was directed to multiple plugins, I
read a lot of stack overflow and I spent most of my time searching for resources on google rather than coding... I
eventually got frustrated. Naturally, I didn't want to open the project anymore because I didn't really know what to do
and where to start.

The tutorials and articles I read felt overwhelming and not perfectly suitable to what I am trying
to build. I told myself some unworthy excuses and left the blog alone for a couple of weeks.

> I needed to go back to the basics and start simple.

On a fine Hackday, I was experimenting with Hotwire and Go with my colleagues. We built a simple todo list app with
basic CRUD functionality. Thanks to the fast, non-hassle set up to get a website up and running in Go, once the bare
bones were wired up, we iteratively fleshed out all the routes and basically built the app in a day. The simplicity
inspired me to start building my blog again. The process felt easy, simple and doable. I wanted to restart building it
using the same strategy and foundation.

### Second attempt

Before my second attempt, I reflected on my previous approach/process and learned from my mistakes. Here were my
learnings:

1. Remember the _purpose_ of what I am trying to build. As a programmer, we should focus on WHAT, WHY and HOW we are
   going to solve a specific problem.
    - **What**: I want to build a blog where I can share articles of topics that I am interested in and I am learning
      about.
    - **Why**: I want to build a network and connect with people who have similar interests, share knowledge that other
      people might find useful and practice communicating through writing effectively.
    - **How**: I made a decision to not use a blog engine and to build the blog from scratch in Go in an _iterative_ way. Whilst a blog engine would have easily enabled me to achieve my goals, I would not have gained this invaluable learning experience.

2. Decide on the _MVP_: I wrote down 2 features which I think I need for my blog to be functional.
    - A home page which has a list of my blogs in descending order by date.
    - The blogs are links that open up on a separate page.

3. Write _tests_ or you will regret it in the future.
    - Go through [Learn Go With Tests](https://github.com/quii/learn-go-with-tests) to learn how to test the code _correctly_ in Go.
    - What would be included in my acceptance test? Do I need one?

4. Have a _process_ and develop it to be more systematic.
    - **Github Issues**:  I used these to keep a backlog of tasks, bugs and general thoughts that need doing especially
      when they crop up in the middle of something else.
    - **Continuous Integration**:
        - I tried to apply some processes from work to make it as easy as possible to continuously integrate my code to
          the live site. I created a github action that builds, tests and deploys the code, every time I push to master. I
          always work on master.
        - I took this opportunity to learn more about docker. I used docker to build my code in a container and this is
          used by Heroku when deploying my site.

5. Don't bother following a tutorial. Figure out what you _want_ to build in the simplest way possible. I started with a
   hard coded blog list on the site and then aimed to replace it with non-hard coded blogs.

In retrospect, I would have _still_ done things differently if I took another shot at creating this blog again. I guess
that's the best (or worst) thing about software development - nothing is perfect and there is _always_ room to improve.
Looking back at my second attempt at building this blog, here is what I would've done differently:

1. TDD it. No, I mean _actually_ TDD it. Yes, I realised that writing _no tests_ would hurt me in the future and make me
   less confident in my code, but I still did not _strictly_ follow TDD. I added a couple of tests with respect to my
   router after writing it up and I realised I probably would've created a better solution if I had tested it first.
2. Don't get obsessed over the CSS - this can be done over time. The perfectionist in me was running around in circles.
   Instead of focusing on the most important task, I wanted to make the site look pixel perfect.
3. Get feedback on the code and the design as soon as possible.

In hindsight, building this blog has been a fantastic learning experience! I felt technically challenged, and I learned
about the process of creating a website from scratch to live. If you are interested in the technicalities, have a read through [this complimentary blog](https://www.riyadattani.com/blog/A%20technical%20approach%20on%20building%20this%20blog%20from%20scratch%20) to view some code.