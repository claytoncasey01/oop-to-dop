use chrono::{DateTime, Utc};
use uuid::Uuid;
use crate::oop::post::find_post_by_id;

pub struct Posts {
    pub ids: Vec<Uuid>,
    pub titles: Vec<String>,
    pub bodies: Vec<String>,
    pub published: Vec<bool>,
    pub updated_ats: Vec<DateTime<Utc>>,
    pub author_idxs: Vec<usize>,
}

pub struct PostUpdate {
    pub id: Uuid,
    pub title: Option<String>,
    pub body: Option<String>,
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

    // TODO: Probably don't need the entire self struct, really only need ids, titles, updated_ats,
    // TODO: and bodies. Check performance difference when benchmarking.
    pub fn update(&mut self, updated_post: &PostUpdate) {
        let post_idx = Self::find_by_id(updated_post.id, &self.ids);

        match post_idx {
            Some(idx) => {
                // Handle each case to make sure we have a value
                if let Some(updated_title) = &updated_post.title {
                    self.titles[idx] = updated_title.to_string();
                }

                if let Some(updated_body) = &updated_post.body {
                    self.bodies[idx] = updated_body.to_string();
                }
                self.updated_ats[idx] = Utc::now();
            },
            None => return
        }
    }

    pub fn publish(id: Uuid, ids: &Vec<Uuid>, published: &mut Vec<bool>) {
        let found_idx = Self::find_by_id(id, ids);
        if let Some(idx) = found_idx {
            published[idx] = true;
        }
    }

    pub fn delete(&mut self, id: Uuid) {
        // If we found a post with the id, remove all the data for it.
        // TODO: Test this swap_remove might be fine here. Remove is o(n) worst case because
        // TODO: We have to shift all elements to preserve order. However we might not care due to
        // TODO: swap_remove replacing the removed with the last element meaning we should still
        // TODO: be fine.
        if let Some(idx) = Self::find_by_id(id, &self.ids) {
            self.ids.remove(idx);
            self.titles.remove(idx);
            self.bodies.remove(idx);
            self.updated_ats.remove(idx);
            self.published.remove(idx);
            self.author_idxs.remove(idx);
        }
    }
}
