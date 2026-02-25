# Day 16: Chronal Classification

This was so fun! This was basically a simplified version of [Day 24 of 2024](https://adventofcode.com/2024/day/24), or maybe even [Day 17 of 2024](https://adventofcode.com/2024/day/17), as far as I was concerned. Extremely simple (comparatively) because generally they had you rebuild lower level computer instructions. This was cool because while the instructions were easy to understand, I had to use a little deduction to determine which opcode did what. To do this, I created some helper functions for me so I could deduce the reasoning myself with some good old fashioned pen and paper. Well a small whiteboard, but same idea.

First part was simple enough, just create all of the functions and determine how many had 3 or more _potential_ instructions. That didn't take very long to create at all.

The second part was where I had to force myself to be clever. I created a helper function that essentially helped me visualize just how many possible instructions each opcode could pertain to. It ran through every single instruction set which showed a before and after of every register set that I could use to determine validity. Then I created a really complicated (yet extremely helpful) datatype that consisted of the following:

```go
type opCodeMap map[int]map[string]int
```

Which I wanted to structure like `[opcode][{function name: <how many times this opcode returned a valid output>}]`, which would create a list that came out to the following:

```text
Opcode map:
Opcode 11:
        banr: 47
        seti: 47
Opcode 1:
        mulr: 40
        muli: 40
        addr: 40
        addi: 40
Opcode 13:
        gtrr: 46
        addr: 46
        addi: 46
        mulr: 46
        muli: 46
        banr: 46
        bani: 46
        gtir: 46
        gtri: 46
        borr: 46
        bori: 46
        setr: 46
        seti: 46
Opcode 6:
        addi: 59
        banr: 59
        borr: 59
        eqrr: 59
        mulr: 59
        bori: 59
        setr: 59
        gtir: 59
        gtri: 59
Opcode 3:
        addi: 59
        muli: 59
Opcode 14:
        bani: 55
        seti: 29
        banr: 26
Opcode 7:
        addi: 16
        mulr: 56
        muli: 40
Opcode 0:
        addr: 42
        borr: 51
        mulr: 9
Opcode 4:
        bani: 54
        bori: 54
        setr: 54
        eqri: 54
        gtir: 14
        muli: 54
        banr: 40
Opcode 5:
        bori: 20
        muli: 46
        seti: 26
        mulr: 46
        setr: 46
        gtri: 46
        gtrr: 46
        eqri: 26
        eqrr: 39
        addi: 20
        banr: 46
        bani: 46
        gtir: 39
        eqir: 46
        addr: 7
        borr: 7
Opcode 10:
        muli: 56
Opcode 15:
        mulr: 13
        muli: 13
        borr: 43
        bori: 43
        addr: 30
        addi: 30
Opcode 8:
        bani: 30
        bori: 57
        setr: 57
        muli: 7
        banr: 39
        borr: 57
        seti: 57
        addr: 14
        addi: 20
        mulr: 4
Opcode 9:
        addr: 59
        addi: 59
        borr: 59
        bori: 59
        gtir: 59
Opcode 12:
        addr: 37
        addi: 52
        borr: 52
        bori: 52
        seti: 52
        mulr: 15
Opcode 2:
        borr: 44
        bori: 44
        setr: 44
        gtir: 44
        mulr: 44
        muli: 44
        banr: 44
        bani: 44
        eqri: 44
        eqrr: 44
```

I broke out my little whiteboard and made a list of `0` to `15`, then went through every item above. If an opcode had only one possible function, that would be my starting point. In this case, I knew right off the rip that `OpCode 10 == muli`, since only one function worked for it. Then I would cycle through and try to eliminate more. So the next step was:

```text
Opcode 3:
        addi: 59
        muli: 59
```

Since I knew that `OpCode 10` is `muli`, then `OpCode 3 == addi`, since it couldn't be `muli`. So Now, `OpCode 10 == muli` and `OpCode 3 == addi`. Next...

```text
Opcode 0:
        addr: 42
        borr: 51
        mulr: 9
```

Another easy win, because `borr` bubbled up to the top with the amount of times it was correct, so then `OpCode 0 == borr`. Now I have `3` Opcodes. I continued this cycle, determining the opcodes until I had all 16. Once I did that, I hard-coded this operation into the `RunProgram()` function, and I was able to get register `0` from that once I set the beginning registers to `[0, 0, 0, 0]`.

That one was cool!