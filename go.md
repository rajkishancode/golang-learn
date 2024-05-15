#GO is OBJECT ORIENTED
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
            