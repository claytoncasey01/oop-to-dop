use chrono::{DateTime, Utc};
use uuid::Uuid;

pub struct Posts {
    pub ids: Vec<Uuid>,
    titles: Vec<String>,
    bodies: Vec<String>,
    published: Vec<bool>,
    updated_ats: Vec<DateTime<Utc>>,
    author_idxs: Vec<usize>
}