package data

import "webengine/data/model"

type DB struct {
	DatabaseName string
	Connection   *model.Connection
	CopySession  bool
}
