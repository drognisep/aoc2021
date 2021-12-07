This one seemed pretty straight forward, but I suspect there's a much more elegant solution.
The strategy I used was to traverse all numbers in the range 1 to max/2 and use that number as the base.
Then the base was compared to each number in the set to determine the cost to move from that position.

I used max/2 because the numbers are rather evenly distributed, but if that were unknown or not the case I would have to consider the full range.

For the second part, the difference was to traverse the number line to the distance from the base and add each value to the cost instead of considering only the linear distance.
