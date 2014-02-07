# Scrabble Suggest

Scrabble Suggest will take as input a list of letters and output a list of Scrabble words that can be made with those letters.

This was an exercise to help me learn a bit of Go. Any feedback on improving my code is very much welcomed.

## Usage

1. Build:

        ./build.sh
This will create two executables, `scrabble-indexer` and `scrabble-suggester`
2. Index:

        ./scrabble-indexer <path to word list>
This builds the index file. A standard word list is included at `data/words.txt`. You can pass it a different word list if you have one. This only needs to be done once (unless you change the word list).
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

## Author

Brad Greenlee <brad@footle.org>
