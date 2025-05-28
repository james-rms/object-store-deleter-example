use object_store::ObjectStore;

#[tokio::main]
async fn main() {
    let bucket_name = std::env::var("WRITABLE_BUCKET").unwrap();
    let store = object_store::gcp::GoogleCloudStorageBuilder::from_env()
        .with_bucket_name(bucket_name)
        .build()
        .unwrap();
    store
        .delete(&object_store::path::Path::parse("%01.txt").unwrap())
        .await
        .unwrap();
}
