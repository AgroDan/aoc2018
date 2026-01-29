# Day 2: Inventory Management System

Pretty straightforward. Look for words that have 2 or 3 letters repeated, then add the amount of IDs that satisfy that criteria. Just wrote a little struct with corresponding functions that count those out.

Then compared every single ID against every ID after it looking for a string of characters that are differing by exactly one letter. Since this is so early in the advent of code challenge I don't have to be terribly efficient yet, so I'll just let it run through the whole thing without too many shortcuts. Ran pretty quickly, just over `1ms`.