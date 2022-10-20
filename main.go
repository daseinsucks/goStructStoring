package main

import (
	"fmt"
	"strconv"

	"encoding/json"

	pogreb "github.com/akrylysov/pogreb"
)

/* type user struct {
	name    string
	surname string
} */

type telegramUser struct {
	Chatid        int64
	Dialog_status int64
	Repo          string
	Dao           string
}

func main() {

	undefineduser := telegramUser{0, 0, "", ""}

	usertemplate := telegramUser{0, 0, "0", "0"}
	//initialising the database. It's stored locally in the "database" folder
	/* 	db, _ := pogreb.Open("database", nil)

	   	//db.Put puts key:value pair into the database
	   	user1 := user{"John", "Doe"}
	   	db.Put([]byte("1"), []byte(user1.name+":"+user1.surname))

	   	//db.Get gets the value for the given key. Note that it's received in bytes
	   	thing, err := db.Get([]byte("1"))
	   	thingstring := string(thing)
	   	//fmt.Println(thingstring)
	   	//log.Printf("%s", thing)

	   	//example of string manipulation
	   	foo := strings.Split(thingstring, ":")
	   	user2 := user{foo[0], foo[1]}

	   	//prints "Doe"
	   	fmt.Println(user2.surname) */

	tgbd, _ := pogreb.Open("telegramdb", nil)
	/*
		//chatstring := fmt.Sprint(user.Chatid)
		//statusstring := fmt.Sprint(user.Dialog_status)

		//tgbd.Put([]byte(chatstring), []byte(statusstring+":"+user.Repo+":"+user.dao))

		userId := "125"

		user := restoreUserViaString(tgbd, userId)
		if user == undefineduser {
			fmt.Println("User not found! Creating an empty template for that tgid.")
			tgbd.Put([]byte(userId), []byte("0"+":"+"0"+":"+"0"))
		} else {
			fmt.Println("Found a record for this user!", user)
		}
	*/

	//////////////////////////////////////////////////////////////
	//    Example of user recovery using json format            //
	/////////////////////////////////////////////////////////////

	/* exampleUser := telegramUser{1337, 22, "link2", "0x732648732646386486332"}

	user2idtostring := fmt.Sprint(exampleUser.Chatid)
	userjson, _ := json.Marshal(exampleUser)

	tgbd.Put([]byte(user2idtostring), userjson) */

	userIdToCheck := "1245443"

	restoredUser := restoreUserViaJson(tgbd, userIdToCheck)

	if restoredUser == undefineduser {
		fmt.Println("User not found! Creating an empty template for that tgid.")
		newUser := usertemplate
		useridtoint, _ := strconv.ParseInt(userIdToCheck, 10, 64)
		newUser.Chatid = useridtoint
		newUserjson, _ := json.Marshal(newUser)
		tgbd.Put([]byte(userIdToCheck), newUserjson)

	} else {
		fmt.Println("Found a record for this user!", restoredUser)
	}

}

func restoreUserViaJson(database *pogreb.DB, chatid string) telegramUser {
	defer handlePanic()

	thing, _ := database.Get([]byte(chatid))

	var userstruct telegramUser
	json.Unmarshal(thing, &userstruct)
	return userstruct
}

func handlePanic() {

	// detect if panic occurs or not
	a := recover()

	if a != nil {
		fmt.Println("RECOVER", a)
	}

}
