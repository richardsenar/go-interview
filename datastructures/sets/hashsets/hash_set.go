package hashsets

// New factory that creates a hash set
func New(values ...interface{}) *HashSet {
	set := HashSet{data: make(map[interface{}]struct{})}
	set.Add(values...)
	return &set
}

// HashSet datastructure
type HashSet struct {
	data map[interface{}]struct{}
}

// Add adds values to the set
func (s *HashSet) Add(values ...interface{}) {
	for _, value := range values {
		s.data[value] = struct{}{}
	}
}

// Remove removes values from the set
func (s *HashSet) Remove(values ...interface{}) {
	for _, value := range values {
		delete(s.data, value)
	}
}

// Contains checks if the value is in the set
func (s *HashSet) Contains(value interface{}) bool {
	_, exists := s.data[value]
	return exists
}

// ContainsAll checks if all values are in the set
func (s *HashSet) ContainsAll(values ...interface{}) bool {
	for _, value := range values {
		_, exists := s.data[value]
		if !exists {
			return false
		}
	}
	return true
}

// ContainsAny checks if any of the values are in the set
func (s *HashSet) ContainsAny(values ...interface{}) bool {
	for _, value := range values {
		_, exists := s.data[value]
		if exists {
			return true
		}
	}
	return false
}

// Merge the two sets
func (s *HashSet) Merge(sets ...*HashSet) {
	for _, set := range sets {
		for _, value := range set.GetValues() {
			s.Add(value)
		}
	}
}

// Clear clears set
func (s *HashSet) Clear() {
	s.data = make(map[interface{}]struct{})
}

// GetValues returns values
func (s *HashSet) GetValues() []interface{} {
	values := make([]interface{}, 0)
	for key := range s.data {
		values = append(values, key)
	}
	return values
}

// IsEmpty checks if the set is empty
func (s *HashSet) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the set
func (s *HashSet) Size() int {
	return len(s.data)
}
