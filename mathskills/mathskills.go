package mathskills
import "math"
func Average(data []float64) float64 {
	sum := 0.0
	count := float64(len(data))
	for _, num := range data {
		sum += num
	}
	return sum / count
}
func StandardDeviation(data []float64) float64 {
	return math.Sqrt(Variance(data))
}
func Variance(data []float64) float64 {
	mean := Average(data)
	var sqDiff float64
	for _, n := range data {
		diff := n - mean
		sqDiff += diff * diff
	}
	return sqDiff / float64(len(data))
}

func LinearRegLine(inp []float64) (float64, float64) {
	meanY := Average(inp)
	n := float64(len(inp))

	
	// Calculate mean of x (indices)
	meanX := (n - 1) / 2 // Since indices go from 0 to n-1, use (n-1)/2

	// Calculate e_xy and esqr_x
	eXY := 0.0  // covariance are they covariance or  not 0-1
	esqrX := 0.0
	
	for i := 0; i < len(inp); i++ {
		x := float64(i) // x value based on index
		eXY += (inp[i] - meanY) * (x - meanX) // cpovariance between the inp and there indice
		esqrX += (x - meanX) * (x - meanX)
	}

	// Calculate coefficients
	coefK := eXY / esqrX
	coefB := meanY - coefK*meanX
	
	return coefK, coefB
}
func PearsonCoef(inp []float64) float64 {
	meanY := Average(inp)
	n := float64(len(inp))

	// Calculate mean of x (indices)
	meanX := (n - 1) / 2 // Indices range from 0 to n-1

	// Calculate variance of x and standard deviation
	varX := 0.0
	for i := 0; i < len(inp); i++ {
		x := float64(i)
		varX += (meanX - x) * (meanX - x)
	}
	varX /= n
	stdX := math.Sqrt(varX)
	stdY := StandardDeviation(inp)

	// Calculate covariance
	eXY := 0.0
	for i := 0; i < len(inp); i++ {
		eXY += (inp[i] - meanY) * (float64(i) - meanX)
	}
	cov := eXY / n

	// Calculate Pearson correlation coefficient
	r := cov / (stdX * stdY)
	return r
}