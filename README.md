#Some code challenges

##Description

###TASK 1.
Given a non-empty string s and a list wordList containing a list of non-empty tokens, determine if s can be represented as a concatenation of tokens from the list (where each token may be used several times). You may assume the dictionary does not contain duplicate tokens.

s = "whataniceday"
wordList = ["a", "what", "an", "nice", "day"].
->  true

s = "dawhaty",
wordList = ["a", "what", "an", "nice", "day"].
-> false

s = "abc",
wordList = ["a", "ab", "bc"].
-> true

###TASK 2
Given a collection of intervals, merge all overlapping intervals.
For example,
Given [2,6],[8,10],[1,3],[15,18],[18,21] 
return [1,6],[8,10], [15, 21]

###Project structure explanation

The project has common GOPATH-independent structure allowing to place its files anywhere on FS.

Contains functional packages dedicated from executable wrappers (accords to common `cmd directory` convention)

minimum go version: 9.1


##Build

```
export GOPATH=`pwd` && \
cd ./src/tasks && dep ensure && cd - && \
go install tasks/cmd/...
```

##Test

```
go test tasks/...

```


##Run

```
./bin/tokens --list 'a nice green frog jump s on the road' --phrase anicegreenfrogjumpsontheroad
```

Expected output: `true`


```
./bin/tokens --list "a ab bc" --phrase abcaabc
```

Expected output: `true`

```
./bin/tokens --list "a ab bc" --phrase abcaabcc
```

Expected output: `false`


```
./bin/intervals --list '[[2,6],[8,10],[1,3],[15,18],[18,21]]'
```

Expected output: `[[1 6] [8 10] [15 21]]`

```
./bin/intervals --list '[[1,2],[3,4],[7,12],[9,14],[8,15],[20,25],[18,30],[35,40],[37,38],[42,43],[45,46]]'
```

Expected output: `[[1 4] [7 15] [18 30] [35 40] [42 43] [45 46]]`
