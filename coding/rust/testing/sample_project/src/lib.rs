/// Простенька функція для прикладу
pub fn add(a: i32, b: i32) -> i32 {
    a + b
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_add() {
        assert_eq!(add(2, 2), 4);
    }

    #[test]
    #[should_panic]
    fn test_fail() {
        assert_eq!(add(2, 2), 5);
    }
}
