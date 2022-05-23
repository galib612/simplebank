several way to create CURD REestFul api
    1.) database/sql (low-level)
        very fast & straightforward
        manual mapping SQL fields to variables
        easy to make mistakes, not caught until runtime.

    2.) GORM (high-level)
        CRUD functions alreadby implemented
        very short production code
        must learn to write queries using gorm's functions
        run slowly on high load

    3.) SQLX (mid level)
        quite fast & easy to use
        fields mapping via query text & struct tags
        failure won't occur until runtime

    4.) SQLC
        very fast & easy to use
        Automatic code generation
        catch SQL query errors before generating codes
        Full support Postgres. Mysql is experimental {so if u are using postgres SQLC is right tool} 