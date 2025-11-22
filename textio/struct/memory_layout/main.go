package memorylayout

// Before
/*
type contact struct {
	sendingLimit int32
	userID       string
	age          int32
}

type perms struct {
	canSend         bool
	canReceive      bool
	permissionLevel int
	canManage       bool
}
*/

// After, re-order fields in struct to reduce size of its
type contact struct {
	userID       string
	sendingLimit int32
	age          int32
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}
