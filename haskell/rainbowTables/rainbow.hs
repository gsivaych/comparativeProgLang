import RainbowAssign
import qualified Data.Map as Map
import Data.Maybe

pwLength, nLetters, width, height :: Int
filename :: FilePath
pwLength = 8            -- length of each password
nLetters = 5            -- number of letters to use in passwords: 5 -> a-e
width = 40              -- length of each chain in the table
height = 1000           -- number of "rows" in the table
filename = "table.txt"  -- filename to store the table

rebase :: (Integral a) => a -> a -> [a]
rebase newbase inum
        | newbase < 2 = error "invalid base !"
        | abs inum >= newbase = [inum `mod` newbase] ++ (rebase newbase (inum `div` newbase))
        | abs inum < newbase = [abs inum] ++ take pwLength (if (inum < 0) then(repeat (newbase-1)) else (repeat 0))
        | otherwise = []

pwReduce :: Hash -> Passwd
pwReduce hash = map toLetter (reverse (take pwLength (rebase nLetters (fromEnum hash))))

nthHash :: Int -> Passwd -> Hash
nthHash n pw0
        | n == 0 = pwHash pw0
        | n > 0 = pwHash (pwReduce (nthHash (n-1) pw0))
        | otherwise = error "Couldn't compute !"

rainbowTable :: Int -> [Passwd] -> Map.Map Hash Passwd
rainbowTable width' pwList = Map.fromList (zip (map (nthHash width') pwList) pwList)

generateTable :: IO ()
generateTable = do
  table <- buildTable rainbowTable nLetters pwLength width height
  writeTable table filename

look4myPasswd :: Hash -> Map.Map Hash Passwd -> Maybe Passwd
look4myPasswd inhash table = Map.lookup inhash table

look4passwd0 :: Hash -> Int -> Map.Map Hash Passwd -> Maybe Passwd
look4passwd0 inhash n table =
         if (look4myPasswd inhash table == Nothing && n > 0)
           then look4passwd0 (pwHash (pwReduce inhash)) (n-1) table
         else look4myPasswd inhash table

look4acPasswd :: Passwd -> Hash -> Int ->Maybe Passwd
look4acPasswd passwd0 inhash n =
         if (pwHash passwd0 /= inhash)
           then look4acPasswd (pwReduce (pwHash passwd0)) inhash (n-1)
         else Just passwd0

findPassword :: Map.Map Hash Passwd -> Int -> Hash -> Maybe Passwd
findPassword table width'' hash =
         if passwd0val /= Nothing && look4acPasswd (fromJust passwd0val) hash width'' /= Nothing
          then look4acPasswd (fromJust passwd0val) hash width''
         else Nothing
             where passwd0val = look4passwd0 hash height table

test :: Int -> IO ([Passwd], Int)
test n = do
  table <- readTable filename
  pws <- randomPasswords nLetters pwLength n
  let hs = map pwHash pws
  let result = mapMaybe (findPassword table width) hs
  return (result, length result)

main :: IO ()
main = do
  generateTable
  res <- test 1
  print res
