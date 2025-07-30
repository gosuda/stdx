package maps

// Map 인터페이스는 키-값 쌍을 저장하는 자료구조의 기본 연산을 정의합니다.
type Map[K comparable, V any] interface {
	// Put은 키-값 쌍을 맵에 저장합니다. 기존 키가 있으면 값을 업데이트하고 이전 값을 반환합니다.
	Put(key K, value V) (previousValue V, exists bool)

	// Get은 키에 해당하는 값을 반환합니다.
	Get(key K) (value V, exists bool)

	// Remove는 키에 해당하는 항목을 제거합니다.
	Remove(key K) (value V, exists bool)

	// ContainsKey는 키가 맵에 존재하는지 확인합니다.
	ContainsKey(key K) bool

	// ContainsValue는 값이 맵에 존재하는지 확인합니다.
	ContainsValue(value V) bool

	// Size는 맵의 크기를 반환합니다.
	Size() int

	// IsEmpty는 맵이 비어있는지 확인합니다.
	IsEmpty() bool

	// Clear는 맵의 모든 항목을 제거합니다.
	Clear()

	// Keys는 맵의 모든 키를 슬라이스로 반환합니다.
	Keys() []K

	// Values는 맵의 모든 값을 슬라이스로 반환합니다.
	Values() []V

	// Entries는 맵의 모든 키-값 쌍을 반환합니다.
	Entries() []Entry[K, V]

	// ForEach는 맵의 모든 키-값 쌍에 대해 함수를 실행합니다.
	ForEach(fn func(key K, value V))
}

// Entry는 키-값 쌍을 나타냅니다.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}
