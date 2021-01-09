package session

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
	"github.com/spf13/viper"
	"time"
)

const (
	UserKey = "userID"
)

// 初始化Session 方式 始终会返回一个数据存储方式，默认为 cookie
func NewSessionStore() *sessions.Sessions {
	//typeOf := viper.GetString("session.type")
	secret := viper.GetString("session.secret")
	//name := viper.GetString("session.name")
	cookie_name := viper.GetString("session.name")
	// 判断
	var store sessions.Database
	//switch typeOf {
	//case "redis":
	//
	//case "cookie":
	//	store = cookieStore(secret)
	//	return sessions.Sessions(name, store)
	//default:
	//	store = cookieStore(secret)
	//	return sessions.Sessions(name, store)
	//}
	store = redisStore(secret)
	sess := sessions.New(sessions.Config{
		Cookie:          cookie_name,
		Expires:         0, // defaults to 0: unlimited life. Another good value is: 45 * time.Minute,
		AllowReclaim:    true,
		CookieSecureTLS: true,
	})
	sess.UseDatabase(store)
	return sess
}

// Redis
func redisStore(secret string) *redis.Database {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	pass := viper.GetString("redis.pass")
	db := redis.New(redis.Config{
		Network:   "tcp",
		Addr:      host + ":" + port,
		Timeout:   time.Duration(30) * time.Second,
		MaxActive: 10,
		Password:  pass,
		Database:  "",
		Prefix:    "session:",
		Driver:    redis.Redigo(), // defaults.
	})
	return db
}

// 登录Session 中间件
func AuthSessionMiddle() iris.Handler {

	return func(c iris.Context) {
		session := NewSessionStore().Start(c)
		userId, err := session.GetInt64(UserKey)
		if userId <= 0 {
			c.Redirect("/admin/login", 302)
			c.Problem("还没有登录")
			return
		}
		if err != nil {
			c.Redirect("/admin/login", 302)
			c.Problem("发生错误")
			return
		}
		c.SetCookieKV(UserKey, string(userId))
		c.Next()
		return
	}
}

// 获取Session
func GetUserSession(c iris.Context) int64 {
	session := NewSessionStore().Start(c)
	userId, err := session.GetInt64(UserKey)
	if err == nil {
		return 0
	}
	return userId
}

// 判断是否有Session
func HadSession(c iris.Context) bool {
	session := NewSessionStore().Start(c)
	userId, err := session.GetInt64(UserKey)
	if err == nil {
		return false
	}
	if userId <= 0 {
		return false
	}
	return true
}

// 清除session
func ClearAuthSession(c iris.Context) {
	session := NewSessionStore().Start(c)
	session.Clear()
}

func GetSession(c iris.Context, key string) interface{} {
	session := NewSessionStore().Start(c)
	get := session.Get(key)
	return get
}

// 保存Session
func SaveSession(c iris.Context, key string, val interface{}) bool {
	session := NewSessionStore().Start(c)
	inrec, _ := json.Marshal(val)
	session.Set(key, inrec)
	return true

}
