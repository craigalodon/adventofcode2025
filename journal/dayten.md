# Day Ten

What a day! The first part was easy enough, so I took my time doing things right.
I wrote tests, cleaned up variables quite a bit, and did some bitwise operations
since I rarely get to do them in real life.

When I read the problem description for part two, I knew I was not going to
solve it quickly. I did some reading on row reduction algorithms (see [here](https://en.wikipedia.org/wiki/Gaussian_elimination)
and [here](https://textbooks.math.gatech.edu/ila/row-reduction.html)). I spent
a lot of time solving the example problems by hand, making sure I understood
the mathematical procedure. In the back of my mind, I knew that once I figured
all that out, I'd still have my work cut out for me just setting up the data
and writing the problem-specific logic.

I was right. You can see that as of this writing, `jolt()` really needs to
be broken down into helper functions. I made a decision early on to use 
floating point numbers instead of Go's rational number library, and that caused 
me quite a bit of debugging pain as I was approaching the finish line
(but at least I didn't have to learn another library).

Still, I'm pretty proud of this solution, and feel like cleaning it up would
be straight forward. I think it is performant, given that many people on the 
Reddit thread were describing run times somewhat longer than mine, though it 
is hard to judge without knowing what devices they are using. Mostly, I
learned a lot and persevered through each the challenges as they came,
and what more can you ask from another day of the Advent of Code?
