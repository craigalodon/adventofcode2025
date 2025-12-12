# Day Four

Today was again straightforward. I don't think my solution is optimal. Probably
the best element of it is that I store the coordinates as a one-dimensional array
making the math operations pretty simple. But the use of multiple hashmaps, which
end up storing most of the data ingested, seems bad. Likewise, re-iterating through
all the rolls each time seems inefficient to me. I also don't like the rat's nest of
if statements and manual calculations in the `getAdjacent` function.

Still, it was very easy to modify from part one to part two. I have some doubts
about the way they do it—they complete a whole pass before moving any rolls, whereas
I think you could just remove rolls as you go, but I had stuff to do, so I went with
the solution presented here.
