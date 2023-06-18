package scyna_const

const KEYSPACE = "scyna2"
const BASEPATH = "/scyna2"

const (
	GEN_GET_ID_URL     = BASEPATH + "/generator/id"
	GEN_GET_SN_URL     = BASEPATH + "/generator/sn"
	SESSION_CREATE_URL = BASEPATH + "/session/create"
	SETTING_WRITE_URL  = BASEPATH + "/setting/write"
	SETTING_READ_URL   = BASEPATH + "/setting/read"
	SETTING_REMOVE_URL = BASEPATH + "/setting/remove"
	START_TASK_URL     = BASEPATH + "/task/start"
	STOP_TASK_URL      = BASEPATH + "/task/stop"
)

const (
	SESSION_UPDATE_CHANNEL = KEYSPACE + ".session.update"
	SESSION_END_CHANNEL    = KEYSPACE + ".session.end"
	LOG_CREATED_CHANNEL    = KEYSPACE + ".log"
	TRACE_CREATED_CHANNEL  = KEYSPACE + ".trace"
	TAG_CREATED_CHANNEL    = KEYSPACE + ".tag"
	ENDPOINT_DONE_CHANNEL  = KEYSPACE + ".tag.endpoint"
	SETTING_UPDATE_CHANNEL = KEYSPACE + ".setting.updated."
	SETTING_REMOVE_CHANNEL = KEYSPACE + ".setting.removed."
	APP_UPDATE_CHANNEL     = KEYSPACE + ".application.updated"
	CLIENT_UPDATE_CHANNEL  = KEYSPACE + ".client.updated"
	SETTING_KEY            = KEYSPACE + ".module.config"
)
