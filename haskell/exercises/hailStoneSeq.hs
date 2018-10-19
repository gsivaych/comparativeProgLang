--hailstone sequence generator
hailStoneSeq :: Int -> [Int]
hailStoneSeq 1 = [1]
hailStoneSeq n = n : hailStoneSeq(next)
    where next
            | even n   = n `div` 2
            | otherwise = 3*n + 1
