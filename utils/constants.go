package utils

const (
	ADMIN      = "admin"
	USER       = "user"
	MAINTAINER = "maintainer"
)

var ACCEPT_ALL_ROLES = []string{ADMIN, MAINTAINER, USER}
var ACCEPT_ADMIN_ROLES = []string{ADMIN}
var ACCEPT_MAINTAINER_ROLES = []string{ADMIN, MAINTAINER}
