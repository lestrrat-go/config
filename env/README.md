# env

Unmarshal values from environment variables

# SYNOPSIS

```go
// TODO: Write an example test code, and paste it here
type Config struct {
    Foo string
    Bar int
}

func LoadConfig(c *Config) error {
    // Loads from MYAPP_FOO, MYAPP_BAR
    return env.NewDecoder(env.System).Prefix("MYAPP").Decode(c)
}
```

# DESCRIPTION

This library can be thought of as a fork of [github.com/kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig). The code was written from scratch, but the goals are the same: We would like to support fetching configuration information from environment variables.

The author initially attempted to use the library above, but there were a few things that needed changing to adapt to the author's needs. However, that would require modifying behavior of a code base that is being used by well established user base, and the author has been around long enough that it is quite hard to chanve existing behavior :)

So here is yet another library, instead of spending days arguing if my change is good for the World or not. If the library above decides to support the differences that I describe below, this library *may* be deprecated/changed/whatever.

## Supported Types

* string
* bool
* int, int8, int16, int32, int64
* uint, uint8, uint16, uint32, utin64
* float32, float64
* time.Time, time.Duration
* slices of the above items
* structs
* pointer to above types
* nested/embedded types
* types that implement env.Unmarshaler interface

See [env_test.go](./env_test.go) for more examples

## DIFFERENCES FROM `envconfig`

### Pointers are not auto-vivified

Using this library, the pointer-to-structs are left as nil given a struct like the following:

```go
type SubConfig struct {
    Foo string
    Bar string
}

type Config struct {
    Sub *SubConfig
}

var c Config

// Make sure that SubConfig cannot be populated
os.Unsetenv("SUB_FOO")
os.Unsetenv("SUB_BAR")

env.Unmarshal(&c)

// If SubConfig is _not_ populated, the struct pointed by
// the c.Sub pointer is left as nil
if c.Sub != nil {
    panic("c.Sub should be nil!")
}
```

The reason for this is to detect unitinialized fields.

### Flexible Source

This is a very minor detail, but you can specify where you get your environment variables from by passing a `Source` to the decoder object.

```go
env.NewDecoder(env.System).Decode(&c)
```

Here, we are using the `System` decoder, which just calls `os.LookupEnv`. But if you implement the `env.Source` interface, you can derive your variables from any source you would like.

This may come in handy if you want to do some elaborate testing.