use std::io;
use std::str::FromStr;
use std::collections::HashMap;
//https://codeforces.com/problemset/problem/237/A
fn main(){
    let mut orders = String::new();
    io::stdin().read_line(&mut orders)
        .unwrap();
    let orders = match u32::from_str(orders.trim()){
        Ok(orders) => orders,
        Err(_) => {println!("tka3brat");return}
    };
    let mut orders_map : HashMap<u32, u32> = HashMap::new(); 
    
    for _i in 0..orders {
        let mut order = String::new();
        io::stdin().read_line(&mut order).expect("failed to read order time");
        
        let mut iter = order.split_whitespace();
        let h: u32 = iter.next().unwrap().parse::<u32>().unwrap();
        let m: u32 = iter.next().unwrap().parse::<u32>().unwrap();
        let t: u32 = h * 60 + m;
        *orders_map.entry(t).or_insert(0) += 1;
    }
    if let Some((_key, value)) = orders_map.iter().max_by_key(|&(_, value)| value) {
        println!("{}", value);
    }
}

