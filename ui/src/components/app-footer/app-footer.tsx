import "./app-footer.sass";

export function AppFooter() {
  return (
    <footer className="app-footer">
      <div className="app-footer__column">
        <span className="app-footer__label">Demo</span>
        <strong className="app-footer__placeholder">Simplified layout placeholder</strong>
      </div>
      <div className="app-footer__column">
        <span className="app-footer__label">Navigation</span>
        <strong className="app-footer__placeholder">Future product links</strong>
      </div>
      <div className="app-footer__column">
        <span className="app-footer__label">Status</span>
        <strong className="app-footer__placeholder">Demo UI, not final design</strong>
      </div>
    </footer>
  );
}
