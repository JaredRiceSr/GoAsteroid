# Comet Contract Verifier

Comet is a formal verifier for the Asteroid programming language.

## Invariants With Comet and Asteroid 

Asteroid enforces invariant conditions. These conditions are implemented as asteroid functions.

```package asteroid 0.0.1
contract Basic {

    internal name string

    external (
        name string
        day int
    )

    external (

        @Comet()
        func setName(n string){
            enforce(n != "Alex")
            name = n
        }

        func executeSwap()() string {
            return name
        }

    )

}
```
