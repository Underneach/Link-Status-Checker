package src

import (
	"bufio"
	"github.com/go-resty/resty/v2"
	"github.com/panjf2000/ants/v2"
	"os"
	"runtime"
	"sync"
	"time"
)

var (
	input        = bufio.NewReader(os.Stdin)
	linksList    []string
	threads      int
	err          error
	client       = resty.New()
	resp         *resty.Response
	resultList   []string
	checkedLinks = 1
	totalLinks   int
	validLinks   int
	invalidLinks int
	startTime    time.Time
	wg           sync.WaitGroup
	workPool, _  = ants.NewMultiPoolWithFunc(
		runtime.NumCPU(),
		threads,
		func(line interface{}) { Worker(line.(string)) },
		ants.RoundRobin,
		ants.WithPreAlloc(false),
	)
)
