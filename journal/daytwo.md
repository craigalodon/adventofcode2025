# Day Two

I had this strong intuition that it should be possible to mathematically generate
the invalid IDs, but due to brain fog, I couldn't ever really work it out. My
initial implementation was to check most of the numbers in the range via
brute force, but use some clever checks to skip over ranges of numbers that
could not contain an invalid ID.

When I got to part two, it became obvious that was not going to work. At that
point, being low on caffeine, I consulted the AI oracle (ChatGPT 5.1),
explaining my intuition that it should be possible to generate the invalid IDs,
and low-and-behold, it is possible.

I cannot really take credit for the core implementation of the solution at all.
I do think I understand it, and it is most definitely more efficient than the
various brute force and regex solutions which I have seen many people describe.

In hindsight, I think I would have needed help with the algebra, but the core
of the algorithm is pretty approachable. It basically consists in finding the
proper divisors of the various decimal lengths present in the ranges, generating
all possible numbers with a decimal length equal to the proper divisors, and then
checking to see if each generated number, repeated the appropriate amount of times,
is present in the ranges.
