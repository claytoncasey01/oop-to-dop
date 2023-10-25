use rand::Rng;
use uuid::Uuid;

use crate::dop::authors::Authors;
use crate::dop::posts::Posts;

pub fn create_authors(amount: usize) -> Authors {
    let mut authors = Authors::new(amount);

    for n in 0..amount {
        authors.ids.push(Uuid::new_v4());
        authors.names.push(format!("Author #{}", n));
        authors.bios.push(format!("Bio for author #{}", n));
    }

    authors
}

pub fn create_posts(amount: usize, author_ids_len: usize) -> Posts {
    let mut posts = Posts::new(amount);
    let mut rng = rand::thread_rng();
    let mut author_idx: usize;

    for n in 0..amount {
        author_idx = if author_ids_len <= 1 {
            0
        } else {
            rng.gen_range(0..author_ids_len - 1)
        };

        posts.add(
            format!("Post #{}", n),
            format!("Body for post #{}", n),
            author_idx,
        );
    }

    posts
}

pub fn create_posts_deterministic(
    amount: usize,
    posts_per_author: usize,
    author_ids_len: usize,
) -> Posts {
    let mut posts = Posts::new(amount);
    let mut author_idx: usize;

    for n in 0..amount {
        author_idx = n / posts_per_author;

        if author_idx < author_ids_len {
            posts.add(
                format!("Post #{}", n),
                format!("Body for post #{}", n),
                author_idx,
            );
        }
    }

    posts
}
