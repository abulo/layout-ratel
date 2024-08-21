package initial

import "github.com/abulo/layout/internal/ants"

func (initial *Initial) InitPool(poolSize int) *Initial {
	// 创建一个 Ants 池
	pool, err := ants.NewAnts(poolSize)
	if err != nil {
		panic(err)
	}
	initial.Pool = pool
	return initial
}
