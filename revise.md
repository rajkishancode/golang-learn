a slice is good for storing many values of a same type
a map is really good for key value lookup
a struct let you aggregate together/ compose together whole bunch of value types of different types.



#packing and unpacking slice of int
ex -->

x := []int{1,2,3,4,5}
s = sum (x...)  // three dots on RIGHT side UNPACKING / spread operator in js

func sum(xi ...int){ // three dots on LEFT side on type PACKING slice of number/ rest operator in js
                    // (...int ) - this is a variadic parameter becomes slice of int
}