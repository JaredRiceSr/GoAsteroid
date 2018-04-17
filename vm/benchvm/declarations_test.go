package benchvm

import (
	"testing"

	"github.com/benchlab/asteroid"
)

func TestBytecodeContractDeclaration(t *testing.T) {
	a := new(Arsonist)
	asteroid.CompileString(a,
		`contract Tester {
            var x = 5
            const y = 10
        }`)
	checkMnemonics(t, a.BVM.Instructions, []string{
		"PUSH", // push string data
		"PUSH", // push hash(x)
		"PUSH", // push offset (0)
		"SET",  // store result in memory at hash(x)[0]
	})
}

func TestBytecodeFuncDeclaration(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a,
		`contract Tester {
            add(a, b int) int {
                return a + b
            }
        }`)
	checkMnemonics(t, a.BVM.Instructions, []string{
		"PUSH",   // push string data
		"PUSH",   // push hash(x)
		"ADD",    // push offset (0)
		"RETURN", // store result in memory at hash(x)[0]
	})
}

func TestBytecodeInterfaceDeclaration(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a,
		`contract Tester {
			interface Animalistic {

			}
		}`)
	checkMnemonics(t, a.BVM.Instructions, []string{})
}

func TestBytecodeClassDeclaration(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a,
		`contract Tester {
			class Animal {

			}
		}`)
	checkMnemonics(t, a.BVM.Instructions, []string{})
}

func TestBytecodeClassDeclarationWithFields(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a,
		`contract Tester {
			class Animal {
				name string
				genus string
			}
		}`)
	checkMnemonics(t, a.BVM.Instructions, []string{})
}

func TestBytecodeClassDeclarationWithMethods(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a,
		`contract Tester {
			class Animal {
				name string
				genus string

				getName() string {
					return name
				}
			}
		}`)
	checkMnemonics(t, a.BVM.Instructions, []string{})
}
