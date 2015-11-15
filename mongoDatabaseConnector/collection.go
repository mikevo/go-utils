package mongoDatabaseConnector

import (
	"errors"

	"gopkg.in/mgo.v2"
)

// Collection represents a collection with
type Collection struct {
	serverURL      string
	databaseName   string
	collectionName string
	database       *Database
	collection     *mgo.Collection
}

// NewCollection sets up a connection to a collection
func NewCollection(serverURL string, databaseName string, collectionName string) (*Collection, error) {
	coll := new(Collection)

	coll.serverURL = serverURL
	coll.databaseName = databaseName
	coll.collectionName = collectionName

	coll.database = NewDatabase(serverURL, databaseName)

	var err error
	coll.collection, err = coll.database.Collection(collectionName)

	if err != nil {
		return nil, err
	}

	return coll, nil
}

// Collection returns the database collection of countries and establishes a database connection if necessary
func (coll *Collection) Collection() (*mgo.Collection, error) {
	if coll.database == nil || coll.collection == nil {
		return nil, errors.New("Collection not setup")
	}

	return coll.collection, nil
}

// Close closes the database connection if established
func (coll *Collection) Close() {
	if coll.database != nil {
		coll.database.Close()
	}
}
