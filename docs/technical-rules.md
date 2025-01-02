# Technical Rules

These are the requirements that set for the system.

The specifics are set arbitrarily, but are common aspects to many scale-out, high-availability services. Some are chosen specifically to inform on the required architecture of the full stack.

## Rules

* **No HTTP Round-Trip (Request+Response) must have 90th percentile of more than 300msec, when measured by the requesting client. The 100th percentile must not be more than 3 sec.**

** A derivative of the requirement is that there must be absolutely no blocking calls/requests.