# Loji (loading-emoji)
A loading emoji plugin for golang

用来生成在控制台展示 Loading 样式的工具

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
        lo.Loading(fmt.Sprintf("冷却: %d", countdown))
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