package data

import "webengine/data/model"

type DB struct {
	DatabaseName string
	Connection   *model.Connection
	CopySession  bool

	Users UserServices
}

type SessionRefresher interface {
	RefreshSession(*model.Connection, string)
}

type UserServices interface {
	SessionRefresher
	GetDetail(id model.Key) (*model.User, error)
}
