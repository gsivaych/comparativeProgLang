-- non-recursive hailstone seq generator
import Data.List
import Data.Maybe
hailStone :: Int -> Int
hailStone 1 = 0
hailStone n
    | even n     = n `div` 2
    | otherwise  = 3*n + 1

hailSeq :: Int -> [Int]
hailSeq n = unfoldr (\x -> if x == 1 then Nothing else Just (x, hailStone x)) n ++ [1]

-- joining strings again
join :: [Char] -> [[Char]] -> [Char]
join separator list = foldr (\a b-> a ++ if b=="" then b else separator ++ b) "" list

-- merge sort
  -- function merge
merge :: Ord a => [a] -> [a] -> [a]
merge [] [] = []
merge xs [] = xs
merge [] ys = ys
merge (x:xs) (y:ys)
    | x < y     = x : merge xs (y:ys)
    | otherwise = y : merge (x:xs) ys

divide :: Ord a => [a] -> ([a], [a])
divide list
    | even len = (take (len `div` 2) list,drop (len `div` 2) list)
    | otherwise = (take ((len+1) `div` 2) list,drop ((len+1) `div` 2) list)
        where len = length list

mergeSort :: Ord a => [a] -> [a]
mergeSort [] = []
mergeSort [x] = [x]
mergeSort list = merge (mergeSort left) (mergeSort right)
                   where (left,right) = divide list

-- searching -- maybe
findElt :: Ord a => a -> [a] -> Maybe Int
findElt element list = case (listToMaybe list) of
                          Nothing -> Nothing
                          Just e | e == element -> Just 0
                          Just _ -> Just (1 + (fromMaybe 0 (findElt element (tail list))))
