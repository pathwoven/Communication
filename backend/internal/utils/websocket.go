package utils

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 存储user_id和websocket连接的映射
var websocketStore = struct {
	sync.RWMutex
	connections map[uint32]*websocket.Conn
}{connections: make(map[uint32]*websocket.Conn)}

// 升级器
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func UpgradeWebsocket(userID uint32, c *gin.Context) (error, *websocket.Conn) {
	// 升级websocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return err, nil
	}
	websocketStore.Lock()
	websocketStore.connections[userID] = conn
	websocketStore.Unlock()
	return nil, conn
}

func CloseWebsocket(userID uint32) {
	websocketStore.Lock()
	defer websocketStore.Unlock()
	conn, ok := websocketStore.connections[userID]
	if ok {
		conn.Close()
		delete(websocketStore.connections, userID)
	}
}

func WsSendMessage(userID uint32, data []byte) error {
	websocketStore.RLock()
	defer websocketStore.RUnlock()
	conn, ok := websocketStore.connections[userID]
	if !ok {
		return fmt.Errorf("用户 %d 未连接", userID)
	}
	return conn.WriteMessage(websocket.BinaryMessage, data)
}
