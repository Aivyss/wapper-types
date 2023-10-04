package types

type Set[K comparable] Map[K, Bool]

func (s *Set[K]) Raw() Map[K, Bool] {
	return Map[K, Bool](*s)
}

func (s *Set[K]) Set(key K) {
	s.Raw()[key] = true
}

func (s *Set[K]) Contains(key K) Bool {
	raw := s.Raw()
	value, flag := raw.GetSafe(key)

	return flag.And(value)
}
