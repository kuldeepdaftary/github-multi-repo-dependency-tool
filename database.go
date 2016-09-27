package main

import (
    "log"
    "fmt"
    "github.com/boltdb/bolt"
    "net/http"
)

func setupDatabase() (*bolt.DB, error){
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

func updateDatabase(prUrl string, depUrl string) {
    db, _ := setupDatabase()
    defer db.Close()

    db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("dependencyMapBucket"))
        err := b.Put([]byte(depUrl), []byte(prUrl))
        return err
    })

    db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("dependencyMapBucket"))
        prUrl := b.Get([]byte(depUrl))
        fmt.Printf("Value Stored: %s\n", append(prUrl, (" With Key: " + depUrl)...))
        return nil
    })
}

func checkDatabase(depUrl string) (bool, string) {
    db, _ := setupDatabase()
    defer db.Close()

    var prUrl = ""
    
    fmt.Printlin(depUrl)

    db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("dependencyMapBucket"))
        val := b.Get([]byte(depUrl))
        prUrl = string(val[:])
        fmt.Printlin(depUrl)
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
