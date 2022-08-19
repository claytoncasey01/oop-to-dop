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
    let author_idx = rng.gen_range(0..authors.len() - 1);

    for n in 0..amount {
        posts.push(Post::new(format!("{}{}", title.clone(), n), format!("{}{}", body.clone(), n), &authors[author_idx]));
    }

    posts
}