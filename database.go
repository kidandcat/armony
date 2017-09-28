package armony

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/dgraph-io/badger"
)

var kv *badger.KV

var databaseLoaded = false

// LoadDatabase : Load database in "database" Dir
func LoadDatabase() error {
	opt := badger.DefaultOptions
	dir := "database"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModeDir)
	}
	opt.Dir = dir
	opt.ValueDir = dir
	auxKv, err := badger.NewKV(&opt)
	kv = auxKv
	databaseLoaded = true

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Println("Signal received:", sig)
			CloseDatabase()
			os.Exit(0)
		}
	}()

	return err
}

func kvSet(key string, value string) {
	kv.Set([]byte(key), []byte(value), 0x00)
}

// CloseDatabase : Gracefully closes the database
func CloseDatabase() {
	if databaseLoaded {
		kv.Close()
	}
}

func kvGet(key string) string {
	var item badger.KVItem
	if err := kv.Get([]byte(key), &item); err != nil {
		fmt.Printf("Error while getting key: %q", key)
		return ""
	}
	var val []byte
	err := item.Value(func(v []byte) error {
		val = make([]byte, len(v))
		copy(val, v)
		return nil
	})
	if err != nil {
		fmt.Printf("Error while getting value for key: %q", key)
		return ""
	}

	return string(val)
}
