package logx

import (
	"sync/atomic"
	"time"
	"fmt"
)

func poller() {
	atomic.SwapUint32(&logger.look, uint32(coreRunning))
	if err := logger.loadCurLogFile(); err != nil {
		fmt.Println()
		if err = logger.createFile(); err != nil {
			panic(err)
		}
	}
	go logger.signalHandler()
	ticker := time.NewTicker(time.Millisecond * time.Duration(pollerinterval))
	now := time.Now()
	next := now.Add(time.Hour * 24)
	next = time.Date(
		next.Year(),
		next.Month(),
		next.Day(),
		0, 0, 0, 0,
		next.Location())
DONE:
	for {
		select {
		case <-logger.closeSignal:
			ticker.Stop()
			break DONE
		case <-ticker.C:
			if logger.fileWriter.Buffered() > 0 {
				logger.sync()
			}
		case n := <-logger.bucket:
			logger.fileWriter.Write(n.Bytes())
			logger.fileActualSize += n.Len()
			if logger.rotate() {
				logger.fileWriter.Reset(logger.file)
			}
			logger.release(n)
		}
	}
}
