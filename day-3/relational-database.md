# Relational Database

Entitas satu dengan yang lain bisa berhubungan (relation)
eFishery menggunakan postgresql (oracle tidak bisa auto increment)

## Data Definition Language
Define structure like schema, db, tables, constraints etc. e.g. create, drop and alter statements

## COMMANDS (POSTGRES)
- psql -h localhost -p 5432 -U postgres --> log into postgre
- create database <db_name>; --> to create db
- \l --> list all dbs
- \c --> connect to db
- create table <table_name> (<column_names> <data_types>) --> create table users (name char(5), age int); --> to create table
- alter table <table_name> rename <new_table_name> --> modifying table structure or properties
- drop table <table_name> --> to delete a table

DATA TYPES POSTGRES --> see documentation for detail

## Data Manipulation Language
Manipulate data inside a table insert into, update, delete

## Data Control Language 
Syntax similar to a computer programming language used to control access to data stored in a database
(also to isolate database for security reasons) e.g GRANT and REVOKE

### SQL JOIN
left join, right join, full outer inclusive, inner join (default), full outer exclusive

### AGGREGATION
sum, min, max, count, avg

### SUBQUERY
e.g. update products set stock = subquery.stock - 2 from (select id, stock from products where id = 1) as subquery where products.id = 1;

### FUNCTION
Create function to simplify query

CREATE FUNCTION kurangi_stock(INT, INT) RETURNS product AS 'update products set stock = subquery.stock - $2 from (select id, stock from products where id = $1) as subquery where products.id = $1;
select \* from products where id = $1'
LANGUAGE 'sql';

to invoke the function 'SELECT <function_name>
