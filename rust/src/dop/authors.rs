use uuid::Uuid;

pub struct Authors {
    pub ids: Vec<Uuid>,
    names: Vec<String>,
    bios: Vec<String>,
}