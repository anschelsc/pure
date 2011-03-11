import Data.Char

data Func = S | K | I | S1 Func | S2 Func Func | K1 Func | Simple Char | Pair Func Func deriving (Eq)

instance Show Func where
	show S = "s"
	show K = "k"
	show I = "i"
	show (S1 x) = "`s" ++ show x
	show (S2 x y) = "``s" ++ show x ++ show y
	show (K1 x) = "`k" ++ show x
	show (Simple c) = [c]
	show (Pair x y) = "`" ++ show x ++ show y

apply :: Func -> Func -> Func
apply S x = S1 x
apply K x = K1 x
apply I x = x
apply (S1 x) y = S2 x y
apply (S2 x y) z = apply (apply x z) (apply y z)
apply (K1 x) _ = x
apply x@(Simple _) y = Pair x y
apply x@(Pair _ _) y = Pair x y

validStart :: Int -> String -> Bool
validStart 0 "" = True
validStart _ "" = False
validStart 0 _ = False
validStart n ('`':xs) = validStart (n+1) xs
validStart n (x:xs) = validStart (n-1) xs

valid :: String -> Bool
valid = validStart 1

splitStart :: Int -> String -> String -> (String, String)
--left is reversed
splitStart 0 left right = (reverse left, right)
splitStart n left ('`':right) = splitStart (n+1) ('`':left) right
splitStart n left (x:right) = splitStart (n-1) (x:left) right

split :: String -> (String, String)
split = splitStart 1 ""

parseChar :: Char -> Func
parseChar 's' = S
parseChar 'k' = K
parseChar 'i' = I
parseChar x = Simple x

parse :: String -> Func
parse ('`':rest) = apply (parse $ fst splat) (parse $ snd splat)
	where splat = split rest
parse [x] = parseChar x

main = do
	raw <- getContents
	let stripped = filter (not . isSpace) raw
	if valid stripped
		then print $ parse stripped
		else putStrLn "Invalid program."