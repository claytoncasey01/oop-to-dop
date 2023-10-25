use criterion::{black_box, criterion_group, criterion_main, Criterion};
use rust_oop_to_dop::dop::posts::{PostUpdate, Posts};
use rust_oop_to_dop::util::dop;

const AUTHOR_AMOUNT: usize = 100;
const POST_AMOUNT: usize = 10000;

fn add_bench(c: &mut Criterion) {
    let authors = dop::create_authors(AUTHOR_AMOUNT);
    let mut posts = dop::create_posts(POST_AMOUNT, authors.ids.len());
    let title = String::from("New Post");
    let body = String::from("New post body");

    c.bench_function("dop_add", |b| {
        b.iter(|| posts.add(black_box(title.clone()), black_box(body.clone()), 0));
    });
}

fn find_by_id_bench(c: &mut Criterion) {
    let authors = dop::create_authors(AUTHOR_AMOUNT);
    let posts = dop::create_posts(POST_AMOUNT, authors.ids.len());
    let post_id = posts.data[posts.data.len() / 2].id.clone();

    c.bench_function("dop_find_by_id", |b| {
        b.iter(|| posts.find_by_id(black_box(post_id)));
    });
}

fn find_by_title_bench(c: &mut Criterion) {
    let authors = dop::create_authors(AUTHOR_AMOUNT);
    let posts = dop::create_posts(POST_AMOUNT, authors.ids.len());
    let title = String::from("Post #40");

    c.bench_function("dop_find_by_title", |b| {
        b.iter(|| posts.find_by_title(black_box(title.as_str())));
    });
}

fn find_by_author_name_bench(c: &mut Criterion) {
    let authors = dop::create_authors(AUTHOR_AMOUNT);
    let posts = dop::create_posts(POST_AMOUNT, authors.ids.len());
    let author_name = String::from("Author #10");

    c.bench_function("dop_find_by_author_name", |b| {
        b.iter(|| {
            Posts::find_by_author_name(
                black_box(author_name.as_str()),
                black_box(&authors.names),
                black_box(&posts.author_idxs),
            )
        });
    });
}

fn update_bench(c: &mut Criterion) {
    let authors = dop::create_authors(AUTHOR_AMOUNT);
    let mut posts = dop::create_posts(POST_AMOUNT, authors.ids.len());
    let post_update = &PostUpdate {
        id: posts.data[50].id,
        title: Some(String::from("Update Title")),
        body: Some(String::from("Update Body")),
    };

    c.bench_function("dop_update", |b| {
        b.iter(|| posts.update(black_box(post_update)));
    });
}

fn publish_bench(c: &mut Criterion) {
    let authors = dop::create_authors(AUTHOR_AMOUNT);
    let mut posts = dop::create_posts(POST_AMOUNT, authors.ids.len());

    c.bench_function("dop_publish", |b| {
        b.iter(|| posts.publish(black_box(posts.data[50].id)));
    });
}

fn delete_bench(c: &mut Criterion) {
    let authors = dop::create_authors(AUTHOR_AMOUNT);
    let mut posts = dop::create_posts(POST_AMOUNT, authors.ids.len());
    let post_to_delete = posts.data[100].id.clone();

    c.bench_function("dop_delete", |b| {
        b.iter(|| posts.delete(black_box(post_to_delete)));
    });
}

criterion_group!(
    benches,
    add_bench,
    find_by_id_bench,
    find_by_title_bench,
    find_by_author_name_bench,
    update_bench,
    publish_bench,
    delete_bench
);
criterion_main!(benches);
