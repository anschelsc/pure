import Data.Char
import Text.ParserCombinators.Parsec
import System (getArgs)

main = do
	raw <- getContents
	args <- getArgs
	let stripped = filter (not . isSpace) raw
	case parse parser "" stripped of
		Right tree -> if null args
			then print $ eval tree
			else print $ foldr elim tree $ args !! 0
		Left err -> putStrLn $ "Parse error at " ++ show err

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

instance Show Tree where
	show (Tree left right) = "`" ++ show left ++ show right
	show (Leaf c) = [c]

apply :: Func -> Func -> Func
apply S x = S1 x
apply K x = K1 x
apply I x = x
apply (S1 x) y = S2 x y
apply (S2 x y) z = apply (apply x z) (apply y z)
apply (K1 x) _ = x
apply x@(Simple _) y = Pair x y
apply x@(Pair _ _) y = Pair x y

parser' :: Parser Tree
parser' = do
		char '`'
		left <- parser'
		right <- parser'
		return $ Tree left right
	<|> (letter >>= return . Leaf)

parser = do
		tree <- parser'
		eof
		return tree

evalChar :: Char -> Func
evalChar 's' = S
evalChar 'k' = K
evalChar 'i' = I
evalChar x = Simple x

eval :: Tree -> Func
eval (Tree left right) = apply (eval left) (eval right)
eval (Leaf x) = evalChar x

has :: Tree -> Char -> Bool
(Leaf x) `has` y = x==y
(Tree left right) `has` c = (left `has` c) || (right `has` c)

elim :: Char -> Tree -> Tree
elim c (Leaf d)
	| c==d = Leaf 'i'
elim c (Tree left (Leaf d))
	| (c==d) && (not $ left `has` c) = left
elim c t
	| not $ t `has` c = Tree (Leaf 'k') t
elim c (Tree left right) = Tree (Tree (Leaf 's') $ elim c left) $ elim c right