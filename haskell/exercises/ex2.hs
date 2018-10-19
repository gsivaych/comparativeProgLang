hailstone a
    | a `mod` 2 == 0 = a `div` 2 -- a is even => a/2
    | otherwise = 3*a+1          -- a is odd = > 3*a+1

hailLen 1 = 0
hailLen n = 1 + hailLen(hailstone n)

-- Divisors and Primes
divisors :: Int -> [Int]
divisors n = [i | i <- [2..(n `div` 2)], n `mod` i == 0]
primes :: Int -> [Int]
primes n = [i | i <- [2..n], divisors i == []]

-- Joining Strings
join :: [Char] -> [[Char]] -> [Char]
join separator [] = ""
join separator [last] = last
join separator (head:tail) = head ++ separator ++ (join separator tail)

-- Pythagorean Triples
pythagorean :: Int -> [(Int, Int, Int)]
pythagorean n = [(a, b, c) | a <- [1..n], b <- [1..n], c <- [1..n], a^2+b^2==c^2, a<b,b<c]

