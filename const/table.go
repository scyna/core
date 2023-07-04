package scyna_const

const (
	MODULE_TABLE              = KEYSPACE + ".module"
	CLIENT_TABLE              = KEYSPACE + ".client"
	CLIENT_USE_ENDPOINT_TABLE = KEYSPACE + ".client_use_endpoint"
	SETTING_TABLE             = KEYSPACE + ".setting"
	GEN_ID_TABLE              = KEYSPACE + ".gen_id"
	GEN_SN_TABLE              = KEYSPACE + ".gen_sn"
	SESSION_TABLE             = KEYSPACE + ".session"
	SESSION_LOG_TABLE         = KEYSPACE + ".session_log"
	TODO_TABLE                = KEYSPACE + ".todo"
	DOING_TABLE               = KEYSPACE + ".doing"
	TASK_TABLE                = KEYSPACE + ".task"
	MODULE_HAS_TASK_TABLE     = KEYSPACE + ".module_has_task"
	TRACE_TABLE               = KEYSPACE + ".trace"
	CLIENT_HAS_TRACE_TABLE    = KEYSPACE + ".client_has_trace"
	SPAN_TABLE                = KEYSPACE + ".span"
	TAG_TABLE                 = KEYSPACE + ".tag"
	LOG_TABLE                 = KEYSPACE + ".log"
	APPLICATION_TABLE         = KEYSPACE + ".application"
	APP_HAS_TRACE_TABLE       = KEYSPACE + ".app_has_trace"
)
