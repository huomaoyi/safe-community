/**
 * @Description: 
 * @Version: 1.0.0
 * @Author: liteng
 * @Date: 2020-02-02 14:33
 */

package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync/atomic"
	"time"
)

const (
	KBSize = 1024 << (10 * iota)
	MBSize
	GBSize
)

const (
	defaultMaxFileSize   int64 = 20 * MBSize // 20MB
	defaultMaxFolderSize int64 = 5 * GBSize  // 5GM
)

type filewriter struct {

	// path to save all logs here
	folderPath    string
	maxFileSize   int64
	maxFolderSize int64

	writeChan  chan []byte
	writeReply chan struct{}
}

func newLogFile(path string) (*os.File, error) {
	if dir, err := os.Stat(path); err == nil {
		if !dir.IsDir() {
			return nil, fmt.Errorf("open %s: not a directory", path)
		}
	} else if os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0740); err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	return os.OpenFile(filepath.Join(path,
		time.Now().Format("2006-01-02_15.04.05"))+".log",
		os.O_RDWR|os.O_CREATE, 0640)
}

func NewFileWriter(path string, maxFileSize, maxFolderSize int64) *filewriter {

	w := &filewriter{
		folderPath:    path,
		maxFileSize:   defaultMaxFileSize,
		maxFolderSize: defaultMaxFolderSize,
		writeChan:     make(chan []byte),
		writeReply:    make(chan struct{}),
	}

	if maxFileSize > 0 {
		w.maxFileSize = maxFileSize * MBSize
	}
	if maxFolderSize > 0 {
		w.maxFolderSize = maxFolderSize * GBSize
	}

	if w.maxFolderSize < w.maxFileSize {
		panic("Max folder size must be greater than max file size")
	}

	go w.writeHandler()

	return w
}

func (w *filewriter) Write(buf []byte) (int, error) {
	w.writeChan <- buf
	<-w.writeReply
	return len(buf), nil
}

func (w *filewriter) writeHandler() {
	var current *os.File
	var fileSize int64
	var folderSize int64

	files, _ := ioutil.ReadDir(w.folderPath)
	for _, f := range files {
		folderSize += f.Size()
	}

	for {
		buf := <-w.writeChan
		var bufLen = int64(len(buf))

		if atomic.AddInt64(&fileSize, bufLen) >= w.maxFileSize || current == nil {

			// create new log file
			file, err := newLogFile(w.folderPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "New log file %s, err %v\n", w.folderPath, err)
			}

			// close previous log file in an another goroutine
			go func(file *os.File) {
				if file != nil {
					if err := file.Close(); err != nil {
						fmt.Fprintf(os.Stderr, "Close log file %s, err %v\n", file.Name(), err)
					}
				}
			}(current)

			current = file
			atomic.StoreInt64(&fileSize, 0)
		}

		// write buf to file
		if _, err := current.Write(buf); err != nil {
			fmt.Fprintf(os.Stderr, "Write log file %s, err %v\n", current.Name(), err)
		}

		w.writeReply <- struct{}{}

		if atomic.AddInt64(&folderSize, bufLen) > w.maxFolderSize {
			var total int64
			files, _ := ioutil.ReadDir(w.folderPath)
			for _, f := range files {
				total += f.Size()
			}

			// get the oldest log file
			oldestLogFile := files[0]
			// remove it
			err := os.Remove(filepath.Join(w.folderPath, oldestLogFile.Name()))
			if err != nil {
				atomic.StoreInt64(&folderSize, total)
				fmt.Fprintf(os.Stderr, "Remove log file %s, err %v\n", oldestLogFile.Name(), err)
			} else {
				atomic.StoreInt64(&folderSize, total-oldestLogFile.Size())
			}
		}
	}
}