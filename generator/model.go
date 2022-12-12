package generator

type context struct {
}

type command struct {
}

type event struct {
}

type domain struct {
	Name     string
	Contexts []*context
}

var Domain domain

func InitDomain(name string) {
	Domain.Name = name
}

func Context(name string) *context {
	/**/
	return nil
}

func GetContext(name string) *context {
	return nil
}

func (ctx *context) Command(name string) *command {
	/**/
	return nil
}

func (ctx *context) Event(name string) {
	/**/
}

func (cmd *command) Request(name string) {
	/**/
}

func Test() {
	InitDomain("iot")
	account := Context("Account")

	account.Command("CreateAccount").
		Request("")

}
