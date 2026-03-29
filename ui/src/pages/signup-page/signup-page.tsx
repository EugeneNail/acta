import type { FormEvent } from "react";
import { useState } from "react";
import { Link } from "react-router-dom";
import "./signup-page.sass";

type SignupResponse = {
  uuid: string;
};

type ValidationErrors = Record<string, string>;

export function SignupPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [passwordConfirmation, setPasswordConfirmation] = useState("");
  const [errors, setErrors] = useState<ValidationErrors>({});
  const [createdUserUuid, setCreatedUserUuid] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setIsSubmitting(true);
    setErrors({});
    setCreatedUserUuid("");

    try {
      const response = await fetch("/api/v1/auth/signup", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email,
          password,
          passwordConfirmation,
        }),
      });

      const payload = (await response.json()) as SignupResponse | ValidationErrors;

      if (!response.ok) {
        setErrors(payload as ValidationErrors);
        return;
      }

      setCreatedUserUuid((payload as SignupResponse).uuid);
      setPassword("");
      setPasswordConfirmation("");
    } catch (_error) {
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

        {createdUserUuid !== "" && (
          <div className="signup-page__success">
            <span className="signup-page__success-label">User created</span>
            <strong className="signup-page__success-value">{createdUserUuid}</strong>
          </div>
        )}

        <Link className="signup-page__link" to="/login">
          Login
        </Link>
      </section>
    </main>
  );
}
