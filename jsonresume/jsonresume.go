package jsonresume

type JSONResume struct {
	Basics       Basics
	Work         []Work
	Volunteer    []Volunteer
	Education    []Education
	Awards       []Award
	Publications []Publication
	Skills       []Skill
	Languages    []Language
	Interests    []Interest
	References   []Reference
}
