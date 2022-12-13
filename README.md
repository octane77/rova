ROVA Coding assignment

The assessment consists of an API to be used for opening a new “current account” of already existing customers.

Requirements

• The API will expose an endpoint which accepts the user information (customerID,
initialCredit).

• Once the endpoint is called, a new account will be opened connected to the user whose ID
is customerID.

• Also, if initialCredit is not 0, a transaction will be sent to the new account.

• Another Endpoint will output the user information showing Name, Surname, balance, and
transactions of the accounts.

Bonuses

• Accounts and Transactions are different services.

• Frontend (simple one is OK). Open API / Swagger UI will be accepted as well.

• Attention to CI/CD

Constraints

Feel free to use any open source tool or framework.

For storing the information, the data can be saved in memory and not actually persisted to
an external database, so that we can test the solution easier. However, remember this is a
backend assignment and consider layers, abstractions, testability and enterprise-level
architecture carefully.

The programming language by default is Golang, but we are flexible if you let us know which
other language you prefer.

Expectations

The expected deliverable is the source code published on GitHub or Gitlab and instructions
on how to execute and test it.

We hope to see a demonstration of good GIT practice and workflow; show us how you work
in a team.

Testability will also be assessed.

Show your knowledge beyond boilerplate endpoints!