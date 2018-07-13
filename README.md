# PathRouter

Package pathrouter is forked from [httprouter](github.com/julienschmidt/httprouter).

```go
router := pathrouter.New()
router.Register("GET", "/users/:id/pets", "cfg for this path")
extra, _, _ := router.Lookup("GET", "/users/666/pets")
if cfg, ok := extra.(string); ok {
		fmt.Println(cfg)
}

```

## TODO

test
