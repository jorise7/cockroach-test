package code

import (
    "github.com/nogio/noggo"
    . "github.com/nogio/noggo/base"
    "time"
    "math/rand"
)

func init () {

    noggo.Data.Table("numbers", Map{
        "name": "numbers", "text": "numbers",
        "field": Map{

            "id": Map{
                "type": "int", "must": false, "name": "id", "text": "id",
            },

            "value": Map{
                "type": "int", "must": true, "name": "value", "text": "value",
                "filter": Map{ "query": true },
            },

            "status": Map{
                "type": "enum", "must": false, "name": "status", "text": "status",
                "filter": Map{ "group": true },
                "enum": Map{
                    "removed":   "removed",
                },
            },
            "changed": Map{
                "type": "datetime", "must": true, "auto": time.Now, "name": "changed", "text": "changed",
            },
            "created": Map{
                "type": "datetime", "must": true, "auto": time.Now, "name": "created", "text": "created",
            },
        },
    })

}


func PutNumbers() (error) {
    db := noggo.Data.Base("db"); defer db.Close()
    db.Begin()

    //
    _,err := db.Table("numbers").Query(Map{
        "status": nil, "value": Map{ GTE: 70, LTE: 77 },
    })
    if err != nil {
        return err
    }



    //just insert is ok.
    for i:=0; i<3;i++ {
        rand.Seed(time.Now().UnixNano())
        value := rand.Intn(500)

        _,err := db.Table("numbers").Create(Map{
            "value": value,
        })
        if err != nil {
            return err
        }
    }

    return db.Submit()
}



func GetNumbers() ([]Map) {
    db := noggo.Data.Base("db"); defer db.Close()
    _,items,_ := db.Table("numbers").Limit(0, 20, Map{
        "status": nil,
    }, Map{
        "id": DESC,
    })
    return items
}




