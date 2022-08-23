use chrono::{DateTime, Utc};
use uuid::Uuid;

use crate::oop::author::Author;

#[derive(Debug, Clone)]
pub struct Post<'a> {
    pub id: Uuid,
    title: String,
    body: String,
    published: bool,
    updated_at: DateTime<Utc>,
    author: &'a Author,
}

impl<'a> Post<'a> {
    pub fn new(title: String, body: String, author: &'a Author) -> Self {
        Post {
            id: Uuid::new_v4(),
            title,
            body,
            published: false,
            updated_at: Utc::now(),
            author,
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
// TODO: Move these into Post impl so we can do Post::function
pub fn find_post_by_id<'a>(id: Uuid, posts: &'a Vec<Post>) -> Option<&'a Post<'a>> {
    posts.into_iter().find(|post| post.id == id)
}

pub fn find_post_by_title<'a>(title: String, posts: &'a Vec<Post>) -> Option<&'a Post<'a>> {
    posts.into_iter().find(|post| post.title == title)
}

pub fn find_posts_by_author_name<'a>(author_name: String, posts: &'a Vec<Post>) -> Vec<&'a Post<'a>> {
    posts.into_iter().filter(|post| post.author.name == author_name).collect()
}

pub fn delete(id: Uuid, posts: &mut Vec<Post>) {
    if let Some(index) = posts.into_iter().position(|post| *post.id.to_string() == id.to_string()) {
        posts.remove(index);
    }
}

#[cfg(test)]
mod test {
    use crate::oop::author::Author;
    use crate::util::*;
    use crate::util::oop::create_authors;

    use super::*;
    
    const AUTHOR_AMOUNT: usize = 50;
    const POST_AMOUNT: usize = 100;

    #[test]
    fn test_new_post() {
        let author = Author::new(String::from("Author 1"), String::from("I am the bio for author 1"));
        let title = String::from("Test Post");
        let body = String::from("Hello World");
        let post = Post::new(title.clone(), body.clone(), &author);

        assert_eq!(post.title, title);
        assert_eq!(post.body, body);
        assert_eq!(post.published, false);
        assert_eq!(post.author.id, author.id);
    }

    #[test]
    fn test_find_post_by_id() {
        let authors = oop::create_authors(AUTHOR_AMOUNT);
        let posts = oop::create_posts(POST_AMOUNT, &authors);
        let expected_id = posts[50].id.clone();
        let actual_id = find_post_by_id(expected_id.clone(), &posts).unwrap().id;
        assert_eq!(actual_id.to_string(), expected_id.to_string());
    }

    #[test]
    fn test_find_post_by_title() {
        let authors = oop::create_authors(AUTHOR_AMOUNT);
        let posts = oop::create_posts(POST_AMOUNT, &authors);
        let expected_title = posts[50].title.clone();
        let actual_title = find_post_by_title(expected_title.clone(), &posts).unwrap().title.clone();
        assert_eq!(actual_title, expected_title);
    }

    #[test]
    fn test_find_posts_by_author_name() {
        let authors = oop::create_authors(AUTHOR_AMOUNT);
        let posts = oop::create_posts_deterministic(100, 10, &authors);
        let expected_length = 10;
        let found_posts = find_posts_by_author_name(String::from("Author #0"), &posts);

        assert_eq!(expected_length, found_posts.len());
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
        let authors = create_authors(AUTHOR_AMOUNT);
        let mut posts = oop::create_posts(POST_AMOUNT, &authors);
        let post_id = posts[50].id.clone();

        delete(post_id, &mut posts);
        assert_eq!(posts.len(), 99);
        assert_ne!(posts[50].id.clone(), post_id);
    }
}