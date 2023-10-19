use std::str;

pub fn run() {
    let cipher =
        hex::decode("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
            .unwrap();

    let text = find_most_probable_text(&cipher);
    println!("Challenge 3: {}", str::from_utf8(&text).unwrap());
}

pub fn find_most_probable_text(cipher: &[u8]) -> Vec<u8> {
    let mut key: Vec<u8> = vec![0; cipher.len()];
    let mut result: Vec<u8> = Vec::with_capacity(0);
    let mut highest = 0;

    for i in 1..=255 {
        key.fill(i);
        let text = crate::utils::xor_bytes(&cipher, &key);

        let count = crate::utils::common_letter_count(&text);
        if count >= highest {
            highest = count;
            result = text;
        }
    }
    return result;
}