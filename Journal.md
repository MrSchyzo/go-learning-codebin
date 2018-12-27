# steps done so far

1. Go here <https://golang.org/doc/install> and install everything
1. Go here <https://github.com/alco/gostart>
1. Go here <https://golang.org/doc/>
1. Go here <https://tour.golang.org/>

---
---

## What I see worthy of taking notes during "A Tour of Go"

1. Go works on a single workspace, defined by the environment variable GOPATH
1. Definition statement inside IF conditions define a variable available only inside the branches
1. Functions may return multiple values (tuples...?)
1. Functions may contain "naked" return if there are named return values, I think they represent a local scoped variable
1. Zero values exist for basic types
1. There is a built-in type for complex numbers
1. Short variable declaration (:=) or explicit declaration (var), you cannot use the first one outside functions
1. It looks like you can shift bits up to 2^511 - 1
1. fmt.Printf contains more than a single flag, look them up
1. UTF-8 support!
1. While can be expressed by using For without semicolons
1. It looks like empty guards are allowed and are equal to "true"
1. -=, += still exist. Can I assume that ++ and -- exist too?
1. Switch-case don't need "break;" as Go puts it for you automatically
1. Switch-cases can be checked against arbitrary expressions, not only constants or variables
1. Switch guard can be empty in order to "emulate" an else-if cascade
1. Defer looks the most interesting for now: function call is executed when the surrounding function returns **while** the arguments are immediately evaluated
1. Defer follows LIFO: that means multiple Defer in the same function are executed in reversed order
1. There are pointers like C, except for the lack of pointer arithmetic
1. There is an automatic dereference when you access a pointed Struct field, in that case s.X == (*s).X
1. Struct constants: you can declare Struct values in-line
1. Array notation is "inverted" on declarations: you write "\[n]Type" and not "Type\[n]"
1. Array literals are of the form: "array := \[n]Type{t0, t1, ..., tn-1}
1. Slices are a view of a piece of array, there is NOT a copy: any side effect on the slice will be propagated to the original array!
1. Slice literals are declared like arrays except for the length: an array will be created and then the slice representing the entire array.
1. Slice and array are different! You cannot assign a slice to an array variable.
1. Range indexing of slices: \["begin":"exclusiveEnd"\], both bounds can be omitted.
1. Slices have length (len(s)) and capacity (cap(s))
1. Slice's zero value is "nil" but it is appendable
1. make(\[]type, length, capacity) is used to create a slice of "type" elements
1. append(s, item1, ..., itemn) returns a new slice with new items appended
1. range "map/slice" can used in For loops to iterate over
1. Map types are identified as "map\[keytype]valuetype"
1. Map's zero value is "nil" and it is unusable
1. make(map\[keytype]valuetype) initializes a usable map
1. Maps' set is as usual as other programming languages
1. delete(map, key) to remove a K-V matching
1. getting a map entry returns two values: value (zero value of type if not found) and a boolean to tell if there was a miss or a hit
1. Zero values on miss can be useful to emulate a map's "Up-sert"
1. Anonymous functions are similar to lambdas and can be passed as values
1. Anonymous functions have the following type schema: "func(argtypes) returntypes"
1. Go does not have classes but you can define something like "extension methods" found in C#, except for the syntax
1. Receiver argument must be written before the function name, something like this: "func (obj Type) MethodName(args) retType"
1. You can "extend" types only on the same package
1. Argument are passed by value, this means that a pointer to a Struct must be passed in order to edit its state
1. (Duality with automatic dereference) - calling methods with pointer receiver automatically forces the receiving Struct as reference: this means v.MethodWithPtrReceiver() == (&v).MethodWithPtrReceiver() if v is a Struct and not a *Struct
1. (In line with automatic dereference) - calling methods with value receiver automatically forces the receiving Struct as value: this means v.MethodWithPtrReceiver() == (*v).MethodWithPtrReceiver() if v is a *Struct and not a Struct

(TODO: continue from 9/26 of "Methods, Interfaces, etc...")
