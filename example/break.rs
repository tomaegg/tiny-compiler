fn main() ->i32 {
    let mut a = 1;
    loop {
        break;
    }
    // break; invalid

    while a {
        break;
        while a{
            break;
        }
        break;
    }

    return 0;
}
