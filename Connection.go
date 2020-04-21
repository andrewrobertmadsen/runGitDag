package main

type Connection struct {
	Host string
	Port int
	Type ConnectionType
}

type ConnectionType string

const(
	Hive ConnectionType = "Hive"
	Impala = "Impala"
	Presto = "Presto"
	SparkSql = "SparkSql"
)