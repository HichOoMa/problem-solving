use std::io;
fn main() {
    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("failed to read input");

    let mut iter = input.split_whitespace();
    let m: u64 = iter.next().unwrap().parse::<u64>().unwrap();
    let n: u64 = iter.next().unwrap().parse::<u64>().unwrap();
    let a: u64 = iter.next().unwrap().parse::<u64>().unwrap();
    
    let mut xa = m/a;
    let mut ya = n/a;

    if a * xa != m {
        xa = xa + 1;
    }

    if a * ya != n {
        ya = ya + 1;
    }

    println!("{}",xa*ya);
}
