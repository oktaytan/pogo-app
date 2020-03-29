BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "user_likes" (
	"id"	INTEGER NOT NULL,
	"post_id"	INTEGER NOT NULL,
	"user_id"	INTEGER NOT NULL,
	"liked"	INTEGER NOT NULL DEFAULT 0,
	FOREIGN KEY("post_id") REFERENCES "posts"("id"),
	PRIMARY KEY("id"),
	FOREIGN KEY("user_id") REFERENCES "users"("id")
);
CREATE TABLE IF NOT EXISTS "users" (
	"id"	INTEGER NOT NULL,
	"username"	TEXT NOT NULL,
	"email"	TEXT NOT NULL,
	"password"	TEXT NOT NULL,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "comments" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"comment"	TEXT NOT NULL,
	"post_id"	INTEGER NOT NULL,
	"user_id"	INTEGER NOT NULL,
	FOREIGN KEY("post_id") REFERENCES "posts"("id"),
	FOREIGN KEY("user_id") REFERENCES "users"("id")
);
CREATE TABLE IF NOT EXISTS "posts" (
	"id"	INTEGER NOT NULL,
	"title"	TEXT NOT NULL,
	"body"	TEXT NOT NULL,
	"created_at"	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at"	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"user_id"	INTEGER NOT NULL,
	"likes"	INTEGER NOT NULL DEFAULT 0,
	PRIMARY KEY("id"),
	FOREIGN KEY("user_id") REFERENCES "users"("id")
);
INSERT INTO "users" VALUES (1,'cawis','cawis@test.com','$2a$04$/SG.WkZNblznY.KDD3QniO.qZl6mEujExUoF5jHydV0waQWzT8De2');
INSERT INTO "users" VALUES (2,'zidak','zidak@test.com','$2a$04$w7VDhWfpSRDhLyodnJeI2OAZ4/eDouDunYPj7efniCnPkRROYhZrq');
INSERT INTO "users" VALUES (3,'maya','maya@test.com','$2a$04$ZX0MqEMqwlGq8VPdnU8DSev9ravCg.Wn/UMvY5OHz5MI60SiT9NCK');
INSERT INTO "comments" VALUES (0,'Comment Eight',94235010,1);
INSERT INTO "comments" VALUES (1,'Comment One',11666145821,2);
INSERT INTO "comments" VALUES (2,'Comment Two',11666145821,2);
INSERT INTO "comments" VALUES (3,'Comment Three',16287113937,2);
INSERT INTO "comments" VALUES (4,'Comment Four',16287113937,1);
INSERT INTO "comments" VALUES (5,'Comment Five',16287113937,1);
INSERT INTO "posts" VALUES (11666145821,'Post Title Three','Pellentesque rhoncus dolor at odio egestas pharetra. Nullam auctor, nisl vitae cursus convallis, ligula ipsum consectetur nunc, et ornare orci erat vel diam. Etiam felis metus, mollis nec tortor vitae, lacinia aliquam justo. Donec felis nunc, viverra venenatis tellus eu, egestas convallis ligula. Duis accumsan neque tincidunt, feugiat est sit amet, convallis massa. Phasellus sed imperdiet enim. Sed luctus cursus est, quis eleifend nisi faucibus vel. Aenean pharetra, leo rutrum pharetra sagittis, ipsum lorem bibendum ex, eu tincidunt mauris felis ut erat. Suspendisse potenti. Donec gravida risus eget elementum porta. Donec eget lacinia magna, sit amet vulputate velit. Vestibulum id condimentum libero.','2020-03-22 20:09:39','2020-03-22 20:09:39',1,34);
INSERT INTO "posts" VALUES (16287113937,'Post Title Five','Integer id sapien auctor, dignissim ante et, commodo ipsum. Pellentesque condimentum erat eu neque aliquam dapibus. Vivamus sit amet tincidunt augue, vitae pellentesque justo. Ut sagittis sollicitudin mauris sit amet interdum. Morbi laoreet lacus in pharetra imperdiet. Duis venenatis gravida leo, vitae vehicula arcu porta nec. Cras scelerisque consectetur quam. Ut laoreet volutpat quam, eu convallis mauris luctus sit amet.','2020-03-22 20:10:47','2020-03-22 20:10:47',2,22);
INSERT INTO "posts" VALUES (23082153551,'Post Title Two','Nullam eu mauris ultricies, tincidunt felis vitae, scelerisque velit. Ut ligula nisi, dapibus in rutrum in, placerat at leo. Duis a est pellentesque erat posuere posuere eu quis nisl. Aenean vel interdum urna. Morbi hendrerit eget turpis vitae semper. Donec vitae odio euismod, vulputate eros ut, consequat dolor. Donec at tempor libero. Maecenas ac ultrices dui. Quisque vel cursus sapien. Aliquam vel lorem erat.','2020-03-22 20:09:20','2020-03-22 20:09:20',1,456);
INSERT INTO "posts" VALUES (91947779410,'Post Title One','Cras euismod erat et ante volutpat imperdiet. Suspendisse non libero at orci malesuada viverra ac elementum eros. Suspendisse potenti. Donec eget fermentum velit. Etiam mattis nibh non arcu imperdiet suscipit. Nunc ultrices, sem aliquet scelerisque dignissim, felis lectus convallis tellus, vel fermentum libero risus eu mauris. Phasellus quis lectus sed diam porta laoreet. Pellentesque blandit imperdiet ultricies.','2020-03-22 20:08:57','2020-03-22 20:08:57',1,350);
INSERT INTO "posts" VALUES (94235010051,'Post Title Four','In tempor metus ut magna aliquet ultricies. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Praesent faucibus mauris ac tempus sollicitudin. Ut sagittis porta nulla, eget maximus neque. Integer lacinia rutrum justo, eu rutrum enim euismod ut. Quisque in arcu sit amet mi consectetur varius ut non ligula. Pellentesque id ex risus. Donec mattis facilisis magna, at vulputate ante gravida eget. Praesent eget aliquet justo. Suspendisse eget neque quam.','2020-03-22 20:10:13','2020-03-22 20:10:13',2,353);
COMMIT;
