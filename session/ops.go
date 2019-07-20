package session
//专门产生session的方法，为用户分配session，检测用户的session是否过期
import (
	"go-micro-server/dbop"
	"go-micro-server/defs"
	"go-micro-server/utils"
	"sync"
	"time"
)


//用go的sync.map做缓存处理，1.9之后支持并发的读写,实现了线程的安全
//读效果特别好，写时效率低下
var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbop.DeleteSession(sid)
}

func loadSessionFromDB() {
	r, err := dbop.RetrieveAllSessions()
	if err != nil {
		return
	}

	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.Session)
		sessionMap.Store(k,ss)
		return true
	})

}

func GenerateNewSessionID(un string) string {
	id, _ := utils.NewUUID()
	ct := time.Now().UnixNano()/1000000   //生成ms级别的时间
	ttl := ct + 30 * 60 * 1000
	ss := &defs.Session{UserName:un, TTL:ttl}
	sessionMap.Store(id, ss)
	dbop.InsertSession(id, ttl, un)

	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := time.Now().UnixNano()/1000000
		if ss.(*defs.Session).TTL < ct {
			//delete expired session
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.Session).UserName, false
	}
	return "", true
}
