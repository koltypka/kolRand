package kolRand

import (
	"math"
	"math/rand"
	"time"
)

type KolRandom struct {
	seed *rand.Rand //генератор случайных чисел
}

func NewKolRandom() KolRandom {
	return KolRandom{rand.New(rand.NewSource(time.Now().UnixMicro()))}
}

func (rnd *KolRandom) makeUniformFloat64() float64 {
	return rnd.seed.Float64()
}

//генератор случайных чисел работает только внутри тела функции
//генератор случайных чисел распределенных по показательному закону
func (rnd *KolRandom) MakeExp(lambda float64) uint16 {
	returnUnit16 := 1 / (rnd.seed.ExpFloat64() / lambda)

	return uint16(returnUnit16)
}

func (rnd *KolRandom) MakeUniform(n int32) uint16 {
	return uint16(rnd.seed.Int31n(n))
}

func (rnd *KolRandom) MakeUniformFloatFromZeroToOne(n, m float32) float32 {
	return (float32(rnd.MakeUniformRange(int32(n)*10, int32(m)*10)) / 10)
}

func (rnd *KolRandom) MakeUniformRange(n, m int32) uint16 {
	return uint16((rnd.seed.Int31() % (m - n)) + n)
}

func (rnd *KolRandom) MakePoisson(lambda float64) uint16 {
	var P float64 = 0
	var counter, uint16Result uint16 = 0, 0

	for true {
		if P > lambda {
			break
		}

		P = P + float64(rnd.MakeExp(1))
		counter = counter + 1
	}

	uint16Result = counter
	return uint16Result
}

func (rnd *KolRandom) MakePoissonMultiply(lambda float64) uint16 {
	var P, newP float64 = rnd.makeUniformFloat64(), 0
	var counter, uint16Result uint16 = 0, 0

	for true {
		newP = P * rnd.makeUniformFloat64()

		if newP == 0 {
			newP = rnd.makeUniformFloat64()
		}

		if P >= math.Exp(-math.Abs(lambda)) && (newP < math.Exp(-math.Abs(lambda))) {
			break
		}

		P = newP
		counter = counter + 1
	}
	uint16Result = counter

	return uint16Result
}
