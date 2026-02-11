# Day 11: Chronal Charge

Man, I really wanted to use a runemap here. I didn't, but I could have. Instead of kinda created one for part two when I realized I needed to start building calculations based on layers of squares.

Part one I tried to over-simplify, basically just create a dinky little cell object that is basically a `Coord{}` object in my utilities, just without most of the fancy functions that work with it. Then I used the set of instructions for the challenge as the way of determining the power level. Then I created a `neighbors()` function which returned a list of coordinates surrounding the cell position, then tallied up the total amounts of every cell there based on the formula. I crunched through every possible position and got my answer after a hearty round of debugging. Still not _quite_ sure why the X/Y coordinates were `+1` what they should have been, but whatever it worked.

Part two was a weird memoization method to build what I needed. Now the size of the boxes drawn were arbitrary, only that _they were always square_. Rectangular boxes were right out, so at least I had that. My method was that I was going to start with a point, get the total power level of that point, then expand by one layer, get the total amount from there, expand another layer, etc -- all the way until I hit the border. If I hit the border, stop -- that's the most I can get out of that point. Then move onto the next point.

This quickly turned into a really large dataset to work with, so I decided to pre-compute as much as I could before delivering an answer. Basically, for every point in the grid, created _another_ dimension that was used for the layer. Where `Point[x][y][0]` was Layer 1, so I just needed the power level total of `[x][y]`. Then get the power level total of its neighbors for Layer 2, and since the total power levels of those neighbors were already calculated, just add them all up and set that to `Point[x][y][1]`. Then for Layer 3, add to the total already calculated in `Point[x][y][1]` and get the total for every bordering point around that pre-computed box. Continue until we hit a border.

Once that was all calculated, go through every single item and check for the largest score. Note the position, enter the coordinates, done.

This took some crunch time! Completed in `4.1754892s`.