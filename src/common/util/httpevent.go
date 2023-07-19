package util

import (
	"common/tlog"
	"context"
	"io"
	"net/http"
	"strings"
)

type HttpEvent struct {
	Event string
	Data  string
}

func (e *HttpEvent) String() string {
	var sb strings.Builder
	if len(e.Event) > 0 {
		sb.WriteString("event: ")
		sb.WriteString(e.Event)
		sb.WriteByte('\n')
	}

	if len(e.Data) > 0 {
		sb.WriteString("data: ")
		sb.WriteString(e.Data)
		sb.WriteByte('\n')
	}
	sb.WriteByte('\n')
	return sb.String()
}

func HttpSendEvents(ctx context.Context, w http.ResponseWriter, out <-chan HttpEvent, finishSig chan struct{}) {
	h := w.Header()
	flusher := w.(http.Flusher)

	h.Set("Content-Type", "text/event-stream; charset=utf-8")
	h.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	h.Set("Connection", "keep-alive")
	//h.Set("Access-Control-Allow-Origin", "*")
	//h.Set("Content-Encoding", "gzip")
	w.WriteHeader(http.StatusOK)
	flusher.Flush()

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case ev, ok := <-out:
			if !ok {
				break loop
			}
			_, err := io.WriteString(w, ev.String())
			if err != nil {
				tlog.Error(err)
				break loop
			}
			flusher.Flush()
		}
	}

	close(finishSig)
}
