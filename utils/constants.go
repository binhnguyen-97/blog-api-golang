package utils

const (
	ADMIN      = "admin"
	USER       = "user"
	MAINTAINER = "maintainer"
	// Mail Mime
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

var ACCEPT_ALL_ROLES = []string{ADMIN, MAINTAINER, USER}
var ACCEPT_ADMIN_ROLES = []string{ADMIN}
var ACCEPT_MAINTAINER_ROLES = []string{ADMIN, MAINTAINER}
