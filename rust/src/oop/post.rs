use uuid::Uuid;

use chrono::{DateTime, offset, Utc};

#[derive(Debug)]
pub struct Post {
    id: Uuid,
    title: String,
    body: String,
    published: bool,
    updated_at: DateTime<Utc>,
}

impl Post {
    pub fn new(title: String, body: String) -> Self {
        Post {
            id: Uuid::new_v4(),
            title,
            body,
            published: false,
            updated_at: offset::Utc::now(),
        }
    }

    pub fn update(&mut self, title: String, body: String) {
        self.title = title;
        self.body = body;
        self.updated_at = offset::Utc::now();
    }

    pub fn publish(&mut self) {
        self.published = true;
    }
}

pub fn find_by_id(id: Uuid, posts: Vec<Post>) -> Option<Post> {
    posts.into_iter().find(|post| post.id == id)
}

pub fn find_by_title(title: String, posts: Vec<Post>) -> Option<Post> {
    posts.into_iter().find(|post| post.title == title)
}

pub fn delete(id: Uuid, posts: &mut Vec<Post>) {
    if let Some(index) = posts.into_iter().position(|post| *post.id.to_string() == id.to_string()) {
        posts.remove(index);
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_new_post() {
        let title = String::from("Test Post");
        let body = String::from("Hello World");
        let post = Post::new(title.clone(), body.clone());

        assert_eq!(post.title, title);
        assert_eq!(post.body, body);
        assert_eq!(post.published, false);
    }

    #[test]
    fn test_find_by_id() {
        let mut posts: Vec<Post> = Vec::new();
        let title = String::from("Test Post");
        let body = String::from("Some post content for post");

        for n in 1..101 {
            posts.push(Post::new(format!("{} {}", title.clone(), n), format!("{} {}", body.clone(), n)));
        }
        let expected_id = posts[50].id.clone();
        let actual_id = find_by_id(expected_id.clone(), posts).unwrap().id;
        assert_eq!(actual_id.to_string(), expected_id.to_string());
    }

    #[test]
    fn test_find_by_title() {
        let mut posts: Vec<Post> = Vec::new();
        let title = String::from("Test Post");
        let body = String::from("Some post content for post");

        for n in 1..101 {
            posts.push(Post::new(format!("{} {}", title.clone(), n), format!("{} {}", body.clone(), n)));
        }
        let expected_title = posts[50].title.clone();
        let actual_title = find_by_title(expected_title.clone(), posts).unwrap().title.clone();
        assert_eq!(actual_title, expected_title);
    }

    #[test]
    fn test_update() {
        let mut post = Post::new(String::from("Test Post"), String::from("Some post content for post"));
        let expected_title = "Test Post Updated";
        let expected_body = "Some post content for post updated";
        post.update(expected_title.to_string().clone(), expected_body.to_string().clone());
        assert_eq!(post.title.clone(), expected_title);
        assert_eq!(post.body.clone(), expected_body);
    }

    #[test]
    fn test_publish() {
        let mut post = Post::new(String::from("Test Post"), String::from("Some post content for post"));
        post.publish();
        assert_eq!(post.published, true);
    }

    #[test]
    fn test_delete() {
        let mut posts: Vec<Post> = Vec::new();
        let title = String::from("Test Post");
        let body = String::from("Some post content for post");

        for n in 1..101 {
            posts.push(Post::new(format!("{} {}", title.clone(), n), format!("{} {}", body.clone(), n)));
        }
        delete(posts[50].id.clone(), &mut posts);
        assert_eq!(posts.len(), 99);
    }
}