package mongoDatabaseConnector

import (
	"errors"

	"gopkg.in/mgo.v2"
)

var connection = make(map[string]*mgo.Session)

// Database represents a database in Database
type Database struct {
	serverURL    string
	databaseName string

	session  *mgo.Session
	database *mgo.Database
}

// NewDatabase generates a new Database object and tryes to reuese connection pools
func NewDatabase(serverURL string, db string) *Database {
	mongoDB := new(Database)

	mongoDB.serverURL = serverURL
	mongoDB.databaseName = db

	return mongoDB
}

// Connect connets the Database object to the Database server
func (mongoDB *Database) Connect() error {
	if mongoDB.serverURL == "" {
		return errors.New("empty server URL")
	}

	if mongoDB.databaseName == "" {
		return errors.New("empty database name")
	}

	if connection[mongoDB.serverURL] == nil {
		var err error

		connection[mongoDB.serverURL], err = mgo.Dial(mongoDB.serverURL)

		if err != nil {
			connection[mongoDB.serverURL].Close()
			return err
		}

		mongoDB.session = connection[mongoDB.serverURL]
	} else {
		mongoDB.session = connection[mongoDB.serverURL].Copy()
	}

	mongoDB.database = mongoDB.session.DB(mongoDB.databaseName)

	return nil
}

// Close closes the connection to the Database server
func (mongoDB *Database) Close() {
	if mongoDB.session != nil {
		mongoDB.session.Close()
	}
}

// Collection returns a collection of the database. It establishes a connection if there is no connection.
func (mongoDB *Database) Collection(collection string) (*mgo.Collection, error) {
	if mongoDB.database == nil {
		err := mongoDB.Connect()
		if err != nil {
			return nil, err
		}
	}

	return mongoDB.database.C(collection), nil
}
