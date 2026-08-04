package LLD

import "fmt"

func main() {

	s := []int{1, 2, 3}

	t := s
	t[0] = 100

	fmt.Println(s, t)

	// both pointing to same array

	// deep copy vs shallow copy

	// nil slice vs empty slice
	var s1 []int
	// len,cap = 0 and nil = true

	var s2 = make([]int, 2)
	// len, cap = 0 but nil = false, empty slice 

}


func Append(s []int, x int) []int {

	fmt.Println(cap(s), len(s))
	t := make([]int, len(s)+1)
	copy(t,s)
	t[len(s)] = x
	return t
}


with make command - 
Map Header

↓

Buckets allocated

↓

Ready for insertions


var m sync.Map
	m.Store(1,2)
	if val, ok := m.Load(1); ok {
		fmt.Print(val)
	}
	m.Delete(1)
	// mutex lock and unlock not required as it is in built for sync.Map


1. Why are concurrent map writes unsafe but concurrent reads are generally safe?
Concurrent map writes are unsafe because maps in Go are not designed to handle simultaneous write operations 
from multiple goroutines. If two or more goroutines attempt to write to the same map at the same time, 
it can lead to race conditions, data corruption, or

2. Why can't you take the address of a map element?
You cannot take the address of a map element because map elements are not addressable. 
This is because maps in Go are implemented as hash tables, and the memory layout does not allow 
for direct addressing of individual elements.

3. Why is map iteration intentionally randomized?
Map iteration is intentionally randomized to prevent developers from writing code that depends on the order of iteration, 
which could lead to non-deterministic behavior.

4. When would you choose sync.Map over map + RWMutex?
You would choose sync.Map over map + RWMutex when you have a high-concurrency scenario 
with many read operations and fewer write operations, as sync.Map is optimized for such use cases.

5. Why are slices invalid as map keys?
Slices are invalid as map keys because they are not comparable in Go. 
A slice's elements can be modified after creation, making it impossible to determine if two slices are equal.

6. What does make(map, n) actually guarantee?
make(map, n) guarantees that the map will be initialized with a capacity of at least n buckets, 
which can help improve performance by reducing the number of hash collisions and resizing operations.


******************************************************

package main

import "fmt"

type User struct{}

func (*User) Hello() {
	fmt.Println("Hello")
}

func main() {
	var u User
	u.Hello()
}
| Receiver         | Method belongs to      |
| ---------------- | ---------------------- |
| `func (u User)`  | `User` **and** `*User` |
| `func (u *User)` | only `*User`           |

*************************************************


type Speaker interface {
	Speak()
}

type Dog struct{}

func (*Dog) Speak(){}

func main(){

    d:=Dog{}

    var s Speaker=d

}

s = d will fail bcz dog does not import Speak(), *Dog does
so s = &d wil work 


package main

import "fmt"

type Dog struct{}

func main() {

	var d *Dog = nil

	var i interface{} = d

	fmt.Println(i == nil)

}

i is not fully nil bcz its value is nil but type is Dog
interface is only nil when both type and value is nil 

Interface stores - Type+Method Table

func B(){

    x:=10

    var i interface{}=x

    fmt.Println(i)

}

it will go to excape 

Dependecny injection - 

// 1. Define the behavior your component needs
type Database interface {
    GetUserName(id int) (string, error)
}

// 2. Accept the interface in your struct
type UserStore struct {
    db Database 
}

// 3. Inject the dependency via the constructor
func NewUserStore(db Database) *UserStore {
    return &UserStore{db: db}
}


===================


defer works in LIFO
defer1() defer2() defer3() then it will execute 
defer3() then defer2() defer1()

func main() {
	x := 10

	defer func() {
		fmt.Println(x)
	}()

	x = 20
}

o/p - 20

func main() {
	x := 10

	defer fmt.Println(x)

	x = 20
}

o/p - 10



func test() int {

	x := 10

	defer func() {
		x = 20
	}()

	return x
}

o/p- 10

func test() (x int) {

	x = 10

	defer func() {
		x = 20
	}()

	return
}
o/p - 20



func main() {

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}

o/p - 2,1,0

func main() {

	for i := 0; i < 3; i++ {

		defer func() {
			fmt.Println(i)
		}()
	}
}

o/p - 3,3,3


=============================


return fmt.Errorf("DB failed: %w", err)
%w wraps the original error.

Sentinel
var ErrNotFound = errors.New("not found")

Typed
type ValidationError struct {
	Field string
}

difference between errors.Is and errors.As:
errors.Is checks if an error is equal to a specific error value, while errors.As checks if an error can be cast to a specific error type.	
errors.As is useful when you want to extract additional information from an error, while errors.Is is useful for checking if an error matches a specific sentinel error.



==================

what is GMP?
G = Goroutine
M = Machine (OS Thread)
P = Processor (Logical Processor)

              CPU Core 1

        +----------------+
        |       M1       |
        +----------------+
                |
                |
              owns
                |
        +----------------+
        |       P1       |
        +----------------+
           |    |    |
           G1   G2   G3


              CPU Core 2

        +----------------+
        |       M2       |
        +----------------+
                |
        +----------------+
        |       P2       |
        +----------------+
          G4 G5 G6


number of P controlled by runtime.GOMAXPROCS()
Every P has - cache, load balancer, queue 

Work Stealing 
suppose P1 has g1,g2,g3,g4 and P2 has 0, then for balancing g2 will steal few g from P1

Goroutine Lifecycle
new -> runnable -> running -> waiting/blocked -> dead

blocked go routine - does not consume CPU time but in consumes memory, stack, heap, etc.

variable x will go to stack or heap based on escape analysis, if it escapes then it will go to heap otherwise stack
for example, if a variable is returned from a function or referenced by a goroutine, it will escape to the heap. 
If it is only used within the function and does not escape, it will be allocated on the stack.	
from heap it has to be garbage collected, from stack it will be automatically cleaned up when the function returns.

why stack preferred over heap?
1. Stack allocation is faster than heap allocation because it involves simple pointer arithmetic, while heap allocation requires more complex memory management.
2. Stack memory is automatically managed, while heap memory requires manual management and garbage collection.
3. Stack memory has better cache locality, which can improve performance.

what triggers GC?
Heap Growth: When the heap grows beyond a certain threshold, the garbage collector is triggered to reclaim memory.
GOGC - Environment variable that controls the garbage collection target percentage.

why native maps are unsafe for concurrent writes?
native maps in Go are not safe for concurrent writes because they are not designed to handle simultaneous write operations from multiple goroutines. If two or more goroutines attempt to write to the same map at the same time, it can lead to race conditions,
data corruption, or program crashes. The internal data structures of maps are not protected by locks, so concurrent writes can interfere with each other and cause unexpected behavior.

m map[int]int and sync.Map are different in that m is a native Go map, which is not safe for concurrent writes, 
while sync.Map is a concurrent map implementation provided by the Go standard library that is designed to be safe for concurrent access. 
sync.Map uses internal locking mechanisms to ensure that multiple goroutines can safely read from and write to the map 
without causing race conditions or data corruption.

go routines shared memory but thread s have their own memory space, so they do not share memory by default.

commands - 
1. race condition - go race -race
2. deadlock - go run -race
3. memory leak, CPU profiling, blocking profiling, heap profiling, goroutine leakage, mutex profiling, channel profiling - go tool pprof
4. escape analysis - go build -gcflags="-m"





