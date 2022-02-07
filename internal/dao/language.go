package dao

// go:generate reform
// reform:languages
type Language struct {
	ID     int64  `reform:"id,pk"`
	Locale string `reform:"locale"`
	Name   string `reform:"name"`
}
