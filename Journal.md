# steps done so far

1. Go here <https://golang.org/doc/install> and install everything
2. Go here <https://github.com/alco/gostart>
3. Go here <https://golang.org/doc/>
4. Go here <https://tour.golang.org/>
5. Go here <https://developers.google.com/sheets/api/quickstart/go> for an API example

---
---

## What I see worthy of taking notes during "A Tour of Go"

1. Go works on a single workspace, defined by the environment variable GOPATH
2. Definition statement inside IF conditions define a variable available only inside the branches
3. Functions may return multiple values (tuples...?)
4. Functions may contain "naked" return if there are named return values, I think they represent a local scoped variable
5. Zero values exist for basic types
6. There is a built-in type for complex numbers
7. Short variable declaration (:=) or explicit declaration (var), you cannot use the first one outside functions
8. It looks like you can shift bits up to 2^511 - 1
9. fmt.Printf contains more than a single flag, look them up
10. UTF-8 support!
11. While can be expressed by using For without semicolons
12. It looks like empty guards are allowed and are equal to "true"
13. -=, += still exist. Can I assume that ++ and -- exist too?
14. Switch-case don't need "break;" as Go puts it for you automatically
15. Switch-cases can be checked against arbitrary expressions, not only constants or variables
16. Switch guard can be empty in order to "emulate" an else-if cascade
17. Defer looks the most interesting for now: function call is executed when the surrounding function returns **while** the arguments are immediately evaluated
18. Defer follows LIFO: that means multiple Defer in the same function are executed in reversed order
19. There are pointers like C, except for the lack of pointer arithmetic
20. There is an automatic dereference when you access a pointed Struct field, in that case s.X == (*s).X
21. Struct constants: you can declare Struct values in-line
22. Array notation is "inverted" on declarations: you write "\[n]Type" and not "Type\[n]"
23. Array literals are of the form: "array := \[n]Type{t0, t1, ..., tn-1}
24. Slices are a view of a piece of array, there is NOT a copy: any side effect on the slice will be propagated to the original array!
25. Slice literals are declared like arrays except for the length: an array will be created and then the slice representing the entire array.
26. Slice and array are different! You cannot assign a slice to an array variable.
27. Range indexing of slices: \["begin":"exclusiveEnd"\], both bounds can be omitted.
28. Slices have length (len(s)) and capacity (cap(s))
29. Slice's zero value is "nil" but it is appendable
30. make(\[]type, length, capacity) is used to create a slice of "type" elements
31. append(s, item1, ..., itemn) returns a new slice with new items appended
32. range "map/slice" can used in For loops to iterate over
33. Map types are identified as "map\[keytype]valuetype"
34. Map's zero value is "nil" and it is unusable
35. make(map\[keytype]valuetype) initializes a usable map
36. Maps' set is as usual as other programming languages
37. delete(map, key) to remove a K-V matching
38. getting a map entry returns two values: value (zero value of type if not found) and a boolean to tell if there was a miss or a hit
39. Zero values on miss can be useful to emulate a map's "Up-sert"
40. Anonymous functions are similar to lambdas and can be passed as values
41. Anonymous functions have the following type schema: "func(argtypes) returntypes"
42. Go does not have classes but you can define something like "extension methods" found in C#, except for the syntax
43. Receiver argument must be written before the function name, something like this: "func (obj Type) MethodName(args) retType"
44. You can "extend" types only on the same package
45. Argument are passed by value, this means that a pointer to a Struct must be passed in order to edit its state
46. (Duality with automatic dereference) - calling methods with pointer receiver automatically forces the receiving Struct as reference: this means v.MethodWithPtrReceiver() == (&v).MethodWithPtrReceiver() if v is a Struct and not a *Struct
47. (In line with automatic dereference) - calling methods with value receiver automatically forces the receiving Struct as value: this means v.MethodWithPtrReceiver() == (*v).MethodWithPtrReceiver() if v is a *Struct and not a Struct

(TODO: continue from 9/26 of "Methods, Interfaces, etc...")