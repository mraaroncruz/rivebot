# Rivebot

Ok. This is an attempt of using [rivescript](https://www.rivescript.com/) from any programming language by build in little crappy script in Golang.

It uses redis as a session store, so there's a hard dependency for you. I plan on just implementing a file store but as I said, this is a time saving process. My alternative was to write a parser in Ruby and I don't want to do that.

The next step would be to wrap some Go in an FFI backed Ruby C extension, but I'll see how far this will get me first.

## Use

```bash
# Install binary
go get github.com/pferdefleisch/rivebot

# Make sure redis is running
redis-server

# Run bot brain
rivebot -rs="./path/to/rivescript/directory" -message="this is a false statement" -session="12346"
```

This will print out your bot's response to stdout

## License

MIT
