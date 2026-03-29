import { Link } from "react-router-dom";
import "./navigation.sass";

export function Navigation() {
  return (
    <header className="topbar">
      <Link className="brand" to="/">
        Acta UI
      </Link>
      <nav className="nav">
        <Link to="/">Home</Link>
        <Link to="/signup">Signup</Link>
        <Link to="/login">Login</Link>
        <Link to="/journal">Journal</Link>
      </nav>
    </header>
  );
}
