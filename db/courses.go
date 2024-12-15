package db

import (
	"audit/models"
)

var CoreCourses = map[string]models.Course{
	//MATH courses
	"MATH161": {
		Name:         "Calculus I",
		Credits:      8,
		Prerequisite: "",
	},
	"MATH162": {
		Name:         "Calculus II",
		Credits:      8,
		Prerequisite: "",
	},
	"MATH273": {
		Name:         "Linear Algebra with Applications",
		Credits:      8,
		Prerequisite: "",
	},
	"MATH251": {
		Name:         "Discrete Mathematics",
		Credits:      6,
		Prerequisite: "",
	},
	"MATH321": {
		Name:         "Probability",
		Credits:      6,
		Prerequisite: "",
	},

	//CSCI courses
	"CSCI151": {
		Name:         "Programming for Scientists and Engineering",
		Credits:      8,
		Prerequisite: "",
	},
	"CSCI152": {
		Name:         "Performace and Data Structures",
		Credits:      8,
		Prerequisite: "",
	},
	"CSCI231": {
		Name:         "Computer Systems & Organisations",
		Credits:      6,
		Prerequisite: "",
	},
	"CSCI235": {
		Name:         "Programming Languages",
		Credits:      8,
		Prerequisite: "",
	},
	"CSCI272": {
		Name:         "Formal Languages",
		Credits:      6,
		Prerequisite: "",
	},
	"CSCI270": {
		Name:         "Algorithms",
		Credits:      6,
		Prerequisite: "",
	},
	"CSCI390": {
		Name:         "Artificial Intelligence",
		Credits:      6,
		Prerequisite: "",
	},
	"CSCI341": {
		Name:         "Database Systems",
		Credits:      6,
		Prerequisite: "",
	},
	"CSCI361": {
		Name:         "Software Engineering",
		Credits:      6,
		Prerequisite: "",
	},
	"CSCI333": {
		Name:         "Computer Networks",
		Credits:      6,
		Prerequisite: "",
	},
	"CSCI332": {
		Name:         "Operating Systems",
		Credits:      6,
		Prerequisite: "",
	},
	"CSCI307": {
		Name:         "Research Methods",
		Credits:      6,
		Prerequisite: "",
	},
	"CSCI408": {
		Name:         "Senior Project I",
		Credits:      6,
		Prerequisite: "",
	},
	"CSCI409": {
		Name:         "Senior Project II",
		Credits:      6,
		Prerequisite: "",
	},

	//HST
	"HST100": {
		Name:         "History of Kazakhstan",
		Credits:      6,
		Prerequisite: "",
	},
	"WCS150": {
		Name:         "Rhetorics and Composition",
		Credits:      6,
		Prerequisite: "",
	},

	//BUS
	"BUS101": {
		Name:         "Core Course in Business",
		Credits:      6,
		Prerequisite: "",
	},

	//PHIL
	"PHIL210": {
		Name:         "Ethics",
		Credits:      6,
		Prerequisite: "",
	},

	//PHYS
	"PHYS161": {
		Name:         "Physics I with Lab",
		Credits:      8,
		Prerequisite: "",
	},
	"PHYS162": {
		Name:         "Physics II with Lab",
		Credits:      8,
		Prerequisite: "",
	},

	//ROBT
	"ROBT206": {
		Name:         "Microcontrollers with Lab",
		Credits:      8,
		Prerequisite: "",
	},
}

var TechnicalElectiveCourses = map[string]models.Course{
	// MATH courses
	"MATH351": {
		Name:         "Numerical Methods w/ Applications",
		Credits:      6,
		Prerequisite: "MATH 274 & 263 gr. C",
	},
	"MATH407": {
		Name:         "Intro to Graph Theory",
		Credits:      6,
		Prerequisite: "MATH 273 & 251/301 gr. C",
	},
	"MATH417": {
		Name:         "Cryptography",
		Credits:      6,
		Prerequisite: "MATH 273 & 251/301/355 gr. C",
	},
	// PHYS courses
	"PHYS270": {
		Name:         "Computational Physics w/ Lab",
		Credits:      6,
		Prerequisite: "MATH 263 & PHYS 162/172 & CSCI 151/150 gr. C-",
	},
	// ROBT courses
	"ROBT205": {
		Name:         "Signals & Sensing w/ Lab",
		Credits:      6,
		Prerequisite: "MATH 162 & PHYS 162/172 gr. C-",
	},
	"ROBT305": {
		Name:         "Embedded Systems",
		Credits:      6,
		Prerequisite: "ROBT 206 gr. C-",
	},
	"ROBT310": {
		Name:         "Image Processing",
		Credits:      6,
		Prerequisite: "MATH 162 & 273 gr. C-",
	},
	"ROBT407": {
		Name:         "Statistical Methods & Machine Learning",
		Credits:      6,
		Prerequisite: "MATH 273 & 321 & CSCI 152 gr. C-",
	},
	"ROBT414": {
		Name:         "Human-Robot Interaction",
		Credits:      6,
		Prerequisite: "CSCI 152 gr. C-",
	},
	// ELCE courses
	"ELCE202": {
		Name:         "Digital Logic Design",
		Credits:      6,
		Prerequisite: "MATH 162 & PHYS 162 & ENG 101 gr. C-",
	},
	"ELCE203": {
		Name:         "Signals & Systems",
		Credits:      6,
		Prerequisite: "MATH 162 & PHYS 162 gr. C-",
	},
	"ELCE307": {
		Name:         "Digital Signal Processing",
		Credits:      6,
		Prerequisite: "ELCE 203 gr. C-",
	},
	"ELCE308": {
		Name:         "Communication Systems",
		Credits:      6,
		Prerequisite: "ELCE 203 gr. C-",
	},
	"ELCE461": {
		Name:         "Industrial Automation",
		Credits:      6,
		Prerequisite: "ELCE 202 gr. C-",
	},
	"ELCE462": {
		Name:         "Wireless Networks",
		Credits:      6,
		Prerequisite: "ELCE 201 & ELCE 308/304 gr. C-",
	},
	"ELCE465": {
		Name:         "PCB Design & Manufacturing",
		Credits:      6,
		Prerequisite: "N/A",
	},
}

var CommunicationCoreCourses = map[string]models.Course{
	// WCS courses
	"WCS200": {
		Name:         "Communication Skills 200",
		Credits:      6,
		Prerequisite: "N/A",
	},
	"WCS210": {
		Name:         "Communication Skills 210",
		Credits:      6,
		Prerequisite: "N/A",
	},
	"WCS220": {
		Name:         "Communication Skills 220",
		Credits:      6,
		Prerequisite: "N/A",
	},
	"WCS230": {
		Name:         "Communication Skills 230",
		Credits:      6,
		Prerequisite: "N/A",
	},
	"WCS240": {
		Name:         "Communication Skills 240",
		Credits:      6,
		Prerequisite: "N/A",
	},
	"WCS250": {
		Name:         "Communication Skills 250",
		Credits:      6,
		Prerequisite: "N/A",
	},
}

var PrefixesRequired = []string{"BUS", "CSCI", "HST", "MATH", "PHIL", "PHYS", "ROBT", "WCS"}
var PrefixNaturalElectives = []string{"BIOL", "CHEM", "GEOL", "PHYS"}
var PrefixTechnicalElectives = []string{"BENG", "ELCE", "ENG"} //probably can be removed
var PrefixKAZElectives = []string{"KAZ"}
var PrefixOpenElectives = []string{"MCHME"} //should be increased later
var PrefixSocialElectives = []string{"ANT", "ECON", "LING", "PLS", "SOC"}

var ElectiveCourses = map[string]int{
	"Technical":          4,
	"Communication Core": 1,
	"Kazakh Language":    2,
	"Natural Science":    2,
	"Social Science":     1,
	"Open Elective":      1,
}
