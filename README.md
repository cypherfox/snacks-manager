# snacks-manager

This is a demonstration app for cloud native methods and tooling. It is structured as a simple micro-service application for managing a snacks supply.

It consists of the following services:

* **frontend**: a minimal server, written in Go that handles user authentication and serves the web gui.

* **backend**: server written in Go, that enforces the business rules and stores data in a SQLite database.

* **web-gui**: a minimal UI, written in Vue.js, to allow human users to use the snack shop.
