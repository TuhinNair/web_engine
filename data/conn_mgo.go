// +buld mgo

package data

import "webengine/data/model"

func (db *DB) Open(driverName, dataSourceName string) error {
	conn, err := model.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}

	db.CopySession = true

	db.Connection = conn
	return nil
}
