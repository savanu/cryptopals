pub fn run() {
    let text = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal".as_bytes();
    
    let result = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f";
    let mut key_expanded: Vec<u8> = Vec::with_capacity(text.len());
    let key = [b'I', b'C', b'E'];


    for i in 0..text.len() {
        key_expanded.push(key[i % 3]);
    }

    let cipher = crate::utils::xor_bytes(text, &key_expanded);
    let encoded_cipher = hex::encode(cipher);

    assert_eq!(encoded_cipher, result);

    println!("Enc: {}", encoded_cipher);
}