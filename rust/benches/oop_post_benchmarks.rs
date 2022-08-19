use criterion::{BenchmarkId, black_box, Criterion, criterion_group, criterion_main};

use rust_oop_to_dop::oop::post::{find_by_id, find_by_title, Post};
use rust_oop_to_dop::oop::author::{Author};
use rust_oop_to_dop::util::oop;

fn new_post_bench(c: &mut Criterion) {
    let post_title = String::from("Post title");
    let post_body = String::from("Post body");
    let author_name = String::from("Author 1");
    let author_bio = String::from("bio for Author 1");
    let author = Author::new( author_name, author_bio);

    c.bench_function("oop_new_post", |b| {
        b.iter(|| Post::new(black_box(post_title.clone()), black_box(post_body.clone()), black_box(&author)));
    });
}

fn find_by_id_bench(c: &mut Criterion) {
    let authors = oop::create_authors(50);
    let posts = oop::create_posts(100, &authors);
    let post_id = posts[posts.len() / 2].id.clone();

    c.bench_function("oop_find_by_id", |b| {
        b.iter(|| find_by_id(black_box(post_id), black_box(posts.clone())));
    });
}

criterion_group!(benches, new_post_bench, find_by_id_bench);
criterion_main!(benches);
