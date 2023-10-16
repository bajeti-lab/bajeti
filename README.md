# [Bajeti]

## Project Description

Bajeti is designed to help individuals manage their finances effectively. It provides a user-friendly interface to track income, expenses, savings, and budgets across various categories. Key features include:

- **Budget Management**: Users can set budgets for specific time periods, allocating funds to different categories (e.g., groceries, entertainment).

- **Expense Tracking**: Expenses can be logged against predefined categories, providing a clear overview of spending habits.

- **Savings Goals**: Users can create savings goals, specifying the target amount and tracking progress towards achieving them.

- **Income Logging**: Allows users to record their income sources, providing a comprehensive view of their financial situation.

- **Date-Based Entries**: Automatic date stamping for entries, ensuring accurate and organized financial records.

- **Security and User Authentication**: Ensures user data is protected through secure login and authentication mechanisms.

## Technologies Used

- Go (Golang)
- PostgreSQL
- Fyne (for GUI)

## Project Setup

### Docker Compose

1. **Clone the Repository:**

    ```bash
    git clone https://github.com/bajeti-lab/bajeti.git
    ```

2. **Navigate to the Project Directory:**

    ```bash
    cd personal-budgeting-system
    ```

3. **Create an Environment File:**

    Create a file named `.env` in the project root and add the following content:

    ```plaintext
    POSTGRES_USER=myuser
    POSTGRES_PASSWORD=secretpassword
    POSTGRES_DB=mydb
    POSTGRES_HOST=postgresdb
    POSTGRES_PORT=5432
    ```

    Replace the values with your desired PostgreSQL credentials.

4. **Start the Docker Containers:**

    ```bash
    docker-compose up -d
    ```

    This will start both the Go application and the PostgreSQL database in separate containers.

5. **Access the Application:**

    Open your web browser and go to [http://localhost:8000](http://localhost:8000) to access the Personal Budgeting System.

6. **Stopping the Application:**

    To stop the containers, run:

    ```bash
    docker-compose down
    ```

### Manual Installation

1. **Clone the Repository:**

    ```bash
    git clone https://github.com/bajeti-lab/bajeti.git
    ```

2. **Setting up the Database:**

   - Install and configure PostgreSQL.
   - Create a database and update the database configuration in the application.

3. **Running the Application:**

    ```bash
    go run main.go
    ```

    The application will start, and you can access it in your web browser.

## Contributing

We welcome contributions from the community! If you have any suggestions, enhancements, or bug fixes, please open an issue or create a pull request.

## License

This project is licensed under the MIT License 

## Acknowledgments

- The Go Programming Language Community
- PostgreSQL Community
- Fyne GUI Library Contributors
