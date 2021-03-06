package persistence

const (
	KEY_WAITIING_RESPONSE_TO = "user:%d:waiting_response_to:%d"
	KEY_MESSAGE_TYPE = "message:%d:type"

	KEY_EVENT = "group:%d:event:%s"
	KEY_EVENT_PARTICIPANTS = "group:%d:event:%s:participants"

	KEY_CURRENT_EVENT = "group:%d:current_event"

	KEY_GROUP_EVENTS = "group:%d:events"
)