# VisuAlGo

### Notes

[Run tests on terminal as](https://stackoverflow.com/a/78905626)

```
go test -v ./... | sed ''/---\ PASS:/s//"$(printf "\033[32m✅--- PASS:\033[0m")"/'' | sed ''/---\ FAIL:/s//"$(printf "\033[31m❌--- FAIL:\033[0m")"/''
```
