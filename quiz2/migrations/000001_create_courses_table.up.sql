CREATE TABLE IF NOT EXISTS courses (

    course_code  text NOT NULL,  
	course_title  text NOT NULL,
	course_credit integer NOT NULL,
	Version     integer NOT NULL DEFAULT 1 

);