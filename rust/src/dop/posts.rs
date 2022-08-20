use chrono::{DateTime, Utc};
use uuid::Uuid;

pub struct Posts {
    pub ids: Vec<Uuid>,
    pub titles: Vec<String>,
    pub bodies: Vec<String>,
    pub published: Vec<bool>,
    pub updated_ats: Vec<DateTime<Utc>>,
    pub author_idxs: Vec<usize>,
}

impl Posts {
    pub fn new(capacity: usize) -> Self {
        Self { ids: Vec::with_capacity(capacity), titles: Vec::with_capacity(capacity),
            bodies: Vec::with_capacity(capacity), published: vec![false; capacity],
            updated_ats: vec![Utc::now(); capacity], author_idxs: Vec::with_capacity(capacity) }
    }

    pub fn find_by_id(id: Uuid, ids: &Vec<Uuid>) -> Option<usize> {
        ids.into_iter().position(|found_id| *found_id == id)
    }

    pub fn find_by_title(title: String, titles: &Vec<String>) -> Option<usize> {
        titles.into_iter().position(|found_title| *found_title == title)
    }

    pub fn publish(id: Uuid, ids: &Vec<Uuid>, published: &mut Vec<bool>) {
        let found_idx = Self::find_by_id(id, ids);
        match found_idx {
            Some(idx) => published[idx] = true,
            None => {}
        }
    }
}
