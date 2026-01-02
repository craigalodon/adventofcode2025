# Day Nine

Another surprisingly easy day. The challenge was straightforward. Taking a
lesson from [day eight](./dayeight.md), I did not try to optimize the solution
in terms of compute or memory. I used a brute force approach to find candidate
point-pairs and then used a node ring to quickly check if the path crossed
through the rectangle. The solution is remarkably easy to read, IMO.

I'm doubtful that my code is idiomatic Go. I'm not sure how important that is
to me, however. I think Go is a good language for pedagogical activities like
this one, because of its simplicity, but I am drawn to the more expressive
type systems of functional and object-oriented languages. I do like Go's 
lack of syntactic sugar, though.

I continue to feel uncertain about whether I should be writing tests and
how otherwise professional I should make this repository. Today I'm grouping
commits into a single pull request. One problem with writing tests is that
there are [restrictions](https://adventofcode.com/2025/about) on saving the
puzzle inputs or solutions to your repository. But that doesn't mean I 
couldn't write tests for the various helper methods. However, my sense is
that writing tests isn't really a thing for coding challenges. IDK.

## Postscript

A [sharp Redditor](https://www.reddit.com/r/adventofcode/comments/1phywvn/comment/nwprqyf/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button)
pointed out an edge case where the polygon has a large concave angle, which my
original solution did not handle correctly. Unfortunately, that meant doing
some calculus. I ended up adapting solution two from [here](https://www.eecs.umich.edu/courses/eecs380/HANDOUTS/PROJ2/InsidePoly.html).

I am again in doubt about the computational efficiency, but I like that I didn't
have to dramatically change my solution to implement the additional checks.
As I cycle through the nodes, I sum the angles between the current node and each
corner. If I pass through the corner, then it is interior to the polygon.
If the sum of the angles is >= 2π, then the point is inside the polygon.

I ended up actually extracting each of the corner checks into their own pass
through the nodes, because it made for cleaner code and iterating through
the nodes is not a bottleneck. Checking the corners increased the runtime
by <= 10% on my machine.
