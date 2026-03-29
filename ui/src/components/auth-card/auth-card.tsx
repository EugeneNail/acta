import type { ReactNode } from "react";
import "./auth-card.sass";

type Props = {
  title: string;
  subtitle: string;
  children: ReactNode;
};

export function AuthCard({ title, subtitle, children }: Props) {
  return (
    <section className="authLayout">
      <article className="authCard">
        <p className="panelTag">Auth Flow</p>
        <h2>{title}</h2>
        <p className="authSubtitle">{subtitle}</p>
        {children}
      </article>
    </section>
  );
}
