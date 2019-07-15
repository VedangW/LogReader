package main

import (
	"encoding/json"
)

func JsonToDynObj(jsonObj map[string]interface{}) DynJsonEntry {
	var dynEntry DynJsonEntry
	extra := make(map[string]interface{})

	for key, value := range jsonObj {
		switch {
		case (key == "user_id"):
			dynEntry.UserId = uint(value.(float64))
		case (key == "phone_no"):
			dynEntry.PhoneNo = value.(string)
		case (key == "password"):
			dynEntry.Password = value.(string)
		case (key == "username"):
			dynEntry.Username = value.(string)
		case (key == "order_valid"):
			dynEntry.OrderValid = value.(string)
		case (key == "file_id"):
			dynEntry.FileId = uint(value.(float64))
		default:
			extra[key] = value
		}
	}

	data := ""

	if (len(extra) != 0){
		marshalled, _ := json.Marshal(extra)
		data = string(marshalled)
	}

	dynEntry.Extra = data

	return dynEntry
}

func ResultToJson(cred string, result map[string]interface{}, dynentry DynJsonEntry) map[string]interface{} {
	final := make(map[string]interface{})

	for key, value := range result {
		final[key] = value
	}

	final["ID"] = dynentry.ID
	final["CreatedAt"] = dynentry.CreatedAt
	final["UpdatedAt"] = dynentry.UpdatedAt
	final["DeletedAt"] = dynentry.DeletedAt

	final["username"] = dynentry.Username
	final["password"] = dynentry.Password
	final["phone_no"] = dynentry.PhoneNo
	final["user_id"] = dynentry.UserId
	final["file_id"] = dynentry.FileId
	final["order_valid"] = dynentry.OrderValid

	if (cred == "user"){
		final["user_id"] = "XXXXXXXXXX"
		final["phone_no"] = "XXXXXXXXXX"
		final["password"] = "XXXXXXXXXX"
	}

	return final
}