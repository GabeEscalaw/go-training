# GoTraining Final Project | Watch Collection Tracker

Welcome to the Watch Collection Tracker, an application that aims to serve as a virtual checklist for those who plan on starting their very own watch collection.

After running the server, the application will be able to handle multiple users by registering them to the database where users can login in and log out of their accounts. While a user is logged in, they will be able to add watches to their collection by providing the brand, model, width, and price of each watch. There is a counter for each watch collected from the ones they listed which is done by initiating a command to mark whether it is collected or not. The users also have the ability to delete watches from their collection which will adjust the counters as needed.

This application is deployed and ran using Docker which uses a Command Line Interface (CLI) to interact with the application. It uses http API as well as a JSON file (`users.json`) to access as locally stored file-based database. It also comes with a video presentation in the `Client` folder to demonstrate what the application does in detail.

Main Features:

- Application | http API and client CLI

- Storage | file-based (json)

- Multiple Users | Registering and Logging in via username and password

- List of Watches | Each user has a unique collection where they will be able to list the watches they want to collect. Each watch has their own details with a field to mark if the user has been collected already.

- Editable List | Users can Add, Remove, and Mark watches from their collection which will always be displayed for the user.

## Instructions

`-cmd=` command typed before any of the following and their required inputs:

1. `-in=` login with a registered `-u=` and `-p=`
2. `-out=` logout of current account
3. `-reg=` register a new user into the database with `-u=` and `-p=` del = delete an existing user from the database with `-u=`and`-p=`
4. `-list=` lists all the watches the user has without modifying anything
5. `-addW=` add a new watch into the user's collection database with `-brand=` `-model=` `-width=` `-price=` and display current collection
6. `-delW=` delete an existing watch from the user's collection database with `-brand=` `-model=` and display current collection
7. `-mark=` mark an existing watch from the user's collection database with `-brand=` `-model=` to mark it as COLLECTED if not yet collected and vice verse, as well as displaying current collection

| Commands | Required Inputs (after -cmd="command")                      |
| -------- | ----------------------------------------------------------- |
| `in`     | `-u=username` `-p=password`                                 |
| `out`    | ` `                                                         |
| `reg`    | `-u=username` `-p=password`                                 |
| `del`    | `-u=username` `-p=password`                                 |
| `list`   | ` `                                                         |
| `addW`   | `-brand=brand` `-model=model` `-width=width` `-price=price` |
| `delW`   | `-brand=brand` `-model=model`                               |
| `mark`   | `-brand=brand` `-model=model`                               |

Notes:

Enclose inputs that have spaces inbetween them with quotation marks.

Type `-h` for a reminder of these commands in app.

Please provide your own measurements for the `-width=` (e.g. mm) and `-price=` (e.g. $).

Using a command that is not included in this list will result into a Bad Request: Error 404 message.

Examples:

```go
-cmd=reg -u=user -p=pass // Will register a user with "user" and "pass" as their credentials
-cmd=addW -brand="Audemars Piguet" -model="Royal Oak" -width=39mm -price=$32000 // Will register a watch with the following details (Note: spaces in input require quotations)
```

## Databases

For a better understanding of what the locally stored database — `users.json` file contains, the following tables display how they are stored.

UserDatabase - Contains all the Users
| Field | Data Type | Description |
|-------|-----------|-------------|
| mu | sync.Mutex | Mutual Excclusion (not part of json file) |
| Users | []UserInfo | Slice containing all the Users |

UserInfo - Contains all the Users' information
| Field | Data Type | Description |
|-------|-----------|-------------|
| Username | string | Unique username key for Users |
| Password | string | Password key attached to a Username |
| IsLoggedIn | bool | Marks the active user (only one can be active at a time) |
| Watches | []WatchInfo | Slice containing the watches of a User |

WatchInfo - Contains all the details of each Watch in a User's collection
| Field | Data Type | Description |
|-------|-----------|-------------|
| Brand | string | Brand of the watch |
| Model | string | Model of the watch |
| Width | string | Case Size of the watch |
| Price | string | Price of the watch |
| Collected | bool | State of watch if collected or not |

Note:

There's a counter for each watch above the list of watches so if a watch is marked as `"collected":true`, it will be counted from the total (e.g. `Watches Collected: 1/3`)

## Error Handling

Logging in

- Logging in to the same account will display a notification
- Error 404 upon retrieving list of users when database is not found or empty (only for testing)
- Error 500 upon having any problems with encoding Users database
- Error 400 upon having any problems with decoding the json request
- Error 400 upon entering incorrect `Password` of an existing `Username`
- Error 400 upon entering a `Username` that does not exist (applies to blank inputs)
- Error 400 upon attempting to log in on a different account without logging out beforehand
- Error 405 upon attempting to method other than `GET` and `POST`

Logging out

- Logging out while not logged in will display a notification
- Error 404 upon retrieving list of users when database is not found or empty (only for testing)
- Error 500 upon having any problems with encoding Users database
- Error 400 upon having any problems with decoding the json request
- Error 405 upon attempting to method other than `GET` and `POST`

Registering a User

- Error 400 upon having any problems with decoding the json request
- Error 409 upon attempting to register a user whose `Username` already exists in the database
- Error 405 upon attempting to register a user with a blank in either the `Username` or `Password` or both
- Error 405 upon attempting to method other than `POST`

Deleting a User

- Error 404 upon retrieving list of users when database is not found or empty (only for testing)
- Error 500 upon having any problems with encoding Users database
- Error 400 upon having any problems with decoding the json request
- Error 400 upon attempting to delete a user without any users in the database to delete
- Error 400 upon attempting to delete a user with the incorrect `Password`
- Error 400 upon attempting to delete a user whose credentials do not match any user from the database
- Error 405 upon attempting to method other than `GET` and `POST`

List

- Listing while not logged in will display a notification
- Error 404 upon attempting to use this command without an empty user database
- Error 500 upon having any problems with encoding Users database
- Error 405 upon attempting to method other than `GET`

Adding a Watch

- Adding a watch while not logged in will display a notification
- Error 400 upon having any problems with decoding the json request
- Error 405 upon attempting to add a watch that is already part of their collection
- Error 405 upon attempting to add a watch with incomplete details
- Error 405 upon attempting to method other than `POST`

Removing a Watch

- Removing a watch while not logged in will display a notification
- Removing a watch that does not match any watch from the user's database will display a notification
- Error 400 upon having any problems with decoding the json request
- Error 405 upon attempting to method other than `POST`

Marking a Watch

- Marking a watch while not logged in will display a notification
- Marking a watch that does not match any watch from the user's database will display a notification
- Error 400 upon having any problems with decoding the json request
- Error 405 upon attempting to mark a watch with a blank in either the `Username` or `Password` or both
- Error 405 upon attempting to method other than `POST`
