package code

import (
    "github.com/nogio/noggo"
    . "github.com/nogio/noggo/base"
)

func init () {

    noggo.Http.Route("home", Map{
        "uri": "/",
        "route": Map{
            "name": "home", "text": "home",
            "action": func(ctx *noggo.HttpContext) {
                ctx.Text("here is home.")
            },
        },
    })


    noggo.Http.Route("numbers", Map{
        "uri": "/numbers",
        "route": Map{
            "name": "numbers", "text": "numbers",
            "action": func(ctx *noggo.HttpContext) {
                numbers := GetNumbers()
                ctx.Json(numbers)
            },
        },
    })

}
