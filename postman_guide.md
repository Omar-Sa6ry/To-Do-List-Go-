# Postman Guide (To-Do List API)

This guide contains all the endpoints for the project and step-by-step instructions on how to use them in Postman.

---

## 1. Create a New Account (Signup)
- **Method:** `POST`
- **URL:** `http://localhost:8080/signup`
- **Body Type:** `raw` -> `JSON`
- **Body Content:**
    ```json
    {
        "email": "test@example.com",
        "password": "mypassword123"
    }
    ```

---

## 2. Login
- **Method:** `POST`
- **URL:** `http://localhost:8080/login`
- **Body Type:** `raw` -> `JSON`
- **Body Content:**
    ```json
    {
        "email": "test@example.com",
        "password": "mypassword123"
    }
    ```
> **Important Note:** After clicking Send, the API will return a `token`. Copy this Token because you will need it for the upcoming steps.

---

## 🔒 Setup Authentication for Protected Routes
For all upcoming requests related to Todos, you must include the Token in Postman as follows:
1. Go to the **Authorization** tab.
2. Select the Type: **Bearer Token**.
3. Paste the Token you copied from the Login step.

---

## 3. Create a New Task (Create Todo)
- **Method:** `POST`
- **URL:** `http://localhost:8080/todos`
- **Body Type:** `raw` -> `JSON`
- **Body Content:**
    ```json
    {
        "title": "Learn Go Interfaces",
        "description": "Study interfaces from Maximilian's course",
        "status": "pending"
    }
    ```

---

## 4. Get All Tasks (Get All Todos)
- **Method:** `GET`
- **URL:** `http://localhost:8080/todos`
- **Body:** None (Empty)

---

## 5. Get a Specific Task (Get Todo By ID)
- **Method:** `GET`
- **URL:** `http://localhost:8080/todos/1`
> Replace the number `1` in the URL with the ID of the task you want to fetch.
- **Body:** None (Empty)

---

## 6. Update Task Data (Update Todo)
- **Method:** `PUT`
- **URL:** `http://localhost:8080/todos/1`
> Replace the number `1` in the URL with the task ID.
- **Body Type:** `raw` -> `JSON`
- **Body Content:**
    ```json
    {
        "title": "Learn Go Interfaces - Completed!",
        "description": "Finished the interfaces section",
        "status": "done"
    }
    ```

---

## 7. Delete Task (Delete Todo)
- **Method:** `DELETE`
- **URL:** `http://localhost:8080/todos/1`
> Replace the number `1` in the URL with the task ID.
- **Body:** None (Empty)
