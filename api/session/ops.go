package session

import (
	"admin/api/dbops"
	"admin/api/defs"
	"admin/api/utils"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInmilli() int64 {
	return time.Now().UnixNano()/1000000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSession()
	if err != nil {
		return
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

func GenerateNewSesssionId(un string) string {
	uuid := utils.NewUUID()
	ct := nowInmilli()
	ttl := ct + 8*60*60// 过期时间 8小时
	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	// 存储缓存
	sessionMap.Store(uuid, ss)
	// 存储db
	dbops.InsertSession(uuid, ttl, un)
	return uuid
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInmilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			// delete expired serssion
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}
	return "", true
}