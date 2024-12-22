import React, { useState, useEffect } from "react";
import Messages from "./Messages";

const Chat = () => {
    const [input, setInput] = useState("");
    const [messages, setMessages] = useState([]);
    const [ws, setWs] = useState(null);
    const [user, setUser] = useState(null);

    useEffect(() => {
        setUser(Math.floor(Math.random() * 1000));
        const socket = new WebSocket("ws://localhost:3000/ws");
        setWs(socket);

        socket.onmessage = (event) => {
            const newMessage = JSON.parse(event.data);
            setMessages((prevMessages) => [...prevMessages, newMessage]);
        };

        return () => {
            socket.close();
        };
    }, []);

    const sendMessage = () => {
        if (!input) {
            return;
        }

        const message = {
            user: user,
            body: input,
            timestamp: new Date().toISOString(),
        };
        if (ws) {
            ws.send(JSON.stringify(message));
        }
        setInput("");
    };

    return (
        <div>
            <div className="container">
                <div className="messages">
                    <Messages messages={messages} user={user} />
                    <form onSubmit={(e) => e.preventDefault()}>
                        <input
                            type="text"
                            name="text"
                            id="text"
                            value={input}
                            onChange={(e) => {
                                setInput(e.target.value);
                            }}
                        />
                        <button id="send" onClick={sendMessage}>
                            Send
                        </button>
                    </form>
                </div>
            </div>
        </div>
    );
};

export default Chat;
