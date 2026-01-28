# Advent of Code Template

This exists just so I can start up a new repo for the AOC whenever it comes out. This includes the scripts and utils directory.

## NOTE: This is a template, so change things when initializing

As of now, it doesn't look like Github allows for the usage of variables in this template, so certain things that are hard-coded into scripts will unfortunately remain that way. Make sure you do the following:

- Update the year in each of the `prep` scripts
- Update the dev location of your script(s) in each of the `prep` scripts
- Initialize the Go Workspace `go work init`
- Initialize the Utils module `go mod init utils`
- Use the utils module `go work use utils`
- Finally...edit this readme

I would create the utils directory as my own repository elsewhere as Go works well with that, but for the sake of brevity I'd like to keep all the code I use for the Advent of Code in the same repository.