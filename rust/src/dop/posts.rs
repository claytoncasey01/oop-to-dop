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

pub struct PostUpdate {
    pub id: Uuid,
    pub title: Option<String>,
    pub body: Option<String>,
}

impl Posts {
    pub fn new(capacity: usize) -> Self {
        Self {
            ids: Vec::with_capacity(capacity),
            titles: Vec::with_capacity(capacity),
            bodies: Vec::with_capacity(capacity),
            published: vec![false; capacity],
            updated_ats: vec![Utc::now(); capacity],
            author_idxs: Vec::with_capacity(capacity),
        }
    }

    pub fn add(&mut self, title: String, body: String, author_idx: usize) {
        self.ids.push(Uuid::new_v4());
        self.titles.push(title);
        self.bodies.push(body);
        self.published.push(false);
        self.updated_ats.push(Utc::now());
        self.author_idxs.push(author_idx);
    }

    pub fn find_by_id(id: Uuid, ids: &[Uuid]) -> Option<usize> {
        ids.iter().position(|found_id| *found_id == id)
    }

    pub fn find_by_title(title: &str, titles: &[String]) -> Option<usize> {
        titles.iter().position(|found_title| *found_title == title)
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
        let post_idx = Self::find_by_id(updated_post.id, &self.ids);

        if let Some(idx) = post_idx {
            if let Some(ref updated_title) = updated_post.title {
                self.titles[idx] = updated_title.clone();
            }

            if let Some(ref updated_body) = updated_post.body {
                self.bodies[idx] = updated_body.clone();
            }
            self.updated_ats[idx] = Utc::now();
        }
    }

    pub fn publish(id: Uuid, ids: &[Uuid], published: &mut Vec<bool>) {
        if let Some(idx) = Self::find_by_id(id, ids) {
            published[idx] = true;
        }
    }

    pub fn delete(&mut self, id: Uuid) {
        // If we found a post with the id, remove all the data for it.
        if let Some(idx) = Self::find_by_id(id, &self.ids) {
            self.ids.swap_remove(idx);
            self.titles.swap_remove(idx);
            self.bodies.swap_remove(idx);
            self.updated_ats.swap_remove(idx);
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
        assert_eq!(posts.ids.capacity(), expected_capacity);
        assert_eq!(posts.titles.capacity(), expected_capacity);
        assert_eq!(posts.bodies.capacity(), expected_capacity);
        assert_eq!(posts.published.capacity(), expected_capacity);
        assert_eq!(posts.updated_ats.capacity(), expected_capacity);
        assert_eq!(posts.author_idxs.capacity(), expected_capacity);
    }

    #[test]
    fn test_add() {
        let authors = create_authors(1);
        let mut posts = create_posts(2, authors.ids.len());
        let expected_length = posts.ids.len() + 1;
        let expected_title = String::from("An Added Post");
        let expected_body = String::from("An added post body");
        posts.add(
            expected_title.clone(),
            expected_body.clone(),
            posts.author_idxs[0],
        );
        assert_eq!(posts.ids.len(), expected_length);
        assert_eq!(posts.titles[posts.ids.len() - 1], expected_title);
        assert_eq!(posts.bodies[posts.ids.len() - 1], expected_body);
    }

    #[test]
    fn test_find_by_id() {
        let authors = create_authors(10);
        let posts = create_posts(100, authors.ids.len());
        let expected_id = posts.ids[50].clone();
        let id_idx = Posts::find_by_id(expected_id.clone(), &posts.ids).unwrap();
        let actual_id = posts.ids[id_idx].clone();
        assert_eq!(expected_id.to_string(), actual_id.to_string());
    }

    #[test]
    fn test_find_by_title() {
        let authors = create_authors(10);
        let posts = create_posts(100, authors.ids.len());
        let expected_title = String::from("Post #49");
        let title_idx = Posts::find_by_title(expected_title.as_str(), &posts.titles).unwrap();
        let actual_title = posts.titles[title_idx].clone();
        assert_eq!(expected_title.to_string(), actual_title.to_string());
    }

    #[test]
    fn test_find_by_author_name() {
        let authors = create_authors(10);
        let posts = create_posts_deterministic(100, authors.ids.len(), 10);
        let expected_length = 10;
        let found_posts = Posts::find_by_author_name(
            "Author #0",
            &authors.names,
            &posts.author_idxs,
        );
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
            id: posts.ids[0].clone(),
            title: Some(expected_title.clone()),
            body: Some(expected_body.clone()),
        });

        let actual_title = posts.titles[0].clone();
        let actual_body = posts.bodies[0].clone();
        assert_eq!(expected_title, actual_title);
        assert_eq!(expected_body, actual_body);
    }

    #[test]
    fn test_publish() {
        let authors = create_authors(1);
        let mut posts = create_posts(1, authors.ids.len());
        Posts::publish(posts.ids[0].clone(), &posts.ids, &mut posts.published);
        let actual_published = posts.published[0];
        assert_eq!(true, actual_published);
    }

    #[test]
    fn test_delete() {
        let authors = create_authors(1);
        let mut posts = create_posts(10, authors.ids.len());
        let expected_length = 9;
        let expected_title = posts.titles[9].clone();
        let expected_id = posts.ids[9].clone();

        posts.delete(posts.ids[4].clone());

        let actual_title = posts.titles[4].clone();
        let actual_id = posts.ids[4].clone();

        assert_eq!(expected_length, posts.ids.len());
        assert_eq!(expected_title, actual_title);
        assert_eq!(expected_id, actual_id);
    }
}
