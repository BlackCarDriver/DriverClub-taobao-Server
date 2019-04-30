--create database website2;

--psql website2;

CREATE TABLE account(
id serial,
uname VARCHAR(21) NOT NULL,
upassword  VARCHAR(31) NOT NULL,
email	VARCHAR(31) DEFAULT NULL,
state 	INT DEFAULT 1,
PRIMARY KEY (id)
);




CREATE TABLE confirm(
userid 	INT,
intimes INT DEFAULT 0,
token 	VARCHAR(51),
PRIMARY KEY (userid),
FOREIGN KEY (userid) REFERENCES account(id) 
ON DELETE CASCADE
ON UPDATE CASCADE
);

CREATE TABLE usermsg1(
userid INT,
username VARCHAR(21),
signing VARCHAR(61),
sex	VARCHAR(5) CHECK(sex='BOY' OR sex='GIRL'),
grade	VARCHAR(5),
colleage VARCHAR(21),
phone 	VARCHAR(13),
qq 	VARCHAR(13),
imgurl 	VARCHAR(100),
PRIMARY KEY (userid),
FOREIGN KEY(userid) REFERENCES account(id)
ON DELETE CASCADE
ON UPDATE CASCADE
);


CREATE TABLE usermsg2(
userid	INT,
lasttime TIMESTAMP,
rank  INT DEFAULT 0,
score INT DEFAULT 0,
goodsnum INT DEFAULT 0,
visit	INT DEFAULT 0,
care INT DEFAULT 0,
becare INT DEFAULT 0,
likes  INT DEFAULT 0,
collect INT DEFAULT 0,
PRIMARY KEY(userid),
FOREIGN KEY(userid) REFERENCES account(id)
ON DELETE CASCADE
ON UPDATE CASCADE
);



CREATE TABLE goods(
goodsid INT NOT NULL,
ownerid INT NOT NULL,
state INT,
PRIMARY KEY(goodsid),
FOREIGN KEY(ownerid) REFERENCES account(id)
ON DELETE CASCADE
ON UPDATE CASCADE
);



CREATE TABLE goodsmsg(
goodsid INT NOT NULL,
imgurl VARCHAR(100),
price real DEFAULT 0.0,
ldata DATE,
title VARCHAR(50),
state INT DEFAULT 1,
PRIMARY KEY(goodsid),
FOREIGN KEY(goodsid) REFERENCES goods(goodsid)
ON DELETE CASCADE
ON UPDATE CASCADE
);


CREATE TABLE goodstate(
goodsid INT NOT NULL,
seetimes INT DEFAULT 0,
likenum INT DEFAULT 0,
talknum INT DEFAULT 0,
PRIMARY KEY(goodsid),
FOREIGN KEY(goodsid) REFERENCES goods(goodsid)
ON DELETE CASCADE
ON UPDATE CASCADE
);


CREATE TABLE coll_goods(
userid INT,
goodsid INT,
PRIMARY KEY (userid),
FOREIGN KEY (userid) REFERENCES account(id) 
ON DELETE CASCADE
ON UPDATE CASCADE,
FOREIGN KEY (goodsid) REFERENCES goods(goodsid) 
ON DELETE CASCADE
ON UPDATE CASCADE
);


CREATE TABLE coll_user(
userid1	INT,
userid2 INT,
PRIMARY KEY (userid1),
FOREIGN KEY (userid1) REFERENCES account(id) 
ON DELETE CASCADE
ON UPDATE CASCADE,
FOREIGN KEY (userid2) REFERENCES account(id) 
ON DELETE CASCADE
ON UPDATE CASCADE
);


CREATE TABLE statis(
sdata DATE NOT NULL,
visistimes INT DEFAULT 0,
sigtimes   INT DEFAULT 0,
newuser	  INT DEFAULT 0,
nwegoods  INT DEFAULT 0,
subgoods  INT DEFAULT 0,
talkunm	  INT DEFAULT 0,
income 	  INT DEFAULT 0,
PRIMARY KEY (sdata)
);


CREATE TABLE busylog(
sdata DATE NOT NULL,
times0 INT DEFAULT 0,
times1 INT DEFAULT 0,
times2 INT DEFAULT 0,
times3 INT DEFAULT 0,
times4 INT DEFAULT 0,
times5 INT DEFAULT 0,
times6 INT DEFAULT 0,
times7 INT DEFAULT 0,
times8 INT DEFAULT 0,
times9 INT DEFAULT 0,
times10 INT DEFAULT 0,
times11 INT DEFAULT 0,
PRIMARY KEY(sdata),
FOREIGN KEY(sdata) REFERENCES statis(sdata)
);

create table regcode(
npe varchar(100),
code int,
state int default 0,
stime timestamp
);



CREATE OR REPLACE FUNCTION auditlogfunc() RETURNS TRIGGER AS $example_table$
BEGIN
INSERT INTO confirm(userid)VALUES(new.id);
INSERT INTO usermsg1(userid)VALUES(new.id);
INSERT INTO usermsg2(userid)VALUES(new.id);
RETURN NEW;     
END;  

$example_table$ LANGUAGE plpgsql;

CREATE TRIGGER init_user_msg AFTER INSERT ON account
FOR EACH ROW EXECUTE PROCEDURE auditlogfunc();

