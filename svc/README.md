# Objective
Create RESTful endpoint(s) to calculate, store, and retrieve customer rewards data from MongoDB.

# Instructions:
* Design and implement the following endpoints.
    * **Endpoint 1:**
        * Accept a customer's order data: **email adress**  (ex. "customer01@gmail.com") and **order total** (ex. 100.80).
        * Calculate and store the following customer rewards data into MongoDB. For each dollar a customer spends, the customer will earn 1 reward point. For example, a customer whose order total is $100.80 earns 100 points and belongs to Rewards Tier A. Note: a customer can only belong to one rewards tier. For example, a customer with 205 reward points belongs to Rewards Tier B and cannot use the reward from Tier A. Once a customer has reached the top rewards tier, there are no additional rewards the customer can earn.
            * **Email Address:** the customer's email address (ex. "customer01@gmail.com")
            * **Reward Points:** the customer's rewards points (ex. 100)
            * **Rewards Tier:** the rewards tier the customer has reached (ex. "A")
            * **Reward Tier Name:** the name of the rewards tier (ex. "5% off purchase")
            * **Next Rewards Tier:** the next rewards tier the customer can reach (ex. "B")
            * **Next Rewards Tier Name:** the name of next rewards tier (ex. "10% off purchase")
            * **Next Rewards Tier Progress:** the percentage the customer is away from reaching the next rewards tier (ex. 0.5)
    * **Endpoint 2:** Accept a customer's email address, and return the customer's rewards data that was stored in Endpoint 1.
    * **Endpoint 3:** Return the same rewards data as Endpoint 2 but for all customers.
* For bonus points, add error handling and unit tests.

# Routes 
`/order`

`/customers`

`/customers/{email}`

* See `smoke_tests` for usage examples

# Setup
```sh
dep ensure
GOOS=linux GOARCH=386 go build
docker build -t rewards_svc .
docker-compose up -d
```