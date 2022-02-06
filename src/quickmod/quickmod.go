package quickmod

import "math/big"

func FastPow(a, b, p *big.Int) *big.Int {
	var res = big.NewInt(1)
	a.Mod(a, p)
	for b.Int64() != 0 {
		if b.Int64()&1 == 1 {
			res.Mul(res, a).Mod(res, p)
		}
		b = b.Rsh(b, 1)
		a.Mul(a, a).Mod(a, p)
	}
	return res
}

func Mod(rat *big.Rat, p *big.Int) *big.Int {
	if rat.Denom().Int64() == 1 {
		return new(big.Int).Mod(rat.Num(), p)
	}
	res := big.NewInt(1)
	fastPow := FastPow(rat.Denom(), new(big.Int).Sub(p, big.NewInt(2)), p)
	res.Mod(rat.Num(), p).Mul(res, fastPow).Mod(res, p)
	return res
}