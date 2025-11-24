Go programs express errors with error values. An Error is any type that implements the simple built-in error interface:

```go
type error interface {
    Error() string
}
```

When something can go wrong in a function, that function should return an error as its last return value. Any code that calls a function that can return an error should handle errors by testing whether the error is nil.

Sometimes new Go developers look at panic/recover, and think, "This is like try/catch! I like this!" Don't be like them.

I use error values for all "normal" error handling, and if I have a truly unrecoverable error, I use log.Fatal to print a message and exit the program.
