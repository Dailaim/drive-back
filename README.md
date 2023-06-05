# Driver-back

## Local Project Setup

The program assumes that you already have a Postgres database and Go installed.

To run the project in development mode, we must follow these steps:

1. Clone the repository
2. Download dependencies: `go mod tidy`
3. Configure environment variables according to the .env.example
4. Run the program: `go run main.go`
5. The program will be running on port 8080
6. View swagger documentation on http://localhost:8080/swagger/index.html

## Run unit tests in vscode

To run the unit tests, we need to make use of Go's testing library. Make sure you have the necessary testing files in your project. They usually have the format `*_test.go`. Once you've confirmed this, you can run your tests within vscode by right-clicking on the test file and selecting "Run Test".


## Deploy on Heroku

For deployment on Heroku, you need to have the Heroku CLI installed and logged into your Heroku account. If you've not set up your Heroku CLI, you can follow [this guide](https://devcenter.heroku.com/articles/heroku-cli).

Then, follow the steps below:

1. Create a new Heroku app: `heroku create`
2. Add all your changes: `git add .`
3. Commit your changes: `git commit -m "Initial commit"`
4. Push the project to Heroku: `git push heroku master`
5. Configure environment variables according to the .env.example, using the heroku dashboard or the heroku cli:

 `heroku config:set MODE=dev SERVER_PORT=8080 DB_HOST=postgres DB_USER=root DB_PASSWORD=12323 DB_PORT=3000 WOMPI_PUBLIC_KEY= WOMPI_PRIVATE_KEY= WOMPI_INTEGRITY_KEY=`


