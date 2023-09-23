use std::env;
use std::str;

mod utils;
mod challenge_one;
mod challenge_two;
mod challenge_three;
mod challenge_four;
mod challenge_five;

fn main() {
    let mut args = env::args();
    if args.len() != 3 {
        return;
    }

    args.next(); // Ignore exec name

    let set_num_opt = args.next();
    let challenge_num_opt = args.next();

    match (set_num_opt, challenge_num_opt) {
        (Some(_), Some(challenge_num)) => dispatch(&challenge_num),
        (_, _) => println!("Invalid args"),
    }
}

fn dispatch(challenge_num: &str) {
    match challenge_num {
        "1" => challenge_one::run(),
        "2" => challenge_two::run(),
        "3" => challenge_three::run(),
        "4" => challenge_four::run(),
        "5" => challenge_five::run(),
        _ => println!("Invalid args"),
    }
}