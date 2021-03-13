package repository

import (
	"database/sql"
	"fmt"
	"github.com/Windowsfreak/go-mc/bot/path"
	"github.com/google/uuid"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type SnitchPlayer struct {
	UUID     uuid.UUID
	Name     string
	Login    time.Time
	Logout   time.Time
	Seen     time.Time
	Interval time.Duration
	PrevV3   *path.V3
	V3       *path.V3
	Hits     int
}

func (s *SnitchPlayer) String() string {
	if s == nil {
		return ""
	}
	return s.Name
}

type Snitch struct {
	V3    *path.V3
	Cull  float32
	Group string
	Name  string
	Seen  time.Time
	Read  time.Time
	Dead  bool
	Hits  int
	Alert bool
}

type SnitchEvent struct {
	Person *SnitchPlayer
	Snitch *Snitch
	Action string
	Time   time.Time
}

func (s *Snitch) GetName() string {
	if s.Name == "" {
		return "unnamed snitch"
	}
	return s.Name
}
func (s *Snitch) String() string {
	return fmt.Sprintf("[%d, %d, %d] %s %s", s.V3.X, s.V3.Y, s.V3.Z, s.Group, s.GetName())
}

type Repository interface {
	GetName(UUID uuid.UUID) (string, error)
	KillSnitchesOlderThan(time time.Time) error
	StoreUser(user *SnitchPlayer) error
	FetchUserByUUID(UUID string) (user *SnitchPlayer, err error)
	FetchUserByName(name string) (user *SnitchPlayer, err error)
	StoreSnitch(snitch *Snitch) error
	FetchSnitch(v3 *path.V3) (snitch *Snitch, err error)
	StoreEvent(event *SnitchEvent) error
	StoreChat(group string, username string, message string) error
	GetDatabase() *sql.DB
}

type repository struct {
	database                 *sql.DB
	uuidCache                sync.Map
	killSnitchesStatement    *sql.Stmt
	storeUserStatement       *sql.Stmt
	fetchUserByUUIDStatement *sql.Stmt
	fetchUserByNameStatement *sql.Stmt
	storeSnitchStatement     *sql.Stmt
	fetchSnitchStatement     *sql.Stmt
	storeEventStatement      *sql.Stmt
	storeChatStatement       *sql.Stmt
}

func (r *repository) GetDatabase() *sql.DB {
	return r.database
}

func (r *repository) GetName(UUID uuid.UUID) (string, error) {
	name, ok := r.uuidCache.Load(UUID)
	if !ok {
		user, err := r.FetchUserByUUID(UUID.String())
		if err != nil {
			return "", err
		}
		r.uuidCache.Store(UUID, user.Name)
		return user.Name, err
	}
	return name.(string), nil
}

func (r *repository) KillSnitchesOlderThan(time time.Time) error {
	_, err := r.killSnitchesStatement.Exec(time.Unix())
	return err
}

func (r *repository) StoreUser(user *SnitchPlayer) error {
	r.uuidCache.Store(user.UUID, user.Name)
	var u, v, w, x, y, z int64
	if user.PrevV3 != nil {
		u = int64(user.PrevV3.X)
		v = int64(user.PrevV3.Y)
		w = int64(user.PrevV3.Z)
	}
	if user.V3 != nil {
		x = int64(user.V3.X)
		y = int64(user.V3.Y)
		z = int64(user.V3.Z)
	}
	_, err := r.storeUserStatement.Exec(user.UUID.String(), user.Name,
		user.Login.Unix(), user.Logout.Unix(), user.Seen.Unix(), int64(user.Interval.Seconds()),
		u, v, w, x, y, z, int64(user.Hits), user.UUID.String(),
		user.Login.Unix(), user.Logout.Unix(), user.Seen.Unix(), int64(user.Interval.Seconds()),
		u, v, w, x, y, z, int64(user.Hits), user.Name)
	return err
}

func (r *repository) FetchUserByUUID(UUID string) (user *SnitchPlayer, err error) {
	row := r.fetchUserByUUIDStatement.QueryRow(UUID)
	user = &SnitchPlayer{}
	var login, logout, seen, interval, u, v, w, x, y, z, hits int64
	err = row.Scan(&UUID, &user.Name, &login, &logout, &seen, &interval, &u, &v, &w, &x, &y, &z, &hits)
	if err != nil {
		return
	}
	user.UUID, err = uuid.Parse(UUID)
	user.Login = time.Unix(login, 0)
	user.Logout = time.Unix(logout, 0)
	user.Seen = time.Unix(seen, 0)
	user.Interval = time.Duration(interval) * time.Second
	user.PrevV3 = &path.V3{X: int(u), Y: int(v), Z: int(w)}
	user.V3 = &path.V3{X: int(x), Y: int(y), Z: int(z)}
	user.Hits = int(hits)
	r.uuidCache.Store(user.UUID, user.Name)
	return
}

func (r *repository) FetchUserByName(name string) (user *SnitchPlayer, err error) {
	row := r.fetchUserByNameStatement.QueryRow(name)
	user = &SnitchPlayer{}
	var login, logout, seen, interval, u, v, w, x, y, z, hits int64
	var UUID string
	err = row.Scan(&UUID, &user.Name, &login, &logout, &seen, &interval, &u, &v, &w, &x, &y, &z, &hits)
	if err != nil {
		return
	}
	user.UUID, err = uuid.Parse(UUID)
	user.Login = time.Unix(login, 0)
	user.Logout = time.Unix(logout, 0)
	user.Seen = time.Unix(seen, 0)
	user.Interval = time.Duration(interval) * time.Second
	user.PrevV3 = &path.V3{X: int(u), Y: int(v), Z: int(w)}
	user.V3 = &path.V3{X: int(x), Y: int(y), Z: int(z)}
	user.Hits = int(hits)
	r.uuidCache.Store(user.UUID, user.Name)
	return
}

func (r *repository) StoreSnitch(snitch *Snitch) error {
	var x, y, z int64
	if snitch.V3 != nil {
		x = int64(snitch.V3.X)
		y = int64(snitch.V3.Y)
		z = int64(snitch.V3.Z)
	} else {
		return fmt.Errorf("snitch without V3")
	}
	_, err := r.storeSnitchStatement.Exec(x, y, z, snitch.Group, snitch.Name, snitch.Seen.Unix(), snitch.Read.Unix(), snitch.Dead, int64(snitch.Hits), snitch.Cull, snitch.Alert,
		snitch.Group, snitch.Name, snitch.Seen.Unix(), snitch.Read.Unix(), snitch.Dead, int64(snitch.Hits), snitch.Cull, snitch.Alert, x, y, z)
	return err
}

func (r *repository) FetchSnitch(v3 *path.V3) (snitch *Snitch, err error) {
	var x, y, z int64
	if v3 != nil {
		x = int64(v3.X)
		y = int64(v3.Y)
		z = int64(v3.Z)
	}
	row := r.fetchSnitchStatement.QueryRow(x, y, z)
	snitch = &Snitch{}
	var seen, read, hits int64
	err = row.Scan(&x, &y, &z, &snitch.Group, &snitch.Name, &seen, &read, &snitch.Dead, &hits, &snitch.Cull, &snitch.Alert)
	if err != nil {
		return
	}
	snitch.V3 = &path.V3{X: int(x), Y: int(y), Z: int(z)}
	snitch.Seen = time.Unix(seen, 0)
	snitch.Read = time.Unix(read, 0)
	snitch.Hits = int(hits)
	return
}

func (r *repository) StoreEvent(event *SnitchEvent) error {
	var x, y, z int64
	if event.Snitch.V3 != nil {
		x = int64(event.Snitch.V3.X)
		y = int64(event.Snitch.V3.Y)
		z = int64(event.Snitch.V3.Z)
	} else {
		return fmt.Errorf("snitch without V3")
	}
	_, err := r.storeEventStatement.Exec(x, y, z, event.Person.Name, event.Action, event.Time.Unix(), event.Snitch.Alert, event.Snitch.Group, event.Snitch.Name)
	return err
}

func (r *repository) StoreChat(group string, username string, message string) error {
	_, err := r.storeChatStatement.Exec(time.Now().Unix(), group, username, message)
	return err
}

func NewRepository() (*repository, error) {
	database, err := sql.Open("sqlite3", "./snitchLogs.db")
	if err != nil {
		return nil, err
	}
	_, err = database.Exec(`CREATE TABLE IF NOT EXISTS user (
    		uuid TEXT NOT NULL,
    		username TEXT NOT NULL,
    		login INTEGER NOT NULL,
    		logout INTEGER NOT NULL,
    		seen INTEGER NOT NULL,
    		interval INTEGER NOT NULL,
     		u INTEGER NOT NULL,
    		v INTEGER NOT NULL,
    		w INTEGER NOT NULL,
     		x INTEGER NOT NULL,
    		y INTEGER NOT NULL,
    		z INTEGER NOT NULL,
    		hits INTEGER NOT NULL,
    		PRIMARY KEY (username)) WITHOUT ROWID;`)
	if err != nil {
		return nil, err
	}
	_, err = database.Exec(`CREATE TABLE IF NOT EXISTS snitch (
     		x INTEGER NOT NULL,
    		y INTEGER NOT NULL,
    		z INTEGER NOT NULL,
   			nl TEXT NOT NULL,
    		name TEXT NOT NULL,
    		seen INTEGER NOT NULL,
    		read INTEGER NOT NULL,
    		dead BOOLEAN NOT NULL,
    		hits INTEGER NOT NULL,
    		cull FLOAT NOT NULL,
    		alert BOOLEAN NOT NULL,
    		PRIMARY KEY (x, z, y)) WITHOUT ROWID;`)
	if err != nil {
		return nil, err
	}
	_, err = database.Exec(`CREATE TABLE IF NOT EXISTS event (
            rowid INTEGER PRIMARY KEY AUTOINCREMENT,
     		x INTEGER NOT NULL,
    		y INTEGER NOT NULL,
    		z INTEGER NOT NULL,
    		nl TEXT NOT NULL,
    		name TEXT NOT NULL,
   			username TEXT NOT NULL,
   			event TEXT NOT NULL,
   			time INTEGER NOT NULL,
   			alert BOOLEAN NOT NULL);`)
	if err != nil {
		return nil, err
	}
	_, err = database.Exec(`CREATE TABLE IF NOT EXISTS chat (
            rowid INTEGER PRIMARY KEY AUTOINCREMENT,
            time INTEGER NOT NULL,
     		nl TEXT NOT NULL,
    		username TEXT NOT NULL,
    		message TEXT NOT NULL);`)
	if err != nil {
		return nil, err
	}

	_, err = database.Exec(`CREATE INDEX IF NOT EXISTS user_by_name ON user (username ASC);`)
	if err != nil {
		return nil, err
	}
	_, err = database.Exec(`CREATE INDEX IF NOT EXISTS user_by_hits ON user (hits ASC);`)
	if err != nil {
		return nil, err
	}
	_, err = database.Exec(`CREATE INDEX IF NOT EXISTS user_by_seen ON user (seen ASC);`)
	if err != nil {
		return nil, err
	}
	_, err = database.Exec(`CREATE INDEX IF NOT EXISTS snitch_by_hits ON snitch (hits ASC);`)
	if err != nil {
		return nil, err
	}
	_, err = database.Exec(`CREATE INDEX IF NOT EXISTS snitch_by_cull ON snitch (cull ASC);`)
	if err != nil {
		return nil, err
	}
	_, err = database.Exec(`CREATE INDEX IF NOT EXISTS snitch_by_seen ON snitch (seen ASC);`)
	if err != nil {
		return nil, err
	}
	_, err = database.Exec(`CREATE INDEX IF NOT EXISTS snitch_by_read ON snitch (read ASC);`)
	if err != nil {
		return nil, err
	}
	_, err = database.Exec(`CREATE INDEX IF NOT EXISTS event_by_snitch ON event (x ASC, z ASC, y ASC);`)
	if err != nil {
		return nil, err
	}
	_, err = database.Exec(`CREATE INDEX IF NOT EXISTS event_by_user ON event (username ASC);`)
	if err != nil {
		return nil, err
	}
	_, err = database.Exec(`CREATE INDEX IF NOT EXISTS event_by_date ON event (time ASC);`)
	if err != nil {
		return nil, err
	}
	_, err = database.Exec(`CREATE INDEX IF NOT EXISTS chat_by_date ON chat (time ASC);`)
	if err != nil {
		return nil, err
	}
	_, err = database.Exec(`CREATE INDEX IF NOT EXISTS chat_by_user ON chat (username ASC, time ASC);`)
	if err != nil {
		return nil, err
	}
	_, err = database.Exec(`CREATE INDEX IF NOT EXISTS chat_by_nl ON chat (nl ASC, time ASC);`)
	if err != nil {
		return nil, err
	}

	r := &repository{database, sync.Map{},
		nil, nil, nil, nil, nil, nil, nil, nil}

	r.killSnitchesStatement, err = r.database.Prepare(`UPDATE snitch SET dead = true WHERE read < ?`)
	if err != nil {
		return nil, err
	}
	r.storeUserStatement, err = r.database.Prepare(`INSERT INTO user
    	(uuid, username, login, logout, seen, interval, u, v, w, x, y, z, hits)
    	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) ON CONFLICT (username) DO UPDATE SET
			uuid = ?, login = ?, logout = ?, seen = ?, interval = ?, u = ?, v = ?, w = ?, x = ?, y = ?, z = ?, hits = ? WHERE username = ?`)
	if err != nil {
		return nil, err
	}
	r.fetchUserByUUIDStatement, err = r.database.Prepare("SELECT uuid, username, login, logout, seen, interval, u, v, w, x, y, z, hits FROM user WHERE uuid = ?")
	if err != nil {
		return nil, err
	}
	r.fetchUserByNameStatement, err = r.database.Prepare("SELECT uuid, username, login, logout, seen, interval, u, v, w, x, y, z, hits FROM user WHERE username = ?")
	if err != nil {
		return nil, err
	}
	r.storeSnitchStatement, err = r.database.Prepare(`INSERT INTO snitch
    	(x, y, z, nl, name, seen, read, dead, hits, cull, alert)
    	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) ON CONFLICT (x, y, z) DO UPDATE SET
			nl = ?, name = ?, seen = ?, read = ?, dead = ?, hits = ?, cull = ?, alert = ? WHERE x = ? AND y = ? AND z = ?`)
	if err != nil {
		return nil, err
	}
	r.fetchSnitchStatement, err = r.database.Prepare("SELECT x, y, z, nl, name, seen, read, dead, hits, cull, alert FROM snitch WHERE x = ? and y = ? AND z = ?")
	if err != nil {
		return nil, err
	}
	r.storeEventStatement, err = r.database.Prepare(`INSERT INTO event
    	(x, y, z, username, event, time, alert, nl, name)
    	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	r.storeChatStatement, err = r.database.Prepare(`INSERT INTO chat
    	(time, nl, username, message)
    	VALUES (?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}

	return r, err
}
