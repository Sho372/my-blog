package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

var ctx = context.Background()

// Redisクライアントの初期化
var redisClient = redis.NewClient(&redis.Options{
    Addr: "redis:6379",
    Password: "", // パスワードを設定している場合はここに記入
    DB: 0,  // デフォルトDB
})

// Redisストアの初期化
var store, _ = redisstore.NewRedisStore(ctx, redisClient)

type Credentials struct {
    Email string `json:"email"`
    Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // TODO: DBからハッシュ持ってくる
    // パスワードのハッシュ化
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }

    // ユーザー認証（例としてユーザーネームが"testuser"、パスワードが"password"の場合）
    if creds.Email != "test@example.com" || !checkPasswordHash(string(hashedPassword), creds.Password) {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    // セッションの作成
    session, err := store.Get(r, "session-name")
    if err != nil {
        http.Error(w, "Failed to create session", http.StatusInternalServerError)
        return
    }

    // セッションにユーザー情報を保存
    session.Values["authenticated"] = true
    session.Save(r, w)

    w.WriteHeader(http.StatusOK)
}

func Logout(w http.ResponseWriter, r *http.Request) {
    session, err := store.Get(r, "session-name")
    if err != nil {
        http.Error(w, "Failed to retrieve session", http.StatusInternalServerError)
        return
    }

    // セッションの破棄
    session.Values["authenticated"] = false
    session.Save(r, w)

    w.WriteHeader(http.StatusOK)
}

func checkPasswordHash(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        session, err := store.Get(r, "session-name")
        if err != nil {
            http.Error(w, "Failed to retrieve session", http.StatusInternalServerError)
            return
        }

        // セッションの認証状態をチェック
        if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }

        next.ServeHTTP(w, r)
    })
}