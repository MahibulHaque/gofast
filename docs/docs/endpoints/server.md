---
sidebar_position: 1
sidebar_label: Server Endpoints
---

# Server Endpoints

## Testing Endpoints with CURL and WebSocat

Testing endpoints is an essential part of ensuring the correctness and functionality of your app. Depending on what options are used for gofast project creation, you have various endpoints for testing your init application status.

Before proceeding, ensure you have the following tools installed:

- [Curl](https://curl.se/docs/manpage.html): A command-line tool for transferring data with URLs.
- [WebSocat](https://github.com/vi/websocat): A command-line WebSocket client.
You can utilize alternative tools that support the WebSocket protocol to establish connections with the server. WebSocat is an open-source CLI tool, while [Postman](https://www.postman.com/) or [Insomnia](https://insomnia.rest/) serves as a GUI tool specifically designed for testing APIs and WebSocket functionality.

## Hello world Endpoint

To test the Hello World endpoint, execute the following curl command:
```bash
curl http://localhost:<YOUR_APPLICATION_PORT>
```

Sample Output:
```bash
{"message": "Hello World"}
```
