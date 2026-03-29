import type { FormEvent } from "react";
import { useState } from "react";
import { AuthCard } from "../../components/auth-card/auth-card";
import { ErrorList } from "../../components/error-list/error-list";
import { PageLayout } from "../../layouts/page-layout/page-layout";
import "./login-page.sass";

type LoginResponse = {
  accessToken: string;
  refreshToken: string;
};

export function LoginPage() {
  const [email, setEmail] = useState("login.user@example.com");
  const [password, setPassword] = useState("Strong123!");
  const [errors, setErrors] = useState<Record<string, string>>({});
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [tokens, setTokens] = useState<LoginResponse | null>(null);

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setIsSubmitting(true);
    setErrors({});
    setTokens(null);

    try {
      const response = await fetch("/api/v1/auth/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          email,
          password,
        }),
      });

      const payload = (await response.json()) as LoginResponse | Record<string, string>;

      if (!response.ok) {
        setErrors(payload as Record<string, string>);
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
    <PageLayout
      eyebrow="Acta / Login"
      title="Sign in and store tokens in localStorage."
      description="This page uses the current login contract and persists both issued tokens for later journal work."
    >
      <AuthCard
        title="Login"
        subtitle="The auth service returns accessToken and refreshToken."
      >
        <form className="authForm" onSubmit={handleSubmit}>
          <label className="field">
            <span>Email</span>
            <input
              autoComplete="email"
              name="email"
              type="email"
              value={email}
              onChange={(event) => setEmail(event.target.value)}
            />
          </label>

          <label className="field">
            <span>Password</span>
            <input
              autoComplete="current-password"
              name="password"
              type="password"
              value={password}
              onChange={(event) => setPassword(event.target.value)}
            />
          </label>

          <button className="primaryAction submitButton" disabled={isSubmitting}>
            {isSubmitting ? "Logging in..." : "Login"}
          </button>
        </form>

        <ErrorList errors={errors} />

        {tokens !== null && (
          <div className="tokenPanel">
            <div>
              <p>Access token</p>
              <code>{tokens.accessToken}</code>
            </div>
            <div>
              <p>Refresh token</p>
              <code>{tokens.refreshToken}</code>
            </div>
          </div>
        )}
      </AuthCard>
    </PageLayout>
  );
}
