# Day Six

Today was the first day I found myself rolling my eyes at part two. Here's a quick
overview of the issues.

- For the first time, I took the easy route and read the data in line by line
  and stored it as a two-dimensional grid. On all previous days, I either streamed
  the data or stored it in a one-dimensional array.
- To solve the problem, I found it necessary to normalize the data. See
  `leftJustifyGrid()`. While this did greatly simplify the solution, it felt
  like a compromise.
- The control flow feels very messy. There are errors being thrown in almost
  every method, which indicates to me that my model of the problem is wrong.
