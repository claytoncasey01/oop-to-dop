use rust_oop_to_dop::oop::post::{Post};
use rust_oop_to_dop::oop::author::{Author};

fn main() {
    let author = Author::new(String::from("Author 1"), String::from("I am the bio for author 1"));
   let post = Post::new(String::from("First Post"), String::from("First Post Body"), &author);

    println!("{:?}", post);
}
