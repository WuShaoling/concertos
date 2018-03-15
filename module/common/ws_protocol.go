package common

const (
	P_WS_REGISTER_PLAYER        = iota //0
	P_WS_REGISTER_USER                 //1
	P_WS_INSTALL_CONTAINER             //2
	P_WS_START_CONTAINER               //3
	P_WS_START_CONTAINER_RESULT        //4
	P_WS_ERROR                         //5
)

type WebSocketMessage struct {
	MessageType int
	Sender      string
	Receiver    string
	Content     string
}
