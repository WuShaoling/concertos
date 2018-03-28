package common

const (
	P_WS_REGISTER_PLAYER     = iota //0
	P_WS_REGISTER_USER              //1
	P_WS_CONTAINER_INSTALL          //2
	P_WS_CONTAINER_INSTALL_R        //3
	P_WS_CONTAINER_START            //4
	P_WS_CONTAINER_START_R          //5
	P_WS_CONTAINER_STOP             //6
	P_WS_CONTAINER_STOP_R           //7
	P_WS_CONTAINER_REMOVE           //8
	P_WS_CONTAINER_REMOVE_R         //9
	P_WS_ERROR                      //10
)

type WebSocketMessage struct {
	MessageType int
	Sender      string
	Receiver    string
	Content     string
}
