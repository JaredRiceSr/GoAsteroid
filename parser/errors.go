package parser

const (
	errAsteroidInterfacePropertyInvalid   = "Asteroid Comet Parser Error: Everything in an interface must be a func type"
	errAsteroidEnumPropertyInvalid        = "Asteroid Comet Parser Error: Everything in an enum must be an identifier"
	errAsteroidMixedParameterName       = "Asteroid Comet Parser Error: Mixed named and unnamed parameters"
	errAsteroidArraySizeInvalid           = "Asteroid Comet Parser Error: Invalid array size"
	errAsteroidEmptyGroup                 = "Asteroid Comet Parser Error: Group declaration must apply modifiers"
	errAsteroidDanglingExpression         = "Asteroid Comet Parser Error: Expression evaluated but not used"
	errAsteroidConstantHasNoValue       = "Asteroid Comet Parser Error: Constants must have a value"
	errAsteroidUnclosedGroup              = "Asteroid Comet Parser Error: Unclosed group not allowed"
	errAsteroidScopeDeclarationInvalid    = "Asteroid Comet Parser Error: Invalid declaration in scope"
	errAsteroidTypeRequired               = "Asteroid Comet Parser Error: Required %s, found %s"
	errAsteroidAnnotationParameterInvalid = "Asteroid Comet Parser Error: Invalid annotation parameter, must be string"
	errAsteroidIncDecInvalid              = "Asteroid Comet Parser Error: Cannot increment or decrement in this context"
	errAsteroidInvalidTypeAfterCast       = "Asteroid Comet Parser Error: Expected type after cast operator"
	errAsteroidIncompleteExpression       = "Asteroid Comet Parser Error: Incomplete expression"
	errAsteroidInvalidImportPath          = "Asteroid Comet Parser Error: Invalid import path: %s"
	errAsteroidConsecutiveExpression      = "Asteroid Comet Parser Error: No terminator or operator after expression: found %s"
)
