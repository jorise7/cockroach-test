
--numbers
CREATE TABLE "test"."numbers" (
	"id"		    serial NOT NULL PRIMARY KEY,
	"value"     float NOT NULL UNIQUE,
	"status"	  string,
	"changed"	  timestamp NOT NULL DEFAULT now(),
	"created"	  timestamp NOT NULL DEFAULT now()
);
