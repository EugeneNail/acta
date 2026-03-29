import { BrowserRouter, Link, Route, Routes } from "react-router-dom";

function Layout({
  eyebrow,
  title,
  description,
}: {
  eyebrow: string;
  title: string;
  description: string;
}) {
  return (
    <main className="shell">
      <header className="topbar">
        <Link className="brand" to="/">
          Acta UI
        </Link>
        <nav className="nav">
          <Link to="/">Home</Link>
          <Link to="/login">Login</Link>
          <Link to="/journal">Journal</Link>
          <Link to="/settings">Settings</Link>
        </nav>
      </header>

      <section className="hero">
        <p className="eyebrow">{eyebrow}</p>
        <h1>{title}</h1>
        <p className="lede">{description}</p>
        <div className="actions">
          <a className="primaryAction" href="/api/v1/auth/login">
            Auth API
          </a>
          <a className="secondaryAction" href="/api/v1/journal/habits">
            Journal API
          </a>
        </div>
      </section>
    </main>
  );
}

function HomePage() {
  return (
    <Layout
      eyebrow="Acta / Home"
      title="Build daily momentum with a calmer interface."
      description="The UI service is ready. Next steps are wiring real auth state, journal screens, and API calls through the shared proxy."
    />
  );
}

function LoginPage() {
  return (
    <Layout
      eyebrow="Acta / Login"
      title="A focused gateway for authentication flows."
      description="This test route stands in for the future login and token refresh screens."
    />
  );
}

function JournalPage() {
  return (
    <Layout
      eyebrow="Acta / Journal"
      title="A route reserved for journal and habit workspaces."
      description="This test route confirms the router can serve a journal surface independently from backend API paths."
    />
  );
}

function SettingsPage() {
  return (
    <Layout
      eyebrow="Acta / Settings"
      title="One place for profile and application preferences."
      description="This route exists only to verify client-side navigation and direct browser reloads through the proxy."
    />
  );
}

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/journal" element={<JournalPage />} />
        <Route path="/settings" element={<SettingsPage />} />
      </Routes>
    </BrowserRouter>
  );
}
