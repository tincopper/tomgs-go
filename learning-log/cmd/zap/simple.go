package main

import (
    "go.uber.org/zap"
    "time"
)

func main() {
    logger, _ := zap.NewProduction()
    defer logger.Sync() // flushes buffer, if any
    url := "http://www.test.com"
    sugar := logger.Sugar()
    sugar.Infow("failed to fetch URL",
        // Structured context as loosely typed key-value pairs.
        "url", url,
        "attempt", 3,
        "backoff", time.Second,
    )
    sugar.Infof("Failed to fetch URL: %s", url)
}
