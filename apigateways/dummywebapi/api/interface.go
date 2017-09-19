package api

type User struct {
	Username  string
	FirstName string
	LastName  string
	Email     string
	Role      string
}

type Banner struct {
	Name string
	URL  string
	Text string
}

type Menu struct {
	Name    string
	Icon    string
	Active  bool
	Submenu []Menu
}

type Header struct {
	Banner []Banner
	Menu   []Menu
	User   *User
}

type Quote struct {
	Text    string
	Teacher string
}

type Teacher struct {
	Name string
	Id   string
	Type string
}

type Person struct {
	Name string
	Id   string
}

type Team struct {
	Name   string
	Id     string
	Member []Person
}

type Category struct {
	Name string
	Id   string
}

type Transcript struct {
	Title    string
	Date     string
	Teacher  []Teacher
	Receiver Person
	Attendee []Person
	Category []Category
	Text     string
}

type Filter struct {
	Receiver  []Person
	DateRange string
	Category  []Category
	Teacher   []Teacher
	Team      []Team
}

type Search struct {
	Keywords string
	Filter   Filter
	Results  []Transcript
}

type Article struct {
	Title string
	Body  string
}

type Content struct {
	Quote   Quote
	Search  Search
	Article Article
}

type Footer struct {
	Text string
}

type Page struct {
	Header  Header
	Content Content
	Footer  Footer
}
