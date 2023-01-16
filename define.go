package scyna

const (
	GEN_GET_ID_URL         = "/scyna/generator/id"
	GEN_GET_SN_URL         = "/scyna/generator/sn"
	SESSION_CREATE_URL     = "/scyna/session/create"
	SESSION_UPDATE_CHANNEL = "scyna.session.update"
	SESSION_END_CHANNEL    = "scyna.session.end"
	LOG_CREATED_CHANNEL    = "scyna.log"
	TRACE_CREATED_CHANNEL  = "scyna.trace"
	TAG_CREATED_CHANNEL    = "scyna.tag"
	ENDPOINT_DONE_CHANNEL  = "scyna.tag.endpoint"
	SETTING_WRITE_URL      = "/scyna/setting/write"
	SETTING_READ_URL       = "/scyna/setting/read"
	SETTING_REMOVE_URL     = "/scyna/setting/remove"
	SETTING_UPDATE_CHANNEL = "scyna.setting.updated."
	SETTING_REMOVE_CHANNEL = "scyna.setting.removed."
	SETTING_KEY            = "scyna.module.config"
	APP_UPDATE_CHANNEL     = "scyna.application.updated"
	CLIENT_UPDATE_CHANNEL  = "scyna.client.updated"
	AUTH_CREATE_URL        = "/scyna/auth/create"
	AUTH_GET_URL           = "/scyna/auth/get"
	AUTH_LOGOUT_URL        = "/scyna/auth/logout"
	START_TASK_URL         = "/scyna/task/start"
	STOP_TASK_URL          = "/scyna/task/stop"
)

var (
	OK                 = NewError(0, "Success")
	SERVER_ERROR       = NewError(1, "Server Error")
	BAD_REQUEST        = NewError(2, "Bad Request")
	PERMISSION_ERROR   = NewError(4, "Permission Error")
	REQUEST_INVALID    = NewError(5, "Request Invalid")
	MODULE_NOT_EXISTED = NewError(6, "Module Not Existed")
)
