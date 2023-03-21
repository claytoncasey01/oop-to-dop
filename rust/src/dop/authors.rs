use uuid::Uuid;

pub struct Authors {
    pub ids: Vec<Uuid>,
    pub names: Vec<String>,
    pub bios: Vec<String>,
}

impl Authors {
    pub fn new(capacity: usize) -> Self {
        Self {
            ids: Vec::with_capacity(capacity),
            names: Vec::with_capacity(capacity),
            bios: Vec::with_capacity(capacity),
        }
    }
}
