import "./app-footer.sass";

export function AppFooter() {
  return (
    <footer className="app-footer">
      <div className="app-footer__column">
        <span className="app-footer__label">About</span>
        <strong className="app-footer__placeholder">Product placeholder</strong>
      </div>
      <div className="app-footer__column">
        <span className="app-footer__label">Support</span>
        <strong className="app-footer__placeholder">Links placeholder</strong>
      </div>
      <div className="app-footer__column">
        <span className="app-footer__label">Status</span>
        <strong className="app-footer__placeholder">Footer placeholder</strong>
      </div>
    </footer>
  );
}
