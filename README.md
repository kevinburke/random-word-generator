# Random Word Generator

This is a random word generator, which uses /usr/share/dict/words as the
source of truth. Proper words (those which start with capital letters) are not
shown.

```
$ random-word-generator --help
Usage of random-word-generator:
  -count int
    	Number of words to print (default 1)
```

## Installation

I haven't produced binaries, but I probably can if people want. For now, run:

```
go install github.com/kevinburke/random-word-generator
```

## Notes

There are a lot of words you've never heard of in this list.

```
$ random-word-generator --count 10
palaeoeremology
dividualism
cankereat
breakthrough
coexpire
wiretail
serodermatosis
undefatigable
unsentient
polyprotodont
```
