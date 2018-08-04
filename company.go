package hellomock

type Company struct{
	User Talker
}

func NewCompany(t Talker)*Company{
	return &Company{
		User:t,
	}
}

func (c *Company) Meeting(name string) string{
	return c.User.SayHello(name)
}
