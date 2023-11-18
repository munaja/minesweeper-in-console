# minesweeper-in-console
A console version of a welknown game called minesweeper.

This is just a practice to understand more about dynamic programming. EXPECT BUGS!!

## Concept
- The is board with wide of width * height
- the wide represents some boxes count
- Each box has status of revealed and mined
- Each time box is selected through a coordinate, it will be checked for mine
  - If there is no mind check mark as revealed then check the unrevealed boxes around for mine, which means there is a lopp here
  - If there are mines around, label the box with count of the mines around
  - If it hits the mine game is over with lost status
- If there is no more boxes with mine, the game is over with winning status

## Note
- It took less than 1 hour to apply the basic concept regarding box.
- It took around 3 hours to finish the program completely including error checking and drawing nice board.
