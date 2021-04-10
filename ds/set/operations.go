package set

// Union returns a new set with elements from both the given sets
func (s1 *ItemSet) Union(s2 *ItemSet) *ItemSet {
	s3 := New()
	s1.lock.RLock()
	for i := range s1.items {
		s3.items[i] = true
	}
	s1.lock.RUnlock()
	s2.lock.RLock()
	for i := range s2.items {
		_, ok := s3.items[i]
		if !ok {
			s3.items[i] = true
		}
	}
	s2.lock.RUnlock()
	return s3
}

// Intersection returns a new set with elements that exist in both sets
func (s1 *ItemSet) Intersection(s2 *ItemSet) *ItemSet {
	s3 := New()
	s1.lock.RLock()
	s2.lock.RLock()
	defer s1.lock.RUnlock()
	defer s2.lock.RUnlock()
	for i := range s1.items {
		_, ok := s2.items[i]
		if ok {
			s3.items[i] = true
		}
	}
	return s3
}

// Difference returns a new set with all the elements that exist
// in the first set and don't exist in the second set
func (s1 *ItemSet) Difference(s2 *ItemSet) *ItemSet {
	s3 := New()
	s1.lock.RLock()
	s2.lock.RLock()
	defer s1.lock.RUnlock()
	defer s2.lock.RUnlock()
	for i := range s1.items {
		_, ok := s2.items[i]
		if !ok {
			s3.items[i] = true
		}
	}
	return s3
}

// Subset returns true if the first set is a subset of the second
func (s1 *ItemSet) Subset(s2 *ItemSet) bool {
	s1.lock.RLock()
	s2.lock.RLock()
	defer s1.lock.RUnlock()
	defer s2.lock.RUnlock()
	for i := range s1.items {
		_, ok := s2.items[i]
		if !ok {
			return false
		}
	}
	return true
}
