package main

import (
	"context"
	"os"

	"cloud.google.com/go/storage"
)

func main() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(err)
	}
	writableBucket := os.Getenv("WRITABLE_BUCKET")
	if writableBucket == "" {
		panic("provide a writable bucket")
	}
	writer := client.Bucket(writableBucket).Object("\x01.txt").NewWriter(ctx)
	if _, err := writer.Write([]byte("hello\n")); err != nil {
		panic(err)
	}
	if err := writer.Close(); err != nil {
		panic(err)
	}
}
