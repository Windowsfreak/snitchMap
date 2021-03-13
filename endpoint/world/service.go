package world

import (
	"github.com/Windowsfreak/go-mc/bot/repository"
	"github.com/Windowsfreak/go-mc/domain"
	"math"
	"time"
)

type Service interface {
	GetEventsAfter(rowid int64, limit int64) ([]domain.Event, error)
	GetEventsByUserAfter(rowid int64, username string, limit int64) (events []domain.Event, err error)
	GetEventsByRegionAfter(rowid, x1, z1, x2, z2 int64, limit int64) (events []domain.Event, err error)
	GetLastEventRowIdBefore(date int64) (rowid int64, err error)
	GetLastEventRowIdByUserBefore(date int64, username string) (rowid int64, err error)
	GetLastEventRowIdByRegionBefore(date, x1, z1, x2, z2 int64) (rowid int64, err error)
	GetChatsAfter(rowid int64, limit int64) (chats []domain.Chat, err error)
	GetLastChatRowIdBefore(date int64) (rowid int64, err error)
	GetUsersSeenAfter(seen int64) (users []domain.User, err error)
	GetUser(username string) (user domain.User, err error)
	GetSnitches() (snitches []domain.Snitch, err error)
	SetSnitchAlert(x1, z1, x2, z2 int64, alert bool) (err error)
}

type service struct {
	r repository.Repository
}

func NewService(
	repository repository.Repository,
) Service {
	return &service{
		r: repository,
	}
}

func (s *service) GetEventsAfter(rowid int64, limit int64) (events []domain.Event, err error) {
	d := s.r.GetDatabase()
	rows, err := d.Query("SELECT rowid, x, y, z, username, event, time, nl, name, alert FROM event WHERE rowid > ? ORDER BY rowid ASC LIMIT ?", rowid, limit)
	if err != nil {
		return
	}
	events = []domain.Event{}
	for rows.Next() {
		evt := domain.Event{}
		if err = rows.Scan(&evt.Rowid, &evt.X, &evt.Y, &evt.Z, &evt.User, &evt.Event, &evt.Time, &evt.Nl, &evt.Name, &evt.Alert); err != nil {
			return
		}
		events = append(events, evt)
	}
	err = rows.Close()
	return
}

func (s *service) GetEventsByUserAfter(rowid int64, username string, limit int64) (events []domain.Event, err error) {
	d := s.r.GetDatabase()
	rows, err := d.Query("SELECT rowid, x, y, z, username, event, time, nl, name, alert FROM event WHERE rowid > ? AND username = ? ORDER BY rowid ASC LIMIT ?", rowid, username, limit)
	if err != nil {
		return
	}
	events = []domain.Event{}
	for rows.Next() {
		evt := domain.Event{}
		if err = rows.Scan(&evt.Rowid, &evt.X, &evt.Y, &evt.Z, &evt.User, &evt.Event, &evt.Time, &evt.Nl, &evt.Name, &evt.Alert); err != nil {
			return
		}
		events = append(events, evt)
	}
	err = rows.Close()
	return
}

func (s *service) GetEventsByRegionAfter(rowid, x1, z1, x2, z2 int64, limit int64) (events []domain.Event, err error) {
	d := s.r.GetDatabase()
	rows, err := d.Query("SELECT rowid, x, y, z, username, event, time, nl, name, alert FROM event WHERE rowid > ? AND x >= ? AND z >= ? AND x <= ? AND z <= ? ORDER BY rowid ASC LIMIT ?", rowid, x1, z1, x2, z2, limit)
	if err != nil {
		return
	}
	events = []domain.Event{}
	for rows.Next() {
		evt := domain.Event{}
		if err = rows.Scan(&evt.Rowid, &evt.X, &evt.Y, &evt.Z, &evt.User, &evt.Event, &evt.Time, &evt.Nl, &evt.Name, &evt.Alert); err != nil {
			return
		}
		events = append(events, evt)
	}
	err = rows.Close()
	return
}

func (s *service) GetLastEventRowIdBefore(date int64) (rowid int64, err error) {
	if date <= 0 {
		date += time.Now().Unix()
	}
	d := s.r.GetDatabase()
	row := d.QueryRow("SELECT rowid FROM event WHERE time < ? ORDER BY rowid DESC LIMIT 1", date)
	if row == nil {
		return
	}
	err = row.Scan(&rowid)
	if err != nil {
		err = nil
	}
	return
}

func (s *service) GetLastEventRowIdByUserBefore(date int64, username string) (rowid int64, err error) {
	if date <= 0 {
		date += time.Now().Unix()
	}
	d := s.r.GetDatabase()
	row := d.QueryRow("SELECT rowid FROM event WHERE time < ? AND username = ? ORDER BY rowid DESC LIMIT 1", date, username)
	if row == nil {
		return
	}
	err = row.Scan(&rowid)
	if err != nil {
		err = nil
	}
	return
}

func (s *service) GetLastEventRowIdByRegionBefore(date, x1, z1, x2, z2 int64) (rowid int64, err error) {
	if date <= 0 {
		date += time.Now().Unix()
	}
	d := s.r.GetDatabase()
	row := d.QueryRow("SELECT rowid FROM event WHERE time < ? AND x >= ? AND z >= ? AND x <= ? AND z <= ? ORDER BY rowid DESC LIMIT 1", date, x1, z1, x2, z2)
	if row == nil {
		return
	}
	err = row.Scan(&rowid)
	if err != nil {
		err = nil
	}
	return
}

func (s *service) GetChatsAfter(rowid int64, limit int64) (chats []domain.Chat, err error) {
	d := s.r.GetDatabase()
	rows, err := d.Query("SELECT rowid, time, nl, username, message FROM chat WHERE rowid > ? ORDER BY rowid ASC LIMIT ?", rowid, limit)
	if err != nil {
		return
	}
	chats = []domain.Chat{}
	for rows.Next() {
		chat := domain.Chat{}
		if err = rows.Scan(&chat.Rowid, &chat.Time, &chat.Nl, &chat.User, &chat.Text); err != nil {
			return
		}
		chats = append(chats, chat)
	}
	err = rows.Close()
	return
}

func (s *service) GetLastChatRowIdBefore(date int64) (rowid int64, err error) {
	if date <= 0 {
		date += time.Now().Unix()
	}
	d := s.r.GetDatabase()
	row := d.QueryRow("SELECT rowid FROM chat WHERE time < ? ORDER BY rowid DESC LIMIT 1", date)
	if row == nil {
		return
	}
	err = row.Scan(&rowid)
	if err != nil {
		err = nil
	}
	return
}

const degree = 180 / math.Pi

func (s *service) GetUsersSeenAfter(seen int64) (users []domain.User, err error) {
	if seen <= 0 {
		seen += time.Now().Unix()
	}
	d := s.r.GetDatabase()
	rows, err := d.Query("SELECT username, login, logout, seen, interval, u, v, w, x, y, z, hits FROM user WHERE seen > ? ORDER BY seen ASC", seen)
	if err != nil {
		return
	}
	users = []domain.User{}
	for rows.Next() {
		user := domain.User{}
		var u, v, w, x, y, z int64
		if err = rows.Scan(&user.User, &user.Login, &user.Logout, &user.Seen, &user.Interval, &u, &v, &w, &x, &y, &z, &user.Hits); err != nil {
			return
		}
		if user.Interval > 0 {
			angle := math.Atan2(float64(z-w), float64(x-u)) * degree
			if angle != math.NaN() {
				user.Angle = float32(angle)
				user.Speed = float32(math.Sqrt(float64((z-w)*(z-w)+(x-u)*(x-u))) / float64(user.Interval))
			}
		}
		user.X = x
		user.Y = y
		user.Z = z
		users = append(users, user)
	}
	err = rows.Close()
	return
}

func (s *service) GetUser(username string) (user domain.User, err error) {
	d := s.r.GetDatabase()
	row := d.QueryRow("SELECT username, login, logout, seen, interval, u, v, w, x, y, z, hits FROM user WHERE username = ?", username)
	if row == nil {
		return
	}
	user = domain.User{}
	var u, v, w, x, y, z int64
	if err = row.Scan(&user.User, &user.Login, &user.Logout, &user.Seen, &user.Interval, &u, &v, &w, &x, &y, &z, &user.Hits); err != nil {
		return
	}
	if user.Interval > 0 {
		angle := math.Atan2(float64(z-w), float64(x-u)) * degree
		if angle != math.NaN() {
			user.Angle = float32(angle)
			user.Speed = float32(math.Sqrt(float64((z-w)*(z-w)+(x-u)*(x-u))) / float64(user.Interval))
		}
	}
	user.X = x
	user.Y = y
	user.Z = z
	return
}

func (s *service) GetSnitches() (snitches []domain.Snitch, err error) {
	d := s.r.GetDatabase()
	rows, err := d.Query("SELECT x, y, z, nl, name, seen, read, hits, cull, dead, alert FROM snitch")
	if err != nil {
		return
	}
	snitches = []domain.Snitch{}
	for rows.Next() {
		snitch := domain.Snitch{}
		if err = rows.Scan(&snitch.X, &snitch.Y, &snitch.Z, &snitch.Nl, &snitch.Name, &snitch.Seen, &snitch.Read, &snitch.Hits, &snitch.Cull, &snitch.Dead, &snitch.Alert); err != nil {
			return
		}
		snitches = append(snitches, snitch)
	}
	err = rows.Close()
	return
}

func (s *service) SetSnitchAlert(x1, z1, x2, z2 int64, alert bool) (err error) {
	_, err = s.r.GetDatabase().Exec("UPDATE snitch SET alert = ? WHERE x >= ? AND z >= ? AND x <= ? AND z <= ?", alert, x1, z1, x2, z2)
	return
}
