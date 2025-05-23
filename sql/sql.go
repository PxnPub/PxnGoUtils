package sql;

import(
	Strings "strings"
);



func IsDuplicateColumnError(err error) bool {
	if err == nil { return false; }
	if Strings.HasPrefix(err.Error(), "failed to execute query ALTER TABLE ") &&
	Strings.Contains(err.Error(), "Error executing statement: SQLite failure: `duplicate column name: ") {
		return true;
	}
	if Strings.HasPrefix(err.Error(), "Catalog Error: Column with name ") &&
	Strings.HasSuffix(err.Error(), " already exists!") {
		return true;
	}
	return false;
}
