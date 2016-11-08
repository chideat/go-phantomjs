package phantomjs

type ErrorCode int

const (
	ErrorCode_Success                         ErrorCode = 0
	ErrorCode_NoSuchElement                             = 7
	ErrorCode_NoSuchFrame                               = 8
	ErrorCode_UnknownCommand                            = 9
	ErrorCode_StaleElementRefrence                      = 10
	ErrorCode_ElementNotVisible                         = 11
	ErrorCode_InvalidElementState                       = 12
	ErrorCode_UnknownError                              = 13
	ErrorCode_ElementIsNotSelectAble                    = 15
	ErrorCode_JavascriptError                           = 17
	ErrorCode_XPathLoopupError                          = 19
	ErrorCode_Timeout                                   = 21
	ErrorCode_NoSuchWindow                              = 23
	ErrorCode_InvalidCookieDomain                       = 24
	ErrorCode_UnableToSetCookie                         = 25
	ErrorCode_UnexceptedAlertOpen                       = 26
	ErrorCode_NoAlertOpen                               = 27
	ErrorCode_ScriptTimeout                             = 28
	ErrorCode_InvalidElementCoordinates                 = 29
	ErrorCode_ImeNotAvaliable                           = 30
	ErrorCode_ImeEngineActivAtionFailed                 = 31
	ErrorCode_InvalidSelector                           = 32
	ErrorCode_MoveTargetOutOfBounds                     = 34
	ErrorCode_InvalidXpathSelector                      = 51
	ErrorCode_InvalidXpathSelectorReturnTyper           = 52
	ErrorCode_MethodNotAllowed                          = 405
)

func (code ErrorCode) Error() string {
	switch code {
	case 0:
		return "OK"
	case 7:
		return "no such element"
	case 8:
		return "no such frame"
	case 9:
		return "unknown command"
	case 10:
		return "stale elemnt refrence"
	case 11:
		return "element not visible"
	case 12:
		return "invalid element state"
	case 13:
		return "unknown error"
	case 15:
		return "element not selectable"
	case 17:
		return "javascript error"
	case 19:
		return "invalid selector"
	case 21:
		return "timeout"
	case 23:
		return "no such window"
	case 24:
		return "invalid cookie domain"
	case 25:
		return "unable to set cookie"
	case 26:
		return "unexpected a ert open"
	case 27:
		return "no such alert"
	case 28:
		return "script timeout"
	case 29:
		return "invalid element coordinates"
	case 30:
		return "ime not available"
	case 31:
		return "ime engine activation failed"
	case 32:
		return "invalid selecotr"
	case 34:
		return "move target out of bounds"
	case 51:
		return "invalid selector"
	case 52:
		return "invalid seelctor"
	case 405:
		return "unsupported operation"
	default:
		return "unkown error code"
	}
}

func (c ErrorCode) String() string {
	return c.Error()
}
