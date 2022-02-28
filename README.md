# Loji (loading-emoji)
A loading emoji plugin for golang

![pic](./animate.gif)

## Usage

```go
package loading

import (
	"time"
	
	"github.com/chenjinya/loji"
)

func main() {
	l := loji.New()
	l.Loading("loading...")
	time.Sleep(2 * time.Second)
	l.Stop()
	time.Sleep(2 * time.Second)
}

```


## License

MIT