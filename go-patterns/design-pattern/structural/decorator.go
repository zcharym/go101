package structural

import (
    "log"
    "time"
)

func decorator(fn func(s string)) func(s string) {
    return func(s string) {
        start := time.Now()
        time.Sleep(time.Millisecond * 500) // mimic func execution
        fn(s)
        elapsed := time.Since(start)
        log.Printf("func execution time is: %dms", elapsed / time.Millisecond)
    }
}