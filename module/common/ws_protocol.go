package common

const (
	P_WS_REGISTER_PLAYER     = iota //0
	P_WS_REGISTER_USER              //1
	P_WS_INSTALL_CONTAINER          //2
	P_WS_INSTALL_CONTAINER_R        //3
	P_WS_START_CONTAINER            //4
	P_WS_START_CONTAINER_R          //5
	P_WS_STOP_CONTAINER             //6
	P_WS_STOP_CONTAINER_R           //7
	P_WS_REMOVE_CONTAINER           //8
	P_WS_REMOVE_CONTAINER_R         //9
	P_WS_ERROR                      //10
)

type WebSocketMessage struct {
	MessageType int
	Sender      string
	Receiver    string
	Content     string
}
