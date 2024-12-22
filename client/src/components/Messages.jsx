import React from "react";

const Messages = ({ messages, user }) => {
    return (
        <div>
            {messages.map((message, index) => (
                <div key={index}>
                    <div>
                        {" "}
                        {message.user === user ? (
                            <span style={{ color: "green" }}>You</span>
                        ) : (
                            <span style={{ color: "red" }}>{message.user}</span>
                        )}{" "}
                    </div>
                    <div>{message.body}</div>
                </div>
            ))}
        </div>
    );
};

export default Messages;
