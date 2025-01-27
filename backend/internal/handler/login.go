package handler

import (
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// 处理登录逻辑
	w.Write([]byte("Login handler"))
}
