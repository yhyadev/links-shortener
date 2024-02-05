# Links Shortener

## Description
Links Shortener is a web application that allows you to shorten your URL links. It provides an API for link shortening and a website for user interaction.

## Local Hosting
To host the application locally, follow the instructions below.

### API Setup
1. Navigate to the `api` directory:
   ```console
   $ cd api
   ```

2. Install Go and its required packages. Run the following command to install the Go packages:
   ```console
   $ go mod tidy
   ```

3. Set up a MongoDB database and add the database URI to the `.env` file in the `api` directory:
   ```env
   MONGOURI="MONGODB_URI_HERE"
   ```

4. Once you have added the MongoDB URI, start the API server:
   ```console
   $ go run main.go
   ```

### Website Setup
1. Navigate to the `website` directory:
   ```console
   $ cd website
   ```

2. Install Node.js and the required npm packages. Similar to the API setup, run the following command to install the npm packages:
   ```console
   $ npm install
   ```

3. Make a file named `.env` in the `website` directory, then add your API URL:
   ```env
   PUBLIC_API_URL="http://localhost:8080"
   ```

4. Before starting the website server, ensure that the API server is already running. Follow the API setup instructions mentioned above. Once the API server is running, start the website server using the following command:
   ```console
   $ npm run dev
   ```
