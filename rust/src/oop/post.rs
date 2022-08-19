
use chrono::{DateTime, Utc};
use uuid::Uuid;
use crate::oop::author::{Author};

#[derive(Debug, Clone)]
pub struct Post {
    pub id: Uuid,
    title: String,
    body: String,
    published: bool,
    updated_at: DateTime<Utc>,
    author_id: Uuid
}

impl Post {
    pub fn new(title: String, body: String, author: &Author) -> Self {
        Post {
            id: Uuid::new_v4(),
            title,
            body,
            published: false,
            updated_at: Utc::now(),
            author_id: author.id.clone()
        }
    }

    pub fn update(&mut self, title: String, body: String) {
        self.title = title;
        self.body = body;
        self.updated_at = Utc::now();
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

pub fn find_by_author_id(author_id: Uuid, posts: Vec<Post>) -> Option<Post> {
    posts.into_iter().find(|post| post.author_id == author_id)
}

pub fn delete(id: Uuid, posts: &mut Vec<Post>) {
    if let Some(index) = posts.into_iter().position(|post| *post.id.to_string() == id.to_string()) {
        posts.remove(index);
    }
}

#[cfg(test)]
mod test {
    use crate::oop::author::Author;
    use super::*;
    use crate::util::*;

    #[test]
    fn test_new_post() {
        let author = Author::new(String::from("Author 1"), String::from("I am the bio for author 1"));
        let title = String::from("Test Post");
        let body = String::from("Hello World");
        let post = Post::new(title.clone(), body.clone(), &author);

        assert_eq!(post.title, title);
        assert_eq!(post.body, body);
        assert_eq!(post.published, false);
        assert_eq!(post.author_id, author.id);
    }

    #[test]
    fn test_find_by_id() {
        let posts = oop::create_posts(100, 50);
        let expected_id = posts[50].id.clone();
        let actual_id = find_by_id(expected_id.clone(), posts).unwrap().id;
        assert_eq!(actual_id.to_string(), expected_id.to_string());
    }

    #[test]
    fn test_find_by_title() {
        let posts = oop::create_posts(100, 50);

        let expected_title = posts[50].title.clone();
        let actual_title = find_by_title(expected_title.clone(), posts).unwrap().title.clone();
        assert_eq!(actual_title, expected_title);
    }

    #[test]
    fn test_find_by_author_id() {
        let posts = oop::create_posts(100, 50);
        let expected_author_id = posts[50].author_id.clone();
        let actual_author_id = find_by_author_id(expected_author_id.clone(), posts).unwrap().author_id;

        assert_eq!(actual_author_id, expected_author_id);
    }

    #[test]
    fn test_update() {
        let author = Author::new(String::from("Author 1"), String::from("I am the bio for author 1"));
        let mut post = Post::new(String::from("Test Post"), String::from("Some post content for post"), &author);
        let expected_title = "Test Post Updated";
        let expected_body = "Some post content for post updated";
        post.update(expected_title.to_string().clone(), expected_body.to_string().clone());
        assert_eq!(post.title.clone(), expected_title);
        assert_eq!(post.body.clone(), expected_body);
    }

    #[test]
    fn test_publish() {
        let author = Author::new(String::from("Author 1"), String::from("I am the bio for author 1"));
        let mut post = Post::new(String::from("Test Post"), String::from("Some post content for post"), &author);
        post.publish();
        assert_eq!(post.published, true);
    }

    #[test]
    fn test_delete() {
        let mut posts = oop::create_posts(100, 50);
        let post_id = posts[50].id.clone();

        delete(post_id, &mut posts);
        assert_eq!(posts.len(), 99);
        assert_ne!(posts[50].id.clone(), post_id);
    }
}