package entry

import (
	"distsys/privilege"
	usr "distsys/user"
)

type Entry struct {
	//list all models in here
	UserModule      *usr.UserModule
	PrivilegeModule *privilege.PrivilegeModule
}
