package blockchain

import (
	"sync"
	"work_queue"
)

type miningWorker struct {
	// Should implement work_queue.Worker
	MinVal	uint64
	MaxVal	uint64
	WorkBlock	Block
}

func (mw *miningWorker) Run() interface{} {
	var blk *Block
	for i := mw.MinVal; i < mw.MaxVal; i++ {
		blk = new(Block)
		*blk = mw.WorkBlock
		blk.SetProof(i)
		if blk.ValidHash() {
			return MiningResult{Proof: i, Found: true}
		}
	}
	return MiningResult{Proof: uint64(0), Found: false}
}

type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // true if valid proof-of-work was found.
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
	var result MiningResult
	wg := sync.WaitGroup{}
	var chunkSize uint64 = (end - start)/chunks - 1
	q := work_queue.Create(uint(workers), uint(chunks))
	var worker *miningWorker
	for i := start; i <= end; i += chunkSize {
		wg.Add(1)
		go func(lowerBound uint64) {
			defer wg.Done()
			worker = new(miningWorker)
			worker.MinVal = lowerBound
			if worker.MinVal + chunkSize  > end {
				worker.MaxVal = end
			} else {
				worker.MaxVal = worker.MinVal + chunkSize
			}
			worker.WorkBlock = blk
			q.Enqueue(worker)
		}(i)

		r := <- q.Results
		result = r.(MiningResult)
		if result.Found {
			q.Shutdown()
			break
		}
	}
	wg.Wait()

	return result
}

// Call .MineRange with some reasonable values that will probably find a result.
// Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << (8 * blk.Difficulty)) // 4 * 2^(bits that must be zero)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}
