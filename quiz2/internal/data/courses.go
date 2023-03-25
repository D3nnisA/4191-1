package data

//represent one row of data in our courses table

type Courses struct {
	CourseCode   string `json:"Course Code"`
	CourseTitle  string `json:"Course Title"`
	CourseCredit string `json:"Course Credit"`
	Version      int32  `json:"version"`
}
