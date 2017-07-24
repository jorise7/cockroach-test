
--numbers on cockroach
CREATE TABLE "test"."numbers" (
	"id"		    serial NOT NULL PRIMARY KEY,
	"value"     float NOT NULL,
	"status"	  varchar(50),
	"changed"	  timestamp NOT NULL DEFAULT now(),
	"created"	  timestamp NOT NULL DEFAULT now()
);



--numbers on postgresql
CREATE TABLE "public"."numbers" (
	"id"		    serial NOT NULL PRIMARY KEY,
	"value"     float NOT NULL,
	"status"	  varchar(50),
	"changed"	  timestamp NOT NULL DEFAULT now(),
	"created"	  timestamp NOT NULL DEFAULT now()
);
