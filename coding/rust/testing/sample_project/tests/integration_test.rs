use hello_add::add;

#[test]
fn integration_add() {
    assert_eq!(add(10, 5), 15);
}
