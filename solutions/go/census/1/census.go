// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
        Name: name,
        Age: age,
        Address: address,
    }
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if r == nil {
		return false
	}

	// Name required
	if r.Name == "" {
		return false
	}

	// Address must not be nil or empty
	if r.Address == nil || len(r.Address) == 0 {
		return false
	}

	// Each key/value in address must be non-empty
	for k, v := range r.Address {
		if k == "" || v == "" || k == "unknown key" {
			return false
		}
	}
	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Age = 0
    r.Name = ""
    r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	count := 0
    for _, res := range residents {
        if res.HasRequiredInfo() {
            count++
        }
    }
    return count
}
