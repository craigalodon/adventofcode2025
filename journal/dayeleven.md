# Day Eleven

Another easy day. I spent way too much time polishing the parser, and then I could
not believe it when naïve recursive search was all it took to solve part one.
I thought part two would be more complicated and wasted some more time on
various half-baked ideas before simply memoizing the original function, which 
was surprisingly effective. I probably should have done some more exploratory
analysis to confirm that it was a DAG problem, rather than assuming it wasn't
and wasting more time on it.
