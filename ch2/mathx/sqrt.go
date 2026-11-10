package mathx
// Sqrt uses Newton's method. Capitalised names are exported.
func Sqrt(x float64) float64 {
z := 1.0
for range 20 {
z -= (z*z - x) / (2 * z)
}
return z
}