[Interfaces](https://go.dev/tour/methods/9) allow you to focus on what a type does rather than how it's built. They can help you write more flexible and reusable code by defining behaviors (like methods) that different types can share. This makes it easy to swap out or update parts of your code without changing everything else.

Interfaces are just collections of method signatures. A type "implements" an interface if it has methods that match the interface's method signatures.

Interfaces are implemented implicitly.

A type never declares that it implements a given interface. If an interface exists and a type has the proper methods defined, then the type automatically fulfills that interface.

A quick way of checking whether a struct implements an interface is to declare a function that takes an interface as an argument. If the function can take the struct as an argument, then the struct implements the interface.

A type implements an interface by implementing its methods. Unlike in many other languages, there is no explicit declaration of intent, there is no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation. You may add methods to a type and in the process be unknowingly implementing various interfaces, and that's okay.

When working with interfaces in Go, every once-in-awhile you'll need access to the underlying type of an interface value. You can cast an interface to its underlying type using a type assertion.

Interfaces are not classes, they are slimmer.

Interfaces don't have constructors or deconstructors that require that data is created or destroyed.

Interfaces aren't hierarchical by nature, though there is syntactic sugar to create interfaces that happen to be supersets of other interfaces.

Interfaces define function signatures, but not underlying behavior. Making an interface often won't DRY up your code in regards to struct methods. For example, if five types satisfy the fmt.Stringer interface, they all need their own version of the String() function.
