package lists

// List 인터페이스는 순서가 있는 컬렉션의 기본 연산을 정의합니다.
type List[T any] interface {
	// Add는 리스트의 끝에 요소를 추가합니다.
	Add(element T)

	// Insert는 지정된 인덱스에 요소를 삽입합니다.
	Insert(index int, element T) error

	// Get은 지정된 인덱스의 요소를 반환합니다.
	Get(index int) (T, error)

	// Set은 지정된 인덱스의 요소를 새 값으로 설정합니다.
	Set(index int, element T) error

	// Remove는 지정된 인덱스의 요소를 제거합니다.
	Remove(index int) (T, error)

	// RemoveElement는 첫 번째로 일치하는 요소를 제거합니다.
	RemoveElement(element T) bool

	// IndexOf는 요소의 첫 번째 인덱스를 반환합니다.
	IndexOf(element T) int

	// LastIndexOf는 요소의 마지막 인덱스를 반환합니다.
	LastIndexOf(element T) int

	// Contains는 요소가 리스트에 포함되어 있는지 확인합니다.
	Contains(element T) bool

	// Size는 리스트의 크기를 반환합니다.
	Size() int

	// IsEmpty는 리스트가 비어있는지 확인합니다.
	IsEmpty() bool

	// Clear는 리스트의 모든 요소를 제거합니다.
	Clear()

	// ToSlice는 리스트의 모든 요소를 슬라이스로 반환합니다.
	ToSlice() []T

	// ForEach는 리스트의 모든 요소에 대해 함수를 실행합니다.
	ForEach(fn func(element T))
}

// Deque 인터페이스는 양방향 큐의 연산을 정의합니다.
type Deque[T any] interface {
	List[T]

	// AddFirst는 데크의 앞쪽에 요소를 추가합니다.
	AddFirst(element T)

	// AddLast는 데크의 뒤쪽에 요소를 추가합니다.
	AddLast(element T)

	// RemoveFirst는 데크의 첫 번째 요소를 제거하고 반환합니다.
	RemoveFirst() (T, error)

	// RemoveLast는 데크의 마지막 요소를 제거하고 반환합니다.
	RemoveLast() (T, error)

	// PeekFirst는 데크의 첫 번째 요소를 제거하지 않고 반환합니다.
	PeekFirst() (T, error)

	// PeekLast는 데크의 마지막 요소를 제거하지 않고 반환합니다.
	PeekLast() (T, error)
}

// Stack 인터페이스는 스택(LIFO)의 연산을 정의합니다.
type Stack[T any] interface {
	// Push는 스택의 맨 위에 요소를 추가합니다.
	Push(element T)

	// Pop은 스택의 맨 위 요소를 제거하고 반환합니다.
	Pop() (T, error)

	// Peek는 스택의 맨 위 요소를 제거하지 않고 반환합니다.
	Peek() (T, error)

	// Size는 스택의 크기를 반환합니다.
	Size() int

	// IsEmpty는 스택이 비어있는지 확인합니다.
	IsEmpty() bool

	// Clear는 스택의 모든 요소를 제거합니다.
	Clear()

	// ToSlice는 스택의 모든 요소를 슬라이스로 반환합니다 (맨 위부터).
	ToSlice() []T
}

// Queue 인터페이스는 큐(FIFO)의 연산을 정의합니다.
type Queue[T any] interface {
	// Enqueue는 큐의 뒤쪽에 요소를 추가합니다.
	Enqueue(element T)

	// Dequeue는 큐의 앞쪽 요소를 제거하고 반환합니다.
	Dequeue() (T, error)

	// Peek는 큐의 앞쪽 요소를 제거하지 않고 반환합니다.
	Peek() (T, error)

	// Size는 큐의 크기를 반환합니다.
	Size() int

	// IsEmpty는 큐가 비어있는지 확인합니다.
	IsEmpty() bool

	// Clear는 큐의 모든 요소를 제거합니다.
	Clear()

	// ToSlice는 큐의 모든 요소를 슬라이스로 반환합니다 (앞쪽부터).
	ToSlice() []T
}
