# Object-Relational Mapper

## GORM

### GET ALL USERS
    func (r *RepositoryUser) GetAllUsers() ([]model.User, error) {
        var users []model.User

        result := r.db.Order("id asc").Find(&users)

        err := result.Error

        if err != nil {
            log.Println(err)
            return nil, err
        }

        return users, nil
    }

### GET USER BY ID
    func (r *RepositoryUser) GetUserById(id int) (model.User, error) {
        var user model.User

        result := r.db.First(&user, id)

        err := result.Error

        if err != nil {
            log.Println(err)
            return user, err
        }

        return user, nil
    }

## CREATE USER
    func (r *RepositoryUser) CreateUser(name string, age int) (model.User, error) {
        user := model.User{Name: name, Age: age}

        result := r.db.Create(&user)

        err := result.Error

        if err != nil {
            log.Println(err)
            return user, err
        }

        return user, nil
    }

### UPDATE USER
    func (r *RepositoryUser) UpdateUser(id int, name string, age int) (model.User, error) {
        var user model.User

        r.db.First(&user, id)

        user.Name = name
        user.Age = age

        result := r.db.Updates(&user)

        err := result.Error

        if err != nil {
            log.Println(err)
            return user, err
        }

        return user, nil
    }

### DELETE USER
    func (r *RepositoryUser) DeleteUser(id int) (model.User, error) {
        var user model.User

        r.db.First(&user, id)

        result := r.db.Delete(&user)

        err := result.Error

        if err != nil {
            log.Println(err)
            return user, err
        }

        return user, nil
    }


READ DOCS READ DOCS READ DOCS

*Works with Take

*Pagination
## Middlewares
*Middleware --> block of code that is executed before and after each http request

## Environment Variables
*Environment Variables --> .env
default from golang --> os.Getenv("VARIABLE_NAME")

## Clean Architecture
*Clean Architecture --> Robert C. Martin
Software design philosophy that separates the elements of a design into ring levels (onion ring)
Frameworks & Drivers -> Interface Adapters -> Application Business Rules -> Enterprise Business Rules
Pros : Testable, Maintanable, Changeable, Easy to Develop, Easy to Deploy, Independent
Cons : no cons maybe? dunno