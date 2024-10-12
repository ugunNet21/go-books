curl -X POST http://localhost:8080/books/create -d '{"title": "Book Title", "author": "Book Author"}' -H "Content-Type: application/json"
curl http://localhost:8080/books
curl -X POST http://localhost:8080/books/update -d '{"id": 1, "title": "Updated Title", "author": "Updated Author"}' -H "Content-Type: application/json"
curl -X POST http://localhost:8080/books/delete -d '1' -H "Content-Type: application/json"

curl -X POST http://localhost:8080/books/create -H "Content-Type: application/json" -d '{"title": "Book Title", "author": "Author Name"}'


create database dan table ;
CREATE DATABASE go_crud_books;

USE go_crud_books;

CREATE TABLE books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL
);

untuk menjalankan project :
go run . 



create databse di databasel mysql ;

anone@anone1:~$ sudo mysql -u root -p
[sudo] password for anone: 
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 193
Server version: 8.0.39-0ubuntu0.24.04.2 (Ubuntu)

Copyright (c) 2000, 2024, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| phpmyadmin         |
| sys                |
| test_app           |
| web_wareuma        |
| yesjitudb          |
+--------------------+
8 rows in set (0.00 sec)

mysql> create database go_crud_books;
Query OK, 1 row affected (0.01 sec)

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| go_crud_books      |
| information_schema |
| mysql              |
| performance_schema |
| phpmyadmin         |
| sys                |
| test_app           |
| web_wareuma        |
| yesjitudb          |
+--------------------+
9 rows in set (0.00 sec)

mysql> USE go_crud_books;
Database changed
mysql> CREATE TABLE books (
    -> id INT AUTO_INCREMENT PRIMARY KEY,
    -> title VARCHAR(255) NOT NULL,
    -> author VARCHAR(255) NOT NULL
    -> );
Query OK, 0 rows affected (0.04 sec)

mysql> SELECT User, Host FROM mysql.user;
+------------------+-----------+
| User             | Host      |
+------------------+-----------+
| debian-sys-maint | localhost |
| mysql.infoschema | localhost |
| mysql.session    | localhost |
| mysql.sys        | localhost |
| phpmyadmin       | localhost |
| root             | localhost |
| yesjitudb        | localhost |
+------------------+-----------+
7 rows in set (0.00 sec)

mysql> GRANT ALL PRIVILEGES ON go_crud_books.* TO 'root'@'localhost';
Query OK, 0 rows affected (0.01 sec)

mysql> FLUSH PRIVILEGES;
Query OK, 0 rows affected (0.01 sec)

mysql> GRANT ALL PRIVILEGES ON go_crud_books.* TO 'yesjitudb'@'localhost';
Query OK, 0 rows affected (0.00 sec)

mysql> FLUSH PRIVILEGES;
Query OK, 0 rows affected (0.01 sec)

mysql> SHOW TABLES;
+-------------------------+
| Tables_in_go_crud_books |
+-------------------------+
| books                   |
+-------------------------+
1 row in set (0.00 sec)

mysql> SELECT * FROMT books;
ERROR 1064 (42000): You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'FROMT books' at line 1
mysql> SELECT * FROM books;
+----+-------+--------+
| id | title | author |
+----+-------+--------+
|  1 | Ipa   | Ugun   |
+----+-------+--------+
1 row in set (0.00 sec)

mysql> SELECT * FROM books;
+----+-------+--------+
| id | title | author |
+----+-------+--------+
|  1 | Ipa   | Ugun   |
|  2 | tes   | aku    |
+----+-------+--------+
2 rows in set (0.00 sec)

mysql> 

