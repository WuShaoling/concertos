//Name      string
//Command   string
//Describe  string
//BaseImage string
//CPU       int
//Memory    uint64

//Id        string
//PlayerId  string
//State     int

// 启动时设置
//Ip        string


function remove(name) {
    doSend(JSON.stringify({
        "MessageType": 8,
        "Sender": "6",
        "Receiver": "",
        "Content": name
    }))
}

function stop(name) {
    doSend(JSON.stringify({
        "MessageType": 6,
        "Sender": "6",
        "Receiver": "",
        "Content": name
    }));
}

function start(name) {
    doSend(JSON.stringify({
        "MessageType": 4,
        "Sender": "6",
        "Receiver": "",
        "Content": name
    }));
}

function install() {
    var installInfo = JSON.stringify({
        "Name": "nginx-3",
        "User": "6",
        "Describe": "test nginx",
        "BaseImage": "nginx:latest",
        "Command": "NULL",
        "CPU": 1,
        "Memory": 1024
    });
    var wsmInfo = JSON.stringify({
        "MessageType": 2,
        "Sender": "6",
        "Receiver": "",
        "Content": installInfo
    });
    doSend(wsmInfo);
}

function doSend(message) {
    writeToScreen("SENT: " + message);
    websocket.send(message);
}

function writeToScreen(message) {
    var pre = document.createElement("p");
    pre.style.wordWrap = "break-word";
    pre.innerHTML = message;
}


function init() {
    websocket = new WebSocket("ws://127.0.0.1:48081/ws");
    websocket.onopen = function (evt) {
        var loginInfo = JSON.stringify({
            "MessageType": 1,
            "Sender": "6",
            "Receiver": "",
            "Content": ""
        });
        doSend(loginInfo);
    };
    websocket.onclose = function (evt) {
        writeToScreen("DISCONNECTED");
    };
    websocket.onmessage = function (evt) {
        writeToScreen('<span style="color: blue;">RESPONSE: ' + evt.data + '</span>');
    };
    websocket.onerror = function (evt) {
        writeToScreen('<span style="color: red;">ERROR:</span> ' + evt.data);
    };
}

window.addEventListener("load", init, false);
