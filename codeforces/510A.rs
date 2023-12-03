use std::io;

fn main() {
    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("Failed to read line");

    let mut iter = input.split_whitespace();
    let n: usize = iter.next().unwrap().parse().unwrap();
    let m: usize = iter.next().unwrap().parse().unwrap();
for i in 1..n+1 {
        if i%2==1 {
            println!("{}","#".repeat(m));
        }else if i%4==0 {
            println!("{}{}","#",".".repeat(m-1));
        }else{
            println!("{}{}",".".repeat(m-1),"#");
        }
    }
}

