-- recursive iterate
myIterate :: (a -> a) -> a -> [a]
myIterate func x = [x] ++ (myIterate func (func x))

-- recursive takeWhile
myTakeWhile :: (a -> Bool) -> [a] -> [a]
myTakeWhile predicate [] = []
myTakeWhile predicate (head:tail)
    | predicate head == True = [] ++ [head] ++ (myTakeWhile predicate tail)
    | otherwise = myTakeWhile predicate tail

-- pascal's triangle
pascal :: Int -> [Int]
pascal 0 = [1]
pascal n = [1] ++ [(a+b) | (a,b) <- (zip prev (tail prev))] ++ [1]
    where prev = pascal (n-1)

-- point-free addition
addPair :: (Int,Int) -> Int
addPair = uncurry (curry (\(x,y) -> x+y))

--point-free filtering
withoutZeros :: (Num a,Eq a) => [a] -> [a]
withoutZeros = filter (\x -> x /= 0)

-- exploring fibbonaci
fib :: Int -> Int
fib 0 = 0
fib 1 = 1
fib 2 = 1
fib n = (fib (n-1)) + (fib (n-2))

fibs = map fib [0..]

-- something else
things :: [Integer]
things = 0 : 1 : zipWith (+) things (tail things)

