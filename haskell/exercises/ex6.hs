import Data.Ratio

rationalSum :: Int -> [Ratio Int]
rationalSum n = [(x % y) | x <- [1..n], y <- [1..n], x+y == n]

rationalSumLowest :: Int -> [Ratio Int]
rationalSumLowest n = [(x % y) | x <- [1..n], y <- [1..n], (x+y) == n, (gcd x y)==1]

rationals :: [Ratio Int]
rationals = concat (map rationalSumLowest [1..])


sumFile :: IO ()
sumFile = do
    ioString <- readFile "input.txt"
    let stringsList = lines ioString
    let numValList = map read stringsList
    let sumofList = sum numValList
    print $ sumofList
