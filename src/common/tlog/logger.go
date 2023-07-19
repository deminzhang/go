package tlog

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"time"
)

type Logger struct {
	fileSize int64
	fileNum  int
	fileName string
	dir      string
	host     string
	level    LEVEL
	byteBuff bytes.Buffer
	queue    chan *Msg
	f        *os.File
	w        *bufio.Writer
	ticker   *time.Ticker
	end      chan struct{}
}

type Msg struct {
	line  int
	file  string
	level LEVEL
	msg   []byte
}

func newLogger(c *Config) *Logger {
	if c.Debug {
		return nil
	}

	os.MkdirAll(c.Dir, 0755)
	fname := path.Join(c.Dir, c.FileName+".log")
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	l := &Logger{
		fileSize: int64(c.FileSize * 1024 * 1024),
		fileNum:  c.FileNum,
		fileName: fname,
		dir:      c.Dir,
		level:    getLevel(c.Level),
		queue:    make(chan *Msg, 102400),
		f:        f,
		w:        bufio.NewWriterSize(f, 1024*1024),
		ticker:   time.NewTicker(2 * time.Second),
		end:      make(chan struct{}),
	}
	l.host, _ = os.Hostname()

	return l
}

func (l *Logger) run() {
	if l != nil {
		go l.flushLoop()
		go l.writeLoop()
	}
}

func (l *Logger) stop() {
	l.ticker.Stop()
	close(l.queue)
	<-l.end

	if l.w != nil {
		l.w.Flush()
		if l.f != nil {
			l.f.Close()
		}
	}
}

func (l *Logger) writeLoop() {
	for a := range l.queue {
		if a == nil {
			l.w.Flush()
			fileInfo, err := os.Stat(l.fileName)
			if err != nil {
				if os.IsNotExist(err) {
					l.f.Close()
					os.MkdirAll(l.dir, 0755)
					l.f, _ = os.OpenFile(l.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
					l.w.Reset(l.f)
				}
			} else if fileInfo.Size() > l.fileSize {
				l.f.Close()
				os.Rename(l.fileName, l.makeOldName())
				l.f, _ = os.OpenFile(l.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				l.w.Reset(l.f)
				l.rmOldFiles()
			}
		} else {
			l.w.Write(l.makeLog(a))
			l.byteBuff.Reset()
		}
	}

	close(l.end)
}

func (l *Logger) flushLoop() {
	for range l.ticker.C {
		l.queue <- nil
	}
}

func (l *Logger) makeOldName() string {
	t := fmt.Sprintf("%s", time.Now())[:19]
	tt := strings.Replace(
		strings.Replace(
			strings.Replace(t, "-", "", -1),
			" ", "", -1),
		":", "", -1)
	return fmt.Sprintf("%s.%s", l.fileName, tt)
}

func (l *Logger) p(level LEVEL, args ...interface{}) {
	if l == nil {
		file, line := getFileNameAndLine()
		var w bytes.Buffer
		fmt.Fprintf(&w, "%s %s %s:%d", genTime(), levelText[level], file, line)
		for _, arg := range args {
			w.WriteByte(' ')
			fmt.Fprint(&w, arg)
		}
		fmt.Println(w.String())
		return
	}
	if level >= l.level {
		file, line := getFileNameAndLine()
		var w bytes.Buffer
		for _, arg := range args {
			fmt.Fprint(&w, arg)
			w.WriteByte(' ')
		}
		b := w.Bytes()

		select {
		case l.queue <- &Msg{file: file, line: line, level: level, msg: b}:
		default:
		}
	}
}

func (l *Logger) pf(level LEVEL, format string, args ...interface{}) {
	if l == nil {
		file, line := getFileNameAndLine()
		fmt.Println(fmt.Sprintf("%s %s %s:%d", genTime(), levelText[level], file, line),
			fmt.Sprintf(format, args...))
		return
	}
	if level >= l.level {
		file, line := getFileNameAndLine()
		var w bytes.Buffer
		fmt.Fprintf(&w, format, args...)
		b := w.Bytes()

		select {
		case l.queue <- &Msg{file: file, line: line, level: level, msg: b}:
		default:
		}
	}
}

func (l *Logger) makeLog(a *Msg) []byte {
	w := &l.byteBuff
	fmt.Fprintf(w, "%s %s %s %s:%d ", genTime(), l.host, levelText[a.level], a.file, a.line)
	w.Write(a.msg)
	w.WriteByte('\n')
	return w.Bytes()
}

func (l *Logger) rmOldFiles() {
	if out, err := exec.Command("ls", l.dir).Output(); err == nil {
		files := bytes.Split(out, []byte("\n"))
		totol, idx := len(files)-1, 0
		for i := totol; i >= 0; i-- {
			file := path.Join(l.dir, string(files[i]))
			if strings.HasPrefix(file, l.fileName) && file != l.fileName {
				idx++
				if idx > l.fileNum {
					os.Remove(file)
				}
			}
		}
	}
}

func genTime() []byte {
	now := time.Now()
	year, month, day := now.Date()
	hour, minute, second := now.Clock()
	return []byte{
		'2', '0', byte((year%100)/10) + 48, byte(year%10) + 48, '-',
		byte(month/10) + 48, byte(month%10) + 48, '-', byte(day/10) + 48, byte(day%10) + 48, ' ',
		byte(hour/10) + 48, byte(hour%10) + 48, ':', byte(minute/10) + 48, byte(minute%10) + 48, ':',
		byte(second/10) + 48, byte(second%10) + 48}
}

func getFileNameAndLine() (string, int) {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		return "???", 0
	}
	dirs := strings.Split(file, "/")
	sz := len(dirs)
	if sz >= 2 {
		return dirs[sz-2] + "/" + dirs[sz-1], line
	}
	return file, line
}
