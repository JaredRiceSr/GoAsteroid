<p align="center">
  <img src="https://github.com/benchlab/GoAsteroid/raw/master/asteroid-logo.png" width="300px" alt="Asteroid Logo"/>
</p>

# Asteroid 

Note: While this looks like a complete project because of its size, this project is still under development. 

Asteroid is a statically typed, object-oriented programming language for smart contracts on the BenchChain MultiChain. Its syntax primarily derives from [Go](https://golang.org), [Java]() and [Python](), along with many of the blockchain-specific constructs drawn (at least in part) from [Solidity](https://github.com/ethereum/solidity) and [Viper](https://github.com/ethereum/viper). Asteroid works alongside BenchChains's integrate distributed virtual machine layer known as [BenchVM](https://github.com/benchlab/benchvm)

Asteroid is agnostic to the distributed virtual machines it runs on. This means the same code can be used on other distributed virtual machine-based smart contract systems on other blockchains.


| Name | Development Status |
|:-----------:|:----|
| BenchVM | 82% Done |
| EVM | 80% Done |
| NEO BVM | Upcoming |



## Asteroid Roadmap To Production

In no particular order, Asteroid strives to:

- Deterministic execution (85% Done)
- Balance legibility, security and the overall safety of using Asteroid in production. (35% Done)
- Follow other OOP Languages with more features. (75% Done)
- Easy to learn (90% Done)
- Support for ByteCode generation (100% Done)
- Coverage for all Smart Contract platforms (50% Done)
- Built-in Smart Contract Generator for EOS, Ethereum and NEO (10% Done)
- From the design and implementation factors of the Asteroid language to the
website and documents for Astroid, we want it to be exceptional. (30% Done)

## Asteroid Smart Contract Example

Below is an example of a bare-bones smart contract using Asteroid.

```go
contract youBench {

    var name string

    external func executeSwap()() string {
        return name
    }

    external func setName(n string) {
        this.name = n
    }

}
```


## Asteroid Packaging and Version Declarations

What makes Asteroid more powerful than other smart contract platforms is the `package import` feature, very similar to the `package import` feature in GoLang applications. This means, libraries and smart contract features can be developed throughout the community and used in future development projects as the Bench and Asteroid communities grow.

Packages are grouped into logical modals and should only contain a single ```contract```, although not every contract requires a ```contract```.

In order for future versions of Asteroid to include potentially backwards-incompatible changes, each Asteroid file must include a version declaration appended to the package declaration:

```
package orderbook asteroid 0.0.1
```

## Importing packages with Asteroid

Asteroid packages may be imported using the following syntax:

```go

import "coins"

import (
    "a"
    "b"
    "c"
)

import (
    "d" as Dogecoin
    "e" as EOS
    "f" as ForkCoin
)

contract PriceWatcher {

    doThing(){
        // this is a function from the coin package
        coin.watch()
        // this is a function from the 'a' package
        a.sayHi()
        // this is an event from the 'e' package
        EOS.LogEvent()
    }
}

```

##  Typing With Asteroid

Asteroid is strongly and statically typed. In order to compile, code must stick to these standards or it will fail. Although you can defer the compiler from certain types, if need be.

```go
x = 5      // x will have type int
y = x      // y will have type int
x = "hello" // will not compile (x has type int)
y = 5.5 //  will not compile (y has type int)
```

Common types are declared as follows:
```go
var a int
// type inside a package
var b pkg.OrderBook
var c map[string]int
var d []string
var e func(string) string
```

###  Inheritance With Asteroid

Since Asteroid allows for `multiple inheritence`, the following class is therefore valid.

```go
class AtomicSwapMembers inherits UserA, UserB {

}
```

When a two methods are inherited by a `class`, where the names are identical as well as the method's parameters, the methods will `cancel` automatically and will be excluded from the `subclass`.

###  Interfaces With Asteroid

Asteroid uses java-style interface syntax.

```go
interface DEX inherits Exchange {
    trade(amount int)
}

class AtomicSwapMembers inherits UserA, UserB is SwapReady {

}
```

All types which explicitly implement an interface through the ```is``` keyword may be referenced as that interface type. However, there is no go-like implicit implementation, as it can be confusing and serves no particular purpose in a class-based language.

## Asteroid Key Features

### Constructors and Destructors With Asteroid

Asteroid uses ```constructor``` and ```destructor``` `keywords`. Each `class` may contain an arbitrary number of `constructors` and `destructors`. In order to have an arbitrary number of both, you will need a valid signature.

```go
contract Test {

    constructor(name string){

    }

    destructor(){

    }

}
```

By default, the `no-args` `constructor` and `destructor` will be called.

###  Generics With Asteroid

Generics can be specified using Java-like syntax:

```go

// Purchases can only be related to things which are Sellable
// this will be checked at compile time

contract Purchase<T is Sellable> {

    var item T
    var quantity int

    constructor(item T, quantity int){
        this.item = item
        this.quantity = quantity
    }
}
```

To declare several generics at once, use the ```|``` character:

```
class OrderBook <T|S|R> {

}
```

Generics can also be restricted:

```
class OrderBook<T inherits UserB, UserA | S is assetType> {

}
```




###  Iteration With Asteroid

A  `randomized map iteration` is provided by many languages as the only solution. This is the case with GoLang. Clearly, this is not deterministic, as demonstrated by the following example:


```go
myMap["hi"] = 2
myMap["bye"] = 3
count = 0
sum = 0
for k, v in myMap {
    sum += v * count
    count++
}
```

Because of this, the value of sum will be completely different because of the  `iteration order` of the associated `map elements`. In terms of integrating with a blockchain, each node on the network could reach a different consensus about the state of your smart contract. Asteroid `maps` completely solves this issue through `Guaranteed Order Of Insertion` when working with `maps`,

If the above example was built with Asteroid, the sum would always equal 3.

###  Modifiers With Asteroid

Solidity by default utilizes ```Access Modifiers``` to direct access to methods. These access modifiers and functions (if a condition must be duplicated over several methods), should be replaced globally with ```require``` statements. Asteroid is designed to use ```require``` statements instead of ```Access Modifiers```

Solidity Example Using Access Modifiers:

```go
modifier local(Location _loc) {
    require(_loc == loc);
    _;
}

chat(Location loc, string msg) local {

}
```

Asteroid Example Using Require Statements:

```go
enforceLocal(loc Location){
    require(location == loc)
}

chat(loc Location, msg string){
    enforceLocal(loc)
}
```

###  Contract Groups With Asteroid
If you're familiar with GoLang, you're familiar with the ability to create  ```Keyword Groups``` which allows most developers to simplify Variable Declaration or Types where properties are nearly the same. Like Go, Asteroid also lets you group ```Logically-Related Declarations``` together to simplify the process of developing a smart contract.

Declarations can be in the top-level group or below the top-level group:

```go
class (
    OrderBook {

    }

    assetType {

    }
)
```

```go
public (
    a string
    b int
    c string
)
```

In this Asteroid example, we apply groups to the top level of the declaration so that nested groups are possible:

```go
internal (

    func (

        add(a, b int) int {
            return a + b
        }

        sub(a, b int) int {
            return a - b
        }

    )
)
```

## Annotations With Asteroid

Asteroid `annotations` are utilized around `declarations` to detail different properties about the `declaration`. Asteroid annotations can not be converted to `Classes` but are included in the program and implementation of your smart contract application.

As for `native annotations`, there is only one. ```@Builtin(string)```, maps `builtin properties` onto `bytecode-generating` `functions`. To create an `annotation`, all that is required is the following:

```go
func (bvm BVM) Annotations() []Annotation {
    return []Annotation {
        ParseAnnotation("@Builtin(string)", handleBuiltin),
        ParseAnnotation("@CryptoKing(int, func(int, int))", handleCryptoKing)
    }
}

func handleBuiltin(i {}interface, a []*Annotation) {

}

func handleCryptoKing(i {}interface, a []*Annotation) {

}

```

Currently, annotations can only accept string parameters, but this is a flexible limitation and is open to revision in future versions.

## Comet Validator

BVM implementations must conform to the following interface:

```go
type BVM interface {
	Traverse(ast.Node) (bvmCreate.Bytecode, util.Errors)
	Builtins() *ast.ScopeNode
	BaseContract() (*ast.ContractDeclarationNode, util.Errors)
	Primitives() map[string]typing.Type
	Literals() LiteralMap
	BooleanName() string
	ValidExpressions() []ast.NodeType
	ValidStatements() []ast.NodeType
	ValidDeclarations() []ast.NodeType
	Modifiers() []*ModifierGroup
	Annotations() []*typing.Annotation
	BytecodeGenerators() map[string]BytecodeGenerator
	Castable(val *Validator, to, from typing.Type, fromExpression ast.ExpressionNode) bool
	Assignable(val *Validator, to, from typing.Type, fromExpression ast.ExpressionNode) bool
}
```

Where ```LiteralMap``` is an alias for ```map[token.Type]func(*Validator, string) Type``` and ```OperatorMap``` is an alias for ```map[token.Type]func(*Validator, ...Type) Type```.

This interface is how language-specific features are implemented on top of the Asteroid core systems.

## Operators With Asteroid

We know some developers who utilize the Asteroid language will want to choose specific operators that are defined for the specific language they're using. If this turns out to be the case, all operators must be in this form:

```
func operator(v *validator.Validator, ...validator.Type) validator.Type {

}
```

Where the returned ```Type``` is the type produced by the operator in the given context (the operand types provided).

To make this simple, Asteroid provides a few helper functions:

```go
// does a simple type lookup using 'name'
func SimpleOperator(name string) OperatorFunc
// returns the smallest available numeric type
func NumericalOperator() OperatorFunc
// returns the smallest available integer type
func IntegerOperator() OperatorFunc
// returns the smallest available fixed point type
func DecimalOperator() OperatorFunc
```

## Literals With Asteroid

Asteroid has a custom translator built-in to the language compiler to convert `lexer tokens` to `Types`. The hangup is only being able to use the tokens defined in the `lexer` configuration but each `lexer token` can be directly mapped to `contextual type` producing `functions`.  Although this also depends on the string being used.

As an example of `Literal` forms, all `Literals` within Asteroid must look like this:

```go
func literal(v *validator.Validator, data string) validator.Type {

}
```

## Primitives With Asteroid

`Primitive types` are the foundational building block of Asteroid-based smart contract.

```go
var type = &typing.NumericType{Size: 16, Signed: true}
```

## CURRENT BUGS AND ONGOING DEVELOPMENT:

- parser's lookahead needs to be reduced.
- construct parsing is easy but are trying to make it better and faster.
- working on compilation speed

Known bugs:
- Expressions are causing crashes due to unknown operators
