package main

import (
    "fmt"
    "time"
    "path/to/your/project/store" 
)

func main() {

    storage := store.NewStore()

    storage.Set("key1", "value1", 5*time.Second)
    storage.Set("key2", "value2", 10*time.Second)

    if val, found := storage.Get("key1"); found {
        fmt.Println("Found key1:", val)
    } else {
        fmt.Println("key1 not found or expired")
    }

    if val, found := storage.Get("key2"); found {
        fmt.Println("Found key2:", val)
    } else {
        fmt.Println("key2 not found or expired")
    }

    fmt.Println("Waiting for items to expire...")
    time.Sleep(11 * time.Second)

    if val, found := storage.Get("key1"); found {
        fmt.Println("Found key1:", val)
    } else {
        fmt.Println("key1 not found or expired")
    }

    if val, found := storage.Get("key2"); found {
        fmt.Println("Found key2:", val)
    } else {
        fmt.Println("key2 not found or expired")
    }
}
