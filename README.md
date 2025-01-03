# snacks-manager

This is a demonstration app for cloud native methods and tooling. It is structured as a simple micro-service application for managing a snacks supply.

It consists of the following services:

* **frontend**: a minimal server, written in Go that handles user authentication and serves the web gui.

* **backend**: server written in Go, that enforces the business rules and stores data in a SQLite database.

* **web-gui**: a minimal UI, written in Vue.js, to allow human users to use the snack shop.

* **snacker-agent**: an agent representing a user that regularly purchases a random snack from the shop.

* **clerk-agent**: a agent representing clerk, which will hand out the purchased snacks. Purchases from users will block, until the clerk hands of the product.


Potential future expansion:

* **manager-agent**: will order new snacks when stocks are low.

* **snacker-agent v2**: snacker has a budget in its bank and will transfer money to have a budget with the shop.

* **bank-agent**

* **producer-agent**
