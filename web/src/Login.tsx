import { PropsWithChildren, ReactNode, useState } from "react";

// TODO: For build, change this to just '/login'.
const LOGIN_ENDPOINT = "http://localhost:8000/login";
const ERR_MESSAGE = "Invalid username or password.";

async function handleLogin(username: string, password: string): Promise<boolean> {
  const response = await fetch(LOGIN_ENDPOINT, {
    method: "POST",
    body: JSON.stringify({ username, password }),
    headers: {
      "Content-Type": "application/json"
    }
  });
  return response.ok;
}

function Login(props: PropsWithChildren): ReactNode {
  // Global app auth state.
  const [user, setUser] = useState<string>();

  // Local login form state.
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string>();

  if (user !== undefined) return props.children;
  return (
    <div className="">
      {error && (
        <div className="">{error}</div>
      )}
      <div className="">
        <input 
          className=""
          type="text" 
          placeholder="Username" 
          value={username} 
          onChange={e => setUsername(e.target.value)} />
        <input 
          className=""
          type="password" 
          placeholder="Password" 
          value={password} 
          onChange={e => setPassword(e.target.value)} />
        <button 
          className=""
          type="button" 
          onClick={() => handleLogin(username, password)
            .then(ok => ok ? setUser(username) : setError(ERR_MESSAGE))}>
          Login
        </button>
      </div>
    </div>
  );
}

export default Login
