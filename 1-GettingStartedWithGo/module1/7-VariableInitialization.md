# Variable Initialization

## Type Declarations
- defining an alias (alternate name) for a type
- may improve clarity
  - `type Celsius float64`
  - `type IDnum int`
- can declare variables using the type alias instead
  - var temp Celsius
  - var pid IDnum

## Initializing Variables
- initialize in the declaration
  - `var x int = 100`
  - `var x = 100`
- initialize after the declaration
  - `var x int`
  - `x = 100`
- uninitialized variables have a zero value
  - `var x int // x = 0`
  - `var x string // x = ""`
- short variable declarations
  - can perform a declaration and initialization together with the `:=` operator
    - `x:=100`
  - variable is declared as the type of expression on the right hand side
  - can only do this inside a function
