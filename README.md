# CNCS

[![Go](https://github.com/weldsk/casual-nocode-service/actions/workflows/go.yml/badge.svg)](https://github.com/weldsk/casual-nocode-service/actions/workflows/go.yml)

## Environment Variables

Set the following environment variables.

| Name                      | Type      | Description                             |
| ------------------------- | --------- | --------------------------------------- |
| SERVER_DB_LOCAL           | bool(0,1) | Use SQLite                              |
| SERVER_DBPATH             | string    | Database address (MySQL or MariaDB)     |
| SERVER_DBUSER             | string    | Username to access the database         |
| SERVER_DBPASS             | string    | Password to access the database         |
| REACT_APP_API_URL         | string    | API address of the backend side         |
| REACT_APP_PRIVATE_API_URL | string    | Private API address on the backend side |
