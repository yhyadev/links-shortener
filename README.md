# Links Shortener

## Description
Links Shortener is a web application that allows you to shorten your URL links. It provides an API for link shortening and a front-end website for user interaction.

## Local Hosting
To host the application locally, follow the instructions below.

### API Setup
1. Navigate to the `api` directory:
   ```bash
   $ cd api
   ```

2. Install Go and its required packages. Run the following command to install the Go packages:
   ```bash
   $ go mod tidy
   ```

3. Set up a MongoDB database and add the database URI to the `.env` file in the `api` directory:
   ```env
   MONGOURI="MONGODB_URI_HERE"
   ```

4. Once you have added the MongoDB URI, start the API server:
   ```bash
   $ go run main.go
   ```

### Front-End Setup
1. Navigate to the `website` directory:
   ```bash
   $ cd website
   ```

2. Install Node.js and the required npm packages. Similar to the API setup, run the following command to install the npm packages:
   ```bash
   $ npm install
   ```

3. Make a file named `.env` in the `website` directory, then add your API URL:
   ```env
   PUBLIC_API_URL="http://localhost:8080"
   ```

4. Before starting the website server, ensure that the API server is already running. Follow the API setup instructions mentioned above. Once the API server is running, start the website server using the following command:
   ```bash
   $ npm run dev
   ```

## Motivation
I developed this application for the purpose of learning and providing valuable resources to others. It serves as a reference for creating a link shortener in Go and demonstrates how to use Axios in Svelte. Additionally, I created this project for personal use, as I needed to complete more projects efficiently within a shorter timeframe.

## Conclusion
I believe I have done a great job in developing this application, which not only benefits me but also provides resources and references to other developers. Thank you for your time, and have a great day, fellow developers!
