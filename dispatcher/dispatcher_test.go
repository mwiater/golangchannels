package dispatcher_test

import (
	"fmt"
	"testing"

	"github.com/mattwiater/golangchannels/dispatcher"
)

var table = []struct {
	workerCount int
}{
	{workerCount: 1},
	{workerCount: 2},
	{workerCount: 4},
	{workerCount: 8},
}

func BenchmarkRun(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("Worker Count: %d", v.workerCount), func(b *testing.B) {
			dispatcher.Run(1, v.workerCount, 16)
		})
	}
}
