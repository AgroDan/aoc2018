# Day 19: Go With The Flow

Getting flashbacks to AOC 2024, just rebuilding that adder...this was similar I guess, only replace the logic gates with what I can only describe as "bastardized assembly." I'm seeing a theme to this year's AOC, and so far I'm at least learning cool new stuff.

This particular one had me printing a TON of debugging output just to figure out how best to handle this. Part one, super easy. Just recreate the instructions defined in [Day 16](../day16). I modified it a little bit to be a bit more flexible for this challenge. Once I assigned all the registers to `0`, I let it rip and it finished pretty quickly. Part two was deceptively difficult. Instead of initializing the registers to all `0`s, set register `[0]` to `1`, then set everything else to `0`. This silly little change caused a waterfall of emotions.

First, letting it go without changing much unsurprisingly would crunch away and seemingly do nothing. With the test data it looks like it just hits an infinite loop. For the challenge data it took a little bit of reverse engineering to find out exactly what it was doing.

I had it print out states and instructions for each iteration just to make sure it wasn't stuck in a loop, sure enough it was incrementing a value. After looking through each instruction I determined that it seemed to want to compare its values against register `[4]` (remember, index at `0`). Stepping through all of the checks it would run, it would multiply the contents of register `[1]` and register `[2]`, then compare that to the value at register `[4]`. If those values were the same, it would add it to register `[0]`. It would increment register `[1]` by `1` and run this check again. After it would increment greater than the value of register `[4]`, it would increment register `[2]` by 1, set register `[1]` to `0` and begin this loop over and over again. This is extremely computationally expensive. I had to use a shortcut.

Since it was just getting the sum of all the divisors of register `[4]`, I ran it `10,000` times and stopped it, got the value of register `[4]`, used my `GetDivisors()` function and added them all together.

This may have been a weird way of going about it, but I think it worked. Well, I _know_ it worked since I solved it, but I feel like I cheated with that new function I wrote. Oh well, maybe this is the intended route?

Here's a sample of some of my notes:

```
at this state: [0, 30754, 1, 0, 10551293, 6]

1. addi 5 1 5, adds 1 to [5], or the IP. Then increments again.

[0, 30754, 1, 0, 10551293, 8]

2. addi 1 1 1, adds 1 to 30754.

[0, 30755, 1, 0, 10551293, 9]

3. gtrr 1 4 3, is 30755 > 10551293, if so set [3] to 1.

[0, 30755, 1, 0, 10551293, 10]

4. addr 5 3 5, add [5] - 10 and [3] - 0  and set to [5]. Then increment again the IP.
	NOTE: If [1] > [3], then the IP becomes 12!

[0, 30755, 1, 0, 10551293, 11]

5. seti 2 6 5, set value of [2] to [5] (then increment IP)

[0, 30755, 1, 0, 10551293, 3]

6. mulr 2 1 3, multiply 1 and 30755 and set it to [3]

[0, 30755, 1, 30755, 10551293, 4]

7. eqrr 3 4 3, if 30755 == 10551293, set [3] to 1. Otherwise set to 0.

[0, 30755, 1, 0, 10551293, 5]

8. addr 3 5 5, add 0 + 5, set to [5] (then increment the IP

[0, 30755, 1, 0, 10551293, 6]

9. addi 5 1 5, add 1 to 6, then increment the IP

[0, 30755, 1, 0, 10551293, 8]

...this then repeats. Let's see what happens if we increment all the way to 10551293:

[0, 10551294, 1, 0 10551293, 9]

0: gtrr 1 4 3, is 10551294 > 10551293, if so set [3] to 1

[0, 10551294, 1, 1, 10551293, 10]

1: addr 5 3 5, add 10 and 1, then increment IP

[0, 10551294, 1, 1, 10551293, 12]

2: addi 2 1 2, add 1 to 1 and set it to [2]

[0, 10551294, 2, 1, 10551293, 13]

gtrr 2 4 3, is 2 > 10551293, if so set 1 to [3], otherwise 0

[0, 10551294, 2, 0, 10551293, 14]

addr 3 5 5, add 0 and 14 together, set it to [5], then increment ip

[0 10551294, 2, 0, 10551293, 15]

seti 1 2 5, store 10551294 into [5]

end program

value would be 10551294...but let's see what happens when we increment all the way to a known divisor (which is what I suspect it's doing...)

[0, 199081, 53, 0, 10551293, 3]

6. mulr 2 1 3, multiply 53 and 199081 and set it to [3]

[0, 199081, 53, 10551293, 10551293, 4]

7. eqrr 3 4 3, if 10551293 == 10551293, set [3] to 1

[0, 199081, 53, 1, 10551293, 5]

addr 3 5 5, add [3] and [5] and set to [5], then increment IP

[0, 199081, 53, 1, 10551293, 7]

addr 2 0 0, add [2] and [0] and set to [0]

[53, 199081, 53, 1, 10551293, 8]

addi 1 1 1, add [1] and 1, set it to [1]

[53, 199082, 53, 1, 10551293, 9]

gtrr 1 4 3, if [1] > [4], set [3] to 1. 199082 > 10551293

[53, 199082, 53, 0, 10551293, 10]

addr 5 3 5, add 10 and 0, set to [5], then increment IP

[53, 199082, 53, 0, 10551293, 11]

seti 2 6 5, store [2] into [5], then increment IP

[53, 199082, 53, 0, 10551293, 3]
```

Instead of letting it compute the whole thing and me using that `GetDivisors()` function, this runs for `47.8453ms`. I'll take it.