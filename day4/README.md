# Day 4: Repose Record

I really liked this one!

This was cool because it involved reading log files, something I have a tendency to do quite frequently. This means that I have to create an object for the log file entries, then use a sort function to sort them chronologically (I used the [slices.SortFunc()](https://pkg.go.dev/slices#SortFunc) function for that one, came in handy), then loop through each entry noting the state of each guard. Once that was done, I basically just subtracted the difference of the minutes when the guard was asleep versus awake, which thankfully was only within one hour, then created a heatmap of each minute's frequency. Just an array the size of `60`, each index referring to the minute in the hour, and holding a number for the amount of times that particular guard was asleep during that minute.

After that, it was just a matter of doing some metrics work on each guard and finding the frequency. I might do this in [Gravwell](https://www.gravwell.io/) just to see if I can, but that would be for another day.