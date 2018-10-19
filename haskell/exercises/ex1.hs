det a b c = b^2 - 4*a*c -- determinant of quadratic
quadsol1 a b c = (-b - sqrt(det a b c))/2*a -- solution 1
quadsol2 a b c = (-b + sqrt(det a b c))/2*a -- solution 2

take3_a a = take 3 a
third_a a = a !! 2 --indexing starts from 0
third_b (_:b:c) = head c -- using recursion

hailstone 1 = 0
hailstone a
    | a `mod` 2 == 0 = a `div` 2 -- a is even => a/2
    | otherwise = 3*a+1          -- a is odd = > 3*a+1
