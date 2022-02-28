# Loji (loading-emoji)
A loading emoji plugin for golang


## Usage

```go
package loading

import (
	"time"
	
	"github.com/chenjinya/loji"
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