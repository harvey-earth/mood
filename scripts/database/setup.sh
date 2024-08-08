#!/usr/bin/env bash

if command -v sqlite3 &> /dev/null
then
    sqlite3 mood.db < scripts/mood-sqlite.sql
else
    echo "Please install sqlite3"
fi
