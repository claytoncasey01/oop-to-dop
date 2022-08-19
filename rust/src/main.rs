mod oop;

use oop::post::{Post};

fn main() {
   let post = Post::new(String::from("First Post"), String::from("First Post Body"));

    println!("{:?}", post);
}
