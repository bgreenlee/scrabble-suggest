# Scrabble Suggest

Scrabble Suggest will take as input a list of letters and output a list of Scrabble words that can be made with those letters.

This was an exercise to help me learn a bit of Go. Any feedback on improving my code is very much welcomed.

## Usage

1. Build:

        ./build.sh
This will create two executables, `scrabble-indexer` and `scrabble-suggester`
2. Index:

        ./scrabble-indexer <path to word list>
This builds the index file. A public domain word list is included at `data/word_list_moby_crossword-flat/word_list_moby_crossword.flat.txt`. You can pass it a different word list if you have one. This only needs to be done once (unless you change the word list).
3. Run:

        ./scrabble-suggester <letters>
Lists all the Scrabble words, in order of descending point value, that can be made with the given letters.
Example:

        $ ./scrabble-suggester BRAD
        BARD (7)
        DARB (7)
        DRAB (7)
        BRAD (7)
        BAD (6)
        DAB (6)
        BRA (5)
        BAR (5)
        RAD (4)
        BA (4)
        AD (3)
        DA (3)
        AR (2)


To run the tests, just do `go test ./...`

## For the Lawyers

SCRABBLE® Brand Crossword Game is a registered trademark of Hasbro, Inc. in the United States and Canada.  Outside of the United States and Canada, the SCRABBLE® trademark is owned by J.W. Spear and Sons, PLC, a subsidiary of Mattel, Inc.  SCRABBLE® is not a generic term. All intellectual property rights in and to the game are owned in the U.S.A. by Hasbro, Inc.

Neither GitHub, the scrabble-suggest program, nor its author is in any way sponsored by or otherwise affiliated with Hasbro, Inc.
