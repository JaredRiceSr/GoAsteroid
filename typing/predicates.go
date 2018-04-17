package typing

func AssignableTo(left, right Type, allowUnknown bool) bool {

	if left == Unknown() && allowUnknown {
		return true
	}

	if rg, ok := right.(*Generic); ok {
		if rg.Accepts(left) {
			return true
		}
	}

	if left.Compare(right) {
		return true
	}
	if left.implements(right) {
		return true
	}
	if left.inherits(right) {
		return true
	}

	if l, ok := ResolveUnderlying(left).(*NumericType); ok {
		if r, ok := ResolveUnderlying(right).(*NumericType); ok {
			if !r.Signed {
				if !l.Signed {
					// uints --> larger uints
					if l.BitSize >= r.BitSize {
						return true
					}
				} else {
					// uints --> larger ints
					if l.BitSize >= r.BitSize+1 {
						return true
					}
				}
			} else {
				if l.Signed {
					if l.BitSize >= r.BitSize {
						return true
					}
				}
			}
		}
	}

	return false
}
