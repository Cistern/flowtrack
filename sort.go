package flowtrack

import (
	"sort"
)

type flowkeySortable struct {
	keys []flowkey
	m    map[flowkey]uint64
}

func (s flowkeySortable) Len() int {
	return len(s.keys)
}

func (s flowkeySortable) Swap(i, j int) {
	s.keys[i], s.keys[j] = s.keys[j], s.keys[i]
}

func (s flowkeySortable) Less(i, j int) bool {
	return s.m[s.keys[i]] < s.m[s.keys[j]]
}

func sortFlowKeyMap(m map[flowkey]uint64) []flowkey {
	var keys []flowkey
	for key, _ := range m {
		keys = append(keys, key)
	}

	s := flowkeySortable{
		keys: keys,
		m:    m,
	}

	sort.Sort(sort.Reverse(s))

	return s.keys
}

type addrPortKeySortable struct {
	keys []addrPortKey
	m    map[addrPortKey]uint64
}

func (s addrPortKeySortable) Len() int {
	return len(s.keys)
}

func (s addrPortKeySortable) Swap(i, j int) {
	s.keys[i], s.keys[j] = s.keys[j], s.keys[i]
}

func (s addrPortKeySortable) Less(i, j int) bool {
	return s.m[s.keys[i]] < s.m[s.keys[j]]
}

func sortAddrPortKeySortableMap(m map[addrPortKey]uint64) []addrPortKey {
	var keys []addrPortKey
	for key, _ := range m {
		keys = append(keys, key)
	}

	s := addrPortKeySortable{
		keys: keys,
		m:    m,
	}

	sort.Sort(sort.Reverse(s))

	return s.keys
}
