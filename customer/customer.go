package customer

type Customer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

// CreateCustomer creates a new customer
func CreateCustomer(ID, Name, Role, Email, Phone string, Contacted bool) (result *Customer) {
	customer := Customer{
		ID:        ID,
		Name:      Name,
		Role:      Role,
		Email:     Email,
		Phone:     Phone,
		Contacted: Contacted,
	}
	result = &customer
	return result
}

// ModifyCustomer modifies an existing customer
func (c *Customer) ModifyCustomer(ID, Name, Role, Email, Phone string, Contacted bool) {
	c.ID = ID
	c.Name = Name
	c.Role = Role
	c.Email = Email
	c.Phone = Phone
	c.Contacted = Contacted
}
