General Instructions
Please comment your code as much as clear as possible


1.00
The first objective is to create the most simple tcp server (without a client).
In this version, not buffer is required
Define the tcp port in a variable
Create a variable "start_time" that keep at what time the server started working
Create a variable "number_messages" that keep the number of messages that the server has received for the clients.
create a function that allows to write all the incoming messages from clients in a file.
In the file in each line write the variable "start_time" - "number_messages" and the x message
we are going to test this using a simple echo -n "this is my first message :)" | nc localhost 3333
please explain
-what are the advantages of using a buffer, if we are assuming that we are going to receive a large volume of messages
-what are the options to implement a buffer (ex bufio)


This is what I have so far https://gist.github.com/maxantonio/fada382613de86b33bc1600234b84f47
*********************************************************************
1.01
