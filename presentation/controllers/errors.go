package controllers

const (
	messageCantFind      = `message": "cant find `
	cantFindUser         = `user with nickname `
	cantFindThread       = `thread with slug or id `
	cantFindForum        = `forum with slug `
	cantFindParentOrUser = `parent or parent in another thread`
	cantFindPost         = `post with id `
	emailUsed            = ` has already taken by another user`
)

const (
	errorUniqueViolation     = `pq: unique_violation`
	errorPqNoDataFound       = `pq: no_data_found`
	errorForeignKeyViolation = `pq: foreign_key_violation`
)
