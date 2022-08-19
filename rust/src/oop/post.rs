use uuid::{Uuid};
use chrono::{DateTime, Utc, offset};

#[derive(Debug)]
pub struct Post {
    id: Uuid,
    title: String,
    body: String,
    published: bool,
    updated_at: DateTime<Utc>
}

impl Post {
    pub fn new(title: String, body: String) -> Self {
        Post {
            id: Uuid::new_v4(),
            title,
            body,
            published: false,
            updated_at: offset::Utc::now()
        }
    }
}