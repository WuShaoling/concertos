package common

const (
	P_WS_REGISTER_PLAYER   = "register_player"
	P_WS_INSTALL_CONTAINER = "install_container"
	P_WS_ERROR             = "error"
)

type WebSocketMessage struct {
	MessageType string
	Content     []byte
}
