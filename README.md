### Backend README

#### Running with Docker Compose:

1. Ensure Docker and Docker Compose are installed on your system.

2. Clone this repository:
    ```bash
    git clone git@github.com:pramudya3/backend.git && cd backend
    ```

   This command clones the repository from the specified GitHub URL and then navigates into the cloned repository directory.

4. Build and start the containers using Docker Compose:
   ```bash
   docker-compose up --build
   ```

5. Access the services through the defined endpoints.

#### Services:

1. **Account Service**
   - **Running at**: `http://localhost:2000`

   - **Users API endpoints:**
     - `POST /signup`: Create a new user account.
       - JSON payload:
         ```json
         {
           "name": "name",
           "username": "username",
           "email": "email@email.com",
           "password": "pass123"
         }
         ```
     - `POST /login`: Log in to an existing user account.
        - JSON payload:
        ```json
         {
           "id": 1,
           "password": "pass123"
         }
         ```
         Response payload:
         ```json
         {
          "data": "access-token"
         }
         ```

     - `GET user/logout`: Log out from the current user session, protected by JWT Token
        - Bearer token: `access-token`

     - `GET /user/my-profile`: Get user information user, protected by JWT Token
        - Bearer token: `access-token`

   - **Payment Account Endpoints:**
     - `POST /payment-account`: Create a payment account associated with a user.
        - Bearer token: `access-token`
        - JSON payload:
         ```json
         {
           "type":"Debit", // {Debit, Credit}
           "balance":50000 //  
         }
         ```

   - **Payment History Endpoints:**
     - `GET /payment-history`: Retrieve payment history for a user, protected by JWT token
        - Bearer token: `access-token`


2. **Payment Service**
   - **Running at**: `http://localhost:3000`

   - **APIs Endpoints:**
     - `POST /send`: Send money to another user.
        - Bearer token: `access-token`
       
       - JSON payload:
         ```json
         {
           "to_user_id": 2,
           "amount": 50000,
           "payment_method":"Debit"
         }
         ```
     - `POST /withdraw`: Withdraw money from the user's account, protected by JWT token 
        - Bearer token: `access-token`
       - JSON payload:
         ```json
         {
           "amount":5000,
           "payment_method":"Debit"
         }
         ```

#### Tech Stack:

- Golang
- Supertokens
- GORM
- PostgreSQL
- Gin Gonic
