package util

import (
	"golang.org/x/net/websocket"
	"net"
	"net/http"
	"strings"
)

type WebsocketConn struct {
	Request *http.Request
	Conn    *websocket.Conn
	h       *wsHandler
}

type WebsocketProcesser func(wc *WebsocketConn)

func NewWebsocketHandler(fn WebsocketProcesser) http.Handler {
	return &wsHandler{fn: fn}
}

func (wc *WebsocketConn) GetHeader(key string) string {
	return wc.Request.Header.Get(key)
}

func (wc *WebsocketConn) Host() string {
	host := wc.Request.Host
	pos := strings.Index(host, ":")
	if pos >= 0 {
		host = host[:pos]
	}
	return strings.ToLower(host)
}

func (wc *WebsocketConn) ClientIP() string {
	req := wc.Request
	clientIP := req.Header.Get("X-Forwarded-For")
	clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])
	if clientIP == "" {
		clientIP = strings.TrimSpace(req.Header.Get("X-Real-Ip"))
	}
	if clientIP != "" {
		return clientIP
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(req.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

func (wc *WebsocketConn) getHandshakeFunc() func(*websocket.Config, *http.Request) error {
	return func(conf *websocket.Config, req *http.Request) error {
		wc.Request = req
		return nil
	}
}

func (wc *WebsocketConn) getWebsocketHandler() websocket.Handler {
	return func(conn *websocket.Conn) {
		wc.Conn = conn
		wc.h.fn(wc)
	}
}

///////////////////////////////////////////////////////////////////
type wsHandler struct {
	fn     WebsocketProcesser
	config websocket.Config
}

func (wh *wsHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	wc := WebsocketConn{h: wh}
	s := websocket.Server{Config: wh.config, Handler: wc.getWebsocketHandler(), Handshake: wc.getHandshakeFunc()}
	s.ServeHTTP(w, req)
}
