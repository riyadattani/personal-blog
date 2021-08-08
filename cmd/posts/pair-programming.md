Deep dive into pair programming
2021-Aug-08
pair-programming-rafiki.png
pair-programming
-----

Before I switched careers into software development, I didn't have a clue of what a complex system looked like, how they are created and the effort spent building them. I had a misconception that programmers spent a lot of time writing code independently. When I started learning how to code and speaking to developers about their experiences, I was quite astonished at the collaborative culture within software development. In my job, I continually interact with other developers, the product owner, the UX/UI designers and sometimes the stakeholders of the product we are working to build. I am encouraged to not only share my thoughts on technical aspects, but on UX/UI prototypes and on the value of the features we are building for our users. My role as a software developer is much more than just coding. In this blog however, I want to focus on how developers collaborate through pair programming.

### Pair programming

Pair programming is where two developers work together to solve a problem. Whilst there are advantages to pairing most of the time, I think there are some occasions when it is better to work individually.

_When_ it is suitable to pair? Well, it depends on what the specific task entails. Is it a laborious task of deleting dead code or is it a new search feature? There are some tasks that are just not worth the time of two developers. One of the developers is likely to sit and stare at the screen watching the other developer continue with the task. I put together a non-exhaustive list of questions that I would ask myself to determine if its worth pairing or not:

- Is it beneficial for another member of the team to have context of the code that will be added (or refactored)? Knowledge sharing is encouraged so that every developer is familiar with the entire code base and is comfortable enough to pick up any ticket on the board.
- Is this a learning opportunity for another developer? The work may not be something new and shiny but say you had a new joiner and you were touching some part of the code that would help them understand the project better.
- Can the work be broken down into extremely small tasks where it makes sense to divide and conquer rather than pair? For example, it might make more sense to divide a bunch of small UI tweaks if they are relatively straight forward.
- Is this a growth opportunity for a budding developer to work on a task individually? It is liberating and exciting for a novice to tackle a challenging task themselves. It is an opportunity for them to make mistakes, learn about their thought process when analysing a technical problem and get feedback on their work.
- Are you in the _mood_ to pair? This is a tough one because it is so subjective. Some developers would love to just plug in their headphones and get on with their task. Some need a social break because pairing can be extremely exhausting since you have to communicate for long periods of time over difficult code stuff. Pairing contributes to overall productivity but _up to a certain point_. I think you can experience diminishing returns when it becomes tiring and you cannot optimally contribute. I think most people can sense when this happens and they should just be honest about wanting to branch off and work on something solo.

### Types of pair programming:

#### 1. Driver-navigator

This is a classic pair programming style. One developer is the person on the wheel (i.e. keyboard) and their role is to type code. The other acts as the "observer" and their role is to review the code on-the-go and navigate the driver on _what_ to code. This exercise should involve the navigator explaining _how_ they plan to solve a problem and prompt the driver to codify that idea. The navigator should avoid dictating the code word for word to the driver. I think this pairing practice is useful when a more experienced developer is coaching another developer on how to code.

In my opinion, 20 minutes or less is a reasonable amount of time to write some useful code. These short sessions encourage developers to make small, iterative changes because you would ideally want to integrate that code often. Switching between the two roles is important because the perspective on the code is different. The drivers job is to think about the nitty gritty of the code at hand. The navigator can think more strategically about the bigger picture such as potential obstacles that might arise.

> The key to pairing is to switch roles frequently.

In practice, these guidelines are not always strictly followed. There have been a number of times in my own experience where a developer and I have not switched roles for an hour or even more sometimes. Those sessions drained us and rendered us unmotivated. Ultimately, since the work still needs to be done, the solution is less thought through and important steps such as refactoring could be missed out of laziness. Switching roles increases the participation from each developer and thus the energy levels.

#### 2. Ping pong

Ping pong pair programming is a fun way to program and it revolves around writing tests. In this style, developer X writes a test and developer Y writes the code to pass that test. Developer Y then writes the next test and developer X has to write the code to pass the test. This is a great technique because you are following Test Driven Development (TDD). TDD is a whole other topic in itself but, writing tests that reflect the desired _behaviour_ is a good _tool_ to design good, modular software. Read [this post](https://quii.dev/The_Why_of_TDD) written by my colleague if you are interested in a convincing article about TDD.

This is my favourite pairing style. A major plus point is that the developer writing the test is focusing on writing a useful test without the solution in mind because the solution will be implemented by their pairing partner. This leads to clear, understandable tests and therefore a good chance of well-designed production code. I like how the code can evolve in an unexpected way when bouncing off each other's code and tests.

The well-known TDD steps should be followed:

- **Red**: write a failing test
- **Green**: write the minimum amount of code to pass the test
- **Refactor**: refactor the code to improve your solution

One of the important steps to remember is the refactoring stage because it is your moment to **design well**! You have reliable tests and working software so refactoring should only improve the quality of your code (if done correctly of course). Can you spot a pattern in your code? Can you extract some code to make it more readable?  Are you injecting dependencies in a function or a class? Are you following [SOLID principles](https://en.wikipedia.org/wiki/SOLID)?

#### 3. Diverge-converge

This is not a very well-known technique but it is something I picked up whilst I was at a coding bootcamp. This technique involves two developers diverging out on their own to figure out how they would solve a problem and then converging back together to discuss and decide what they will implement. I think this technique useful when there is not a concrete plan on how to make something work. This method is suitable if the task requires a lot of research or reading through documentation which is more efficiently done solo rather than in a pair. The developers could perform a spike to see if their solution actually works (a spike is experimenting/exploring with some code). However, once you figure out something viable, it is important to bin all that code and start again using TDD.

### Before and after pairing

There are different pair programming methods to _code_ together. But when picking up a task, there are other important ceremonies before and after coding.

- **Kick-off**: Fully understand the use case by fleshing it out with the rest of the team including any non-devs that can provide more context. I find breaking down the scenarios into a test case structure e.g. "given, when, then" is quite helpful. Think about any dependencies you may have beforehand e.g. how are you going to integrate with an external service?
- **Start pair programming**
- **Definition of done**: Once you finish a story, run through a sanity checklist (approved by the team) to determine if you have sufficiently done the task to an agreed standard with your pairing partner.

### Silos

We talked about _how_ developers can work together but I think it's interesting to note _why_ working in silos can damage overall output.

- A developer can be shoehorned into building one specific part of the system just because they are most familiar with it. What happens when they leave or are on holiday for an extended period of time? Other team members are likely to take a long time to familiarise themselves with the code and question some decisions (especially if they are not documented).
- A developer can miss out on instant feedback and may need to wait for a pull request (PR) review depending on the teams workflow. Having another pair of eyes providing instant feedback and understanding the thought process is more efficient than reading a PR and requesting changes (and then reviewing subsequent changes on top of that). In fact through pairing, there would not be a need to raise a PR since your pairing partner has worked on the code with you. If you pair regularly and trust the developers on your team, why not consider trunk-based development? Make small iterative changes and _really_ practice continuous integration by working on the most recent codebase.
- Two heads are better than one. To design code that is sustainable, scaleable and all the lovely things developers want, you need to think about it. Discussing ideas, possibilities and architecture with another developer usually always produces a better outcome.

### Benefits to pair programming

- Knowledge sharing. Mitigating knowledge gaps leads to better decision making, increases the ability to work smarter and faster as a team and increases engagement.
- Coaching and mentoring opportunities. More experienced developers have a chance to nurture less experienced developers and hone in on their coaching or teaching skills.
- Increased communication. It is an opportunity to ask questions and discuss the reasoning behind the code. If you can explain the code to another developer well, it probably makes sense. If you can explain the code to a non-techie, then you are definitely doing something right - it's a very useful skill to explain code in lament terms.
- Create and build relationships with other developers. You are likely to talk about other non-work related things and connect with your colleagues on a deeper level. Having this strong synergy and an open, comfortable relationship within a team will foster a fun culture where people enjoy coming to work, trust each other and thus accelerate productivity.

Overall, I think pair programming can be quite enjoyable if done correctly. Here are my top tips to pairing:

1. Don't hog the keyboard. **Switch roles** often and work in small steps. Maybe try to use the [Pomodoro technique](https://en.wikipedia.org/wiki/Pomodoro_Technique#Description)?
2. **Take breaks!** Thinking, talking and typing can get exhausting so it's refreshing to take frequent short breaks.
3. Are you in a **"pairing mood"?** If not, be honest about it and maybe work on something that can be done by yourself for a while.
4. **Rotate the pairs**. Anchor a story to one developer but change their pair partner everyday. This spreads knowledge and forces the developer anchored to the story to reevaluate and reinforce their decisions by explaining them again.

Sources: [Pair programming illustration by Storyset](https://storyset.com/web)