package main

import (
    "log"
    "fmt"
    "github.com/boltdb/bolt"
    "net/http"
)

func setupDatabase() {
    db, err := bolt.Open("dependencyMap.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

    db.Update(func(tx *bolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists([]byte("dependencyMapBucket"))
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        }
        return nil
    })
}

func updateDatabase(prUrl string, pullUrl string)  {
    db, err := bolt.Open("dependencyMap.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("dependencyMapBucket"))
        err := b.Put([]byte(prUrl), []byte(pullUrl))
        return err
    })

    db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("dependencyMapBucket"))
        depPullUrl := b.Get([]byte(prUrl))
        fmt.Printf("Value Stored: %s\n", append(depPullUrl, (" With Key: " + prUrl)...))
    return nil
    })
}

func flushDatabase(w http.ResponseWriter, r *http.Request) {
    db, err := bolt.Open("dependencyMap.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    db.Update(func(tx *bolt.Tx) error {
        tx.DeleteBucket([]byte("dependencyMapBucket"))
        return err
    })

    db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("dependencyMapBucket"))
    fmt.Printf("BUCKET: %s\n", b)
    return nil
})
}
