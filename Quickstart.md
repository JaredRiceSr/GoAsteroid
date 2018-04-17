# Get Started With Asteroid
This is a simple quickstart guide to getting started with Asteroid's smart contract programming syntax and launch your first contract.

## Build Your First Asteroid Contract

Let's start by defining a simple contract:

```go
contract Exchange {

    buyOrderAssetID string

    external assetId(name string){
        assetID = name
    }

}
```

This contract creates a public (```exported```) function interface ```sayHi```, which anyone can call with one provided ```string``` parameter. Whoever last called this function will have their ```name``` stored as the ```lastPersonGreeted```.
