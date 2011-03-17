import Data.Char

main = do
	raw <- getContents
	let stripped = filter (not . isSpace) raw
	if valid stripped
		then print $ eval $ parse stripped
		else putStrLn "Invalid program."

data Tree = Tree Tree Tree | Leaf Char
	deriving (Eq)

data Func = S | K | I | S1 Func | S2 Func Func | K1 Func | Simple Char | Pair Func Func
	deriving (Eq)

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

valid :: String -> Bool
valid s = count s == Just 1

up :: Int -> Maybe Int
up x = Just (x+1)

down :: Int -> Maybe Int
down 1 = Nothing
down x = Just (x-1)

adjust :: Char -> Maybe Int -> Maybe Int
adjust '`' acc = acc >>= down
adjust c acc = acc >>= up

count :: String -> Maybe Int
count = foldr adjust (Just 0)

splitStart :: Int -> String -> String -> (String, String)
--left is reversed
splitStart 0 left right = (reverse left, right)
splitStart n left ('`':right) = splitStart (n+1) ('`':left) right
splitStart n left (x:right) = splitStart (n-1) (x:left) right

split :: String -> (String, String)
split = splitStart 1 ""

evalChar :: Char -> Func
evalChar 's' = S
evalChar 'k' = K
evalChar 'i' = I
evalChar x = Simple x

parse :: String -> Tree
parse ('`':rest) = Tree (parse $ fst splat) (parse $ snd splat)
	where splat = split rest
parse [x] = Leaf x

eval :: Tree -> Func
eval (Tree left right) = apply (eval left) (eval right)
eval (Leaf x) = evalChar x
