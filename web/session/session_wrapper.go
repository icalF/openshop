package session

import (
	"time"

	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/boltdb"
	"github.com/kataras/iris"
)

type Wrapper struct {
	sess *sessions.Sessions
}

func NewSessionWrapper() Wrapper {
	db, _ := boltdb.New("./sessions.db", 0666, "users")
	db.Async(true)

	iris.RegisterOnInterrupt(func() {
		db.Close()
	})

	sess := sessions.New(sessions.Config{
		Cookie:  "SHOPSESS_ID",
		Expires: 45 * time.Minute,
	})
	sess.UseDatabase(db)

	return Wrapper{sess}
}

func (w *Wrapper) GetSession() *sessions.Sessions {
	return w.sess
}
