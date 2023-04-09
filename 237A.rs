use std::io;
use std::str::FromStr;

fn main(){
    let mut orders = String::new();
    io::stdin().read_line(&mut orders)
        .unwrap();
    //let orders :u8 = orders.parse::<u8>().unwrap();
    let orders = match u8::from_str(orders.trim()){
        Ok(orders) => orders,
        Err(_) => {println!("tka3brat");return}
    };
    let mut cashes: u8 = 0 ;
    let mut orders_vec: Vec<(u8, u8)> = Vec::new(); 

    for i in 0..orders {
        if cashes == 0 {
            cashes = 1 ;
        }
        let mut order = String::new();
        io::stdin().read_line(&mut order).expect("failed to read order time");
        
        let mut iter = order.split_whitespace();
        let h: u8 = iter.next().unwrap().parse::<u8>().unwrap();
        let m: u8 = iter.next().unwrap().parse::<u8>().unwrap();
        orders_vec.push((h, m));

        if orders_vec.len() > 1 {
            let index: usize = i as usize;
            let  test_cashes: u8 = time_less_than_minute(&orders_vec, index, index-1, 1);
            if test_cashes > cashes {
                cashes = test_cashes;
            }
        }
    }
    println!("{}", cashes);
}

fn time_less_than_minute(orders_vec: &Vec<(u8, u8)>, i:usize, iv:usize, test_cashes: u8) -> u8{
    let (h1, m1) = orders_vec[i];
    let (h2, m2) = orders_vec[iv];
    let (ht, mt) =orders_vec[iv+1];
    let hour_diff;
    if h1 > h2 {
         hour_diff = h1-h2;
    }else{
         hour_diff = h2-h1;
    }

    if iv+1 != i && (h2 > ht || (h2 == ht && m2 > mt)) {
        
    }else{
        return test_cashes;
    }

    if hour_diff > 1 && hour_diff != 23{
        return test_cashes;
    } else if hour_diff == 1 || hour_diff == 23{
        if m1!= 00 || m2 != 59 {
            return test_cashes;
        } else {
            if iv == 0{
                return test_cashes + 1 ;
            }else{
                return time_less_than_minute(&orders_vec, i, iv-1, test_cashes + 1);
            }
        }
    } else if hour_diff == 0 && m1-m2 > 1 {
            return test_cashes;
    } else {
        if iv == 0{
            return test_cashes + 1 ;
        } else {
            return time_less_than_minute(&orders_vec, i, iv-1, test_cashes + 1);
        }
    }
}
