package main

import (
    "github.com/nogio/noggo"
    _ "github.com/nogio/noggo/core"
    "github.com/nogio/noggo/middler"
    "github.com/nogio/noggo/driver/data-cockroach"
    _ "./code"
)

func main() {

    noggo.Data.Driver("cockroach", data_cockroach.Driver())

    noggo.Use(middler.HttpLogger())
    noggo.Launch(":8080")
}
