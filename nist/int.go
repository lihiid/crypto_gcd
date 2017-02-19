






// Set to the gcd of a and its module
func (i *Int) Gcd(a abstract.Scalar) abstract.Scalar {
	ai := a.(*Int)
	i.M = ai.M
	i.V.GCD(nil, nil, &ai.V, i.M)
	return i
}
