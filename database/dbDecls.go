package main

type dbParams struct {
	string host,
	int port,
	string dbName,
	string username,
	string pwd
}

type Users struct {
	string username,
	string age,
	string address
}