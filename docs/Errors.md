### Error Handling

Errors are one of the most important aspects, 
the way you handle errors impacts the performance 
of the application in many ways such as:

1. Consistency
2. Traceability
3. Debuggability
4. Maintainability

In Go, errors are values. 
For example, you can assign an error to a 
variable in the same way you assign an integer
to a variable. Consider the following example:

```go
i, err := strconv.Atoi("95")
if err != nil {
    fmt.Print("Error: ", err)
}
````

We can handle Go errors mainly in two ways:

- We can panic (Will crash the application)
- We can handle gracefully. (Can log the error and return)

#### When to panic
When some unexpected issue happens, panic can be used. 
Mostly panic is being used to fail the application 
in case of any issue which interrupts 
the normal operation of the application.
For example, we can think of a program which uses 
a MySQL database to store data. Usually, 
the application would try to establish
a connection with the MySQL database when initializing. 
But if the application fails to establish 
the connection with the database, 
the application can’t continue to function properly.
So in this kind of scenarios, 
the application should panic.

```go
db, err := sql.Open("mysql", "username:pwd@tcp(localhost:3306)/db_name")
if err != nil {
  panic(err.Error())
}
```

Panic will result in a stack trace which 
will allow us to trace the error. Here is an example:

```bash
goroutine 1 [running]:
log.Panicln(0xc000117db8, 0x1, 0x1)
        /snap/go/7360/src/log/log.go:368 +0xae
github.com/aasumitro/svc-tgbotgo/config.AppConfig.InitDatabaseConnection()
        /home/aasumitro/Documents/Projects/Telkom/telegram-webhook/config/database_conf.go:31 +0x272
main.init.0()
        /home/aasumitro/Documents/Projects/Telkom/telegram-webhook/main.go:17 +0x105
exit status 2
```

#### When not to panic
But consider an application which allows users to login. 
What if a user tries to log in with an email which
doesn't exist in the database. 
In this kind of scenarios, we can’t panic.
We have to handle the error gracefully. 
We can log the error with the login details that 
the user entered and return an error response to the user.

There is no hard and fast rule for handling 
errors gracefully in Golang. So here is 
a better way to handle errors in your application gracefully.

Thanks to <br>
Sudaraka J.