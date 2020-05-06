package initialize

import (
	"github.com/gocql/gocql"
	"log"
	"strings"
	"voyageone.com/dp/infrastructure/model/config"
	"voyageone.com/dp/infrastructure/model/global"
)

func InitCqlSession() {
	var err error
	global.CqlSession, err = initCqlSession(global.DPConfig.Cassandra)
	if err != nil {
		log.Fatal(err)
	}
}

func initCqlSession(cassandraConfig config.CassandraConfig) (*gocql.Session, error) {
	urlsSlice := strings.Split(cassandraConfig.HostsUrls, ",")
	cluster := gocql.NewCluster(urlsSlice...)
	cluster.Consistency = gocql.Quorum
	cluster.NumConns = 3
	return cluster.CreateSession()
}
