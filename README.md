# CLI Auction house

This project was done in order to improve my **Golang** skills.

It simulates an auction house where you can create rooms and place bids for certain rewards.

Since it's an auction house, it must support several clients. Golang makes it simple to do so by using *goroutines*.

Clients connect to the server using TCP connection. This could also be done by using UDP or HTTP.

## Usage

To start the server:
1. Open a command line window or terminal and navigate to the **server** folder.
2. If on Windows, simply run: <code>server *{PORT}*</code> where **_{PORT}_** is the port you desire.
3. Otherwise you must have **Golang** installed and run <code>go build</code>

To start a client:
1. Open a command line window or terminal and navigate to the **client** folder
2. If on Windows, simply run <code>client *{IP}*:*{PORT}*</code> where *{IP}* is **_localhost_** or **_127.0.0.1_** and **_{PORT}_** is the one specified in the server
3. Otherwise you must have **Golang** installed and run <code>go build</code>
4. Repeat the process for as many clients as you want!
