CREATE TABLE Test.Tester (
    tid int PRIMARY KEY,
    first_name varchar(255),
    last_name varchar(255)
)

CREATE TABLE Test.submissions (
    sid int PRIMARY KEY,
    tid int,
    file_path varchar(255),
    CONSTRAINT fk_tidsub FOREIGN KEY(tid) REFERENCES Tester(tid)
)

CREATE TABLE Test.PlagiarismResult (
    rid int PRIMARY KEY,
    date_time DATETIME,
    sid1 int,
    sid2 int,
    similarity_score float,
    CONSTRAINT fk_sid1 FOREIGN KEY(sid1) REFERENCES Test.submissions(sid),
    CONSTRAINT fk_sid2 FOREIGN KEY(sid2) REFERENCES Test.submissions(sid)
)

INSERT INTO Test.Tester VALUES (60050200, "Harley", "Greene")
INSERT INTO Test.Tester VALUES (60050201, "Tilly", "Frazier")
INSERT INTO Test.Tester VALUES (60050202, "Gloria", "Hull")

INSERT INTO Test.submissions VALUES (5001, 60050200, "~/sms/ExitExam200.zip/")
INSERT INTO Test.submissions VALUES (5002, 60050201, "~/sms/exit_201.rar/")
INSERT INTO Test.submissions VALUES (5003, 60050202, "~/sms/60050202_Exit.7z/")

INSERT INTO Test.PlagiarismResult VALUES (1,"2020-03-28T16:12:24", 5001, 5002, 0.20)
INSERT INTO Test.PlagiarismResult VALUES (2,"2020-03-28T16:18:01", 5001, 5003, 0.19)
INSERT INTO Test.PlagiarismResult VALUES (2,"2020-03-28T16:18:02", 5002, 5003, 0.78)