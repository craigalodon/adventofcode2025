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
