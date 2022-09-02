GORM

var user User
var users []User

*Works because destination struct is passed in
db.First(&user) --> select * from `users` order by `users`.`id` limit 1

*Works because model is specified using `db.Modesl()`
result := map[string]interface{}{}
db.Model(&User{}).First(&result) --> select * from `users` order by `users`.`id` limit 1

*Doesn't work
result := map[string]interface{}{}
db.Table("users").First(&result)

READ DOCS READ DOCS READ DOCS

*Works with Take

*Pagination

*Middleware --> block of code that is executed before and after each http request

*Environment Variables --> .env
default from golang --> os.Getenv("VARIABLE_NAME")

*Clean Architecture --> Robert C. Martin
Software design philosophy that separates the elements of a design into ring levels (onion ring)
Frameworks & Drivers -> Interface Adapters -> Application Business Rules -> Enterprise Business Rules
Pros : Testable, Maintanable, Changeable, Easy to Develop, Easy to Deploy, Independent
Cons : no cons maybe? dunno