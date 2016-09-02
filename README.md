# Github Multi Repo Dependency Tool

## Synopsis

## Usage

## API Reference 

The application contains four public API routes: 

- GET: "/" (Server Test Ping)
- POST: "/" (Main Repo Hook)
- POST: "/dep" (Dependency Repo Hook) 
- GET: "/flush" (Flush The Database)

###Server Test Ping: 

Sending a GET request to the root API("/") will ping the server and print out `Server is up and running`

###Main Repo Hook: 

Sending a POST request to the root API("/") will trigger the functionality that checks if the Pull Request has a dependency, then locks it if it does. The POST request expects data in the format of Githubs Pull Request Hook. 

###Dependency Repo Hook:

Sending a POST request to the "/dep" of the API will trigger the functionality that unlocks the Pull Request from the main repo, allowing it to be merged in. The POST request expects data in the format of Githubs Pull Request Hook.

###Flush The Database: 

Sending a GET request to the "/flush" of the API will clear the database of all it's stored values. (**WARNING:** Don't flush the database if you have a Pull Request open that is currently being blocked by a dependency)

## Deployment Via Heroku

## License
