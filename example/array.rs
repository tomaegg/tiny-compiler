fn main() -> i32{
    let mut a0:[i32;4];

    let a1 = [1,2,3,4,5];

    let mut a2 = [[1,2],[4,5]];
    let a3 = [[1,2,4,5],[4,5,1,2],[1,2,3,4]];
    let a4: [[[i32;2];4];6];
    return a1[4] * a1[4] + 1;
}

fn case1() {
    let a = 1;
    let mut a0 = [a,a,a,a];
}

fn case2() {
    let mut a0;
    a0 = [1,1,1,1];
    a0[1];

    let mut a1 = [[[1,2],[2,3]], [[1,2],[2,3]]];
    a1[1][1][1];

    let mut b = a1[1][1][1];
}
