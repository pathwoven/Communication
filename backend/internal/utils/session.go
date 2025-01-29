package utils

import (
	"Communication/config"
	"encoding/base64"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/sessions"
)

var sessionStore sessions.Store

func InitSessionStore() {
	sessionStore = sessions.NewCookieStore(
		[]byte(decodeBase64(config.AppConfig.SecurityConfig.SessionConfig.HashKey)),
		[]byte(decodeBase64(config.AppConfig.SecurityConfig.SessionConfig.BlockKey)),
	)
}

func decodeBase64(s string) []byte {
	bytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatalf("解码密钥失败: %v", err)
	}
	return bytes
}

// 全局 map 用于存储 user_id 和 session的映射
var userSessions = struct {
	sync.RWMutex
	m map[uint32]*sessions.Session
}{m: make(map[uint32]*sessions.Session)}

// getSession 获取 session
func getSession(r *http.Request, name string) (*sessions.Session, error) {
	return sessionStore.Get(r, name)
}
func getUserSession(r *http.Request) (*sessions.Session, error) {
	return getSession(r, "user_session")
}

// saveSession 保存 session到响应中
func saveSession(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	return session.Save(r, w)
}

// 删除session
func DeleteSession(r *http.Request, w http.ResponseWriter, session *sessions.Session, userID uint32) {
	session.Options.MaxAge = -1
	session.Save(r, w)
	unregisterUserSession(userID, session)
}

// // SetSessionValue 设置 session 值
// func SetSessionValue(session *sessions.Session, key string, value interface{}) {
// 	session.Values[key] = value
// }

// // GetSessionValue 获取 session 值
// func GetSessionValue(session *sessions.Session, key string) interface{} {
// 	return session.Values[key]
// }

// 将session写入响应，且存储user_id
func SetUserSession(r *http.Request, w http.ResponseWriter, userID uint32) error {
	// 获取session
	session, err := getSession(r, "user_session")
	if err != nil {
		return err
	}
	// 设置user_id
	setUserID(session, userID)
	// 配置Cookie
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400,                 // Session有效期（秒）
		HttpOnly: true,                  // 防止XSS攻击
		Secure:   true,                  // 如果是HTTPS，设置为true
		SameSite: http.SameSiteNoneMode, // 支持跨域
	}

	// 保存session
	err = saveSession(r, w, session)
	if err != nil {
		unregisterUserSession(userID, session)
	}
	return err
}

// GetUserID 从 http 中获取session，从而返回 user_id
func GetUserID(r *http.Request) (uint32, bool) {
	session, err := getUserSession(r)
	log.Printf("Session : %v", session)
	if err != nil {
		return 0, false
	}
	log.Printf("Session values: %v", session.Values)
	userID, ok := session.Values["user_id"].(uint32)
	return userID, ok
}

// setUserID 在 session 中设置 user_id
func setUserID(session *sessions.Session, userID uint32) {
	session.Values["user_id"] = userID
	log.Printf("Session values: %v", session.Values)
	registerUserSession(userID, session) // 保存映射
}

// 注册 user_id 和 session的映射
func registerUserSession(userID uint32, session *sessions.Session) {
	userSessions.Lock()
	defer userSessions.Unlock()
	userSessions.m[userID] = session
}

// 注销 user_id 和 session  的映射
func unregisterUserSession(userID uint32, session *sessions.Session) {
	userSessions.Lock()
	defer userSessions.Unlock()
	if currentSessionID, ok := userSessions.m[userID]; ok && currentSessionID == session {
		delete(userSessions.m, userID)
	}
}

// 获取 user_id 对应的 session
func GetUserIDSession(userID uint32) (*sessions.Session, bool) {
	userSessions.RLock()
	defer userSessions.RUnlock()
	session, ok := userSessions.m[userID]
	return session, ok
}
