import type { ReactNode } from "react";
import { Navigation } from "../../components/navigation/navigation";
import "./page-layout.sass";

type Props = {
  eyebrow: string;
  title: string;
  description: string;
  children?: ReactNode;
};

export function PageLayout({ eyebrow, title, description, children }: Props) {
  return (
    <main className="shell">
      <Navigation />

      <section className="hero">
        <p className="eyebrow">{eyebrow}</p>
        <h1>{title}</h1>
        <p className="lede">{description}</p>
      </section>

      {children}
    </main>
  );
}
