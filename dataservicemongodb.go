package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net"

	mgo "gopkg.in/mgo.v2"
)

/*
{
 "default_database": "d42879e",
 "hosts": [
  "d42879e-mongod-initial-master-0.node.dc1.consul:27017"
 ],
 "password": "a9sae29de3e96688a16846565223032239788dd7a2e",
 "uri": "mongodb://a9s-brk-usr-028bfd40c069da5cbac0930f581f393748a8ccc9:a9sae29de3e96688a16846565223032239788dd7a2e@d42879e-mongod-initial-master-0.node.dc1.consul:27017/d42879e",
 "username": "a9s-brk-usr-028bfd40c069da5cbac0930f581f393748a8ccc9"
}
*/

type DataServiceMongoDB struct {
}

// Connect the drive to a MongoDB Instance with the given URL
func (d DataServiceMongoDB) getSession() (*mgo.Session, error) {
	log.Println("print credentials")
	for k, v := range d.credentials {
		log.Printf("%s----%s", k, v)
	}

	//d.PrintCredentials()

	log.Printf("Create MongoDB session for [%s].\n", d.credentials["uri"].(string))
	dialInfo, err := mgo.ParseURL(d.credentials["uri"].(string))

	if err != nil {
		panic(err)
	}

	cacert := d.credentials["cacrt"].(string)
	if len(cacert) > 0 {
		tlsConfig := &tls.Config{}

		roots := x509.NewCertPool()
		roots.AppendCertsFromPEM([]byte(cacert))
		tlsConfig.RootCAs = roots
		/*} else {
			tlsConfig.InsecureSkipVerify = true
		}*/

		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	return session, err
}

func (d DataServiceMongoDB) getDatabase(session *mgo.Session) (*mgo.Database, error) {
	log.Println("Get MongoDB database.")
	database := session.DB(d.credentials["default_database"].(string))
	err := database.Login(d.credentials["username"].(string), d.credentials["password"].(string))
	if err != nil {
		return nil, err
	}

	return database, nil
}

func (d DataServiceMongoDB) GetStatus(id string, credentials string) (output int, err error) {
	log.Println("1")
	session, err := d.getSession()
	if err != nil {
		return 2, err
	}
	log.Println("2")
	defer session.Close()
	database, err := d.getDatabase(session)
	if err != nil {
		return 2, err
	}
	log.Println("3")
	err = d.insert(database, "testvalue")
	if err != nil {
		return 2, err
	}
	log.Println("4")
	exists, err := d.exists(database, "testvalue")
	if err != nil {
		return 2, err
	}
	log.Println("5")
	err = d.delete(database, "testvalue")
	if err != nil {
		return 2, err
	}
	log.Println("6")
	if exists {
		return 0, nil
	}
	log.Println("7")
	return 1, nil
}

func (d DataServiceMongoDB) SetCredentials(id string, credentials string) error {
	d.credentials = make(map[string]interface{})

	creds, _ := d.SetCredentialsImpl(id, credentials)

	for k, v := range *creds {
		d.credentials[k] = v
	}
	return nil
}

func (d DataServiceMongoDB) Insert(id string, credentials string, value string) (err error) {
	session, err := d.getSession()
	if err != nil {
		return err
	}

	defer session.Close()
	database, err := d.getDatabase(session)
	if err != nil {
		return err
	}

	return d.insert(database, value)
}

func (d DataServiceMongoDB) Exists(id string, credentials string, value string) (exists bool, err error) {
	session, err := d.getSession()
	if err != nil {
		return false, err
	}

	defer session.Close()
	database, err := d.getDatabase(session)
	if err != nil {
		return false, err
	}

	return d.exists(database, value)
}

func (d DataServiceMongoDB) Delete(id string, credentials string, value string) (err error) {
	session, err := d.getSession()
	if err != nil {
		return err
	}

	defer session.Close()
	database, err := d.getDatabase(session)
	if err != nil {
		return err
	}

	return d.delete(database, value)
}

func (d DataServiceMongoDB) insert(database *mgo.Database, value string) (err error) {
	coll := database.C("collection")

	return coll.Insert(value)
}

func (d DataServiceMongoDB) exists(database *mgo.Database, value string) (exists bool, err error) {
	coll := database.C("collection")
	count, err := coll.Find(value).Count()
	if err != nil {
		return false, err
	}
	log.Printf("Count [%d]", count)

	return (count > 0), nil
}

func (d DataServiceMongoDB) delete(database *mgo.Database, value string) (err error) {
	coll := database.C("collection")

	return coll.Remove(value)
}
