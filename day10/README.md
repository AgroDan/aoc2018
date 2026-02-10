# Day 10: The Stars Align

Score one for AI on this one. I used a lot of fun things for this puzzle, specifically my `aocutils.Coord` object and my favorite `Runemap` utility. I typically use it for puzzles where the input string is a gigantic 2-Dimensional map, but for this one I made kind of an "Ad-Hoc" runemap so I could print it appropriately. Unfortunately the only way I had to create a runemap on the fly was to hand it a string slice (`[]string`) so it could run it, and instead of re-engineering that method I just created that exact string slice to create a runemap, then just added the points to the map as expected. Looking back, I should probably create a function that generates a runemap of a specific length filled with blank runes, but that's for future Dan.

Anyway, so AI to the rescue. Not in doing everything for me, but rather for debugging my crap logic! I resorted to an old puzzle from 2024, [Restroom Redoubt](https://github.com/AgroDan/aoc2024/tree/main/day14), which was a bunch of robot vacuums scattered around a room that would eventually create a Christmas Tree picture. This was fun because I had to check for adjacency to determine if it would create a picture, and this was no different. Only problem is for whatever reason I couldn't detect two blocks that were adjacent because...I was only looking at blocks diagonal from them.

This line:

```go
if aocutils.Abs(bc[i].X-bc[j].X) == 1 && aocutils.Abs(bc[i].Y-bc[j].Y) == 1 {
    isAdjacent = true
    break
}
```

Should rather be

```go
if aocutils.Abs(bc[i].X-bc[j].X) <= 1 && aocutils.Abs(bc[i].Y-bc[j].Y) <= 1 {
    isAdjacent = true
    break
}
```

Such a minor bug amounts to so much headache. I was happy to see that this was the only issue with my code though, so that was nice at least.

Part 1 was just the building of the light beams, and Part 2 was just "how many seconds passed before the message appears?"

I just put in the number I already printed...I did part 1 and part 2 without even realizing I was doing them both. Always sweet when that happens.