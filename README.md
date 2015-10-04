# CMPE273-Assignment1-2015
JSON-RPC client server yahoo API program
client will send command line arguments to server
server will get info from yahoo finance api and return the result to client
we can check number of stocks of each company we can buy for percentage of budget we specify in input and their respective price of each stock 
first run server program
then run client program
ex:go run client.go GOOG:50%,YHOO:50% 6000
TO check if their is change in stock price we have already checked we have to run client program with tradeid returned by server as parameter(server should be runnning while making both scenario)
ex: go run client.go 1
