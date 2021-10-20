package main

func serve(path, method string, reqBody []byte) (res string, ok bool) {
	switch path {
	case "get":
		item, errStr := getState(stateStoreName)
		if errStr != "" {
			return errStr, false
		}
		return string(item.Value), true
	case "put":
		errStr := putState(stateStoreName, string(reqBody))
		if errStr != "" {
			return errStr, false
		}
		return "OK", true
	case "del":
		errStr := delState(stateStoreName)
		if errStr != "" {
			return errStr, false
		}
		return "OK", true
	default:
		return "method not impl", false
	}
}
