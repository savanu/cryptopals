use std::str;
use std::fs::File;
use std::io::{self, BufRead};

pub fn run() {
    let file = match File::open("../files/challenge_4_data.txt") {
        Err(reason) => panic!("Couldn't open file: {}", reason),
        Ok(file) => file,
    };

    let mut i = 0;
    for line in io::BufReader::new(file).lines() {
        let cipher = hex::decode(line.unwrap()).unwrap();
        let text = crate::challenge_three::find_most_probable_text(&cipher);
        match str::from_utf8(&text) {
            Ok(t) => println!("Line {}, Text: {}", i, t),
            _ => (),
        }

        i += 1;
    }
}