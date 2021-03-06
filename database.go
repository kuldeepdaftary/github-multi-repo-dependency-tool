package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
)

func setupDatabase() (*bolt.DB, error) {
	db, err := bolt.Open("dependencyMap.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("Database Setup failure")
	} else {
		fmt.Printf("Database Setup success")
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("dependencyMapBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	return db, err
}

func updateDatabase(StoredPRData *StoredPRData, depUrl string) {
	db, _ := setupDatabase()
	defer db.Close()

	prData, _ := json.Marshal(StoredPRData)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("dependencyMapBucket"))
		err := b.Put([]byte(depUrl), []byte(prData))
		return err
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("dependencyMapBucket"))
		prData := b.Get([]byte(depUrl))
		fmt.Printf("Value Stored: %s\n", append(prData, (" With Key: "+depUrl)...))
		return nil
	})
}

func checkDatabase(depUrl string) (bool, string) {
	db, _ := setupDatabase()
	defer db.Close()

	var prUrl = ""

	fmt.Println(depUrl)

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("dependencyMapBucket"))
		val := b.Get([]byte(depUrl))
		prUrl = string(val[:])
		fmt.Println(depUrl)
		return nil
	})

	fmt.Printf(prUrl)

	if prUrl != "" {
		return true, prUrl
	} else {
		return false, prUrl
	}
}

func removeKey(depUrl string) {
	db, _ := setupDatabase()
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("dependencyMapBucket"))
		err := b.Delete([]byte(depUrl))
		return err
	})
}

func flushDatabase(w http.ResponseWriter, r *http.Request) {
	db, _ := setupDatabase()
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		tx.DeleteBucket([]byte("dependencyMapBucket"))
		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("dependencyMapBucket"))
		fmt.Printf("BUCKET: %s\n", b)
		return nil
	})
}
