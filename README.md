### object_store paths with control characters reproducer

To reproduce, follow these steps:

```bash
# set WRITABLE_BUCKET to some GCS bucket that your application default credentials have write access to
export WRITABLE_BUCKET=...
# the Go program creates an object in the bucket named \x01.txt, addressable at https://storage.cloud.google.com/$WRITABLE_BUCKET/%01.txt
cd go
go run main.go 

# The rust program cannot delete that object, because there is no valid Path object that can address it:
cd ../rust
cargo run 
```
