pub fn run() {
    let answer = "746865206b696420646f6e277420706c6179";
    let a = "1c0111001f010100061a024b53535009181c";
    let b = "686974207468652062756c6c277320657965";

    let a = hex::decode(&a).unwrap();
    let b = hex::decode(&b).unwrap();

    let result = crate::utils::xor_bytes(&a, &b);
    let result = hex::encode(result);

    assert_eq!(answer, result);
    println!("Challenge 2: {result}");
}