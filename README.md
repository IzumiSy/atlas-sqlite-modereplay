# How to reproduce

1. Run migration
```bash
$ go run -mod=mod ariga.io/atlas/cmd/atlas@master migrate apply --dir="file://ent/migrate/migrations" --url="sqlite3://main.db?_fk=1"
Migrating to version 20221024083708 (2 migrations in total):

  -- migrating version 20221024083312
    -> CREATE TABLE `users` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `name` text NOT NULL);
  -- ok (133.214826ms)

  -- migrating version 20221024083708
    -> ALTER TABLE `users` ADD COLUMN `age` integer NOT NULL;
  -- ok (132.619177ms)

  -------------------------
  -- 266.676089ms
  -- 2 migrations
  -- 2 sql statements
```

2. Tries generating new migration, then it fails.
```bash
$ go run -mod=mod ent/migrate/main.go add_job
2022/10/24 17:49:30 failed generating migration file: sql/migrate: connected database is not clean: found table "atlas_schema_revisions"
exit status 1
```
