window.onload = function() {
    // Open websocket connectiont to repl
    document.connection = new WebSocket("ws://" + window.location.host + "/user/repl");

    // Get the prommpt
    document.prompt = document.getElementById("prompt");
    // Focus the user on prompt
    document.prompt.focus();

    // Log messages from connection
    document.connection.onmessage = function(event) {
        putMessage(JSON.parse(event.data).output)
        console.log(JSON.parse(event.data))
    };

    // Listen for enter key from prompt
    document.prompt.addEventListener("keypress", function (e) {
        var key = e.which || e.keyCode;
        if (key === 13) { // 13 is enter
            var msg = {
                "command": document.prompt.value,
            };
            document.connection.send(JSON.stringify(msg));
            document.prompt.value = "";
        }
    });

    document.content = document.getElementById("content")

    putMessage = function(str) {
        var div = document.createElement("div");
        div.className = "output";
        var p = document.createElement("p");
        p.className = "info-text";
        p.innerHTML = str;

        div.appendChild(p);
        document.content.appendChild(div);
    }
}
