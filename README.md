# loading-emoji
A loading emoji plugin for golang


## Usage

```go
package loading

import (
	"time"
	
	loji "github.com/chenjinya/loading-emoji"
)

func main() {
	l := loji.NewLoading("")
	l.Loading("loading...")
	time.Sleep(2 * time.Second)
	l.Stop()
	time.Sleep(2 * time.Second)
}

```


## License

MIT