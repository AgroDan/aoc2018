# Day 6: Chronal Coordinates

This tripped me up quite a bit, but all in all it was a good challenge. What caused me so much pain was that in part one, once I got all of the calculations set up properly I pulled a bone-headed move with my [`FindClosestPoint()`](chronalcoords/miscpoints.go) function, which was supposed to loop through every single coordinate provided and find the closest point (as the name suggests). The problem is that the first time I found two coordinates that were equidistant I exited and marked the point as having more than one coordinate that was closest. I should have checked every single coordinate's distance _first_ before I made that assumption, because there are plenty of points on this map that share the same distance between two provided coordinates...they just aren't the closest. Oops.

You'll see some of my debugging as I try and figure out why the numbers didn't add up. I'll leave it in as a testament to my unholy shame.

Part two was mercifully easier than part one. Mostly because I had basically done most of the work already, I just needed to specify a new function to calculate another metric. That took me all of 5 minutes to get.

Since we're dealing with "infinites," or rather a map that doesn't have a beginning or end, I had to create a boundary so I could work with some solid numbers. Basically a border around all the coordinates allows me to work with a fairly large map area, but nothing too serious. Dealing with tens of thousands of coordinates rather than hundreds of trillions, which would be a showstopper for my general "brute-forcing all the coordinates" method. Luckily we're still early enough on in the challenge that I can feasibly do exactly that and not tax my CPU too much.

This challenge was pretty easy to calculate since all it really did was some simple math on every single coordinate. Managed to solve in `21.9984ms`.