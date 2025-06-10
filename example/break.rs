fn while_case_1() {
    let mut a: i32;
    a = 1;
    // while with return
    while a + 1{
        if a < 0 {
            return;
        }
        // if a {
        //     break;
        // }
    }
}

fn while_case_2() {
    let mut a;
    a = 1;
    // empty while
    while a {
    }
}

fn while_case_3() {
    let mut a;
    a = 1;
    // while with useless code
    while a {
        break;
        while a{
            break;
        }
        break;
    }
}

fn while_case_4() {
    let mut a;
    a = 1;
    while a {
        continue;
    }
}

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
            break;
        }
        loop {
            a = 1;
        }
    }

    return 0;
}
