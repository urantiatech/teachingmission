package page

type User struct {
	Username string
	Email    string
	Role     string
}

type Menu struct {
	Title   string
	Url     string
	SubMenu *Menu
}

type Banner struct {
	Url  string
	Text string
}

type Quote struct {
	Text    string
	Teacher string
}

type Category struct {
	Id   string
	Name string
}

type Filter struct {
	Category Category
	Teacher  string
	Receiver string
	Team     string
	DateFrom string
	DateTo   string
}

type Teacher struct {
	Id   string
	Name string
	Type string
}

type Person struct {
	Id   string
	Name string
}

type Section struct {
	TocTitle  string
	Heading   string
	Paragraph []string
}

type Transcript struct {
	Id       string
	Title    string
	Category Category
	Date     string
	Location string
	Teacher  []Teacher
	Receiver Person
	Attendee []Person
	Body     []Section
}

type InternalResult struct {
	Title    string
	Teacher  string
	Uri      string
	Excerpts string
}

type ExternalResult struct {
	Title    string
	Url      string
	Excerpts string
}

type Search struct {
	Keywords        string
	Filters         Filter
	Results         []InternalResult
	ExternalResults []ExternalResult
}

type Article struct {
	Title string
	Id    string
	Text  string
}

type Meta struct {
	Keywords    string
	Description string
}

type Page struct {
	Title      string
	Meta       *Meta
	User       *User
	Menu       *Menu
	Banner     *Banner
	Quote      *Quote
	Search     *Search
	Transcript *Transcript
	Article    *Article
}
