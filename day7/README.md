# Day 7: The Sum of Its Parts

Part two was bananas. But not too bad all in all.

The first part got me a bit twisted up in data structure hell. How do I properly represent this data in such a way that would be useful? I have a tendency to over-engineer things, which is something I consider a plus, personally. Better to future-proof, assuming the world will be asked of you come part 2. I'm usually wrong in my predictions 95% of the time, but if I'm ever right, oh it's that good stuff right there.

Anyways, the hardest part for part 1 was probably figuring out how best to handle this data structure. Should each step be an object? Should I just create a map and do nothing special with its contents? Either of these routes were fairly complicated for my little pea-brain to wrap my mind around, but ultimately I settled on two actual data structures:

```go
type step struct {
	name         rune
	requirements []rune
}
```

and

```go
type Instructions struct {
	manual map[rune]*step
}
```

With assigned functions for each. An `Instructions` object is a superset of `step`s. This may be overcomplicating it quite a bit, but it allows me to sweep through the list of instructions and just refer to functions to do things to maintain the list of next steps.

Part two was a bit more involved. I created one of my patented megafunctions(tm) to handle the concept of doling out tasks to all of the workers. I initially thought I could parallelize this but honestly this didn't need that level of complication. I just sequentially ran down the list of workers for every single "tick" of the clock. If a worker completed the task, it took the rune it was working on and appended it to the retval, then marked itself as idle. Then it runs an earlier function to find the next available instruction, pops it off the list. Repeat. An oversimplification but I annotated the function as much as I could.

Finished in `1.025ms`. Boom.