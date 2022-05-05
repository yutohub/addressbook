# Database based addressbook

This database based addressbook can be controlled from the command line.

## Function

| Option | Value | Description |
| --- | --- | --- |
| -mode | show | Show all registrations (default) |
|   | insert | Register new address |
|   | update | Update existing address |
|   | delete | Delete existing address |
|   | reset | Delete addressbook |

## Example

- Register new address
    ```
    $ go run main.go -mode insert
    ```

    ```
    Name > A
    TEL > 11122223333
    ```
- Show all registrations
    ```
    $ go run main.go -mode show
    ```
    or
    ```
    $ go run main.go
    ```
    ```
    ----------------------------------------
    ID: 1   Name: A TEL: 11122223333
    ----------------------------------------
    ```
- Update existing address
    ```
    $ go run main.go -mode update
    ```
    ```
    ID > 1
    Name > B
    TEL > 22233334444
    ```
    ```
    ----------------------------------------
    ID: 1   Name: B TEL: 22233334444
    ----------------------------------------
    ```

## Specification

- Tools that can retain ID, name, and phone number
- ID is AUTOINCREMENT
- Data is retained even after the program is restarted
- Command line arguments
- SQLite is used for the database
- Data can be updated by specifying ID
- Data can be deleted by specifying ID

## Reference
- [gohandson](https://github.com/gohandson/gacha-ja)

## Other
I created this for my own Golang study.  
If anyone has any advice, feel free to share it with me.