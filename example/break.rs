fn main() ->i32 {
    let mut a ;
    // loop 1
    loop {
        a = 1;
        break;
        a = 1;
        break;
        break;
    }
    // break; invalid

    // loop2
    loop {
        a = 1;
        if a {
            break;
        }
    }

    // loop in loop
    loop {
        a = 1;
        loop {
            a = 1;
            break
        }
        loop {
            a = 1;
        }
    }

    while a {
        break;
        while a{
            break;
        }
        break;
    }

    return 0;
}
