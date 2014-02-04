#!/bin/sh

cd indexer
go build -o ../scrabble-indexer
cd ../suggester
go build -o ../scrabble-suggester
