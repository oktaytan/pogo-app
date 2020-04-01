BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "user_likes" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"post_id"	INTEGER NOT NULL,
	"user_id"	INTEGER NOT NULL,
	"liked"	INTEGER NOT NULL DEFAULT 0,
	FOREIGN KEY("user_id") REFERENCES "users"("id") ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY("post_id") REFERENCES "posts"("id") ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "comments" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"comment"	TEXT NOT NULL,
	"post_id"	INTEGER NOT NULL,
	"user_id"	INTEGER NOT NULL,
	FOREIGN KEY("user_id") REFERENCES "users"("id") ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY("post_id") REFERENCES "posts"("id") ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "posts" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"title"	TEXT NOT NULL,
	"body"	TEXT NOT NULL,
	"created_at"	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at"	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"user_id"	INTEGER NOT NULL,
	"likes"	INTEGER NOT NULL DEFAULT 0,
	FOREIGN KEY("user_id") REFERENCES "users"("id") ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "users" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"username"	TEXT NOT NULL,
	"email"	TEXT NOT NULL,
	"password"	TEXT NOT NULL
);
INSERT INTO "user_likes" VALUES (1,3,6,1);
INSERT INTO "user_likes" VALUES (2,4,6,1);
INSERT INTO "user_likes" VALUES (3,5,6,1);
INSERT INTO "user_likes" VALUES (4,9,5,1);
INSERT INTO "user_likes" VALUES (5,3,5,1);
INSERT INTO "user_likes" VALUES (6,2,5,1);
INSERT INTO "user_likes" VALUES (7,8,4,1);
INSERT INTO "user_likes" VALUES (8,9,4,1);
INSERT INTO "user_likes" VALUES (9,5,4,1);
INSERT INTO "comments" VALUES (1,'Tricies pharetra. Suspendisse potenti. Fusce eget laoreet arcu, vitae viverra magna. Vestibulum at gravida metus. Cras lacinia odio sit amet tincidunt',3,6);
INSERT INTO "comments" VALUES (2,'Enim vel commodo facilisis, tellus diam semper felis, vel laoreet velit eros eget dolor. Phasellus id sapien rhoncus, feugiat risus sit amet, porttitor velit. Fusce augue eros, imperdiet sit amet eleme',5,6);
INSERT INTO "comments" VALUES (3,'Nulla facilisi. Mauris eu imperdiet lacus. In sit amet elit sagittis, pellentesque ex et, pretium sem. Fusce lobortis eget lorem ac laoreet.',5,5);
INSERT INTO "comments" VALUES (4,'Fusce augue eros, imperdiet sit amet elementum ut, sollicitudin at sapien. In laoreet laoreet volutpat. Praesent efficitur nisi id tempor facilisis. Nullam gravida libero et dolor ultrices, sed finibus',2,5);
INSERT INTO "comments" VALUES (5,'Suspendisse potenti. Fusce eget laoreet arcu, vitae viverra magna. Vestibulum at gravida metus. Cras lacinia odio sit amet tincidunt faucibus. Praesent volutpat lorem eget magna aliquet, vitae faucibus',4,5);
INSERT INTO "comments" VALUES (6,'Mauris id neque arcu. Aliquam metus tortor, facilisis sed pharetra a, dignissim sed nisi. Donec mi massa, facilisis non erat eu, bibendum facilisis est. Nunc tincidunt nisl justo, ac efficitur nibh eff',4,4);
INSERT INTO "comments" VALUES (7,'Quisque posuere, elit in pretium sagittis, mi purus tincidunt ante, bibendum luctus nisi massa sed eros. Sed tincidunt tincidunt molestie. ',9,4);
INSERT INTO "comments" VALUES (8,'Mauris id neque arcu. Aliquam metus tortor, facilisis sed pharetra a, dignissim sed nisi. Donec mi massa, facilisis non erat eu, bibendum facilisis est. Nunc tincidunt nisl justo, ac efficitur nibh eff',5,4);
INSERT INTO "comments" VALUES (9,'Nunc nec lorem tincidunt, sollicitudin ligula sed, fringilla sapien. Nullam feugiat accumsan lacus, vitae dapibus leo vulputate sed. Orci varius natoque penatibus et magnis dis parturient montes, nasce',2,4);
INSERT INTO "posts" VALUES (1,'Nam non pharetra odio.','Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas nibh dui, lacinia ut dignissim a, consectetur a tellus. Integer consequat ex vel gravida porta. Donec mauris nunc, lobortis ac efficitur vel, finibus id felis. Aliquam erat volutpat. In ultrices venenatis facilisis. Fusce ultricies eu metus pharetra consequat. Sed consequat rhoncus .','2020-03-31 23:47:09','2020-03-31 23:47:09',4,22);
INSERT INTO "posts" VALUES (2,'In bibendum porta','Fusce facilisis ultrices nulla, id imperdiet justo ornare in. In id velit eget massa iaculis pharetra. Vivamus vitae lorem in justo varius mattis. Quisque augue libero, tempor ut tincidunt sodales, ornare id arcu. Mauris id neque arcu. Aliquam metus tortor, facilisis sed pharetra a, dignissim sed nisi. Donec mi massa, facilisis non erat eu, bibendu.','2020-03-31 23:47:48','2020-03-31 23:47:48',4,44);
INSERT INTO "posts" VALUES (3,'Etiam faucibus diam sit amet tellus fermentum','Nam sagittis urna eu sem pretium gravida. Quisque posuere, elit in pretium sagittis, mi purus tincidunt ante, bibendum luctus nisi massa sed eros. Sed tincidunt tincidunt molestie. Aenean posuere dolor a ultricies pharetra. Suspendisse potenti. Fusce eget laoreet arcu, vitae viverra magna. Vestibulum at gravida metus.','2020-03-31 23:48:14','2020-03-31 23:48:14',4,324);
INSERT INTO "posts" VALUES (4,'Vivamus vel finibus mi','Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque eu orci felis. Nam a sodales augue. Maecenas vestibulum nisl ante, ac tempus ex pellentesque et. Phasellus ultricies est vel mauris varius, non mollis massa fringilla. Curabitur non ex id eros lacinia rutrum eget nec dui. Ut pellentesque, sapien ac pretium porttitor, velit nibh sodale.','2020-03-31 23:53:22','2020-03-31 23:53:22',5,775);
INSERT INTO "posts" VALUES (5,'Mi magna pulvinar era','Vulputate, vulputate id elit. Nulla facilisi. Mauris eu imperdiet lacus. In sit amet elit sagittis, pellentesque ex et, pretium sem. Fusce lobortis eget lorem ac laoreet. Donec accumsan at ante nec ornare. Aliquam posuere felis nulla, at dignissim ligula ornare eget. Integer tempor ligula nec suscipit aliquam. Vivamus velit ligula, tristique ut blan','2020-03-31 23:53:47','2020-03-31 23:53:47',5,456);
INSERT INTO "posts" VALUES (6,'Amet elementum ut','Ultrices nulla, id imperdiet justo ornare in. In id velit eget massa iaculis pharetra. Vivamus vitae lorem in justo varius mattis. Quisque augue libero, tempor ut tincidunt sodales, ornare id arcu. Mauris id neque arcu. Aliquam metus tortor, facilisis sed pharetra a, dignissim sed nisi. Donec mi massa, facilisis non erat eu, bibendum facilisis est. ','2020-03-31 23:54:13','2020-03-31 23:54:13',5,34);
INSERT INTO "posts" VALUES (7,'Scelerisque','Aliquam fermentum metus sit amet laoreet lacinia. Vestibulum eu nulla id lorem sagittis mattis. Suspendisse potenti. In egestas suscipit tortor quis hendrerit. Aenean vitae tempor ipsum, eget consequat purus. Ut at arcu ipsum. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquam nibh ac elit.','2020-03-31 23:55:06','2020-03-31 23:55:06',6,98);
INSERT INTO "posts" VALUES (8,'Cras at iaculis libero','Fusce vel nibh mi. Integer non odio in magna maximus ornare vitae vitae turpis. Aenean volutpat sed eros at tristique. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas porta, felis eu vestibulum malesuada, mauris urna consectetur nibh, ac facilisis nulla eros vel dolor. Ut sollicitudin turpis eu ipsum.','2020-03-31 23:55:29','2020-03-31 23:55:29',6,245);
INSERT INTO "posts" VALUES (9,'Pretium porttitor, veli','Sed ipsum nulla, laoreet id euismod vulputate, vulputate id elit. Nulla facilisi. Mauris eu imperdiet lacus. In sit amet elit sagittis, pellentesque ex et, pretium sem. Fusce lobortis eget lorem ac laoreet. Donec accumsan at ante nec ornare. Aliquam posuere felis nulla, at dignissim ligula ornare eget. Integer tempor ligula nec suscipit aliquam.','2020-03-31 23:55:56','2020-03-31 23:55:56',6,54);
INSERT INTO "users" VALUES (4,'cawis','cawis@test.com','$2a$04$FQDIF3UoeN2QH1lp9g47AuNMLppDti0QFj/xdOs0weBt8uk1nBGnu');
INSERT INTO "users" VALUES (5,'zidak','zidak@test.com','$2a$04$xumtM5.b4aWRtdmplEtnU.WGdSQr95aP0wpveew4K37qbs.AElXgu');
INSERT INTO "users" VALUES (6,'maya','maya@test.com','$2a$04$eX1eEgny9OLV8HRBdaJj3uEjX7h9VypshP5Gm6rt3fMIEGFpYXyqO');
COMMIT;
