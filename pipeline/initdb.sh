#! /bin/sh

touch db.sqlite
cat sql/create_tables.sql | sqlite3 db.sqlite
