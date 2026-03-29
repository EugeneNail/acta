import type { FormEvent } from "react";
import { useState } from "react";
import { AuthCard } from "../../components/auth-card/auth-card";
import { ErrorList } from "../../components/error-list/error-list";
import { PageLayout } from "../../layouts/page-layout/page-layout";
import "./signup-page.sass";

type SignupResponse = {
  uuid: string;
};

export function SignupPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [passwordConfirmation, setPasswordConfirmation] = useState("");
  const [errors, setErrors] = useState<Record<string, string>>({});
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [createdUserUuid, setCreatedUserUuid] = useState("");

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setIsSubmitting(true);
    setErrors({});
    setCreatedUserUuid("");

    try {
      const response = await fetch("/api/v1/auth/signup", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          email,
          password,
          passwordConfirmation,
        }),
      });

      const payload = (await response.json()) as SignupResponse | Record<string, string>;

      if (!response.ok) {
        setErrors(payload as Record<string, string>);
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
    <PageLayout
      eyebrow="Acta / Signup"
      title="Create a user through the real auth backend."
      description="This page sends the current signup payload shape and renders validation errors returned by the service."
    >
      <AuthCard
        title="Registration"
        subtitle="The backend expects email, password, and passwordConfirmation."
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
              autoComplete="new-password"
              name="password"
              type="password"
              value={password}
              onChange={(event) => setPassword(event.target.value)}
            />
          </label>

          <label className="field">
            <span>Repeat password</span>
            <input
              autoComplete="new-password"
              name="passwordConfirmation"
              type="password"
              value={passwordConfirmation}
              onChange={(event) => setPasswordConfirmation(event.target.value)}
            />
          </label>

          <button className="primaryAction submitButton" disabled={isSubmitting}>
            {isSubmitting ? "Creating account..." : "Create account"}
          </button>
        </form>

        <ErrorList errors={errors} />

        {createdUserUuid !== "" && (
          <div className="successPanel">
            <p>Account created.</p>
            <strong>{createdUserUuid}</strong>
          </div>
        )}
      </AuthCard>
    </PageLayout>
  );
}
