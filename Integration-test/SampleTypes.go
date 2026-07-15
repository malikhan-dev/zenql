package Integration_test

type ForeignAddress struct {
	Country string
}
type Address struct {
	Street string
	City   string
	State  string
	Zip    string
	No     int
	Id     int
	Flag   bool
}
type User struct {
	Name     string
	Age      int
	Id       int
	Addr     []Address
	ParentId int
}

type ForeignUser struct {
	Name     string
	Age      int
	Id       int
	Addr     Address
	ParentId int
}
type city struct {
	Name   string
	Id     int
	Active bool
}

type Student struct {
	Name     string
	Age      int
	Id       int
	Pressent bool
}
type Person struct {
	Name       string
	LastName   string
	Identifier int
	Mail       string
	Active     bool
}
type InternalEmp struct {
	FullName string
	Dep      string
}
type Employee struct {
	Name       string
	Department string
	Age        int
}
type ComplexObjectToSearch struct {
	Name string
	Age  int
	Id   int
	Flag bool
}

type Addr struct {
	City string
}
type Persons struct {
	Name       string
	LastName   string
	Identifier int
	Mail       string
	Active     bool
	Address    []Addr
}

type SysUser struct {
	FName   string
	LName   string
	Id      int
	Email   string
	Enabled bool
	Address string
}
