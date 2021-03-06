package cluster

import (
	"bytes"
	"encoding/json"
	"strings"

	log "github.com/Sirupsen/logrus"

	kv "github.com/portworx/kvdb"
)

func readDatabase() (Database, error) {
	kvdb := kv.Instance()

	db := Database{Cluster: Info{Status: StatusInit},
		Nodes: make(map[string]Node)}

	kv, err := kvdb.Get("cluster/database")
	if err != nil && !strings.Contains(err.Error(), "100: Key not found") {
		log.Warn("Warning, Could not read cluster database")
		goto done
	}

	if kv == nil || bytes.Compare(kv.Value, []byte("{}")) == 0 {
		log.Info("Cluster is uninitialized...")
		err = nil
		goto done
	} else {
		err = json.Unmarshal(kv.Value, &db)
		if err != nil {
			log.Warn("Fatal, Could not parse cluster database ", kv)
			goto done
		}
	}

done:
	return db, err
}

func writeDatabase(db *Database) error {
	kvdb := kv.Instance()
	b, err := json.Marshal(db)
	if err != nil {
		log.Warn("Fatal, Could not marshal cluster database to JSON")
		goto done
	}

	_, err = kvdb.Put("cluster/database", b, 0)
	if err != nil {
		log.Warn("Fatal, Could not marshal cluster database to JSON")
		goto done
	}

done:
	if err != nil {
		log.Println(err)
	}
	return err
}
