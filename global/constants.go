package global

const (
	BOT_NAME = "EventTrackingBot"
	TOKEN = "133198388:AAHHbnm7cNHMEF6hmdehKTCMRDrFMN46n-U"
	BASE_URL = "https://api.telegram.org/bot" + TOKEN + "/"

	EVENT_PROPERTY_NAME = "name"
	EVENT_PROPERTY_DATE = "date"
	EVENT_PROPERTY_PLACE = "place"

	MESSAGE_TYPE_EVENT_PROVIDE_NAME = "event_provide_name"
	MESSAGE_TYPE_EVENT_PROVIDE_DATE = "event_provide_date"
	MESSAGE_TYPE_EVENT_PROVIDE_PLACE = "event_provide_place"
	MESSAGE_TYPE_DELETE_EVENT_PROVIDE_INDEX = "delete_event_provide_index"
	MESSAGE_TYPE_SELECT_CURRENT_EVENT = "select_current_event"

	MESSAGE_PARSE_MODE_MARKDOWN = "Markdown"
	MESSAGE_PARSE_MODE_HTML = "HTML"

	COMMAND_NAME_CREATE_EVENT = "create_event"
	COMMAND_NAME_SET_WHEN = "set_when"
	COMMAND_NAME_SET_WHERE = "set_where"
	COMMAND_NAME_DELETE_EVENT = "delete_event"

	GET_UPDATES_DEFAULT_TIMEOUT = 0
	GET_UPDATES_DEFAULT_LIMIT = 100

	PERSISTENCE_RETRIES = 2
)

var (
	PROPERTY_TO_MESSAGE_TYPE_MAP = map[string]string {
		EVENT_PROPERTY_NAME: MESSAGE_TYPE_EVENT_PROVIDE_NAME,
		EVENT_PROPERTY_DATE: MESSAGE_TYPE_EVENT_PROVIDE_DATE,
		EVENT_PROPERTY_PLACE: MESSAGE_TYPE_EVENT_PROVIDE_PLACE,
	}
)