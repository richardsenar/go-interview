package hashmultisets

import "sort"

// New factory that creates a new Hash Multi Set
func New(values ...interface{}) *HashMultiSet {
	set := HashMultiSet{data: make(map[interface{}]int)}
	set.Add(values...)
	return &set
}

// MultiSetPair a set's key/count pair
type MultiSetPair struct {
	Key   interface{}
	Count int
}

// HashMultiSet a data structure representing a set with counts
type HashMultiSet struct {
	data map[interface{}]int
}

// Merge merge multiple sets
func (s *HashMultiSet) Merge(sets ...*HashMultiSet) {
	for _, set := range sets {
		for _, value := range set.List() {
			s.IncrementBy(value, set.GetCount(value))
		}
	}
}

// Add adds a value to the set
func (s *HashMultiSet) Add(values ...interface{}) {
	for _, value := range values {
		s.IncrementBy(value, 1)
	}
}

// IncrementBy increments a value's count by a number
func (s *HashMultiSet) IncrementBy(value interface{}, count int) {
	existingCount := s.data[value]
	s.data[value] = existingCount + count
}

// List returns a list of the set's values
func (s *HashMultiSet) List() []interface{} {
	values := make([]interface{}, 0)
	for key := range s.data {
		values = append(values, key)
	}
	return values
}

// Contains checks if a value is in the set
func (s *HashMultiSet) Contains(value interface{}) bool {
	_, exists := s.data[value]
	return exists
}

// ContainsAll checks if all values are in the set
func (s *HashMultiSet) ContainsAll(values ...interface{}) bool {
	for _, value := range values {
		_, exists := s.data[value]
		if !exists {
			return false
		}
	}
	return true
}

// ContainsAny checks if any values are in the set
func (s *HashMultiSet) ContainsAny(values ...interface{}) bool {
	for _, value := range values {
		_, exists := s.data[value]
		if exists {
			return true
		}
	}
	return false
}

// GetCount returns count associated with the value
func (s *HashMultiSet) GetCount(value interface{}) int {
	return s.data[value]
}

// Remove removes a value
func (s *HashMultiSet) Remove(values ...interface{}) {
	for _, value := range values {
		delete(s.data, value)
	}
}

// Clear clears the set
func (s *HashMultiSet) Clear() {
	s.data = make(map[interface{}]int)
}

// IsEmpty checks if the set is empty
func (s *HashMultiSet) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the set
func (s *HashMultiSet) Size() int {
	return len(s.data)
}

// GetTopValues returns values ordered in descending order
func (s *HashMultiSet) GetTopValues() []MultiSetPair {
	setPairs := make([]MultiSetPair, 0)
	for key, count := range s.data {
		setPairs = append(setPairs, MultiSetPair{Key: key, Count: count})
	}

	sort.SliceStable(setPairs, func(i, j int) bool {
		return setPairs[i].Count > setPairs[j].Count
	})
	return setPairs
}
