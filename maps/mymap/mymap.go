package mymap

func SetUserSet(settings map[string]map[string]string, user, key, value string) {
	
	for userItem := range settings {
		if userItem == user {
			settings[user][key] = value
		} else {
			settings[user] = map[string]string{
				key: value,
			}
		}
	}
	
}