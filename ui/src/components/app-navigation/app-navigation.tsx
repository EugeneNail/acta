import { Link, useLocation } from "react-router-dom";
import "./app-navigation.sass";

type ViewMode = "desktop" | "mobile";

type Props = {
  viewMode: ViewMode;
};

const items = [
  { to: "/", label: "Overview", icon: "OV" },
  { to: "/journal", label: "Journal", icon: "JR" },
  { to: "/habits", label: "Habits", icon: "HB" },
  { to: "/gallery", label: "Photos", icon: "PH" },
  { to: "/signup", label: "Signup", icon: "SU" },
  { to: "/login", label: "Login", icon: "LI" },
  { to: "/account", label: "Account", icon: "AC" },
];

export function AppNavigation({ viewMode }: Props) {
  const location = useLocation();

  return (
    <nav className={`app-navigation app-navigation--${viewMode}`}>
      <div className="app-navigation__brand">
        <div className="app-navigation__logo">
          <div className="app-navigation__logo-mark" />
          <span className="app-navigation__logo-placeholder">logo placeholder</span>
        </div>
        {viewMode === "desktop" && (
          <div className="app-navigation__brand-copy">
            <strong className="app-navigation__brand-name">acta</strong>
            <span className="app-navigation__brand-text">
              daily journal, goals, and photo memories
            </span>
          </div>
        )}
      </div>

      <div className="app-navigation__links">
        {items.map((item) => {
          const isActive = location.pathname === item.to;

          return (
            <Link
              key={item.to}
              className={`app-navigation__link${isActive ? " app-navigation__link--active" : ""}`}
              to={item.to}
            >
              <span className="app-navigation__icon">{item.icon}</span>
              {viewMode === "desktop" && (
                <span className="app-navigation__label">{item.label}</span>
              )}
            </Link>
          );
        })}
      </div>
    </nav>
  );
}
