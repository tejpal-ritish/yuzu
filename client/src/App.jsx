import "./App.css";
import {
    Routes,
    Route
} from "react-router-dom"
import Chat from "./components/Chat";
import Login from "./components/Login";

function App() {
    return (
        <div>
            <h1>Chat are we cooked????</h1>
            <Routes>
                <Route index element={<Login />} />
                <Route path="/chat" element={<Chat />} />
            </Routes>
        </div>
    );
}

export default App;
