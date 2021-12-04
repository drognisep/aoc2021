This was a fun one! :)

As someone that has made tic-tac-toe a few times to learn new languages, this seemed like well trodden territory.
Even so, I got my row iteration logic wrong the first pass.
I ended up finding the problem by adding additional test cases, so I was very thankful I took the time to make them.

I separated all game board logic to its own type to make things simpler from the perspective of simulating the game state.
It also helped me to consistently represent different ways of interacting with the board that I couldn't have done with a lot of shared state.

The general board data pattern revolves around 2 structures:
* The data array holds whether that place on the board has been marked 
* The index map keeps track of what numbers have been placed on the board, and the data array at which they reside.

Using this structure, it's easy and efficient to determine if a win condition has been met and to tally score.

[Back](../README.md)
