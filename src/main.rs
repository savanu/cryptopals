use std::env;
use std::fs::File;
use std::io::{self, BufRead};
use std::str;

mod utils;

fn main() {
    let mut args = env::args();
    if args.len() != 3 {
        return;
    }

    args.next();

    match (args.next(), args.next()) {
        (Some(a), Some(b)) => match (a.as_str(), b.as_str()) {
            ("1", x) => set_one(x),
            (_, _) => println!("Invalid args"),
        },
        (_, _) => println!("Invalid args"),
    }
}

fn set_one(challenge_num: &str) {
    match challenge_num {
        "1" => challenge_1(),
        "2" => challenge_2(),
        "3" => challenge_3(),
        "4" => challenge_4(),
        "5" => challenge_5(),
        _ => println!("Invalid args"),
    }
}

fn challenge_5() {
    let text = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal".as_bytes();
    
    let result = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f";
    let mut key_expanded: Vec<u8> = Vec::with_capacity(text.len());
    let key = [b'I', b'C', b'E'];


    for i in 0..text.len() {
        key_expanded.push(key[i % 3]);
    }

    let cipher = utils::xor_bytes(text, &key_expanded);
    let encoded_cipher = hex::encode(cipher);

    assert_eq!(encoded_cipher, result);

    println!("Enc: {}", encoded_cipher);
}

fn challenge_4() {
    let file = match File::open("../files/challenge_4_data.txt") {
        Err(reason) => panic!("Couldn't open file: {}", reason),
        Ok(file) => file,
    };

    let mut i = 0;
    for line in io::BufReader::new(file).lines() {
        let cipher = hex::decode(line.unwrap()).unwrap();
        let text = find_most_probable_text(&cipher);
        match str::from_utf8(&text) {
            Ok(t) => println!("Line {}, Text: {}", i, t),
            _ => (),
        }

        i += 1;
    }
}

fn challenge_3() {
    let cipher =
        hex::decode("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
            .unwrap();

    let text = find_most_probable_text(&cipher);
    println!("Challenge 3: {}", str::from_utf8(&text).unwrap());
}

fn find_most_probable_text(cipher: &[u8]) -> Vec<u8> {
    let mut key: Vec<u8> = vec![0; cipher.len()];
    let mut result: Vec<u8> = Vec::new();
    let mut highest = 0;

    for i in 1..=255 {
        key.fill(i);
        let text = utils::xor_bytes(&cipher, &key);

        let count = utils::common_letter_count(&text);
        if count >= highest {
            highest = count;
            result = text;
        }
    }
    return result;
}

fn challenge_2() {
    let answer = "746865206b696420646f6e277420706c6179";
    let a = "1c0111001f010100061a024b53535009181c";
    let b = "686974207468652062756c6c277320657965";

    let a = hex::decode(&a).unwrap();
    let b = hex::decode(&b).unwrap();

    let result = utils::xor_bytes(&a, &b);
    let result = hex::encode(result);

    assert_eq!(answer, result);
    println!("Challenge 2: {result}");
}

fn challenge_1() {
    let answer = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t";
    let payload = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d";
    let result = utils::hex_to_base64(&payload);

    assert_eq!(answer, result);
    println!("Challenge 1: {result}");
}