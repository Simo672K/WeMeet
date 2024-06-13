import { useEffect } from "react";

function App() {
  useEffect(() => {
    const socket = new WebSocket("ws://localhost:8080/room/id");

    socket.addEventListener("open", (event: Event) => {
      socket.send("hello, world! This is a message from react to golang");
    });

    // socket.addEventListener("message", (event: MessageEvent) => {
    //   console.log("Message from server ", event.data);
    // });
  }, []);
  return <h1>hello there</h1>;
}

export default App;
