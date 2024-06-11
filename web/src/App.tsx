import { useEffect, useState } from "react";

function App() {
  const [data, setData] = useState("")
  const msg = {
    content: "hello",
  };
  

  useEffect(() => {
    const socket = new WebSocket("http://localhost:8080/room/somerandomid");

    socket.addEventListener("open", (event: Event) => {
      socket.send(JSON.stringify(msg))
    });

    socket.addEventListener("message", (event: MessageEvent) => {
      console.log("Message from server ", event.data);
      setData(JSON.parse(event.data).message)
    });

  }, []);
  return <>hello there, the message is: "{data}"</>;
}

export default App;
