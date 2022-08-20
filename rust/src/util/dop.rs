use rand::Rng;
use uuid::Uuid;
use crate::dop::posts::Posts;
use crate::dop::authors::Authors;


fn create_authors(amount: usize) -> Authors {
    let mut authors = Authors::new(amount);

    for n in 0..amount {
        authors.ids.push(Uuid::new_v4());
        authors.names.push(format!("Author #{}", n));
        authors.bios.push(format!("Bio for author #{}", n));
    }

    authors
}

fn create_posts(amount: usize, author_ids: &Vec<String>) -> Posts {
    let mut posts = Posts::new(amount);
    let mut rng = rand::thread_rng();
    let mut author_idx;

    for n in 0..amount {
        author_idx = rng.gen_range(0..author_ids.len() - 1);
        posts.ids.push(Uuid::new_v4());
        posts.titles.push(format!("Post #{}", n));
        posts.bodies.push(format!("Body for post #{}", n));
        posts.author_idxs.push(author_idx);
    }

    posts
}