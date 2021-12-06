package store

type FishStore struct {
	store map[int]uint64
}

func NewFishStore(fishes []int) *FishStore {
	dataMap := map[int]uint64{
		-1: 0,
		0:  0,
		1:  0,
		2:  0,
		3:  0,
		4:  0,
		5:  0,
		6:  0,
		7:  0,
		8:  0,
	}
	for _, fishTimer := range fishes {
		data := dataMap[fishTimer]
		data++
		dataMap[fishTimer] = data
	}
	return &FishStore{store: dataMap}
}

func (s *FishStore) DecrementFish() {
	for i := 0; i <= 8; i++ {
		s.store[i-1] = s.store[i]
	}
	newFishes := s.store[-1]
	s.store[-1] = 0

	resetFish := s.store[6]
	resetFish += newFishes
	s.store[6] = resetFish

	s.store[8] = newFishes
}

func (s *FishStore) CountFish() uint64 {
	var sum uint64
	for _, v := range s.store {
		sum += v
	}
	return sum
}
