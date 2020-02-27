# vmix-go Examples

## scenes-list
Super-Simple list sample.  
Lists current available inputs,and Fades into first scene.
```
go run ./examples/scenes-list/main.go -addr "http://localhost:8088/"
```

## interactive
This is interactive shell example.
You can Fade or cut scene by interacting command shells that suggests input lists.
```
go run ./examples/interactive/main.go -addr "http://localhost:8088/"
```