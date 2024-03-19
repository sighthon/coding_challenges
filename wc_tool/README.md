Problem: https://codingchallenges.fyi/challenges/challenge-wc/

#### Installation
- go build .
- ./wc_tool -option file_name
- Available options:
    - c: for byte count
    - l: for line count
    - w: for word count
    - m: for character count


#### Potential improvements
- needs to read entire file to memory, can this be improved?

#### Learnings:
- Golang
    - flag parsing
    - read file bytes
    - Rune v/s bytes
- UTF-8 encoding v/s ASCII encoding
    - For files encoded in single-byte character sets like ASCII, these two counts will be the same. However, for files encoded in multibyte character sets like UTF-8, where a single character can be represented by multiple bytes, wc -m will give a more accurate count of characters.
    - character v/s byte


#### Results
```
╰─⠠⠵ wc -m test.txt
  339292 test.txt
╰─⠠⠵ wc -c test.txt
  342190 test.txt
╰─⠠⠵ wc -l test.txt
    7145 test.txt
╰─⠠⠵ wc -w test.txt
   58164 test.txt
╰─⠠⠵ wc test.txt
    7145   58164  342190 test.txt


╰─⠠⠵ ./wc_tool -m test.txt
339292 test.txt
╰─⠠⠵ ./wc_tool -c test.txt
342190 test.txt
╰─⠠⠵ ./wc_tool -l test.txt
7145 test.txt
╰─⠠⠵ ./wc_tool -w test.txt
58164 test.txt
╰─⠠⠵ ./wc_tool test.txt
7145 58164 342190 test.txt
```


#### References:
- https://blog.ippon.tech/no-framework-approach-to-building-a-cli-with-go/
- https://pkg.go.dev/flag
- https://www.scaler.com/topics/golang/golang-read-file/
- Rune: https://medium.com/@golangda/golang-quick-reference-runes-2f3f117987a6#:~:text=Runes%20in%20Go%20represent%20Unicode,various%20languages%20and%20symbol%20sets.