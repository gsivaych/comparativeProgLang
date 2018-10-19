-- imports at top : mandatory
import Data.Time.Calendar
import Data.Time.Calendar.OrdinalDate

-- merging
merge :: Ord a => [a] -> [a] -> [a]
merge [] [] = []
merge xs [] = xs
merge [] ys = ys
merge (x:xs) (y:ys)
    | x < y     = x : merge xs (y:ys)
    | otherwise = y : merge (x:xs) ys


-- tail_recursive_hailstone
hailStone :: Int -> Int
hailStone 1 = 0
hailStone n
    | even n    = n `div` 2
    | otherwise = 3*n + 1
hailLen :: Int -> Int
hailLen n = hailTail 0 n
  where
    hailTail a 1 = a
    hailTail a n = hailTail (a+1) (hailStone n)

--factorials
fact :: Int -> Int
fact 0 = 1
fact n = n * fact (n-1)

fact' :: Int -> Int
fact' n = foldl (*) n [1..n-1]

--haskell library and date
daysInYear :: Integer -> [Day]
daysInYear y = [jan1..dec31]
    where jan1 = (fromGregorian y 01 01)
          dec31 = (fromGregorian y 12 31)

isFriday :: Day -> Bool
isFriday day
    | snd (mondayStartWeek day) == 5 = True
    | otherwise                       = False

getDay (y,m,d) = d
isPrime n = null [ x | x <- [2..n-1], n `mod` x  == 0]
isPrimeDay :: Day -> Bool
isPrimeDay day = isPrime (getDay (toGregorian day))

primeFridays :: Integer -> [Day]
primeFridays year = [day | day <- daysInYear year, isFriday day, isPrimeDay day]
