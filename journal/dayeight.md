# Day Eight

I wrote a lot more code today. I was able to leverage some interesting data
structures and algorithms which I have previously had opportunity to utilize.
The *k*-dimensional tree and nearest neighbor are particularly interesting to
me as they allow you to avoid brute force but they require some fine-tuning.
Basically, I repeatedly ran the program, incrementing the *k* argument, until
the return value of `calculateCircuitVolume` became stable. If *k* is too small,
some of the closest point-pairs may be omitted, because it arbitrarily caps
how many point-pairs to return for each point. In a case where a point
has fifteen nearest point-pairs which should also be within the list of closest
point-pairs, but *k* is set to ten, then five of them will not be returned,
and therefore will be omitted from the list of closest point-pairs.

From the standpoint of programming in **Go**, today highlighted some of its strengths
and weaknesses. Part of why so much code had to be written is due to the
clunky implementation of "heap". What is one line of code in **C#** is twenty
three in **Go**. And for your efforts, you are rewarded with methods that return
`any`, forcing you to write yet more boilerplate to get back out what you put in.
On the other hand, the constraints of the type system at times were helpful.
Early on, I found myself trying to mimic constants and static members from **C#**.
The result was so unsatisfactory that I came up with something much better—
the `Point` interface is remarkably simple and the implementation of `KDTree`
is therefore both quite general and straightforward.

Finally, I wanted to remark upon the use of AI for today's challenge. The
core set of procedures was entirely my own idea, but I found **Claude**
indispensable as a discussion partner when dealing with the vagaries
of **Go**'s type system. As the length of the program increased, I also
found myself more and more inclined to simply tell **Claude** what to do
rather than type everything out myself.

## Postscript

As we come to the final third of the challenge, I expect the difficulty
to only increase. Despite that fact, I find myself bothered by various
discrepancies in my work, exemplified by today's solution. There is
significant duplication between days with regards to file loading. The
error handling scheme is not really thought through. Should `JunctionBox`
be treated as a value type or a reference type? I went with value type,
but then call `toPointers` to build the *k*-dimensional tree.

It is the nature of this challenge that I will not have lots of time to
work on it or revisit day-by-day, but I wanted to note my discomfort
at least, in case readers are in any doubt.
