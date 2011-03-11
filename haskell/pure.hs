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