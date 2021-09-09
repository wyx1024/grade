package limiter

import (
	"math"
	"time"
)

type BucketLimit struct {
	rate       float64
	bucketSize float64
	unixNano   int64
	curWater   float64
}

func NewBucketLimit(rate float64, bucketSize int64) *BucketLimit {
	return &BucketLimit{
		rate:       rate,
		bucketSize: float64(bucketSize),
		unixNano:   time.Now().Unix(),
		curWater:   0,
	}
}

func (b *BucketLimit) reflesh()  {
	now := time.Now().Unix()
	diffSec := float64(now-b.unixNano)/1000/1000/1000
	b.curWater = math.Max(0, b.curWater - b.rate*diffSec)
	b.unixNano = now
}

func (b *BucketLimit) Allow() bool {
	b.reflesh()
	if b.curWater < b.bucketSize{
		return true
	}
	return false
}
