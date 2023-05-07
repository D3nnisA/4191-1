CREATE TABLE IF NOT EXISTS courses (
id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    CourseCode  text NOT NULL,  
	CourseTitle  text NOT NULL,
	CourseCredit text NOT NULL,
	version integer NOT NULL DEFAULT 1

);
