use uuid::Uuid;

#[derive(Debug)]
pub struct Author {
    pub id: Uuid,
    pub name: String,
    pub bio: String,
}

impl Author {
    pub fn new(name: String, bio: String) -> Author {
        Author {
            id: Uuid::new_v4(),
            name,
            bio,
        }
    }
}