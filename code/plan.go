package code

import (
    "github.com/nogio/noggo"
    . "github.com/nogio/noggo/base"
)

func init () {


    noggo.Plan.Route("numbers", Map{
        "time": "*/1 * * * * *",
        "route": Map{
            "name": "numbers", "text": "write numbers",
            "action": func(ctx *noggo.PlanContext) {
                err := PutNumbers()
                noggo.Logger.Info("putnumbers", err)
                ctx.Finish()
            },
        },
    })


}
