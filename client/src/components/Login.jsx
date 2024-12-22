import React, { useState } from "react";
import axios from "axios";
import { Navigate, redirect } from "react-router-dom";

const Login = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const submitForm = (e) => {
        e.preventDefault();
        if (username === "" || password === "") {
            return;
        }

        axios
            .post(
                "http://localhost:3000/login",
                {
                    username: username,
                    password: password,
                },
                {
                    headers: {
                        "Content-Type": "application/x-www-form-urlencoded",
                    },
                }
            )
            .then((res) => {
                if (res.status === 200) {
                    redirect("/chat");
                } else {
                    console.log(res.data);
                }
            })
            .catch((err) => {
                console.error(err);
            });
    };

    return (
        <div>
            <h2>Login</h2>
            <form method="post">
                <label htmlFor="username">Username: </label>
                <input
                    type="text"
                    name="username"
                    id="username"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                />
                <br />
                <br />

                <label htmlFor="password">Password: </label>
                <input
                    type="password"
                    name="password"
                    id="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <br />
                <br />

                <button type="submit" onClick={submitForm}>
                    Login
                </button>
            </form>
        </div>
    );
};

export default Login;
