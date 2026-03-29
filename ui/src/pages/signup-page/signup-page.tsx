import type { FormEvent } from "react";
import axios from "axios";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { httpClient, saveAuthTokens } from "../../infrastructure/http-client/http-client";
import "./signup-page.sass";

type LoginResponse = {
  accessToken: string;
  refreshToken: string;
};

type ValidationErrors = Record<string, string>;

export function SignupPage() {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [passwordConfirmation, setPasswordConfirmation] = useState("");
  const [errors, setErrors] = useState<ValidationErrors>({});
  const [isSubmitting, setIsSubmitting] = useState(false);

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setIsSubmitting(true);
    setErrors({});

    try {
      await httpClient.post("/api/v1/auth/signup", {
        email,
        password,
        passwordConfirmation,
      });

      const loginResponse = await httpClient.post<LoginResponse>("/api/v1/auth/login", {
        email,
        password,
      });

      saveAuthTokens(loginResponse.data.accessToken, loginResponse.data.refreshToken);
      navigate("/");
    } catch (error) {
      if (axios.isAxiosError<ValidationErrors>(error) && error.response?.data !== undefined) {
        setErrors(error.response.data);
        return;
      }

      setErrors({ request: "Request failed. Try again." });
    } finally {
      setIsSubmitting(false);
    }
  }

  return (
    <main className="signup-page">
      <section className="signup-page__card">
        <h1 className="signup-page__title">Create account</h1>

        <form className="signup-page__form" onSubmit={handleSubmit}>
          <label className="signup-page__field">
            <span className="signup-page__label">Email</span>
            <input
              className="signup-page__input"
              autoComplete="email"
              name="email"
              type="email"
              value={email}
              onChange={(event) => setEmail(event.target.value)}
            />
          </label>

          <label className="signup-page__field">
            <span className="signup-page__label">Password</span>
            <input
              className="signup-page__input"
              autoComplete="new-password"
              name="password"
              type="password"
              value={password}
              onChange={(event) => setPassword(event.target.value)}
            />
          </label>

          <label className="signup-page__field">
            <span className="signup-page__label">Repeat password</span>
            <input
              className="signup-page__input"
              autoComplete="new-password"
              name="passwordConfirmation"
              type="password"
              value={passwordConfirmation}
              onChange={(event) => setPasswordConfirmation(event.target.value)}
            />
          </label>

          <button className="signup-page__submit" disabled={isSubmitting} type="submit">
            {isSubmitting ? "Creating account..." : "Create account"}
          </button>
        </form>

        {Object.keys(errors).length > 0 && (
          <ul className="signup-page__errors">
            {Object.entries(errors).map(([field, message]) => (
              <li className="signup-page__error" key={`${field}-${message}`}>
                <span className="signup-page__error-field">{field}</span>
                <strong className="signup-page__error-message">{message}</strong>
              </li>
            ))}
          </ul>
        )}

        <Link className="signup-page__link" to="/login">
          Login
        </Link>
      </section>
    </main>
  );
}
