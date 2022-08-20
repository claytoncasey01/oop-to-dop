use rand::Rng;

use crate::oop::author::Author;
use crate::oop::post::Post;

pub fn create_authors(amount: usize) -> Vec<Author> {
    let mut authors = Vec::with_capacity(amount);
    let name = String::from("Author #");
    let bio = String::from("I am the bio for author #");

    for n in 0..amount {
        authors.push(Author::new(format!("{}{}", name, n), format!("{}{}", name, bio)));
    }

    authors
}

pub fn create_posts(amount: usize, authors: &Vec<Author>) -> Vec<Post> {
    let mut posts = Vec::with_capacity(amount);
    let title = String::from("Post #");
    let body = String::from("Body for post #");
    let mut rng = rand::thread_rng();
    let mut author_idx = 0;

    for n in 0..amount {
        author_idx = rng.gen_range(0..authors.len() - 1);
        posts.push(Post::new(format!("{}{}", title.clone(), n),
                             format!("{}{}", body.clone(), n), &authors[author_idx]));
    }

    posts
}

pub fn create_posts_deterministic(amount: usize, posts_per_author: usize, authors: &Vec<Author>) -> Vec<Post> {
    let mut posts = Vec::with_capacity(amount);
    let title = String::from("Post #");
    let body = String::from("Body for post #");
    let mut author_idx: usize;

    for n in 0..amount {
        author_idx = n / posts_per_author;

        if author_idx < authors.len() {
            posts.push(Post::new(format!("{}{}", title.clone(), n), format!("{}{}", body.clone(), n), &authors[author_idx]));
        }
    }
    posts
}