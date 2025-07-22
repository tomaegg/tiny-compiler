fn p_2_3_1() {
    let mut a:i32 = 1;
    let mut b = 1;
}

fn p_2_3_2() {
    let a = 1;
    let mut b:i32 = a;
}

fn p_2_3_3() {
    // shallow
    let mut a:i32 = 1;
    let mut a = 2;
    let mut a:i32 = 3;
    let mut a;
    a = 100;
    a = a + 1;
    let a;
    a = 1;
    // let mut a:i32 = b;
}