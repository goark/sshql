# [sshql] -- Go SQL drivers over SSH

[![check vulns](https://github.com/goark/sshql/workflows/vulns/badge.svg)](https://github.com/goark/sshql/actions)
[![lint status](https://github.com/goark/sshql/workflows/lint/badge.svg)](https://github.com/goark/sshql/actions)
[![GitHub license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/goark/sshql/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/goark/sshql.svg)](https://github.com/goark/sshql/releases/latest)

This package is forked from [github.com/mattn/pqssh](https://github.com/mattn/pqssh) package.

## Usage

### PostgreSQL over SSH

```go
package main

import (
    "database/sql"
    "fmt"
    "os"

    "github.com/goark/sshql"
    "github.com/goark/sshql/pgdrv"
)

func main() {
    dialer := &sshql.Dialer{
        Hostname:   "sshserver",
        Port:       22,
        Username:   "remoteuser",
        Password:   "passphraseforauthkey",
        PrivateKey: "/home/username/.ssh/id_eddsa",
    }
    pgdrv.New(dialer).Register()

    db, err := sql.Open(pgdrv.DriverName, "postgres://dbuser:dbpassword@localhost:5432/example?sslmode=disable")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    defer dialer.Close()
    defer db.Close()

    rows, err := db.Query("SELECT id, name FROM example ORDER BY id")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    for rows.Next() {
        var id int64
        var name string
        if err := rows.Scan(&id, &name); err != nil {
            fmt.Fprintln(os.Stderr, err)
            break
        }
        fmt.Printf("ID: %d  Name: %s\n", id, name)
    }
    rows.Close()
}
```

### MySQL over SSH

```go
package main

import (
    "database/sql"
    "fmt"
    "os"

    "github.com/goark/sshql"
    "github.com/goark/sshql/mysqldrv"
)

func main() {
    dialer := &sshql.Dialer{
        Hostname:   "sshserver",
        Port:       22,
        Username:   "remoteuser",
        Password:   "passphraseforauthkey",
        PrivateKey: "/home/username/.ssh/id_eddsa",
    }
    mysqldrv.New(dialer).RegisterDial()

    db, err := sql.Open("mysql", fmt.Sprintf("dbuser:dbpassword@%s(localhost:3306)/dbname", mysqldrv.DialName))
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    defer dialer.Close()
    defer db.Close()

    rows, err := db.Query("SELECT id, name FROM example ORDER BY id")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    for rows.Next() {
        var id int64
        var name string
        if err := rows.Scan(&id, &name); err != nil {
            fmt.Fprintln(os.Stderr, err)
            break
        }
        fmt.Printf("ID: %d  Name: %s\n", id, name)
    }
    rows.Close()
}
```

## Modules Requirement Graph (UML)

![](./uml.png)

[sshql]: https://github.com/goark/sshql "goark/sshql: Go SQL drivers over SSH"
