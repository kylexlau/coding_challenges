package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	u, e := units[unit]
	if !e {
		return false
	}
	i, e := bill[item]
	if e {
		bill[item] = i + u
		return true
	} else {
		bill[item] = u
		return true
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	i, e := bill[item]
	if !e {
		return false
	}
	u, e := units[unit]
	if !e {
		return false
	}

	new_i := i - u
	if new_i < 0 {
		return false
	} else if new_i == 0 {
		delete(bill, item)
		return true
	} else {
		bill[item] = new_i
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	i, e := bill[item]
	if !e {
		return 0, false
	} else {
		return i, true
	}
}
