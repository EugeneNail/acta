import type { ReactNode } from "react";
import { useEffect, useState } from "react";
import { AppFooter } from "../../components/app-footer/app-footer";
import { AppNavigation } from "../../components/app-navigation/app-navigation";
import "./adaptive-layout.sass";

type ViewMode = "desktop" | "mobile";

type Props = {
  eyebrow: string;
  title: string;
  description: string;
  children: ReactNode;
};

function resolveViewMode(): ViewMode {
  if (typeof window === "undefined") {
    return "desktop";
  }

  return window.innerWidth / window.innerHeight > 1 ? "desktop" : "mobile";
}

export function AdaptiveLayout({
  eyebrow,
  title,
  description,
  children,
}: Props) {
  const [viewMode, setViewMode] = useState<ViewMode>(resolveViewMode);

  useEffect(() => {
    function handleResize() {
      setViewMode(resolveViewMode());
    }

    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, []);

  return (
    <div className={`adaptive-layout adaptive-layout--${viewMode}`}>
      <AppNavigation viewMode={viewMode} />
      <div className="adaptive-layout__body">
        <main className="adaptive-layout__content">
          <section className="adaptive-layout__hero">
            <span className="adaptive-layout__eyebrow">{eyebrow}</span>
            <h1 className="adaptive-layout__title">{title}</h1>
            <p className="adaptive-layout__description">{description}</p>
          </section>

          {children}
        </main>

        <AppFooter />
      </div>
    </div>
  );
}
