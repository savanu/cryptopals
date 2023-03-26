use std::iter::zip;

use base64::{engine::general_purpose, Engine as _};

pub fn xor_bytes(a: &[u8], b: &[u8]) -> Vec<u8> {
    assert!(a.len() == b.len(), "XOR arrays are not equal length. Len a: {}, Len b: {}", a.len(), b.len());
    let mut result = Vec::with_capacity(a.len());
    for (x, y) in zip(a, b) {
        result.push(x ^ y);
    }
    return result;
}

pub fn hex_to_base64(hex_str: &str) -> String {
    let arr = hex::decode(hex_str).unwrap();
    return general_purpose::STANDARD.encode(arr);
}

pub fn common_letter_count(x: &[u8]) -> i32 {
    let mut count = 0;
    for a in x.iter() {
        if is_common_letter(a) {
            count += 1;
        }
    }
    return count;
}

fn is_common_letter(x: &u8) -> bool {
    match x {
        b'E' => true,
        b'T' => true,
        b'A' => true,
        b'O' => true,
        b'I' => true,
        b'N' => true,
        b'S' => true,
        b'H' => true,
        b'R' => true,
        b'D' => true,
        b'L' => true,
        b'U' => true,
        b'C' => true,
        b'M' => true,
        b'e' => true,
        b't' => true,
        b'a' => true,
        b'o' => true,
        b'i' => true,
        b'n' => true,
        b's' => true,
        b'h' => true,
        b'r' => true,
        b'd' => true,
        b'l' => true,
        b'u' => true,
        b'c' => true,
        b'm' => true,
        _ => false,
    }
}