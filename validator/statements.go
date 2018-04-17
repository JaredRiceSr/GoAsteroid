package validator

import (
	"github.com/benchlab/asteroid/typing"

	"github.com/benchlab/asteroid/ast"
)

func (v *Validator) validateStatement(node ast.Node) {
	switch n := node.(type) {
	case *ast.AssignmentStatementNode:
		v.validateAssignment(n)
		break
	case *ast.ForStatementNode:
		v.validateForStatement(n)
		break
	case *ast.IfStatementNode:
		v.validateIfStatement(n)
		break
	case *ast.ReturnStatementNode:
		v.validateReturnStatement(n)
		break
	case *ast.SwitchStatementNode:
		v.validateSwitchStatement(n)
		break
	case *ast.ForEachStatementNode:
		v.validateForEachStatement(n)
		break
	case *ast.ImportStatementNode:
		v.validateImportStatement(n)
		return
	case *ast.PackageStatementNode:
		v.validatePackageStatement(n)
		return
	}
	v.finishedImports = true
}

func (v *Validator) validateAssignment(node *ast.AssignmentStatementNode) {

	for _, l := range node.Left {
		if l == nil {
			v.addError(node.Start(), errAsteroidUnknown)
			return
		} else {
			switch l.Type() {
			case ast.CallExpression, ast.Literal, ast.MapLiteral,
				ast.ArrayLiteral, ast.SliceExpression, ast.FuncLiteral:
				v.addError(l.Start(), errAsteroidNonValidExpressionLeft)
			}
		}
	}

	leftTuple := v.ExpressionTuple(node.Left)
	rightTuple := v.ExpressionTuple(node.Right)
	if len(leftTuple.Types) > len(rightTuple.Types) && len(rightTuple.Types) == 1 {
		right := rightTuple.Types[0]

		for _, left := range leftTuple.Types {
			if !v.bvm.Assignable(v, left, right, node.Right[0]) {
				v.addError(node.Left[0].Start(), errAsteroidAssignmentInvalid, typing.WriteType(left), typing.WriteType(right))
			}
		}

		for i, left := range node.Left {
			if leftTuple.Types[i] == typing.Unknown() {
				if id, ok := left.(*ast.IdentifierNode); ok {
					ty := rightTuple.Types[0]
					id.Resolved = ty
					id.Resolved.SetModifiers(nil)
					ignored := "_"
					if id.Name != ignored {
						v.declareVar(id.Start(), id.Name, id.Resolved)
					}

				}
			}
		}

	} else {
		if len(leftTuple.Types) == len(rightTuple.Types) {

			count := 0
			remaining := 0
			for i, left := range leftTuple.Types {
				right := rightTuple.Types[i]
				if !v.bvm.Assignable(v, left, right, node.Right[count]) {
					v.addError(node.Start(), errAsteroidAssignmentInvalid, typing.WriteType(leftTuple), typing.WriteType(rightTuple))
					break
				}
				if remaining == 0 {
					if node.Right[count] == nil {
						count++
					} else {
						switch a := node.Right[count].ResolvedType().(type) {
						case *typing.Tuple:
							remaining = len(a.Types) - 1
							break
						default:
							count++
						}
					}

				} else {
					remaining--
				}

			}
		} else {
			v.addError(node.Start(), errAsteroidAssignmentInvalid, typing.WriteType(leftTuple), typing.WriteType(rightTuple))
		}

		if len(node.Left) == len(rightTuple.Types) {
			for i, left := range node.Left {
				if leftTuple.Types[i] == typing.Unknown() {
					if id, ok := left.(*ast.IdentifierNode); ok {
						id.Resolved = rightTuple.Types[i]
						if id.Name != "_" {
							v.declareVar(id.Start(), id.Name, id.Resolved)
						}
					}
				}
			}
		}
	}
}

func (v *Validator) validateIfStatement(node *ast.IfStatementNode) {

	v.openScope(nil, nil)

	if node.Init != nil {
		v.validateAssignment(node.Init.(*ast.AssignmentStatementNode))
	}

	for _, cond := range node.Conditions {
		v.requireType(cond.Condition.Start(), typing.Boolean(), v.resolveExpression(cond.Condition))
		v.validateScope(node, cond.Body)
	}

	if node.Else != nil {
		v.validateScope(node, node.Else)
	}

	v.closeScope()
}

func (v *Validator) validateSwitchStatement(node *ast.SwitchStatementNode) {

	switchType := typing.Boolean()

	if node.Target != nil {
		switchType = v.resolveExpression(node.Target)
	}

	for _, node := range node.Cases.Sequence {
		if node.Type() == ast.CaseStatement {
			v.validateCaseStatement(switchType, node.(*ast.CaseStatementNode))
		}
	}

}

func (v *Validator) validateCaseStatement(switchType typing.Type, clause *ast.CaseStatementNode) {
	for _, expr := range clause.Expressions {
		t := v.resolveExpression(expr)
		if !v.bvm.Assignable(v, switchType, t, expr) {
			v.addError(clause.Start(), errAsteroidSwitchTargetInvalid, typing.WriteType(switchType), typing.WriteType(t))
		}

	}
	v.validateScope(clause, clause.Block)
}

func (v *Validator) validateReturnStatement(node *ast.ReturnStatementNode) {
	for c := v.scope; c != nil; c = c.parent {
		if c.context != nil {
			switch a := c.context.(type) {
			case *ast.FuncDeclarationNode:
				results := a.Resolved.(*typing.Func).Results
				returned := v.ExpressionTuple(node.Results)
				if (results == nil || len(results.Types) == 0) && len(returned.Types) > 0 {
					v.addError(node.Start(), errAsteroidReturnInvalidFromVoid, typing.WriteType(returned), a.Signature.Identifier)
					return
				}
				if !typing.AssignableTo(results, returned, false) {
					v.addError(node.Start(), errAsteroidReturnInvalid, typing.WriteType(returned), a.Signature.Identifier, typing.WriteType(results))
				}
				return
			case *ast.FuncLiteralNode:
				results := a.Resolved.(*typing.Func).Results
				returned := v.ExpressionTuple(node.Results)
				if (results == nil || len(results.Types) == 0) && len(returned.Types) > 0 {
					v.addError(node.Start(), errAsteroidReturnInvalidFromVoid, typing.WriteType(returned), "literal")
					return
				}
				if !typing.AssignableTo(results, returned, false) {
					v.addError(node.Start(), errAsteroidReturnInvalid, typing.WriteType(returned), "literal", typing.WriteType(results))
				}
				return
			}
		}
	}
	v.addError(node.Start(), errAsteroidReturnStatementInvalidOutsideFunction)
}

func (v *Validator) validateForEachStatement(node *ast.ForEachStatementNode) {

	v.openScope(nil, nil)

	gen := v.resolveExpression(node.Producer)
	var req int
	switch a := gen.(type) {
	case *typing.Map:
		req = 2
		if len(node.Variables) != req {
			v.addError(node.Begin, errAsteroidForEachVarInvalid, len(node.Variables), req)
		} else {
			v.declareVar(node.Start(), node.Variables[0], a.Key)
			v.declareVar(node.Start(), node.Variables[1], a.Value)
		}
		break
	case *typing.Array:
		req = 2
		if len(node.Variables) != req {
			v.addError(node.Start(), errAsteroidForEachVarInvalid, len(node.Variables), req)
		} else {
			v.declareVar(node.Start(), node.Variables[0], v.LargestNumericType(false))
			v.declareVar(node.Start(), node.Variables[1], a.Value)
		}
		break
	default:
		v.addError(node.Start(), errAsteroidForEachTypeInvalid, typing.WriteType(gen))
	}

	v.validateScope(node, node.Block)

	v.closeScope()

}

func (v *Validator) validateForStatement(node *ast.ForStatementNode) {

	v.openScope(nil, nil)

	if node.Init != nil {
		v.validateAssignment(node.Init)
	}

	v.requireType(node.Cond.Start(), typing.Boolean(), v.resolveExpression(node.Cond))

	if node.Post != nil {
		v.validateStatement(node.Post)
	}

	v.validateScope(node, node.Block)

	v.closeScope()
}

func (v *Validator) createPackageType(path string) *typing.Package {
	scope, errs := ValidatePackage(v.bvm, path)
	if errs != nil {
		v.errs = append(v.errs, errs...)
	}
	pkg := new(typing.Package)
	pkg.Variables = scope.variables
	pkg.Types = scope.types
	return pkg
}

func trimPath(n string) string {
	lastSlash := 0
	for i := 0; i < len(n); i++ {
		if n[i] == '/' {
			lastSlash = i
		}
	}
	return n[lastSlash:]
}

func (v *Validator) validateImportStatement(node *ast.ImportStatementNode) {
	if v.finishedImports {
		v.addError(node.Start(), errAsteroidImportsFinished)
	}
	if node.Alias != "" {
		v.declareType(node.Start(), node.Alias, v.createPackageType(node.Path))
	} else {
		v.declareType(node.Start(), trimPath(node.Path), v.createPackageType(node.Path))
	}
}

func (v *Validator) validatePackageStatement(node *ast.PackageStatementNode) {
	if node.Name == "" {
		v.addError(node.Start(), errAsteroidPackageNameInvalid, node.Name)
		return
	}
	if v.packageName == "" {
		v.packageName = node.Name
	} else {
		if v.packageName != node.Name {
			v.addError(node.Start(), errAsteroidDuplicatePackageName, node.Name, v.packageName)
		}
	}
}
