# Pay Connect

Payment application build for MNC Client Test By Enimga Camp

## How to run the application

- create `.env` file. You can use `.env.example` file as reference
- Ensure that you have `PostgreSQL` installed
- Execute this command in your terminal

```bash
go run main.go
```

## Endpoint

- /api/v1/auth

  - POST /register
    - Request
      - body
      ```json
      {
        "username": "example",
        "email": "example@example.com",
        "password": "@Example123!"
      }
      ```
    - Response
    ```json
    {
      "status": "Success",
      "message": "Registration Successfull"
    }
    ```
  - POST /login

    - Request
      - body
      ```json
      {
        "username": "example",
        "password": "@Example123!"
      }
      ```
    - Response

    ```json
    {
      "status": "Success",
      "message": "Logged in",
      "data": {
        "token": "jwt_token"
      }
    }
    ```

- /api/v1/payments

  - POST /

    - header
      - Authorization: Bearer jwt_token
    - Request

      - body

      ```json
      {
        "amount": 10000,
        "receiver_customer_id": 1
      }
      ```

    - Response

    ```json
    {
      "status": "Success",
      "message": "Payment created successfully",
      "data": {
        "ID": 3,
        "Amount": 10000,
        "SenderCustomerID": 1,
        "ReceiverCustomerID": 2,
        "Status": "pending",
        "CreatedAt": "2024-10-05T18:28:51.986952112+07:00",
        "UpdatedAt": "2024-10-05T18:28:51.986952112+07:00"
      }
    }
    ```

  - PATCH /:id/complete

        - header
          - Authorization: Bearer jwt_token
        - Response

        ```json
        {
        "Status": "Success",
        "Message": "Payment status updated successfully"
        }
        ```
