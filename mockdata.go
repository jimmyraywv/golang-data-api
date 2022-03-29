package main

import (
	"encoding/json"
	Log "github.com/sirupsen/logrus"
	"jimmyray.io/data-api/utils"
)

const mockData string = `{
	"218000": {
	  "id": "218000",
	  "fname": "Indrajit",
	  "lname": "Raney",
	  "sex": "M",
	  "dob": "1964-08-04T00:00:00Z",
	  "hireDate": "1989-08-31T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 64633,
	  "dept": { "id": "d006", "name": "Quality Management", "mgrId": "110854" },
	  "address": {
		"street": "4200 Old Us Highway395n N",
		"city": "Washoe Valley",
		"county": "Washoe",
		"state": "NV",
		"zipcode": "89704"
	  }
	},
	"218001": {
	  "id": "218001",
	  "fname": "Hausi",
	  "lname": "Jansch",
	  "sex": "F",
	  "dob": "1952-04-03T00:00:00Z",
	  "hireDate": "1996-09-08T00:00:00Z",
	  "position": "Staff",
	  "salary": 53819,
	  "dept": { "id": "d002", "name": "Finance", "mgrId": "110114" },
	  "address": {
		"street": "2519 Broadway",
		"city": "Rockford",
		"county": "Winnebago",
		"state": "IL",
		"zipcode": "61108"
	  }
	},
	"218002": {
	  "id": "218002",
	  "fname": "Dietrich",
	  "lname": "Masamoto",
	  "sex": "F",
	  "dob": "1960-03-29T00:00:00Z",
	  "hireDate": "1992-11-04T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 58582,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "295 State Route 10 E",
		"city": "Succasunna",
		"county": "Morris",
		"state": "NJ",
		"zipcode": "7876"
	  }
	},
	"218003": {
	  "id": "218003",
	  "fname": "Hideyuki",
	  "lname": "Crouzet",
	  "sex": "M",
	  "dob": "1959-05-22T00:00:00Z",
	  "hireDate": "1992-06-11T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "315 Clay St",
		"city": "Cedar Falls",
		"county": "Black Hawk",
		"state": "IA",
		"zipcode": "50613"
	  }
	},
	"218004": {
	  "id": "218004",
	  "fname": "Odinaldo",
	  "lname": "Hanratty",
	  "sex": "M",
	  "dob": "1953-12-29T00:00:00Z",
	  "hireDate": "1996-03-21T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 64339,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "24 Frank Lloyd Wright Dr",
		"city": "Ann Arbor",
		"county": "Washtenaw",
		"state": "MI",
		"zipcode": "48105"
	  }
	},
	"218005": {
	  "id": "218005",
	  "fname": "Koldo",
	  "lname": "Greibach",
	  "sex": "F",
	  "dob": "1964-09-02T00:00:00Z",
	  "hireDate": "1989-08-10T00:00:00Z",
	  "position": "Staff",
	  "salary": 61255,
	  "dept": { "id": "d002", "name": "Finance", "mgrId": "110114" },
	  "address": {
		"street": "185 Old Colony Ave",
		"city": "Boston",
		"county": "Suffolk",
		"state": "MA",
		"zipcode": "2127"
	  }
	},
	"218006": {
	  "id": "218006",
	  "fname": "Ger",
	  "lname": "Colorni",
	  "sex": "F",
	  "dob": "1954-12-29T00:00:00Z",
	  "hireDate": "1988-03-19T00:00:00Z",
	  "position": "Staff",
	  "salary": 60947,
	  "dept": { "id": "d001", "name": "Marketing", "mgrId": "110039" },
	  "address": {
		"street": "424 S Minnesota Ave",
		"city": "Saint Peter",
		"county": "Nicollet",
		"state": "MN",
		"zipcode": "56082"
	  }
	},
	"218007": {
	  "id": "218007",
	  "fname": "Radoslaw",
	  "lname": "Benaini",
	  "sex": "F",
	  "dob": "1963-01-09T00:00:00Z",
	  "hireDate": "1990-03-01T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "13 E Eau Claire St",
		"city": "Rice Lake",
		"county": "Barron",
		"state": "WI",
		"zipcode": "54868"
	  }
	},
	"218008": {
	  "id": "218008",
	  "fname": "Kristina",
	  "lname": "Baalen",
	  "sex": "F",
	  "dob": "1958-05-22T00:00:00Z",
	  "hireDate": "1997-08-27T00:00:00Z",
	  "position": "Staff",
	  "salary": 46938,
	  "dept": { "id": "d009", "name": "Customer Service", "mgrId": "111939" },
	  "address": {
		"street": "325 3rd St",
		"city": "Iowa City",
		"county": "Johnson",
		"state": "IA",
		"zipcode": "52240"
	  }
	},
	"218009": {
	  "id": "218009",
	  "fname": "George",
	  "lname": "Stanfel",
	  "sex": "M",
	  "dob": "1956-03-03T00:00:00Z",
	  "hireDate": "1988-10-21T00:00:00Z",
	  "position": "Senior Staff",
	  "salary": 65535,
	  "dept": { "id": "d003", "name": "Human Resources", "mgrId": "110228" },
	  "address": {
		"street": "214 W Main St",
		"city": "Fredericksburg",
		"county": "Gillespie",
		"state": "TX",
		"zipcode": "78624"
	  }
	},
	"218010": {
	  "id": "218010",
	  "fname": "Eran",
	  "lname": "Farrow",
	  "sex": "M",
	  "dob": "1959-09-27T00:00:00Z",
	  "hireDate": "1986-08-07T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "340 Ne 11th Ave",
		"city": "Portland",
		"county": "Multnomah",
		"state": "OR",
		"zipcode": "97232"
	  }
	},
	"218011": {
	  "id": "218011",
	  "fname": "Yonghong",
	  "lname": "Wiegley",
	  "sex": "F",
	  "dob": "1952-02-11T00:00:00Z",
	  "hireDate": "1992-03-06T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 56197,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "1995 University Ave #-2",
		"city": "Berkeley",
		"county": "Alameda",
		"state": "CA",
		"zipcode": "94704"
	  }
	},
	"218012": {
	  "id": "218012",
	  "fname": "Eishiro",
	  "lname": "Walstra",
	  "sex": "M",
	  "dob": "1963-09-17T00:00:00Z",
	  "hireDate": "1988-01-23T00:00:00Z",
	  "position": "Assistant Engineer",
	  "salary": 54277,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "530 E Main St",
		"city": "Charlottesville",
		"county": "Charlottesville City",
		"state": "VA",
		"zipcode": "22902"
	  }
	},
	"218013": {
	  "id": "218013",
	  "fname": "Visit",
	  "lname": "Pavlopoulou",
	  "sex": "M",
	  "dob": "1962-07-10T00:00:00Z",
	  "hireDate": "1985-11-01T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 57507,
	  "dept": { "id": "d006", "name": "Quality Management", "mgrId": "110854" },
	  "address": {
		"street": "50 Broadway",
		"city": "Westwood",
		"county": "Bergen",
		"state": "NJ",
		"zipcode": "7675"
	  }
	},
	"218014": {
	  "id": "218014",
	  "fname": "Krister",
	  "lname": "Redmiles",
	  "sex": "M",
	  "dob": "1952-08-05T00:00:00Z",
	  "hireDate": "1991-03-26T00:00:00Z",
	  "position": "Senior Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "157 E New England Ave #-375",
		"city": " Winter Park",
		"county": "Orange",
		"state": "FL",
		"zipcode": "32789"
	  }
	},
	"218015": {
	  "id": "218015",
	  "fname": "Mang",
	  "lname": "Spinelli",
	  "sex": "F",
	  "dob": "1964-10-11T00:00:00Z",
	  "hireDate": "1986-11-09T00:00:00Z",
	  "position": "Staff",
	  "salary": 61974,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "3 A D P Blvd",
		"city": "Roseland",
		"county": "Essex",
		"state": "NJ",
		"zipcode": "7068"
	  }
	},
	"218016": {
	  "id": "218016",
	  "fname": "Jayson",
	  "lname": "Shumilov",
	  "sex": "M",
	  "dob": "1964-11-16T00:00:00Z",
	  "hireDate": "1997-05-24T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "1 N Charles St",
		"city": "Baltimore",
		"county": "Baltimore City",
		"state": "MD",
		"zipcode": "21201"
	  }
	},
	"218017": {
	  "id": "218017",
	  "fname": "Feipei",
	  "lname": "Narlikar",
	  "sex": "M",
	  "dob": "1956-05-20T00:00:00Z",
	  "hireDate": "1986-06-02T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "720 Olive St #-2100",
		"city": "Saint Louis",
		"county": "Saint Louis City",
		"state": "MO",
		"zipcode": "63101"
	  }
	},
	"218018": {
	  "id": "218018",
	  "fname": "Eckart",
	  "lname": "Chandrasekhar",
	  "sex": "M",
	  "dob": "1959-05-14T00:00:00Z",
	  "hireDate": "1986-01-21T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d006", "name": "Quality Management", "mgrId": "110854" },
	  "address": {
		"street": "70080 Exchange St",
		"city": "Binghamton",
		"county": "Broome",
		"state": "NY",
		"zipcode": "13901"
	  }
	},
	"218019": {
	  "id": "218019",
	  "fname": "Kamakshi",
	  "lname": "Gustavson",
	  "sex": "F",
	  "dob": "1964-09-06T00:00:00Z",
	  "hireDate": "1985-02-15T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d009", "name": "Customer Service", "mgrId": "111939" },
	  "address": {
		"street": "101 E Chesapeake Ave",
		"city": " Towson",
		"county": "Baltimore",
		"state": "MD",
		"zipcode": "21286"
	  }
	},
	"218020": {
	  "id": "218020",
	  "fname": "Fumino",
	  "lname": "Cappello",
	  "sex": "M",
	  "dob": "1956-02-20T00:00:00Z",
	  "hireDate": "1985-12-27T00:00:00Z",
	  "position": "Engineer",
	  "salary": 42944,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "1402 Airport Rd",
		"city": "Urbana",
		"county": "Champaign",
		"state": "IL",
		"zipcode": "61801"
	  }
	},
	"218021": {
	  "id": "218021",
	  "fname": "Yuguang",
	  "lname": "Casperson",
	  "sex": "M",
	  "dob": "1957-10-17T00:00:00Z",
	  "hireDate": "1987-11-08T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d009", "name": "Customer Service", "mgrId": "111939" },
	  "address": {
		"street": "16 Church St",
		"city": "Bound Brook",
		"county": "Somerset",
		"state": "NJ",
		"zipcode": "8805"
	  }
	},
	"218022": {
	  "id": "218022",
	  "fname": "Mokhtar",
	  "lname": "Poupard",
	  "sex": "M",
	  "dob": "1965-01-26T00:00:00Z",
	  "hireDate": "1987-06-18T00:00:00Z",
	  "position": "Technique Leader",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "1351 Kuser Rd",
		"city": "Trenton",
		"county": "Mercer",
		"state": "NJ",
		"zipcode": "8619"
	  }
	},
	"218023": {
	  "id": "218023",
	  "fname": "Bowen",
	  "lname": "Hooghiemstra",
	  "sex": "M",
	  "dob": "1964-02-13T00:00:00Z",
	  "hireDate": "1987-03-28T00:00:00Z",
	  "position": "Senior Staff",
	  "salary": 47977,
	  "dept": { "id": "d009", "name": "Customer Service", "mgrId": "111939" },
	  "address": {
		"street": "2663 Farragut Dr",
		"city": "Springfield",
		"county": "Sangamon",
		"state": "IL",
		"zipcode": "62704"
	  }
	},
	"218024": {
	  "id": "218024",
	  "fname": "Maha",
	  "lname": "Madeira",
	  "sex": "M",
	  "dob": "1963-09-30T00:00:00Z",
	  "hireDate": "1994-08-01T00:00:00Z",
	  "position": "Senior Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "320 Penn Ave",
		"city": "Reading",
		"county": "Berks",
		"state": "PA",
		"zipcode": "19611"
	  }
	},
	"218025": {
	  "id": "218025",
	  "fname": "Jungsoon",
	  "lname": "Mansanne",
	  "sex": "M",
	  "dob": "1957-11-22T00:00:00Z",
	  "hireDate": "1993-08-18T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 59092,
	  "dept": { "id": "d006", "name": "Quality Management", "mgrId": "110854" },
	  "address": {
		"street": "534 Main St",
		"city": "Madawaska",
		"county": "Aroostook",
		"state": "ME",
		"zipcode": "4756"
	  }
	},
	"218026": {
	  "id": "218026",
	  "fname": "Shuji",
	  "lname": "Liesche",
	  "sex": "M",
	  "dob": "1956-12-23T00:00:00Z",
	  "hireDate": "1988-03-23T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 52434,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "55 New Orleans Rd #-205",
		"city": "Hilton Head Island",
		"county": "Beaufort",
		"state": "SC",
		"zipcode": "29928"
	  }
	},
	"218027": {
	  "id": "218027",
	  "fname": "Hideyuki",
	  "lname": "Gide",
	  "sex": "M",
	  "dob": "1960-03-21T00:00:00Z",
	  "hireDate": "1992-08-01T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 64577,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "101 W Broadway",
		"city": "San Diego",
		"county": "San Diego",
		"state": "CA",
		"zipcode": "92101"
	  }
	},
	"218028": {
	  "id": "218028",
	  "fname": "Muzhong",
	  "lname": "Pokrovskii",
	  "sex": "M",
	  "dob": "1963-04-03T00:00:00Z",
	  "hireDate": "1987-09-07T00:00:00Z",
	  "position": "Staff",
	  "salary": 53649,
	  "dept": { "id": "d002", "name": "Finance", "mgrId": "110114" },
	  "address": {
		"street": "333 Market St",
		"city": "San Francisco",
		"county": "San Francisco",
		"state": "CA",
		"zipcode": "94105"
	  }
	},
	"218029": {
	  "id": "218029",
	  "fname": "Vivian",
	  "lname": "Lortz",
	  "sex": "M",
	  "dob": "1954-05-03T00:00:00Z",
	  "hireDate": "1986-07-17T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d008", "name": "Research", "mgrId": "111534" },
	  "address": {
		"street": "11221 W Burleigh St",
		"city": "Milwaukee",
		"county": "Milwaukee",
		"state": "WI",
		"zipcode": "53222"
	  }
	},
	"229970": {
	  "id": "229970",
	  "fname": "Tse",
	  "lname": "Berendt",
	  "sex": "F",
	  "dob": "1955-08-16T00:00:00Z",
	  "hireDate": "1988-02-21T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "6051 N Fresno St #-200",
		"city": "Fresno",
		"county": "Fresno",
		"state": "CA",
		"zipcode": "93710"
	  }
	},
	"229971": {
	  "id": "229971",
	  "fname": "Limsoon",
	  "lname": "Gopalakrishnan",
	  "sex": "M",
	  "dob": "1958-01-17T00:00:00Z",
	  "hireDate": "1987-05-27T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d006", "name": "Quality Management", "mgrId": "110854" },
	  "address": {
		"street": "111 E Washington St",
		"city": "Walterboro",
		"county": "Colleton",
		"state": "SC",
		"zipcode": "29488"
	  }
	},
	"229972": {
	  "id": "229972",
	  "fname": "Serge",
	  "lname": "Segond",
	  "sex": "M",
	  "dob": "1960-07-29T00:00:00Z",
	  "hireDate": "1987-05-02T00:00:00Z",
	  "position": "Engineer",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "4 Seagate",
		"city": "Toledo",
		"county": "Lucas",
		"state": "OH",
		"zipcode": "43604"
	  }
	},
	"229999": {
	  "id": "229999",
	  "fname": "Wilmer",
	  "lname": "Schnabel",
	  "sex": "F",
	  "dob": "1964-10-03T00:00:00Z",
	  "hireDate": "1986-06-20T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "309 S Willard St",
		"city": "Burlington",
		"county": "Chittenden",
		"state": "VT",
		"zipcode": "5401"
	  }
	},
	"28120": {
	  "id": "28120",
	  "fname": "Arra",
	  "lname": "Zolotykh",
	  "sex": "F",
	  "dob": "1964-01-28T00:00:00Z",
	  "hireDate": "1992-04-17T00:00:00Z",
	  "position": "Technique Leader",
	  "salary": 53795,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "8400 W 110th St #-450",
		"city": " Overland Park",
		"county": "Johnson",
		"state": "KS",
		"zipcode": "66210"
	  }
	},
	"28121": {
	  "id": "28121",
	  "fname": "Sreekrishna",
	  "lname": "Chiola",
	  "sex": "M",
	  "dob": "1960-05-14T00:00:00Z",
	  "hireDate": "1990-12-25T00:00:00Z",
	  "position": "Staff",
	  "salary": 63435,
	  "dept": { "id": "d009", "name": "Customer Service", "mgrId": "111939" },
	  "address": {
		"street": "15 Avon St",
		"city": "Keene",
		"county": "Cheshire",
		"state": "NH",
		"zipcode": "3431"
	  }
	},
	"28122": {
	  "id": "28122",
	  "fname": "Arve",
	  "lname": "Skafidas",
	  "sex": "M",
	  "dob": "1955-03-24T00:00:00Z",
	  "hireDate": "1994-04-28T00:00:00Z",
	  "position": "Staff",
	  "salary": 62964,
	  "dept": { "id": "d003", "name": "Human Resources", "mgrId": "110228" },
	  "address": {
		"street": "3640 Colonel Glenn Hwy",
		"city": "Dayton",
		"county": "Greene",
		"state": "OH",
		"zipcode": "45435"
	  }
	},
	"28123": {
	  "id": "28123",
	  "fname": "Vincent",
	  "lname": "Tomescu",
	  "sex": "M",
	  "dob": "1955-12-06T00:00:00Z",
	  "hireDate": "1985-06-26T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 57474,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "1625 Broadway #-1600",
		"city": "Denver",
		"county": "Denver",
		"state": "CO",
		"zipcode": "80202"
	  }
	},
	"28124": {
	  "id": "28124",
	  "fname": "Adhemar",
	  "lname": "Eiter",
	  "sex": "M",
	  "dob": "1958-08-31T00:00:00Z",
	  "hireDate": "1988-12-25T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d002", "name": "Finance", "mgrId": "110114" },
	  "address": {
		"street": "1709 Frederica Rd",
		"city": "Saint Simons Island",
		"county": "Glynn",
		"state": "GA",
		"zipcode": "31522"
	  }
	},
	"28125": {
	  "id": "28125",
	  "fname": "Ottavia",
	  "lname": "Marquardt",
	  "sex": "M",
	  "dob": "1960-11-22T00:00:00Z",
	  "hireDate": "1985-04-18T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "666 Walnut St #-2500",
		"city": "Des Moines",
		"county": "Polk",
		"state": "IA",
		"zipcode": "50309"
	  }
	},
	"28126": {
	  "id": "28126",
	  "fname": "Sanjiv",
	  "lname": "Dolinsky",
	  "sex": "M",
	  "dob": "1953-04-26T00:00:00Z",
	  "hireDate": "1993-07-27T00:00:00Z",
	  "position": "Senior Staff",
	  "salary": 55704,
	  "dept": { "id": "d003", "name": "Human Resources", "mgrId": "110228" },
	  "address": {
		"street": "1487 Farnsworth St",
		"city": "Detroit",
		"county": "Wayne",
		"state": "MI",
		"zipcode": "48211"
	  }
	},
	"28127": {
	  "id": "28127",
	  "fname": "Minghong",
	  "lname": "Veevers",
	  "sex": "M",
	  "dob": "1953-07-08T00:00:00Z",
	  "hireDate": "1988-04-15T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d002", "name": "Finance", "mgrId": "110114" },
	  "address": {
		"street": "501 Wyoming Ave",
		"city": "Scranton",
		"county": "Lackawanna",
		"state": "PA",
		"zipcode": "18509"
	  }
	},
	"28128": {
	  "id": "28128",
	  "fname": "Huican",
	  "lname": "Slaats",
	  "sex": "F",
	  "dob": "1954-01-10T00:00:00Z",
	  "hireDate": "1987-01-26T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "800 Greenwich Ave",
		"city": "Warwick",
		"county": "Kent",
		"state": "RI",
		"zipcode": "2886"
	  }
	},
	"28129": {
	  "id": "28129",
	  "fname": "Arco",
	  "lname": "Swiler",
	  "sex": "F",
	  "dob": "1958-07-19T00:00:00Z",
	  "hireDate": "1985-08-09T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d009", "name": "Customer Service", "mgrId": "111939" },
	  "address": {
		"street": "420 S Dixie Hwy #-2r",
		"city": "Miami",
		"county": "Miami-Dade",
		"state": "FL",
		"zipcode": "33146"
	  }
	},
	"28230": {
	  "id": "28230",
	  "fname": "Fan",
	  "lname": "Roohalamini",
	  "sex": "M",
	  "dob": "1958-03-19T00:00:00Z",
	  "hireDate": "1985-05-10T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "6501 E Nevada St",
		"city": "Detroit",
		"county": "Wayne",
		"state": "MI",
		"zipcode": "48234"
	  }
	},
	"28231": {
	  "id": "28231",
	  "fname": "Shuky",
	  "lname": "Anger",
	  "sex": "F",
	  "dob": "1963-01-03T00:00:00Z",
	  "hireDate": "1988-02-07T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "5090 Central Hwy",
		"city": "Merchantville",
		"county": "Camden",
		"state": "NJ",
		"zipcode": "8109"
	  }
	},
	"28232": {
	  "id": "28232",
	  "fname": "Fun",
	  "lname": "Veccia",
	  "sex": "M",
	  "dob": "1956-05-21T00:00:00Z",
	  "hireDate": "1988-07-17T00:00:00Z",
	  "position": "Engineer",
	  "salary": 65535,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "9423 Reseda Blvd",
		"city": "Northridge",
		"county": "LosAngeles",
		"state": "CA",
		"zipcode": "91324"
	  }
	},
	"28233": {
	  "id": "28233",
	  "fname": "Uinam",
	  "lname": "Zolotykh",
	  "sex": "F",
	  "dob": "1960-03-11T00:00:00Z",
	  "hireDate": "1985-06-16T00:00:00Z",
	  "position": "Technique Leader",
	  "salary": 53368,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "1931 W San Carlos St",
		"city": " San Jose",
		"county": "Santa Clara",
		"state": "CA",
		"zipcode": "95128"
	  }
	},
	"28234": {
	  "id": "28234",
	  "fname": "Basem",
	  "lname": "Greibach",
	  "sex": "F",
	  "dob": "1961-12-18T00:00:00Z",
	  "hireDate": "1985-11-21T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d006", "name": "Quality Management", "mgrId": "110854" },
	  "address": {
		"street": "18 N High St",
		"city": "Millville",
		"county": "Cumberland",
		"state": "NJ",
		"zipcode": "8332"
	  }
	},
	"28235": {
	  "id": "28235",
	  "fname": "Shaibal",
	  "lname": "Unno",
	  "sex": "F",
	  "dob": "1960-02-08T00:00:00Z",
	  "hireDate": "1985-08-05T00:00:00Z",
	  "position": "Assistant Engineer",
	  "salary": 52177,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "28985 Golden Lantern",
		"city": "Laguna Niguel",
		"county": "Orange",
		"state": "CA",
		"zipcode": "92677"
	  }
	},
	"28236": {
	  "id": "28236",
	  "fname": "Kazuhira",
	  "lname": "Radivojevic",
	  "sex": "M",
	  "dob": "1954-02-10T00:00:00Z",
	  "hireDate": "1991-04-16T00:00:00Z",
	  "position": "Engineer",
	  "salary": 48354,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "21212 Laguna Hills Dr",
		"city": "Aliso Viejo",
		"county": "Orange",
		"state": "CA",
		"zipcode": "92656"
	  }
	},
	"28237": {
	  "id": "28237",
	  "fname": "Danny",
	  "lname": "Suessmith",
	  "sex": "F",
	  "dob": "1961-06-04T00:00:00Z",
	  "hireDate": "1996-01-23T00:00:00Z",
	  "position": "Staff",
	  "salary": 45335,
	  "dept": { "id": "d009", "name": "Customer Service", "mgrId": "111939" },
	  "address": {
		"street": "21212 Bake Pky #-a",
		"city": "Lake Forest",
		"county": "Orange",
		"state": "CA",
		"zipcode": "92630"
	  }
	},
	"28238": {
	  "id": "28238",
	  "fname": "Rafols",
	  "lname": "Ushiama",
	  "sex": "M",
	  "dob": "1957-03-21T00:00:00Z",
	  "hireDate": "1994-01-20T00:00:00Z",
	  "position": "Staff",
	  "salary": 46809,
	  "dept": { "id": "d001", "name": "Marketing", "mgrId": "110039" },
	  "address": {
		"street": "3951 Fruitvale Ave",
		"city": "Bakersfield",
		"county": "Kern",
		"state": "CA",
		"zipcode": "93308"
	  }
	},
	"28239": {
	  "id": "28239",
	  "fname": "Satosi",
	  "lname": "Farris",
	  "sex": "F",
	  "dob": "1961-01-27T00:00:00Z",
	  "hireDate": "1992-04-20T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "1635 Old Hardin Rd",
		"city": "Billings",
		"county": "Yellowstone",
		"state": "MT",
		"zipcode": "59101"
	  }
	},
	"28241": {
	  "id": "28241",
	  "fname": "Yuguang",
	  "lname": "Baik",
	  "sex": "M",
	  "dob": "1959-01-04T00:00:00Z",
	  "hireDate": "1988-01-28T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "227 Clifton Ave #-4",
		"city": "Darby",
		"county": "Delaware",
		"state": "PA",
		"zipcode": "19023"
	  }
	},
	"28242": {
	  "id": "28242",
	  "fname": "Thodoros",
	  "lname": "Vasanthakumar",
	  "sex": "M",
	  "dob": "1955-03-17T00:00:00Z",
	  "hireDate": "1986-06-15T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "2104 S King St #-203",
		"city": "Honolulu",
		"county": "Honolulu",
		"state": "HI",
		"zipcode": "96826"
	  }
	},
	"28243": {
	  "id": "28243",
	  "fname": "Marsal",
	  "lname": "Jording",
	  "sex": "M",
	  "dob": "1952-09-03T00:00:00Z",
	  "hireDate": "1995-06-14T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 48967,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "2 Adams St",
		"city": "Denver",
		"county": "Denver",
		"state": "CO",
		"zipcode": "80206"
	  }
	},
	"28270": {
	  "id": "28270",
	  "fname": "Leaf",
	  "lname": "Gyorkos",
	  "sex": "M",
	  "dob": "1962-06-19T00:00:00Z",
	  "hireDate": "1994-04-06T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 50618,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "320 Sw Oak St #-500",
		"city": "Portland",
		"county": "Multnomah",
		"state": "OR",
		"zipcode": "97204"
	  }
	},
	"28271": {
	  "id": "28271",
	  "fname": "Sugwoo",
	  "lname": "Wynblatt",
	  "sex": "F",
	  "dob": "1952-05-31T00:00:00Z",
	  "hireDate": "1989-09-18T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "249 N Main St",
		"city": "Manville",
		"county": "Somerset",
		"state": "NJ",
		"zipcode": "8835"
	  }
	},
	"28272": {
	  "id": "28272",
	  "fname": "Shigeichiro",
	  "lname": "Savasere",
	  "sex": "M",
	  "dob": "1963-06-10T00:00:00Z",
	  "hireDate": "1985-09-10T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d009", "name": "Customer Service", "mgrId": "111939" },
	  "address": {
		"street": "185 W F St #-200",
		"city": "San Diego",
		"county": "San Diego",
		"state": "CA",
		"zipcode": "92101"
	  }
	},
	"28273": {
	  "id": "28273",
	  "fname": "Rosine",
	  "lname": "Schade",
	  "sex": "F",
	  "dob": "1960-03-28T00:00:00Z",
	  "hireDate": "1993-06-05T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 55662,
	  "dept": { "id": "d006", "name": "Quality Management", "mgrId": "110854" },
	  "address": {
		"street": "919 3rd Ave",
		"city": "New York",
		"county": "New York",
		"state": "NY",
		"zipcode": "10022"
	  }
	},
	"28274": {
	  "id": "28274",
	  "fname": "Arfst",
	  "lname": "Koblick",
	  "sex": "F",
	  "dob": "1959-08-24T00:00:00Z",
	  "hireDate": "1987-11-05T00:00:00Z",
	  "position": "Senior Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "900 Sun Valley Dr",
		"city": "Roswell",
		"county": "Fulton",
		"state": "GA",
		"zipcode": "30076"
	  }
	},
	"28275": {
	  "id": "28275",
	  "fname": "Zhonghua",
	  "lname": "Setlzner",
	  "sex": "M",
	  "dob": "1961-12-05T00:00:00Z",
	  "hireDate": "1985-05-07T00:00:00Z",
	  "position": "Senior Staff",
	  "salary": 65535,
	  "dept": { "id": "d008", "name": "Research", "mgrId": "111534" },
	  "address": {
		"street": "3601 W Marginal Way Sw",
		"city": "Seattle",
		"county": "King",
		"state": "WA",
		"zipcode": "98106"
	  }
	},
	"28276": {
	  "id": "28276",
	  "fname": "Danil",
	  "lname": "Rahier",
	  "sex": "F",
	  "dob": "1961-06-19T00:00:00Z",
	  "hireDate": "1986-08-28T00:00:00Z",
	  "position": "Technique Leader",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "10 S Main St",
		"city": "Mount Clemens",
		"county": "Macomb",
		"state": "MI",
		"zipcode": "48043"
	  }
	},
	"28277": {
	  "id": "28277",
	  "fname": "Katsuo",
	  "lname": "Litvinov",
	  "sex": "M",
	  "dob": "1959-02-10T00:00:00Z",
	  "hireDate": "1997-04-22T00:00:00Z",
	  "position": "Assistant Engineer",
	  "salary": 44834,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "530 Natchez St",
		"city": "New Orleans",
		"county": "Orleans",
		"state": "LA",
		"zipcode": "70130"
	  }
	},
	"28278": {
	  "id": "28278",
	  "fname": "Shushma",
	  "lname": "Farrow",
	  "sex": "F",
	  "dob": "1961-02-22T00:00:00Z",
	  "hireDate": "1986-10-07T00:00:00Z",
	  "position": "Technique Leader",
	  "salary": 64362,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "300 Main St",
		"city": "Buffalo",
		"county": " Erie",
		"state": "NY",
		"zipcode": "14202"
	  }
	},
	"28279": {
	  "id": "28279",
	  "fname": "Zhenhua",
	  "lname": "Brickell",
	  "sex": "M",
	  "dob": "1957-08-13T00:00:00Z",
	  "hireDate": "1987-02-07T00:00:00Z",
	  "position": "Senior Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "175 S 3rd St",
		"city": "Columbus",
		"county": "Franklin",
		"state": "OH",
		"zipcode": "43215"
	  }
	},
	"28280": {
	  "id": "28280",
	  "fname": "Demos",
	  "lname": "Demke",
	  "sex": "M",
	  "dob": "1961-03-13T00:00:00Z",
	  "hireDate": "1988-03-30T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "50 Division Ave",
		"city": "Millington",
		"county": "Morris",
		"state": "NJ",
		"zipcode": "7946"
	  }
	},
	"28281": {
	  "id": "28281",
	  "fname": "Adhemar",
	  "lname": "Oxenboll",
	  "sex": "M",
	  "dob": "1953-12-28T00:00:00Z",
	  "hireDate": "1985-07-05T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "101 E Chesapeake Ave",
		"city": "Towson",
		"county": "Baltimore",
		"state": "MD",
		"zipcode": "21286"
	  }
	},
	"28282": {
	  "id": "28282",
	  "fname": "Arto",
	  "lname": "Spelt",
	  "sex": "F",
	  "dob": "1953-10-03T00:00:00Z",
	  "hireDate": "1986-03-21T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "2041 E Burnside Rd",
		"city": "Gresham",
		"county": "Multnomah",
		"state": "OR",
		"zipcode": "97030"
	  }
	},
	"28283": {
	  "id": "28283",
	  "fname": "Quingbo",
	  "lname": "Dulay",
	  "sex": "F",
	  "dob": "1952-03-29T00:00:00Z",
	  "hireDate": "1985-11-07T00:00:00Z",
	  "position": "Engineer",
	  "salary": 65535,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "3246 N 16th St",
		"city": "Phoenix",
		"county": "Maricopa",
		"state": "AZ",
		"zipcode": "85016"
	  }
	},
	"28284": {
	  "id": "28284",
	  "fname": "Lene",
	  "lname": "Gurbaxani",
	  "sex": "M",
	  "dob": "1959-01-04T00:00:00Z",
	  "hireDate": "1994-12-28T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "953 Us Highway 202s S",
		"city": "Somerville",
		"county": "Somerset",
		"state": "NJ",
		"zipcode": "8876"
	  }
	},
	"28285": {
	  "id": "28285",
	  "fname": "Mircea",
	  "lname": "Heuter",
	  "sex": "M",
	  "dob": "1956-05-10T00:00:00Z",
	  "hireDate": "1986-05-10T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "1011 2nd St N",
		"city": "Saint Cloud",
		"county": "Stearns",
		"state": "MN",
		"zipcode": "56303"
	  }
	},
	"28286": {
	  "id": "28286",
	  "fname": "Mohan",
	  "lname": "Kuzuoka",
	  "sex": "F",
	  "dob": "1956-12-11T00:00:00Z",
	  "hireDate": "1996-08-20T00:00:00Z",
	  "position": "Staff",
	  "salary": 58707,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "445 S Figueroa St",
		"city": "Los Angeles",
		"county": "Los Angeles",
		"state": "CA",
		"zipcode": "90071"
	  }
	},
	"28287": {
	  "id": "28287",
	  "fname": "Maria",
	  "lname": "Varker",
	  "sex": "F",
	  "dob": "1960-07-01T00:00:00Z",
	  "hireDate": "1988-04-26T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 64408,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "Headquarters",
		"city": "Morristown",
		"county": "Morris",
		"state": "NJ",
		"zipcode": "7960"
	  }
	},
	"28288": {
	  "id": "28288",
	  "fname": "Martijn",
	  "lname": "Niizuma",
	  "sex": "F",
	  "dob": "1953-10-04T00:00:00Z",
	  "hireDate": "1985-04-13T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d009", "name": "Customer Service", "mgrId": "111939" },
	  "address": {
		"street": "11906 Woodberry Pl",
		"city": "Kingsville",
		"county": "Baltimore",
		"state": "MD",
		"zipcode": "21087"
	  }
	},
	"28289": {
	  "id": "28289",
	  "fname": "Ebbe",
	  "lname": "Herber",
	  "sex": "M",
	  "dob": "1960-02-25T00:00:00Z",
	  "hireDate": "1992-12-13T00:00:00Z",
	  "position": "Technique Leader",
	  "salary": 56772,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "6 Gwynns Mill Ct",
		"city": "Owings Mills",
		"county": "Baltimore",
		"state": "MD",
		"zipcode": "21117"
	  }
	},
	"28290": {
	  "id": "28290",
	  "fname": "Hyuncheol",
	  "lname": "Ashish",
	  "sex": "F",
	  "dob": "1960-12-31T00:00:00Z",
	  "hireDate": "1985-09-01T00:00:00Z",
	  "position": "Engineer",
	  "salary": 43772,
	  "dept": { "id": "d009", "name": "Customer Service", "mgrId": "111939" },
	  "address": {
		"street": "4909 Lakawana St",
		"city": "Dallas",
		"county": "Dallas",
		"state": "TX",
		"zipcode": "75247"
	  }
	},
	"28291": {
	  "id": "28291",
	  "fname": "Kankanahalli",
	  "lname": "Montemayor",
	  "sex": "F",
	  "dob": "1959-06-14T00:00:00Z",
	  "hireDate": "1987-01-02T00:00:00Z",
	  "position": "Senior Staff",
	  "salary": 65535,
	  "dept": { "id": "d009", "name": "Customer Service", "mgrId": "111939" },
	  "address": {
		"street": "7513 E Highway 86",
		"city": "Franktown",
		"county": "Douglas",
		"state": "CO",
		"zipcode": "80116"
	  }
	},
	"28292": {
	  "id": "28292",
	  "fname": "Zeljko",
	  "lname": "Wynblatt",
	  "sex": "M",
	  "dob": "1953-05-02T00:00:00Z",
	  "hireDate": "1994-02-22T00:00:00Z",
	  "position": "Engineer",
	  "salary": 59962,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "8500 Frankstown Rd",
		"city": "Pittsburgh",
		"county": "Allegheny",
		"state": "PA",
		"zipcode": "15235"
	  }
	},
	"28293": {
	  "id": "28293",
	  "fname": "Dipankar",
	  "lname": "Ladret",
	  "sex": "F",
	  "dob": "1963-04-11T00:00:00Z",
	  "hireDate": "1993-05-11T00:00:00Z",
	  "position": "Staff",
	  "salary": 41153,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "870 Airport Rd",
		"city": "Chapel Hill",
		"county": "Orange",
		"state": "NC",
		"zipcode": "27514"
	  }
	},
	"28294": {
	  "id": "28294",
	  "fname": "Ebru",
	  "lname": "Candan",
	  "sex": "M",
	  "dob": "1960-08-10T00:00:00Z",
	  "hireDate": "1992-01-24T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "1231 Market St",
		"city": "San Francisco",
		"county": "San Francisco",
		"state": "CA",
		"zipcode": "94103"
	  }
	},
	"28295": {
	  "id": "28295",
	  "fname": "Babette",
	  "lname": "Luef",
	  "sex": "M",
	  "dob": "1959-07-28T00:00:00Z",
	  "hireDate": "1988-02-06T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 64466,
	  "dept": { "id": "d006", "name": "Quality Management", "mgrId": "110854" },
	  "address": {
		"street": "1322 E Shaw Ave",
		"city": "Fresno",
		"county": "Fresno",
		"state": "CA",
		"zipcode": "93710"
	  }
	},
	"28296": {
	  "id": "28296",
	  "fname": "Zengping",
	  "lname": "Matzke",
	  "sex": "F",
	  "dob": "1958-11-29T00:00:00Z",
	  "hireDate": "1991-01-06T00:00:00Z",
	  "position": "Staff",
	  "salary": 64649,
	  "dept": { "id": "d008", "name": "Research", "mgrId": "111534" },
	  "address": {
		"street": "Panama City",
		"city": "Panama City",
		"county": "Bay",
		"state": "FL",
		"zipcode": " 32401"
	  }
	},
	"28297": {
	  "id": "28297",
	  "fname": "Salvador",
	  "lname": "Mahnke",
	  "sex": "F",
	  "dob": "1964-12-11T00:00:00Z",
	  "hireDate": "1992-06-06T00:00:00Z",
	  "position": "Engineer",
	  "salary": 64300,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "560 Work St",
		"city": "Salinas",
		"county": "Monterey",
		"state": "CA",
		"zipcode": "93901"
	  }
	},
	"28298": {
	  "id": "28298",
	  "fname": "Holgard",
	  "lname": "Ritcey",
	  "sex": "F",
	  "dob": "1952-06-04T00:00:00Z",
	  "hireDate": "1988-07-03T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "Crossroads",
		"city": "Hartwell",
		"county": "Hart",
		"state": "GA",
		"zipcode": "30643"
	  }
	},
	"28299": {
	  "id": "28299",
	  "fname": "Tse",
	  "lname": "Matheson",
	  "sex": "M",
	  "dob": "1962-08-31T00:00:00Z",
	  "hireDate": "1996-08-30T00:00:00Z",
	  "position": "Engineer",
	  "salary": 49262,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "2735 E Southern Ave",
		"city": "Phoenix",
		"county": "Maricopa",
		"state": "AZ",
		"zipcode": "85040"
	  }
	},
	"296098": {
	  "id": "296098",
	  "fname": "Hidekazu",
	  "lname": "Rosis",
	  "sex": "F",
	  "dob": "1955-10-02T00:00:00Z",
	  "hireDate": "1990-07-07T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 50631,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "8000 Manchester Rd",
		"city": "Saint Louis",
		"county": "Saint Louis",
		"state": "MO",
		"zipcode": "63144"
	  }
	},
	"296099": {
	  "id": "296099",
	  "fname": "Constantine",
	  "lname": "Byoun",
	  "sex": "F",
	  "dob": "1963-06-22T00:00:00Z",
	  "hireDate": "1990-09-01T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 61673,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "300 E 42nd St",
		"city": "NewYork",
		"county": "New York",
		"state": "NY",
		"zipcode": "10017"
	  }
	},
	"86220": {
	  "id": "86220",
	  "fname": "Nidapan",
	  "lname": "Danecki",
	  "sex": "M",
	  "dob": "1953-01-20T00:00:00Z",
	  "hireDate": "1989-09-07T00:00:00Z",
	  "position": "Assistant Engineer",
	  "salary": 58006,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "349 E Main St",
		"city": "Elmsford",
		"county": "Westchester",
		"state": "NY",
		"zipcode": "10523"
	  }
	},
	"86221": {
	  "id": "86221",
	  "fname": "Mohan",
	  "lname": "Decleir",
	  "sex": "M",
	  "dob": "1957-11-19T00:00:00Z",
	  "hireDate": "1992-05-16T00:00:00Z",
	  "position": "Engineer",
	  "salary": 43599,
	  "dept": { "id": "d008", "name": "Research", "mgrId": "111534" },
	  "address": {
		"street": "925 E 85th St",
		"city": "Kansas City",
		"county": "Jackson",
		"state": "MO",
		"zipcode": "64131"
	  }
	},
	"86222": {
	  "id": "86222",
	  "fname": "Alexius",
	  "lname": "Marquardt",
	  "sex": "F",
	  "dob": "1954-12-21T00:00:00Z",
	  "hireDate": "1985-02-16T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d001", "name": "Marketing", "mgrId": "110039" },
	  "address": {
		"street": "320 N State St",
		"city": "Harrison",
		"county": "Hamilton",
		"state": "OH",
		"zipcode": "45030"
	  }
	},
	"86223": {
	  "id": "86223",
	  "fname": "Kaijung",
	  "lname": "Peroz",
	  "sex": "M",
	  "dob": "1964-07-03T00:00:00Z",
	  "hireDate": "1995-04-29T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 49902,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "2712 Lawndale Dr",
		"city": "Greensboro",
		"county": "Guilford",
		"state": "NC",
		"zipcode": "27408"
	  }
	},
	"86224": {
	  "id": "86224",
	  "fname": "Shaunak",
	  "lname": "Ishibashi",
	  "sex": "M",
	  "dob": "1954-02-05T00:00:00Z",
	  "hireDate": "1985-04-04T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d008", "name": "Research", "mgrId": "111534" },
	  "address": {
		"street": "219 Taunton Ave",
		"city": "East Providence",
		"county": "Providence",
		"state": "RI",
		"zipcode": "2914"
	  }
	},
	"86225": {
	  "id": "86225",
	  "fname": "Goncalo",
	  "lname": "Serdy",
	  "sex": "F",
	  "dob": "1953-03-14T00:00:00Z",
	  "hireDate": "1991-02-14T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "485 22nd Pl",
		"city": "Vero Beach",
		"county": "Indian River",
		"state": "FL",
		"zipcode": "32960"
	  }
	},
	"86226": {
	  "id": "86226",
	  "fname": "Collette",
	  "lname": "Fargier",
	  "sex": "F",
	  "dob": "1959-08-07T00:00:00Z",
	  "hireDate": "1985-07-06T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "3095 Bay Settlement Rd",
		"city": "Green Bay",
		"county": "Brown",
		"state": "WI",
		"zipcode": "54311"
	  }
	},
	"86227": {
	  "id": "86227",
	  "fname": "Mohammad",
	  "lname": "Pillow",
	  "sex": "F",
	  "dob": "1956-03-17T00:00:00Z",
	  "hireDate": "1995-04-04T00:00:00Z",
	  "position": "Technique Leader",
	  "salary": 47430,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "19333 E Grant Hwy",
		"city": "Marengo",
		"county": "McHenry",
		"state": "IL",
		"zipcode": "60152"
	  }
	},
	"86228": {
	  "id": "86228",
	  "fname": "Lansing",
	  "lname": "Rosaz",
	  "sex": "F",
	  "dob": "1958-10-20T00:00:00Z",
	  "hireDate": "1997-06-15T00:00:00Z",
	  "position": "Engineer",
	  "salary": 53344,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "2401 Columbia Pike",
		"city": "Arlington",
		"county": "Arlington",
		"state": "VA",
		"zipcode": "22204"
	  }
	},
	"86229": {
	  "id": "86229",
	  "fname": "Visit",
	  "lname": "Przulj",
	  "sex": "M",
	  "dob": "1961-10-18T00:00:00Z",
	  "hireDate": "1989-07-15T00:00:00Z",
	  "position": "Senior Staff",
	  "salary": 65535,
	  "dept": { "id": "d002", "name": "Finance", "mgrId": "110114" },
	  "address": {
		"street": "529 14th St Nw",
		"city": "Washington",
		"county": "District of Columbia",
		"state": "DC",
		"zipcode": "20045"
	  }
	},
	"86250": {
	  "id": "86250",
	  "fname": "Nakhoon",
	  "lname": "Weedman",
	  "sex": "F",
	  "dob": "1963-12-02T00:00:00Z",
	  "hireDate": "1985-12-26T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "800 Lasalle Ave",
		"city": "Minneapolis",
		"county": "Hennepin",
		"state": "MN",
		"zipcode": "55402"
	  }
	},
	"86251": {
	  "id": "86251",
	  "fname": "Raimond",
	  "lname": "Malinowski",
	  "sex": "M",
	  "dob": "1958-03-22T00:00:00Z",
	  "hireDate": "1989-08-02T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 48666,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "1785 S Johnson Rd",
		"city": " New Berlin",
		"county": "Waukesha",
		"state": "WI",
		"zipcode": "53146"
	  }
	},
	"86252": {
	  "id": "86252",
	  "fname": "Mahendra",
	  "lname": "Selenyi",
	  "sex": "F",
	  "dob": "1955-08-14T00:00:00Z",
	  "hireDate": "1988-05-23T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "1505 Ford Rd",
		"city": "Bensalem",
		"county": "Bucks",
		"state": "PA",
		"zipcode": "19020"
	  }
	},
	"86253": {
	  "id": "86253",
	  "fname": "Vidya",
	  "lname": "Zhang",
	  "sex": "M",
	  "dob": "1958-07-25T00:00:00Z",
	  "hireDate": "1988-02-12T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "14 W Second St",
		"city": "Media",
		"county": "Delaware",
		"state": "PA",
		"zipcode": "19063"
	  }
	},
	"86254": {
	  "id": "86254",
	  "fname": "Moti",
	  "lname": "Servi",
	  "sex": "M",
	  "dob": "1964-01-10T00:00:00Z",
	  "hireDate": "1987-07-18T00:00:00Z",
	  "position": "Engineer",
	  "salary": 65535,
	  "dept": { "id": "d006", "name": "Quality Management", "mgrId": "110854" },
	  "address": {
		"street": "3201 W Le Fevre Rd",
		"city": "Sterling",
		"county": "Whiteside",
		"state": "IL",
		"zipcode": "61081"
	  }
	},
	"86255": {
	  "id": "86255",
	  "fname": "Kyoichi",
	  "lname": "Zschoche",
	  "sex": "M",
	  "dob": "1956-03-19T00:00:00Z",
	  "hireDate": "1985-06-02T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "2000 Tate Springs Rd",
		"city": " Lynchburg",
		"county": "Lynchburg City",
		"state": "VA",
		"zipcode": "24501"
	  }
	},
	"86256": {
	  "id": "86256",
	  "fname": "Gift",
	  "lname": "Dulay",
	  "sex": "M",
	  "dob": "1955-06-16T00:00:00Z",
	  "hireDate": "1990-07-07T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d009", "name": "Customer Service", "mgrId": "111939" },
	  "address": {
		"street": "1818 N St Nw #-200",
		"city": "Washington",
		"county": " District of Columbia",
		"state": "DC",
		"zipcode": "20036"
	  }
	},
	"86257": {
	  "id": "86257",
	  "fname": "Niteen",
	  "lname": "Klyachko",
	  "sex": "F",
	  "dob": "1961-03-06T00:00:00Z",
	  "hireDate": "1985-11-20T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "111 Sw 5th Ave",
		"city": "Portland",
		"county": "Multnomah",
		"state": " OR",
		"zipcode": "97204"
	  }
	},
	"86258": {
	  "id": "86258",
	  "fname": "Shai",
	  "lname": "Demke",
	  "sex": "M",
	  "dob": "1960-03-02T00:00:00Z",
	  "hireDate": "1993-11-25T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d002", "name": "Finance", "mgrId": "110114" },
	  "address": {
		"street": "205 5th Ave S #-621",
		"city": "La Crosse",
		"county": "La Crosse",
		"state": "WI",
		"zipcode": "54601"
	  }
	},
	"86259": {
	  "id": "86259",
	  "fname": "Yongmin",
	  "lname": "Atchley",
	  "sex": "M",
	  "dob": "1962-03-31T00:00:00Z",
	  "hireDate": "1989-12-04T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "720 Olive St #-2100",
		"city": "Saint Louis",
		"county": "Saint Louis City",
		"state": "MO",
		"zipcode": "63101"
	  }
	},
	"86260": {
	  "id": "86260",
	  "fname": "Aruna",
	  "lname": "Hartvigsen",
	  "sex": "M",
	  "dob": "1964-12-31T00:00:00Z",
	  "hireDate": "1993-06-06T00:00:00Z",
	  "position": "Engineer",
	  "salary": 65535,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "2198 Filbert St",
		"city": "San Francisco",
		"county": "San Francisco",
		"state": "CA",
		"zipcode": "94123"
	  }
	},
	"86261": {
	  "id": "86261",
	  "fname": "Cristinel",
	  "lname": "Thibadeau",
	  "sex": "M",
	  "dob": "1957-04-02T00:00:00Z",
	  "hireDate": "1992-02-28T00:00:00Z",
	  "position": "Staff",
	  "salary": 50218,
	  "dept": { "id": "d002", "name": "Finance", "mgrId": "110114" },
	  "address": {
		"street": "1180 Quail Valley Run",
		"city": "Oakley",
		"county": "Contra Costa",
		"state": "CA",
		"zipcode": "94561"
	  }
	},
	"86262": {
	  "id": "86262",
	  "fname": "Mana",
	  "lname": "Sooriamurthi",
	  "sex": "M",
	  "dob": "1956-10-12T00:00:00Z",
	  "hireDate": "1986-12-07T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "101 E 52nd St",
		"city": "New York",
		"county": "New York",
		"state": "NY",
		"zipcode": "10022"
	  }
	},
	"86263": {
	  "id": "86263",
	  "fname": "Indrajit",
	  "lname": "Alpin",
	  "sex": "M",
	  "dob": "1963-12-04T00:00:00Z",
	  "hireDate": "1992-04-23T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "17000 E State Highway 120",
		"city": "Ripon",
		"county": "San Joaquin",
		"state": "CA",
		"zipcode": "95366"
	  }
	},
	"86264": {
	  "id": "86264",
	  "fname": "Bouchung",
	  "lname": "Butterworth",
	  "sex": "M",
	  "dob": "1962-10-13T00:00:00Z",
	  "hireDate": "1987-03-05T00:00:00Z",
	  "position": "Staff",
	  "salary": 65084,
	  "dept": { "id": "d009", "name": "Customer Service", "mgrId": "111939" },
	  "address": {
		"street": "240 Stockton St",
		"city": "SanFrancisco",
		"county": "San Francisco",
		"state": "CA",
		"zipcode": "94108"
	  }
	},
	"86265": {
	  "id": "86265",
	  "fname": "Georg",
	  "lname": "Peres",
	  "sex": "M",
	  "dob": "1955-12-06T00:00:00Z",
	  "hireDate": "1987-09-18T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d009", "name": "Customer Service", "mgrId": "111939" },
	  "address": {
		"street": "2916 S Reed Ave",
		"city": "Sanger",
		"county": "Fresno",
		"state": "CA",
		"zipcode": "93657"
	  }
	},
	"86266": {
	  "id": "86266",
	  "fname": "Danco",
	  "lname": "Hofman",
	  "sex": "M",
	  "dob": "1956-10-07T00:00:00Z",
	  "hireDate": "1987-01-03T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d003", "name": "Human Resources", "mgrId": "110228" },
	  "address": {
		"street": "25 Main St En",
		"city": "Rochester",
		"county": "Monroe",
		"state": "NY",
		"zipcode": "14614"
	  }
	},
	"86267": {
	  "id": "86267",
	  "fname": "Masato",
	  "lname": "Guting",
	  "sex": "F",
	  "dob": "1953-05-22T00:00:00Z",
	  "hireDate": "1987-08-27T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "324 Broadway",
		"city": "Elizabeth",
		"county": "Union",
		"state": "NJ",
		"zipcode": "7206"
	  }
	},
	"86268": {
	  "id": "86268",
	  "fname": "Angel",
	  "lname": "McDermid",
	  "sex": "M",
	  "dob": "1956-12-31T00:00:00Z",
	  "hireDate": "1987-11-10T00:00:00Z",
	  "position": "Engineer",
	  "salary": 48309,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "21 Griffin Ln",
		"city": "Hopewell Junction",
		"county": "Dutchess",
		"state": "NY",
		"zipcode": "12533"
	  }
	},
	"86269": {
	  "id": "86269",
	  "fname": "Kristinn",
	  "lname": "Swift",
	  "sex": "M",
	  "dob": "1953-04-18T00:00:00Z",
	  "hireDate": "1986-12-08T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "445 E Main St #-100",
		"city": "Hillsboro",
		"county": "Washington",
		"state": "OR",
		"zipcode": "97123"
	  }
	},
	"86270": {
	  "id": "86270",
	  "fname": "Prodip",
	  "lname": "Menhoudj",
	  "sex": "F",
	  "dob": "1954-02-01T00:00:00Z",
	  "hireDate": "1985-08-30T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d002", "name": "Finance", "mgrId": "110114" },
	  "address": {
		"street": "837 N Quincy St",
		"city": "Arlington",
		"county": "Arlington",
		"state": "VA",
		"zipcode": "22203"
	  }
	},
	"86271": {
	  "id": "86271",
	  "fname": "Ipke",
	  "lname": "Llado",
	  "sex": "M",
	  "dob": "1954-04-23T00:00:00Z",
	  "hireDate": "1988-04-06T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d009", "name": "Customer Service", "mgrId": "111939" },
	  "address": {
		"street": "126 W Forest Grove Ave",
		"city": "Phoenix",
		"county": "Maricopa",
		"state": "AZ",
		"zipcode": "85041"
	  }
	},
	"86272": {
	  "id": "86272",
	  "fname": "Moriyoshi",
	  "lname": "Gustavson",
	  "sex": "M",
	  "dob": "1953-05-11T00:00:00Z",
	  "hireDate": "1992-02-29T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d002", "name": "Finance", "mgrId": "110114" },
	  "address": {
		"street": "325 Orient Way",
		"city": "Lyndhurst",
		"county": "Bergen",
		"state": "NJ",
		"zipcode": "7071"
	  }
	},
	"86273": {
	  "id": "86273",
	  "fname": "Tsuneo",
	  "lname": "Pramanik",
	  "sex": "M",
	  "dob": "1964-02-12T00:00:00Z",
	  "hireDate": "1988-09-24T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "38425 15th St E",
		"city": "Palmdale",
		"county": "Los Angeles",
		"state": "CA",
		"zipcode": "93550"
	  }
	},
	"86274": {
	  "id": "86274",
	  "fname": "Gila",
	  "lname": "Valtorta",
	  "sex": "M",
	  "dob": "1963-01-04T00:00:00Z",
	  "hireDate": "1990-03-25T00:00:00Z",
	  "position": "Senior Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "Rr 1 #-430g",
		"city": "Louisa",
		"county": "Louisa",
		"state": "VA",
		"zipcode": " 23093"
	  }
	},
	"86275": {
	  "id": "86275",
	  "fname": "Ipke",
	  "lname": "Genin",
	  "sex": "M",
	  "dob": "1953-04-13T00:00:00Z",
	  "hireDate": "1987-06-20T00:00:00Z",
	  "position": "Engineer",
	  "salary": 60492,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "3311 Stanford Dr Ne",
		"city": "Albuquerque",
		"county": "Bernalillo",
		"state": "NM",
		"zipcode": "87107"
	  }
	},
	"86276": {
	  "id": "86276",
	  "fname": "Nirmal",
	  "lname": "Ratzlaff",
	  "sex": "M",
	  "dob": "1956-08-02T00:00:00Z",
	  "hireDate": "1992-04-23T00:00:00Z",
	  "position": "Engineer",
	  "salary": 52722,
	  "dept": { "id": "d006", "name": "Quality Management", "mgrId": "110854" },
	  "address": {
		"street": "15616 W Main St",
		"city": "CutOff",
		"county": "Lafourche",
		"state": "LA",
		"zipcode": "70345"
	  }
	},
	"86277": {
	  "id": "86277",
	  "fname": "Nakhoon",
	  "lname": "Bodoff",
	  "sex": "F",
	  "dob": "1958-06-11T00:00:00Z",
	  "hireDate": "1994-08-07T00:00:00Z",
	  "position": "Technique Leader",
	  "salary": 52492,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "1791 Victory Blvd",
		"city": "Staten Island",
		"county": "Richmond",
		"state": "NY",
		"zipcode": "10314"
	  }
	},
	"86278": {
	  "id": "86278",
	  "fname": "Lobel",
	  "lname": "Matzel",
	  "sex": "M",
	  "dob": "1959-03-20T00:00:00Z",
	  "hireDate": "1996-08-06T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "2065 Bird St",
		"city": "Oroville",
		"county": "Butte",
		"state": "CA",
		"zipcode": "95965"
	  }
	},
	"86279": {
	  "id": "86279",
	  "fname": "Bojan",
	  "lname": "Biran",
	  "sex": "F",
	  "dob": "1958-04-09T00:00:00Z",
	  "hireDate": "1988-10-17T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "39023 Harper Ave",
		"city": "ClintonTownship",
		"county": "Macomb",
		"state": "MI",
		"zipcode": "48036"
	  }
	},
	"86280": {
	  "id": "86280",
	  "fname": "Mads",
	  "lname": "Heyers",
	  "sex": "M",
	  "dob": "1963-10-18T00:00:00Z",
	  "hireDate": "1988-12-03T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 62249,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "1300 E Pike St",
		"city": "Seattle",
		"county": "King",
		"state": "WA",
		"zipcode": "98122"
	  }
	},
	"86281": {
	  "id": "86281",
	  "fname": "Srinidhi",
	  "lname": "Conde",
	  "sex": "M",
	  "dob": "1954-02-12T00:00:00Z",
	  "hireDate": "1990-03-06T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 59566,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "5561 S Lewis",
		"city": "Tulsa",
		"county": "Tulsa",
		"state": "OK",
		"zipcode": "74105"
	  }
	},
	"86282": {
	  "id": "86282",
	  "fname": "Martien",
	  "lname": "Baak",
	  "sex": "M",
	  "dob": "1954-05-23T00:00:00Z",
	  "hireDate": "1986-07-24T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d001", "name": "Marketing", "mgrId": "110039" },
	  "address": {
		"street": "16525 W Glendale Dr",
		"city": "New Berlin",
		"county": "Waukesha",
		"state": "WI",
		"zipcode": "53151"
	  }
	},
	"86283": {
	  "id": "86283",
	  "fname": "Xinan",
	  "lname": "Parveen",
	  "sex": "F",
	  "dob": "1959-02-24T00:00:00Z",
	  "hireDate": "1985-11-11T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d002", "name": "Finance", "mgrId": "110114" },
	  "address": {
		"street": "190 S Main Rd",
		"city": "Vineland",
		"county": "Cumberland",
		"state": "NJ",
		"zipcode": "8360"
	  }
	},
	"86284": {
	  "id": "86284",
	  "fname": "Fabrizio",
	  "lname": "Wrigley",
	  "sex": "M",
	  "dob": "1954-01-16T00:00:00Z",
	  "hireDate": "1988-10-30T00:00:00Z",
	  "position": "Engineer",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "700 N Olive Ave",
		"city": "West PalmBeach",
		"county": "Palm Beach",
		"state": "FL",
		"zipcode": "33401"
	  }
	},
	"86285": {
	  "id": "86285",
	  "fname": "Moon",
	  "lname": "Figueira",
	  "sex": "M",
	  "dob": "1958-08-07T00:00:00Z",
	  "hireDate": "1985-04-03T00:00:00Z",
	  "position": "Engineer",
	  "salary": 53977,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "1001 Marshall St",
		"city": "Redwood City",
		"county": " San Mateo",
		"state": "CA",
		"zipcode": "94063"
	  }
	},
	"86286": {
	  "id": "86286",
	  "fname": "Somnath",
	  "lname": "Esposito",
	  "sex": "M",
	  "dob": "1954-06-17T00:00:00Z",
	  "hireDate": "1989-06-06T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "111 N Houston St",
		"city": "Fort Worth",
		"county": "Tarrant",
		"state": "TX",
		"zipcode": "76102"
	  }
	},
	"86287": {
	  "id": "86287",
	  "fname": "Masami",
	  "lname": "Liesche",
	  "sex": "F",
	  "dob": "1954-05-03T00:00:00Z",
	  "hireDate": "1987-05-11T00:00:00Z",
	  "position": "Senior Staff",
	  "salary": 65535,
	  "dept": { "id": "d002", "name": "Finance", "mgrId": "110114" },
	  "address": {
		"street": "1700 Southwest Blvd",
		"city": "Tulsa",
		"county": "Tulsa",
		"state": "OK",
		"zipcode": "74107"
	  }
	},
	"86288": {
	  "id": "86288",
	  "fname": "Ayonca",
	  "lname": "Hellwagner",
	  "sex": "F",
	  "dob": "1958-07-29T00:00:00Z",
	  "hireDate": "1989-12-11T00:00:00Z",
	  "position": "Engineer",
	  "salary": 65535,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "8 Woodmont Crossing St",
		"city": "Texarkana",
		"county": "Bowie",
		"state": "TX",
		"zipcode": "75503"
	  }
	},
	"86289": {
	  "id": "86289",
	  "fname": "Patricio",
	  "lname": "Falby",
	  "sex": "F",
	  "dob": "1960-03-30T00:00:00Z",
	  "hireDate": "1985-10-30T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "1914 Plymouth St",
		"city": "Mountain View",
		"county": "Santa Clara",
		"state": "CA",
		"zipcode": "94043"
	  }
	},
	"86290": {
	  "id": "86290",
	  "fname": "Alselm",
	  "lname": "Willoner",
	  "sex": "F",
	  "dob": "1952-05-24T00:00:00Z",
	  "hireDate": "1991-07-19T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65163,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "111 E Washington St",
		"city": "Walterboro",
		"county": "Colleton",
		"state": "SC",
		"zipcode": "29488"
	  }
	},
	"86291": {
	  "id": "86291",
	  "fname": "Kerhong",
	  "lname": "Bashian",
	  "sex": "M",
	  "dob": "1963-11-03T00:00:00Z",
	  "hireDate": "1988-07-15T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "411 S Wells St",
		"city": "Chicago",
		"county": "Cook",
		"state": "IL",
		"zipcode": "60607"
	  }
	},
	"86292": {
	  "id": "86292",
	  "fname": "Shuho",
	  "lname": "Vural",
	  "sex": "F",
	  "dob": "1954-12-17T00:00:00Z",
	  "hireDate": "1993-07-26T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d001", "name": "Marketing", "mgrId": "110039" },
	  "address": {
		"street": "Clough Ave",
		"city": "Superior",
		"county": "Douglas",
		"state": "WI",
		"zipcode": "54880"
	  }
	},
	"86293": {
	  "id": "86293",
	  "fname": "Fatemeh",
	  "lname": "Heering",
	  "sex": "F",
	  "dob": "1958-07-04T00:00:00Z",
	  "hireDate": "1995-12-18T00:00:00Z",
	  "position": "Engineer",
	  "salary": 45637,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "25 Bridge St",
		"city": "Madawaska",
		"county": "Aroostook",
		"state": "ME",
		"zipcode": "4756"
	  }
	},
	"86294": {
	  "id": "86294",
	  "fname": "Heather",
	  "lname": "Jumpertz",
	  "sex": "F",
	  "dob": "1962-07-23T00:00:00Z",
	  "hireDate": "1990-08-03T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "2 Embarcadero Ctr",
		"city": "San Francisco",
		"county": "San Francisco",
		"state": "CA",
		"zipcode": "94111"
	  }
	},
	"86295": {
	  "id": "86295",
	  "fname": "Cheong",
	  "lname": "Hofstetter",
	  "sex": "M",
	  "dob": "1956-11-08T00:00:00Z",
	  "hireDate": "1986-03-28T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "1025 Thomas Dr",
		"city": "Warminster",
		"county": "Bucks",
		"state": "PA",
		"zipcode": "18974"
	  }
	},
	"86296": {
	  "id": "86296",
	  "fname": "Koldo",
	  "lname": "Monkewich",
	  "sex": "F",
	  "dob": "1956-11-13T00:00:00Z",
	  "hireDate": "1988-09-08T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 61263,
	  "dept": { "id": "d008", "name": "Research", "mgrId": "111534" },
	  "address": {
		"street": "100 E Wisconsin Ave",
		"city": "Milwaukee",
		"county": "Milwaukee",
		"state": "WI",
		"zipcode": "53202"
	  }
	},
	"86297": {
	  "id": "86297",
	  "fname": "Susuma",
	  "lname": "Feldmann",
	  "sex": "F",
	  "dob": "1952-12-15T00:00:00Z",
	  "hireDate": "1985-03-31T00:00:00Z",
	  "position": "Technique Leader",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "409 S 17th St",
		"city": "Omaha",
		"county": "Douglas",
		"state": "NE",
		"zipcode": "68102"
	  }
	},
	"86298": {
	  "id": "86298",
	  "fname": "Ziva",
	  "lname": "Farrag",
	  "sex": "F",
	  "dob": "1960-09-03T00:00:00Z",
	  "hireDate": "1994-07-14T00:00:00Z",
	  "position": "Engineer",
	  "salary": 57923,
	  "dept": { "id": "d006", "name": "Quality Management", "mgrId": "110854" },
	  "address": {
		"street": "1000 Michigan National Towe",
		"city": "Lansing",
		"county": "Ingham",
		"state": "MI",
		"zipcode": "48933"
	  }
	},
	"86299": {
	  "id": "86299",
	  "fname": "Emran",
	  "lname": "Brookman",
	  "sex": "M",
	  "dob": "1959-06-08T00:00:00Z",
	  "hireDate": "1987-08-18T00:00:00Z",
	  "position": "Assistant Engineer",
	  "salary": 53922,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "350 Sparta Ave",
		"city": "Sparta",
		"county": "Sussex",
		"state": "NJ",
		"zipcode": "7871"
	  }
	},
	"86330": {
	  "id": "86330",
	  "fname": "Tesuya",
	  "lname": "Shumilov",
	  "sex": "F",
	  "dob": "1962-05-28T00:00:00Z",
	  "hireDate": "1990-05-16T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "9 E Ct",
		"city": "Franklin",
		"county": "Johnson",
		"state": "IN",
		"zipcode": "46131"
	  }
	},
	"86331": {
	  "id": "86331",
	  "fname": "Visit",
	  "lname": "Ramamoorthy",
	  "sex": "F",
	  "dob": "1961-09-20T00:00:00Z",
	  "hireDate": "1988-09-27T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "113 S Lithia Pinecrest Rd",
		"city": "Brandon",
		"county": "Hillsborough",
		"state": "FL",
		"zipcode": "33511"
	  }
	},
	"86332": {
	  "id": "86332",
	  "fname": "Arto",
	  "lname": "Perry",
	  "sex": "M",
	  "dob": "1961-02-18T00:00:00Z",
	  "hireDate": "1992-01-08T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d009", "name": "Customer Service", "mgrId": "111939" },
	  "address": {
		"street": "2019 Linnhurst Dr",
		"city": "Savannah",
		"county": "Chatham",
		"state": "GA",
		"zipcode": "31404"
	  }
	},
	"86333": {
	  "id": "86333",
	  "fname": "Ennio",
	  "lname": "Peak",
	  "sex": "M",
	  "dob": "1958-07-10T00:00:00Z",
	  "hireDate": "1987-07-22T00:00:00Z",
	  "position": "Staff",
	  "salary": 56856,
	  "dept": { "id": "d003", "name": "Human Resources", "mgrId": "110228" },
	  "address": {
		"street": "3605 Franklin Ave",
		"city": "Waco",
		"county": "McLennan",
		"state": "TX",
		"zipcode": "76710"
	  }
	},
	"86334": {
	  "id": "86334",
	  "fname": "Kwun",
	  "lname": "Cappelletti",
	  "sex": "M",
	  "dob": "1964-12-20T00:00:00Z",
	  "hireDate": "1985-09-09T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "2155 Webster St",
		"city": "San Francisco",
		"county": "San Francisco",
		"state": "CA",
		"zipcode": "94115"
	  }
	},
	"86335": {
	  "id": "86335",
	  "fname": "Xinan",
	  "lname": "Bierbaum",
	  "sex": "M",
	  "dob": "1964-09-26T00:00:00Z",
	  "hireDate": "1989-04-15T00:00:00Z",
	  "position": "Staff",
	  "salary": 65535,
	  "dept": { "id": "d002", "name": "Finance", "mgrId": "110114" },
	  "address": {
		"street": "5500 Wayzata Blvd #-730",
		"city": "Minneapolis",
		"county": "Hennepin",
		"state": "MN",
		"zipcode": "55416"
	  }
	},
	"86336": {
	  "id": "86336",
	  "fname": "Mohit",
	  "lname": "Prochazka",
	  "sex": "M",
	  "dob": "1964-11-08T00:00:00Z",
	  "hireDate": "1987-12-15T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d004", "name": "Production", "mgrId": "110420" },
	  "address": {
		"street": "100 Edwards St",
		"city": "Shreveport",
		"county": "Caddo",
		"state": "LA",
		"zipcode": "71101"
	  }
	},
	"86337": {
	  "id": "86337",
	  "fname": "Rafols",
	  "lname": "Munke",
	  "sex": "F",
	  "dob": "1962-03-29T00:00:00Z",
	  "hireDate": "1996-08-16T00:00:00Z",
	  "position": "Senior Staff",
	  "salary": 65535,
	  "dept": { "id": "d007", "name": "Sales", "mgrId": "111133" },
	  "address": {
		"street": "2163 Newcastle Ave #-200",
		"city": "Cardiff by the Sea",
		"county": "San Diego",
		"state": "CA",
		"zipcode": "92007"
	  }
	},
	"86338": {
	  "id": "86338",
	  "fname": "Mooi",
	  "lname": "Lunt",
	  "sex": "M",
	  "dob": "1958-04-14T00:00:00Z",
	  "hireDate": "1989-12-23T00:00:00Z",
	  "position": "Senior Staff",
	  "salary": 60323,
	  "dept": { "id": "d008", "name": "Research", "mgrId": "111534" },
	  "address": {
		"street": "14701 E 38th Ave",
		"city": "Aurora",
		"county": "Adams",
		"state": "CO",
		"zipcode": "80011"
	  }
	},
	"86339": {
	  "id": "86339",
	  "fname": "Raimond",
	  "lname": "Debuse",
	  "sex": "F",
	  "dob": "1955-04-09T00:00:00Z",
	  "hireDate": "1989-05-16T00:00:00Z",
	  "position": "Staff",
	  "salary": 60428,
	  "dept": { "id": "d008", "name": "Research", "mgrId": "111534" },
	  "address": {
		"street": "485 Us Highway 1",
		"city": "Iselin",
		"county": "Middlesex",
		"state": "NJ",
		"zipcode": "8830"
	  }
	},
	"86341": {
	  "id": "86341",
	  "fname": "Angus",
	  "lname": "Zuberek",
	  "sex": "F",
	  "dob": "1953-12-18T00:00:00Z",
	  "hireDate": "1996-06-17T00:00:00Z",
	  "position": "Technique Leader",
	  "salary": 47172,
	  "dept": { "id": "d006", "name": "Quality Management", "mgrId": "110854" },
	  "address": {
		"street": "1515 Broadway",
		"city": "New York",
		"county": "New York",
		"state": "NY",
		"zipcode": "10036"
	  }
	},
	"86342": {
	  "id": "86342",
	  "fname": "Marit",
	  "lname": "Boguraev",
	  "sex": "M",
	  "dob": "1952-12-14T00:00:00Z",
	  "hireDate": "1985-07-02T00:00:00Z",
	  "position": "Senior Engineer",
	  "salary": 65535,
	  "dept": { "id": "d005", "name": "Development", "mgrId": "110567" },
	  "address": {
		"street": "5616 Urbana Pike",
		"city": "Frederick",
		"county": "Frederick",
		"state": "MD",
		"zipcode": "21704"
	  }
	}
  }
  `

func loadMockData() error {
	e := employees{}
	err := json.Unmarshal([]byte(mockData), &e)

	if err != nil {
		utils.Logger.WithFields(Log.Fields{"error": err.Error()}).Debug("")
	}
	utils.Logger.WithFields(Log.Fields{"length": len(e)}).Debug("Employee map length")

	if err == nil {
		l.m.Lock()
		defer l.m.Unlock()

		for k, v := range e {
			l.serviceData[k] = v
		}
	}

	utils.Logger.WithFields(Log.Fields{"length": len(l.serviceData)}).Debug("ServiceData map length")

	return err
}
