import './style.css'
import {createClient} from "@connectrpc/connect";
import {createGrpcWebTransport} from "@connectrpc/connect-web";
import {GreetService} from "./gen/greet/v1/greet_pb.ts";

// 1. Configure the Transport
// This tells Connect to use the gRPC-Web protocol over HTTP
const transport = createGrpcWebTransport({
  baseUrl: "http://localhost:8080",
});

// 2. Create the Client
// This creates a strongly-typed client using the generated service definition
const client = createClient(GreetService, transport);


document.querySelector<HTMLDivElement>('#app')!.innerHTML = `
  <div>
    <h1>Connect-Go Web Client</h1>
    <div class="card">
      <input id="name-input" type="text" placeholder="Enter your name" value="World" />
      <button id="greet-btn" type="button">Greet Server</button>
    </div>
    <p id="response-text" class="read-the-docs"></p>
  </div>    
`

const btn = document.getElementById("greet-btn");
const input = document.getElementById("name-input") as HTMLInputElement;
const responseText = document.getElementById("response-text");

btn?.addEventListener("click", async () => {
  if (!input.value) return;

  try {
    // 4. Make the RPC Call
    const res = await client.greet({
      name: input.value,
    });

    // 5. Display Result
    if (responseText) {
      responseText.innerText = `Server says: ${res.greeting}`;
      responseText.style.color = "green";
    }
  } catch (err) {
    console.error(err);
    if (responseText) {
      responseText.innerText = `Error: ${err}`;
      responseText.style.color = "red";
    }
  }
});