package entity

type User struct {
	Id      int64
	Name    string
	Age     int64
	Version int64
}

var Users = make(map[int64]*User)

func init() {
	u := &User{
		Id:      1,
		Name:    "test",
		Age:     18,
		Version: 0,
	}
	Users[u.Id] = u
}
