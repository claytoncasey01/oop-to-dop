use criterion::{black_box, Criterion, criterion_group, criterion_main};

use rust_oop_to_dop::oop::author::Author;
use rust_oop_to_dop::oop::post::{delete, find_post_by_id, find_post_by_title, find_posts_by_author_name, Post};
use rust_oop_to_dop::util::oop;

const AUTHOR_AMOUNT: usize = 100;
const POST_AMOUNT: usize = 10000;

fn new_post_bench(c: &mut Criterion) {
    let post_title = String::from("Post title");
    let post_body = String::from("Post body");
    let author_name = String::from("Author 1");
    let author_bio = String::from("bio for Author 1");
    let author = Author::new(author_name, author_bio);

    c.bench_function("oop_new_post", |b| {
        b.iter(|| Post::new(black_box(post_title.clone()), black_box(post_body.clone()), black_box(&author)));
    });
}

fn find_by_id_bench(c: &mut Criterion) {
    let authors = oop::create_authors(AUTHOR_AMOUNT);
    let posts = oop::create_posts(POST_AMOUNT, &authors);
    let post_id = posts[posts.len() / 2].id.clone();

    c.bench_function("oop_find_by_id", move |b| {
        b.iter(|| find_post_by_id(black_box(post_id), black_box(&posts)));
    });
}

fn find_by_title_bench(c: &mut Criterion) {
    let authors = oop::create_authors(AUTHOR_AMOUNT);
    let posts = oop::create_posts(POST_AMOUNT, &authors);
    let title = String::from("Post #40");

    c.bench_function("oop_find_by_title", |b| {
        b.iter(|| find_post_by_title(black_box(title.clone()), black_box(&posts)));
    });
}

fn find_by_author_name_bench(c: &mut Criterion) {
    let authors = oop::create_authors(AUTHOR_AMOUNT);
    let posts = oop::create_posts(POST_AMOUNT, &authors);
    let author_name = String::from("Author #10");

    c.bench_function("oop_find_by_author_name", |b| {
        b.iter(|| find_posts_by_author_name(black_box(author_name.clone()), black_box(&posts)));
    });
}

fn update_bench(c: &mut Criterion) {
    let author = Author::new(String::from("Author #1"), String::from("Bio for author #1"));
    let mut post = Post::new(String::from("Post #1"), String::from("Body for post #1"), &author);
    let updated_title = String::from("Updated Title");
    let updated_body = String::from("Updated Body");

    c.bench_function("oop_update", |b| {
        b.iter(|| post.update(black_box(updated_title.clone()), black_box(updated_body.clone())));
    });
}

fn publish_bench(c: &mut Criterion) {
    let author = Author::new(String::from("Author #1"), String::from("Bio for author #1"));
    let mut post = Post::new(String::from("Post #1"), String::from("Body for post #1"), &author);

    c.bench_function("oop_publish", |b| {
        b.iter(|| post.publish());
    });
}

fn delete_bench(c: &mut Criterion) {
    let authors = oop::create_authors(AUTHOR_AMOUNT);
    let mut posts = oop::create_posts(POST_AMOUNT, &authors);
    let post_to_delete = posts[100].id.clone();

    c.bench_function("oop_delete", |b| {
        b.iter(|| delete(post_to_delete, &mut posts));
    });
}

fn add_bench(c: &mut Criterion) {
    let authors = oop::create_authors(AUTHOR_AMOUNT);
    let mut posts = oop::create_posts(POST_AMOUNT, &authors);

    c.bench_function("oop_add", |b| {
        b.iter(|| posts.push(Post::new(String::from("New Post"), String::from("New Post Body"), &authors[0])));
    });
}

criterion_group!(benches, new_post_bench, find_by_id_bench, find_by_title_bench,
    find_by_author_name_bench, update_bench, publish_bench, delete_bench, add_bench);
criterion_main!(benches);
