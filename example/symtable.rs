// 符号表和类型检查相关测试样例

// 正确声明与赋值
fn test_var_declare_assign() {
    let mut x: i32 = 10;
    x = 20;
    let y = x + 5;
}

// 重复声明（同作用域）
fn test_redeclare_same_scope() {
    let a = 1;
    let a = 2;
    let a: i32 = 3;
}

// 不同作用域的声明
fn test_var_scope() {
    let b = 1;
    {
        let b = 2; // 新作用域，允许同名
    }
    let c = b + 1;
}

// 函数参数类型检查 - 应报错
fn foo(x: i32) {}
fn test_func_param_type() {
    foo(); // 错误：类型不匹配
}

// 函数返回类型检查 - 应报错
fn bar() -> i32 {
    return; // 错误：返回类型不匹配
}

// 正确的类型推断
fn test_type_infer() {
    let h = 5;
    let i = h + 2;
}
