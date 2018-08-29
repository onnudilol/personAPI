BEGIN;
CREATE TABLE "person_person" ("id" serial NOT NULL PRIMARY KEY, "displayname" varchar(100) NOT NULL, "is_team" boolean NOT NULL);
CREATE TABLE "person_person_members" ("id" serial NOT NULL PRIMARY KEY, "from_person_id" integer NOT NULL, "to_person_id" integer NOT NULL);
ALTER TABLE "person_person_members" ADD CONSTRAINT "person_person_member_from_person_id_d27091c2_fk_person_pe" FOREIGN KEY ("from_person_id") REFERENCES "person_person" ("id") DEFERRABLE INITIALLY DEFERRED;
ALTER TABLE "person_person_members" ADD CONSTRAINT "person_person_members_to_person_id_85fa173b_fk_person_person_id" FOREIGN KEY ("to_person_id") REFERENCES "person_person" ("id") DEFERRABLE INITIALLY DEFERRED;
ALTER TABLE "person_person_members" ADD CONSTRAINT person_person_members_from_person_id_to_person_id_c8ade548_uniq UNIQUE ("from_person_id", "to_person_id");
CREATE INDEX "person_person_members_from_person_id_d27091c2" ON "person_person_members" ("from_person_id");
CREATE INDEX "person_person_members_to_person_id_85fa173b" ON "person_person_members" ("to_person_id");
COMMIT;

INSERT INTO person_person (displayname, is_team)
VALUES ('Alice', FALSE);

INSERT INTO person_person (displayname, is_team)
VALUES ('Bob', FALSE);

INSERT INTO person_person (displayname, is_team)
VALUES ('Carlos', FALSE);

INSERT INTO person_person (displayname, is_team)
VALUES ('Carol', FALSE);

INSERT INTO person_person (displayname, is_team)
VALUES ('Charlie', FALSE);

INSERT INTO person_person (displayname, is_team)
VALUES ('Chuck', FALSE);

INSERT INTO person_person (displayname, is_team)
VALUES ('Dave', FALSE);

INSERT INTO person_person (displayname, is_team)
VALUES ('Eve', FALSE);

INSERT INTO person_person (displayname, is_team)
VALUES ('Mallory', FALSE);

INSERT INTO person_person (displayname, is_team)
VALUES ('Peggy', FALSE);

INSERT INTO person_person (displayname, is_team)
VALUES ('Trent', FALSE);

INSERT INTO person_person (displayname, is_team)
VALUES ('Victor', FALSE);

INSERT INTO person_person (displayname, is_team)
VALUES ('Walter', FALSE);

INSERT INTO person_person (displayname, is_team)
VALUES ('The A-Team', TRUE);

INSERT INTO person_person (displayname, is_team)
VALUES ('The B-Team', TRUE);

INSERT INTO person_person (displayname, is_team)
VALUES ('The C-Team', TRUE);

INSERT INTO person_person_members (from_person_id, to_person_id)
VALUES (14, 1);

INSERT INTO person_person_members (from_person_id, to_person_id)
VALUES (14, 2);

INSERT INTO person_person_members (from_person_id, to_person_id)
VALUES (14, 3);

INSERT INTO person_person_members (from_person_id, to_person_id)
VALUES (15, 10);

INSERT INTO person_person_members (from_person_id, to_person_id)
VALUES (15, 11);

INSERT INTO person_person_members (from_person_id, to_person_id)
VALUES (15, 12);

INSERT INTO person_person_members (from_person_id, to_person_id)
VALUES (16, 5);

INSERT INTO person_person_members (from_person_id, to_person_id)
VALUES (16, 8);

INSERT INTO person_person_members (from_person_id, to_person_id)
VALUES (16, 14);

INSERT INTO person_person_members (from_person_id, to_person_id)
VALUES (14, 15);

INSERT INTO person_person_members (from_person_id, to_person_id)
VALUES (15, 16);