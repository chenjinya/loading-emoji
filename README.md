# Loji (loading-emoji)
A loading emoji plugin for golang

![pic](./animate.gif)

## Usage

```go
import (
    "time"
    
    "github.com/chenjinya/loji"
)

func main() {
    countdown := 60
    lo := loji.NewLoading("loading")
    for true {
        lo.Loading("冷却")
        countdown--
        time.Sleep(1 * time.Second)
        if countdown <= 0 {
            break
        }
    }
    lo.Stop()
}

```


## License

MIT