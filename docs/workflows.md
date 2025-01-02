# Workflows

This is a list of workflows, that are handled by the application.
See [Roles](roles.md) for a detailed description of the roles acting in these workflows.

## Purchase Items



1. Customer orders item from snack manager.
2. Clerk hands out item(s). Clerk is blocked until customer acknowledges acceptance of items
3. Customer acknowledges receipt of items.

Timeout: if customer does not acknowledge reception within set time limit, the mitigation process is triggered. 
Current mitigation strategy: the customer is not billed for items, but items are deducted from stock. (The customer is always right)

Later Expansion:

* Program checks, that

  * there are enough funds in the budget of the Customer.

  * enough stock available to fulfill the request of the customer.

* smarter mitigation strategy. Some possible examples:

  * Does customer re-order the same items (indicating the they actually did not receive them)?
  * Does a single customer have an above average number of acknowledgement timeouts?
  * Does a single clerk have an above average number of acknowledgement timeouts?
