package raw

import (
	"math/rand"
)

func genRandom() bool {
	return 5-rand.Intn(11) > 0
}
