55 day Go
==========

Learning journey with Go programming language - the "Golang" with "gopher" mascot.

Some tips for beginners/novices as a "cheat sheet" to learn quicker.
Mostly language features & stdlib up to Golang 1.18 release (2022 Mar), though 1.20+ is recommended.

### Usage

**Basic Usage**

Build and run (given setup already done $GOPATH $GOROOT etcetera) in each folder (e.g. `src/0-hl`):
```bash
  $ cd src; cd 0-hl
  $ go run .
```

Usually there is one .txt file in each folder in "src" directory, which includes the steps to build and run.
And it even contains the sample output of each instruction/command (mostly)!

### Samples
```bash
  $ cd ../hello
  $ go build hello.go
```
When BUILD OK (no error in standard output/console), we can proceed to running and testing.

#### Running samples

```
  $ ./hello
  $ go run hello.go
  $ go run .
  $ cd ../greetings
  $ go test example.com/greetings
  $ go test . -v --count=1
```

Likewise, you can run "4-slices" or "5-range", it's up to you!

#### References

https://gitlab.com/ducquoc/55-day-go

https://github.com/ducquoc/42-day-kotlin

https://bitbucket.org/ducquoc/dq-training

https://ducquoc.github.io/

https://gobyexample.com/

https://www.digitalocean.com/community/tutorial-series/how-to-code-in-go
