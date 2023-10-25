use std::collections::HashMap;

use chrono::{DateTime, Utc};
use uuid::Uuid;

pub struct PostData {
    pub id: Uuid,
    pub title: String,
    pub body: String,
    pub updated_at: DateTime<Utc>,
}

pub struct Posts {
    pub data: Vec<PostData>,
    pub id_to_index: HashMap<Uuid, usize>,
    pub published: Vec<bool>,
    pub author_idxs: Vec<usize>,
}

pub struct PostUpdate {
    pub id: Uuid,
    pub title: Option<String>,
    pub body: Option<String>,
}

impl Posts {
    pub fn new(capacity: usize) -> Self {
        Self {
            data: Vec::with_capacity(capacity),
            id_to_index: HashMap::with_capacity(capacity),
            published: vec![false; capacity],
            author_idxs: Vec::with_capacity(capacity),
        }
    }

    pub fn add(&mut self, title: String, body: String, author_idx: usize) {
        let id = Uuid::new_v4();
        let post = PostData {
            id,
            title,
            body,
            updated_at: Utc::now(),
        };
        self.id_to_index.insert(id, self.data.len());
        self.data.push(post);
        self.published.push(false);
        self.author_idxs.push(author_idx);
    }

    pub fn find_by_id(&self, id: Uuid) -> Option<usize> {
        self.id_to_index.get(&id).cloned()
    }

    pub fn find_by_title(&self, title: &str) -> Option<usize> {
        self.data
            .iter()
            .position(|post| post.title.as_str() == title)
    }

    pub fn find_by_author_name(
        author_name: &str,
        author_names: &[String],
        post_author_idxs: &[usize],
    ) -> Vec<usize> {
        post_author_idxs
            .iter()
            .enumerate()
            .filter_map(|(i, idx)| {
                if author_names[*idx] == author_name {
                    Some(i)
                } else {
                    None
                }
            })
            .collect()
    }

    pub fn update(&mut self, updated_post: &PostUpdate) {
        if let Some(idx) = self.find_by_id(updated_post.id) {
            let post = &mut self.data[idx];
            if let Some(ref title) = updated_post.title {
                post.title = title.clone();
            }
            if let Some(ref body) = updated_post.body {
                post.body = body.clone();
            }
            post.updated_at = Utc::now();
        }
    }

    pub fn publish(&mut self, id: Uuid) {
        if let Some(idx) = self.find_by_id(id) {
            self.published[idx] = true;
        }
    }

    pub fn delete(&mut self, id: Uuid) {
        // If we found a post with the id, remove all the data for it.
        if let Some(idx) = self.find_by_id(id) {
            self.id_to_index.remove(&id);
            self.id_to_index.remove(&id);
            self.data.swap_remove(idx);
            self.published.swap_remove(idx);
            self.author_idxs.swap_remove(idx);
        }
    }
}

#[cfg(test)]
mod test {
    use crate::util::dop::{create_authors, create_posts, create_posts_deterministic};

    use super::*;

    #[test]
    fn test_new() {
        let expected_capacity = 50;
        let posts = Posts::new(50);
        assert_eq!(posts.data.capacity(), expected_capacity);
        assert_eq!(posts.published.capacity(), expected_capacity);
        assert_eq!(posts.author_idxs.capacity(), expected_capacity);
    }

    #[test]
    fn test_add() {
        let authors = create_authors(1);
        let mut posts = create_posts(2, authors.ids.len());
        let expected_length = posts.data.len() + 1;
        let expected_title = String::from("An Added Post");
        let expected_body = String::from("An added post body");
        posts.add(
            expected_title.clone(),
            expected_body.clone(),
            posts.author_idxs[0],
        );
        assert_eq!(posts.data.len(), expected_length);
        assert_eq!(posts.data[posts.data.len() - 1].title, expected_title);
        assert_eq!(posts.data[posts.data.len() - 1].body, expected_body);
    }

    #[test]
    fn test_find_by_id() {
        let authors = create_authors(10);
        let posts = create_posts(100, authors.ids.len());
        let expected_id = posts.data[50].id.clone();
        let id_idx = posts.find_by_id(expected_id.clone()).unwrap();
        let actual_id = posts.data[id_idx].id.clone();
        assert_eq!(expected_id.to_string(), actual_id.to_string());
    }

    #[test]
    fn test_find_by_title() {
        let authors = create_authors(10);
        let posts = create_posts(100, authors.ids.len());
        let expected_title = String::from("Post #49");
        let title_idx = posts.find_by_title(expected_title.as_str()).unwrap();
        let actual_title = posts.data[title_idx].title.clone();
        assert_eq!(expected_title.to_string(), actual_title.to_string());
    }

    #[test]
    fn test_find_by_author_name() {
        let authors = create_authors(10);
        let posts = create_posts_deterministic(100, authors.ids.len(), 10);
        let expected_length = 10;
        let found_posts =
            Posts::find_by_author_name("Author #0", &authors.names, &posts.author_idxs);
        assert_eq!(expected_length, found_posts.len());
        assert_eq!(found_posts[0], 0);
    }

    #[test]
    fn test_update() {
        let authors = create_authors(1);
        let mut posts = create_posts(1, authors.ids.len());
        let expected_title = String::from("Post #0 Updated");
        let expected_body = String::from("Updated post body");

        posts.update(&PostUpdate {
            id: posts.data[0].id.clone(),
            title: Some(expected_title.clone()),
            body: Some(expected_body.clone()),
        });

        let actual_title = posts.data[0].title.clone();
        let actual_body = posts.data[0].body.clone();
        assert_eq!(expected_title, actual_title);
        assert_eq!(expected_body, actual_body);
    }

    #[test]
    fn test_publish() {
        let authors = create_authors(1);
        let mut posts = create_posts(1, authors.ids.len());
        posts.publish(posts.data[0].id);
        let actual_published = posts.published[0];
        assert_eq!(true, actual_published);
    }

    #[test]
    fn test_delete() {
        let authors = create_authors(1);
        let mut posts = create_posts(10, authors.ids.len());
        let expected_length = 9;
        let expected_title = posts.data[9].title.clone();
        let expected_id = posts.data[9].id.clone();

        posts.delete(posts.data[4].id.clone());

        let actual_title = posts.data[4].title.clone();
        let actual_id = posts.data[4].id.clone();

        assert_eq!(expected_length, posts.data.len());
        assert_eq!(expected_title, actual_title);
        assert_eq!(expected_id, actual_id);
    }
}
