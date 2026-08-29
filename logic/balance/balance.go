package balance

import "sync/atomic"

type Balance struct {
	balance atomic.Int64
}

func NewBalance() *Balance {
	return &Balance{}
}

func (b *Balance) Increase(n int64) {
	b.balance.Add(n)
}

func (b *Balance) Decrease(n int64) {
	res := b.balance.Load()
	res -= n
	b.balance.Store(res)
}

func (b *Balance) Get() int64 {
	return b.balance.Load()
}
