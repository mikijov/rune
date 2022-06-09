# Redesign

program
- name -> type
- name -> value reference

reference
- reference to value
- can be read(get) or written(set)
- is read/written/const

value
- anything that has value
- int, real, bool, string, function, interface, slice

type
- kind
- complete?
- element type (array)
- params (function)
- fields (object)
-

int
real
bool
string
- length
- get char
function
- call
object
- get field
- set field
slice
- length
- capacity
- get element
- set element


parsing
- iteratively gather all top level declarations
- gather all top level object
- iteratively gather all nested/object declarations
- gather all implementations
- gather all initialization code
