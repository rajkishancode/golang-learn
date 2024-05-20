# GO is OBJECT ORIENTED
1. Encapsulation
    a. state ("fields")
    b. behavior ("methods")
    c. exported & unexported; viewable & not viewable
2.  Reusability
    a. inheritance
3. Polymorphism
    a.interfaces
4. Overriding
    a. "promotion"


# Traditional OOP
1. Classes
    a. data structure describing a type of object
    b. you can then create "instances/objects" from        class/  bluprint
    c.classes hold both:
        i.state/data/fields
        ii.behavior/methods
    d. public/private

2. Inheritance

# In Go:
    1.you don't create classes , you create a TYPE
    2.you don't instantiate , you create a VALUE of a TYPE
# User defined types
   1. We can declare a new type
   2. foo
       - the underlying type of foo is int
       - int conversion
            -int(myAge)
            -converting type foo to type int

# IT IS A BAD PRACTISE TO ALIAS TYPES
    - one exception
        -if you need to attach methods to a type

# Numeric types
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
            