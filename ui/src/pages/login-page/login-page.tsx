import type { FormEvent } from "react";
import { useState } from "react";
import { Link } from "react-router-dom";
import "./login-page.sass";

type LoginResponse = {
  accessToken: string;
  refreshToken: string;
};

type ValidationErrors = Record<string, string>;

export function LoginPage() {
  const [email, setEmail] = useState("login.user@example.com");
  const [password, setPassword] = useState("Strong123!");
  const [errors, setErrors] = useState<ValidationErrors>({});
  const [tokens, setTokens] = useState<LoginResponse | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setIsSubmitting(true);
    setErrors({});
    setTokens(null);

    try {
      const response = await fetch("/api/v1/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email,
          password,
        }),
      });

      const payload = (await response.json()) as LoginResponse | ValidationErrors;

      if (!response.ok) {
        setErrors(payload as ValidationErrors);
        return;
      }

      const nextTokens = payload as LoginResponse;
      localStorage.setItem("accessToken", nextTokens.accessToken);
      localStorage.setItem("refreshToken", nextTokens.refreshToken);
      setTokens(nextTokens);
    } catch (_error) {
      setErrors({ request: "Request failed. Try again." });
    } finally {
      setIsSubmitting(false);
    }
  }

  return (
    <main className="login-page">
      <section className="login-page__card">
        <h1 className="login-page__title">Login</h1>

        <form className="login-page__form" onSubmit={handleSubmit}>
          <label className="login-page__field">
            <span className="login-page__label">Email</span>
            <input
              className="login-page__input"
              autoComplete="email"
              name="email"
              type="email"
              value={email}
              onChange={(event) => setEmail(event.target.value)}
            />
          </label>

          <label className="login-page__field">
            <span className="login-page__label">Password</span>
            <input
              className="login-page__input"
              autoComplete="current-password"
              name="password"
              type="password"
              value={password}
              onChange={(event) => setPassword(event.target.value)}
            />
          </label>

          <button className="login-page__submit" disabled={isSubmitting} type="submit">
            {isSubmitting ? "Logging in..." : "Login"}
          </button>
        </form>

        {Object.keys(errors).length > 0 && (
          <ul className="login-page__errors">
            {Object.entries(errors).map(([field, message]) => (
              <li className="login-page__error" key={`${field}-${message}`}>
                <span className="login-page__error-field">{field}</span>
                <strong className="login-page__error-message">{message}</strong>
              </li>
            ))}
          </ul>
        )}

        {tokens !== null && (
          <div className="login-page__tokens">
            <div className="login-page__token">
              <span className="login-page__token-label">Access token</span>
              <code className="login-page__token-value">{tokens.accessToken}</code>
            </div>
            <div className="login-page__token">
              <span className="login-page__token-label">Refresh token</span>
              <code className="login-page__token-value">{tokens.refreshToken}</code>
            </div>
          </div>
        )}

        <Link className="login-page__link" to="/signup">
          Register
        </Link>
      </section>
    </main>
  );
}
