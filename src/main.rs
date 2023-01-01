use std::iter::zip;

fn main() {
    challenge_2();
}

fn challenge_2() {
    let answer = "746865206b696420646f6e277420706c6179";
    let a = "1c0111001f010100061a024b53535009181c";
    let b = "686974207468652062756c6c277320657965";

    let a = hex::decode(&a).unwrap();
    let b = hex::decode(&b).unwrap();

    let result = xor_bytes(&a, &b);
    let result = hex::encode(result);

    assert_eq!(answer, result);
    println!("Challenge 2: {result}");
}

fn challenge_1() {
    let answer = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t";
    let payload = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d";
    let result = hex_to_base64(&payload);

    assert_eq!(answer, result);
    println!("Challenge 1: {result}");
}

fn xor_bytes(a: &[u8], b: &[u8]) -> Vec<u8> {
    let mut result = Vec::with_capacity(a.len());
    for (x, y) in zip(a, b) {
        result.push(x ^ y);
    }
    return result;
}

fn hex_to_base64(hex_str: &str) -> String {
    base64::encode(hex::decode(hex_str).unwrap())
}
